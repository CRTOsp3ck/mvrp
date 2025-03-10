-- +migrate Up
CREATE SCHEMA IF NOT EXISTS sale;

CREATE TYPE sale.sales_order_status AS ENUM (
    'pending', 
    'accepted', 
    'declined'
);

CREATE TYPE sale.sales_quotation_status AS ENUM (
    'pending', 
    'accepted', 
    'declined'
);

CREATE TYPE sale.shipping_status AS ENUM (
    'ready_for_pickup',
    'in_transit',
    'shipped'
);

CREATE TABLE sale.sales_order (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    sales_order_number VARCHAR(50) UNIQUE NOT NULL,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    sales_representative_employee_id INT REFERENCES entity.entity(id),
    ship_to_information JSONB,
    ship_from_information JSONB,
    payment_due_date DATE,
    order_status sale.sales_order_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.sales_order_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    sales_order_id INT NOT NULL REFERENCES sale.sales_order(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.delivery_note (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    delivery_note_number VARCHAR(50) UNIQUE NOT NULL,
    sales_order_id INT NOT NULL REFERENCES sale.sales_order(id) ON DELETE CASCADE,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information JSONB NOT NULL,
    ship_from_information JSONB NOT NULL,
    bill_to_information JSONB NOT NULL,
    shipping_date DATE NOT NULL,
    shipping_personnel_information JSONB,
    shipping_status sale.shipping_status NOT NULL,
    received_by JSONB,
    overall_goods_condition TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.delivery_note_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    delivery_note_id INT NOT NULL REFERENCES sale.delivery_note(id) ON DELETE CASCADE,
    goods_condition TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.goods_return_note (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    goods_return_note_number VARCHAR(50) UNIQUE NOT NULL,
    return_date DATE NOT NULL,
    returned_by_customer_id INT REFERENCES entity.entity(id),
    receiving_location_information JSONB,
    received_by_employee_id INT REFERENCES entity.entity(id),
    overall_goods_condition TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.goods_return_note_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    goods_return_note_id INT NOT NULL REFERENCES sale.goods_return_note(id) ON DELETE CASCADE,
    rma_item_id INT REFERENCES inventory.return_merchandise_authorization_item(id),
    sales_order_item_id INT REFERENCES sale.sales_order_item(id) ON DELETE CASCADE,
    invoice_id INT REFERENCES invoice.invoice(id),
    credit_note_id INT REFERENCES invoice.credit_note(id),
    return_quantity NUMERIC(12, 2) DEFAULT 0,
    return_condition TEXT,
    return_reason TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.order_confirmation (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    order_confirmation_number VARCHAR(50) UNIQUE NOT NULL,
    sales_order_id INT NOT NULL REFERENCES sale.sales_order(id) ON DELETE CASCADE,
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information JSONB,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.order_confirmation_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    order_confirmation_id INT NOT NULL REFERENCES sale.order_confirmation(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.sales_quotation (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    sales_quotation_number VARCHAR(50) UNIQUE NOT NULL,
    valid_until_date DATE,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information JSONB,
    requested_by JSONB,
    prepared_by_employee_id INT REFERENCES entity.entity(id),
    quotation_status sale.sales_quotation_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE sale.sales_quotation_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    sales_quotation_id INT NOT NULL REFERENCES sale.sales_quotation(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE VIEW sale.sales_order_item_view AS
SELECT
    soi.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = soi.base_document_item_id
    ) AS base_document_item,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = (
            SELECT bdi.inventory_id
            FROM base.base_document_item bdi
            WHERE bdi.id = soi.base_document_item_id
        )
    ) AS inventory_info
FROM
    sale.sales_order_item soi;

CREATE VIEW sale.sales_order_view AS
SELECT
    so.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = so.base_document_id
    ) AS base_document,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = so.vendor_id
    ) AS vendor_info,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = so.customer_id
    ) AS customer_info,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = so.sales_representative_employee_id
    ) AS sales_representative_info,
    (
        SELECT json_agg(row_to_json(soiv))
        FROM sale.sales_order_item_view soiv
        WHERE soiv.sales_order_id = so.id
    ) AS sales_order_items
FROM
    sale.sales_order so;

CREATE VIEW sale.delivery_note_item_view AS
SELECT
    dni.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = dni.base_document_item_id
    ) AS base_document_item,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = (
            SELECT bdi.inventory_id
            FROM base.base_document_item bdi
            WHERE bdi.id = dni.base_document_item_id
        )
    ) AS inventory_info
FROM
    sale.delivery_note_item dni;

CREATE VIEW sale.delivery_note_view AS
SELECT
    dn.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = dn.base_document_id
    ) AS base_document,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = dn.vendor_id
    ) AS vendor_info,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = dn.customer_id
    ) AS customer_info,
    (
        SELECT json_agg(row_to_json(dniv))
        FROM sale.delivery_note_item_view dniv
        WHERE dniv.delivery_note_id = dn.id
    ) AS delivery_note_items
FROM
    sale.delivery_note dn;

CREATE VIEW sale.goods_return_note_item_view AS
SELECT
    grni.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = grni.base_document_item_id
    ) AS base_document_item,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = (
            SELECT bdi.inventory_id
            FROM base.base_document_item bdi
            WHERE bdi.id = grni.base_document_item_id
        )
    ) AS inventory_info
FROM
    sale.goods_return_note_item grni;

CREATE VIEW sale.goods_return_note_view AS
SELECT
    grn.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = grn.base_document_id
    ) AS base_document,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = grn.returned_by_customer_id
    ) AS customer_info,
    (
        SELECT json_agg(row_to_json(grniv))
        FROM sale.goods_return_note_item_view grniv
        WHERE grniv.goods_return_note_id = grn.id
    ) AS goods_return_note_items
FROM
    sale.goods_return_note grn;

CREATE VIEW sale.order_confirmation_item_view AS
SELECT
    oci.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = oci.base_document_item_id
    ) AS base_document_item
FROM
    sale.order_confirmation_item oci;

CREATE VIEW sale.order_confirmation_view AS
SELECT
    oc.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = oc.base_document_id
    ) AS base_document,
    (
        SELECT json_agg(row_to_json(ociv))
        FROM sale.order_confirmation_item_view ociv
        WHERE ociv.order_confirmation_id = oc.id
    ) AS order_confirmation_items
FROM
    sale.order_confirmation oc;

CREATE VIEW sale.sales_quotation_item_view AS
SELECT
    sqi.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = sqi.base_document_item_id
    ) AS base_document_item,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = (
            SELECT bdi.inventory_id
            FROM base.base_document_item bdi
            WHERE bdi.id = sqi.base_document_item_id
        )
    ) AS inventory_info
FROM
    sale.sales_quotation_item sqi;

CREATE VIEW sale.sales_quotation_view AS
SELECT
    sq.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = sq.base_document_id
    ) AS base_document,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = sq.vendor_id
    ) AS vendor_info,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = sq.customer_id
    ) AS customer_info,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = sq.prepared_by_employee_id
    ) AS prepared_by_employee_info,
    (
        SELECT json_agg(row_to_json(sqiv))
        FROM sale.sales_quotation_item_view sqiv
        WHERE sqiv.sales_quotation_id = sq.id
    ) AS sales_quotation_items
FROM
    sale.sales_quotation sq;

-- +migrate Down
DROP VIEW IF EXISTS sale.sales_quotation_view;
DROP VIEW IF EXISTS sale.sales_quotation_item_view;

DROP VIEW IF EXISTS sale.order_confirmation_view;
DROP VIEW IF EXISTS sale.order_confirmation_item_view;

DROP VIEW IF EXISTS sale.goods_return_note_view;
DROP VIEW IF EXISTS sale.goods_return_note_item_view;

DROP VIEW IF EXISTS sale.delivery_note_view;
DROP VIEW IF EXISTS sale.delivery_note_item_view;

DROP VIEW IF EXISTS sale.sales_order_view;
DROP VIEW IF EXISTS sale.sales_order_item_view;


DROP TABLE IF EXISTS sale.sales_quotation_item;
DROP TABLE IF EXISTS sale.sales_quotation;

DROP TABLE IF EXISTS sale.order_confirmation_item;
DROP TABLE IF EXISTS sale.order_confirmation;

DROP TABLE IF EXISTS sale.goods_return_note_item;
DROP TABLE IF EXISTS sale.goods_return_note;

DROP TABLE IF EXISTS sale.delivery_note_item;
DROP TABLE IF EXISTS sale.delivery_note;

DROP TABLE IF EXISTS sale.sales_order_item;
DROP TABLE IF EXISTS sale.sales_order;

DROP TYPE IF EXISTS sale.shipping_status;
DROP TYPE IF EXISTS sale.sales_quotation_status;
DROP TYPE IF EXISTS sale.sales_order_status;
