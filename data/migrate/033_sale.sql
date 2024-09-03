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

CREATE TABLE sale.sales_order (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    sales_order_number VARCHAR(50) UNIQUE NOT NULL,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    sales_representative_information TEXT,
    ship_to_information TEXT,
    ship_from_information TEXT,
    payment_due_date DATE,
    order_status sale.sales_order_status NOT NULL
);

CREATE TABLE sale.sales_order_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    sales_order_id INT NOT NULL REFERENCES sale.sales_order(id) ON DELETE CASCADE
);

CREATE TABLE sale.delivery_note (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    delivery_note_number VARCHAR(50) UNIQUE NOT NULL,
    sales_order_id INT NOT NULL REFERENCES sale.sales_order(id) ON DELETE CASCADE,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information TEXT,
    ship_from_information TEXT,
    bill_to_information TEXT,
    delivery_date DATE,
    shipping_personnel_information TEXT,
    received_by TEXT,
    goods_condition TEXT
);

CREATE TABLE sale.delivery_note_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    delivery_note_id INT NOT NULL REFERENCES sale.delivery_note(id) ON DELETE CASCADE
);

CREATE TABLE sale.goods_return_note (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    goods_return_note_number VARCHAR(50) UNIQUE NOT NULL,
    sales_order_id INT REFERENCES sale.sales_order(id) ON DELETE CASCADE,
    invoice_id INT REFERENCES invoice.invoice(id),
    credit_note_id INT REFERENCES invoice.credit_note(id),
    rma_id INT REFERENCES inventory.return_merchandise_authorization(id),
    issue_date DATE NOT NULL,
    return_date DATE,
    customer_id INT REFERENCES entity.entity(id),
    receiving_location_information TEXT,
    received_by TEXT,
    overall_goods_condition TEXT,
    return_reason TEXT
);

CREATE TABLE sale.goods_return_note_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    goods_return_note_id INT NOT NULL REFERENCES sale.goods_return_note(id) ON DELETE CASCADE,
    rma_item_id INT REFERENCES inventory.return_merchandise_authorization_item(id),
    return_quantity NUMERIC(12, 2) DEFAULT 0,
    return_condition TEXT
);

CREATE TABLE sale.order_confirmation (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    order_confirmation_number VARCHAR(50) UNIQUE NOT NULL,
    sales_order_id INT NOT NULL REFERENCES sale.sales_order(id) ON DELETE CASCADE,
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information TEXT
);

CREATE TABLE sale.order_confirmation_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    order_confirmation_id INT NOT NULL REFERENCES sale.order_confirmation(id) ON DELETE CASCADE
);

CREATE TABLE sale.sales_quotation (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    sales_quotation_number VARCHAR(50) UNIQUE NOT NULL,
    valid_until_date DATE,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information TEXT,
    requested_by TEXT,
    prepared_by TEXT,
    quotation_status sale.sales_quotation_status NOT NULL
);

CREATE TABLE sale.sales_quotation_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    sales_quotation_id INT NOT NULL REFERENCES sale.sales_quotation(id) ON DELETE CASCADE
);

CREATE VIEW sale.sales_order_item_view AS
SELECT
    soi.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = soi.base_document_item_id
    ) AS base_document_item
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
    ) AS base_document_item
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
    ) AS base_document_item
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
    ) AS base_document_item
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

DROP TYPE IF EXISTS sale.sales_quotation_status;

DROP TYPE IF EXISTS sale.sales_order_status;
