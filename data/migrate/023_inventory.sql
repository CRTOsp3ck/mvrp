-- +migrate Up
CREATE SCHEMA IF NOT EXISTS inventory;

CREATE TYPE inventory.inventory_transaction_type AS ENUM (
    'sale', 
    'purchase', 
    'transfer',
    'issuance',
    'return',
    'shipping',
    'stock_count', 
    'sale_cancellation',
    'purchase_cancellation',
    'transfer_cancellation',
    'issuance_cancellation',
    'return_cancellation',
    'shipping_cancellation',
    'stock_count_cancellation',
    'sale_adjustment',
    'purchase_adjustment',
    'transfer_adjustment',
    'issuance_adjustment',
    'return_adjustment',
    'shipping_adjustment',
    'stock_count_adjustment',
    'general_adjustment',
    'initial_stock'
);

CREATE TABLE inventory.inventory (
    id INT PRIMARY KEY,
    inventory_number VARCHAR(50) UNIQUE NOT NULL,
    item_id INT REFERENCES item.item(id),
    quantity_reserved DECIMAL(20, 5) DEFAULT 0, -- Quantity reserved for orders
    quantity_available DECIMAL(20, 5) DEFAULT 0, -- Quantity available for sale
    quantity_returned DECIMAL(20, 5) DEFAULT 0, -- Quantity returned by customers
    quantity_total_gen DECIMAL(20, 5) DEFAULT 0, -- Current stock level
    cost_per_unit DECIMAL(20, 5) DEFAULT 0,
    price_per_unit DECIMAL(20, 5) DEFAULT 0,
    total_value_on_hand_gen DECIMAL(20, 5) DEFAULT 0,
    reorder_level DECIMAL(20, 5) DEFAULT 0,
    reorder_quantity DECIMAL(20, 5) DEFAULT 0,
    is_discontinued BOOLEAN DEFAULT FALSE,
    remarks TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE inventory.inventory_periodic_record (
    id INT PRIMARY KEY,
    inventory_id INT REFERENCES inventory.inventory(id) ON DELETE CASCADE NOT NULL,
    record_start_date DATE DEFAULT CURRENT_DATE,
    record_start_quantity_reserved DECIMAL(20, 5) DEFAULT 0, -- Quantity reserved for orders
    record_start_quantity_available DECIMAL(20, 5) DEFAULT 0, -- Quantity available for sale
    record_start_quantity_returned DECIMAL(20, 5) DEFAULT 0, -- Quantity returned by customers
    record_start_quantity_total DECIMAL(20, 5),
    record_end_date DATE,
    record_end_quantity_reserved DECIMAL(20, 5) DEFAULT 0, 
    record_end_quantity_available DECIMAL(20, 5) DEFAULT 0,
    record_end_quantity_returned DECIMAL(20, 5) DEFAULT 0, 
    record_end_quantity_total DECIMAL(20, 5),
    remarks TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE inventory.inventory_transaction (
    id INT PRIMARY KEY,
    inventory_id INT REFERENCES inventory.inventory(id),
    transaction_type inventory.inventory_transaction_type NOT NULL,
    quantity DECIMAL(20, 5) NOT NULL, -- include +/- sign
    reason TEXT,
    transaction_date TIMESTAMPTZ DEFAULT (NOW() AT TIME ZONE 'UTC'),
    created_by TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE inventory.goods_issue_note (
    id INT PRIMARY KEY,
    gin_number VARCHAR(50) UNIQUE NOT NULL,
    receipient_id INT REFERENCES entity.entity(id),
    issue_date DATE DEFAULT CURRENT_DATE,
    total_value_gen NUMERIC(15, 2) NOT NULL,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE inventory.goods_issue_note_item (
    id INT PRIMARY KEY,
    gin_id INT REFERENCES inventory.goods_issue_note(id),
    inventory_id INT REFERENCES inventory.inventory(id),
    quantity DECIMAL(20,5) NOT NULL,
    unit_value NUMERIC(15, 2) NOT NULL,
    total_value_gen NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- stock count sheet is used to record the physical count of inventory items
-- discrepenacies is the difference between the physical count and the system count
CREATE TABLE inventory.stock_count_sheet (
    id INT PRIMARY KEY,
    scs_number VARCHAR(50) UNIQUE NOT NULL,
    inventory_id INT REFERENCES inventory.inventory(id),
    count_date DATE DEFAULT CURRENT_DATE,
    total_quantity DECIMAL(20,5) NOT NULL,
    discrepancies DECIMAL(20,5),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE inventory.return_merchandise_authorization (
    id INT PRIMARY KEY,
    rma_number VARCHAR(50) UNIQUE NOT NULL,
    rma_date DATE DEFAULT CURRENT_DATE,
    total_value_gen NUMERIC(15, 2) NOT NULL,
    received_by_employee_id INT REFERENCES entity.entity(id),
    returned_by_customer_id INT REFERENCES entity.entity(id),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE inventory.return_merchandise_authorization_item (
    id INT PRIMARY KEY,
    rma_id INT REFERENCES inventory.return_merchandise_authorization(id),
    inventory_id INT REFERENCES inventory.inventory(id),
    quantity DECIMAL(20,5) NOT NULL,
    unit_value NUMERIC(15, 2) NOT NULL,
    total_value_gen NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE VIEW inventory.inventory_view AS
SELECT
    i.*,
    (
        SELECT row_to_json(it)
        FROM item.item it
        WHERE it.id = i.item_id
    ) AS item,
    (
        SELECT json_agg(row_to_json(it))
        FROM inventory.inventory_transaction it
        WHERE it.inventory_id = i.id
    ) AS transactions
FROM
    inventory.inventory i;

CREATE VIEW inventory.goods_issue_note_item_view AS
SELECT
    gin_item.*,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = gin_item.inventory_id
    ) AS inventory
FROM
    inventory.goods_issue_note_item gin_item;

CREATE VIEW inventory.goods_issue_note_view AS
SELECT
    gin.*,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = gin.receipient_id
    ) AS receipient,
    (
        SELECT json_agg(row_to_json(gin_item))
        FROM inventory.goods_issue_note_item_view gin_item
        WHERE gin_item.gin_id = gin.id
    ) AS items
FROM
    inventory.goods_issue_note gin;

CREATE VIEW inventory.return_merchandise_authorization_item_view AS
SELECT
    rma_item.*,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = rma_item.inventory_id
    ) AS inventory
FROM
    inventory.return_merchandise_authorization_item rma_item;

CREATE VIEW inventory.return_merchandise_authorization_view AS
SELECT
    rma.*,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = rma.returned_by_customer_id
    ) AS returned_by_info,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = rma.received_by_employee_id
    ) AS received_by_info,
    (
        SELECT json_agg(row_to_json(rma_item))
        FROM inventory.return_merchandise_authorization_item_view rma_item
        WHERE rma_item.rma_id = rma.id
    ) AS items
FROM
    inventory.return_merchandise_authorization rma;

-- +migrate Down
DROP VIEW IF EXISTS inventory.return_merchandise_authorization_view;
DROP VIEW IF EXISTS inventory.return_merchandise_authorization_item_view;
DROP VIEW IF EXISTS inventory.goods_issue_note_view;
DROP VIEW IF EXISTS inventory.goods_issue_note_item_view;
DROP VIEW IF EXISTS inventory.inventory_view;
DROP TABLE IF EXISTS inventory.return_merchandise_authorization_item;
DROP TABLE IF EXISTS inventory.return_merchandise_authorization;
DROP TABLE IF EXISTS inventory.stock_count_sheet;
DROP TABLE IF EXISTS inventory.goods_issue_note_item;
DROP TABLE IF EXISTS inventory.goods_issue_note;
DROP TABLE IF EXISTS inventory.inventory_transaction;
DROP TABLE IF EXISTS inventory.inventory_periodic_record;
DROP TABLE IF EXISTS inventory.inventory;
DROP TYPE IF EXISTS inventory.inventory_transaction_type;