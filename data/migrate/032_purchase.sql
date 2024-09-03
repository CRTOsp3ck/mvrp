-- +migrate Up
CREATE SCHEMA IF NOT EXISTS purchase;

CREATE TYPE purchase.purchase_order_status AS ENUM (
    'Pending', 
    'Accepted', 
    'Declined'
);

CREATE TABLE purchase.purchase_order (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    purchase_order_number VARCHAR(50) UNIQUE NOT NULL,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information TEXT,
    payment_due_date DATE,
    order_status purchase.purchase_order_status NOT NULL
);

CREATE TABLE purchase.purchase_order_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    purchase_order_id INT NOT NULL REFERENCES purchase.purchase_order(id) ON DELETE CASCADE
);

CREATE TABLE purchase.goods_receipt_note (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    goods_receipt_note_number VARCHAR(50) UNIQUE NOT NULL,
    purchase_order_id INT NOT NULL REFERENCES purchase.purchase_order(id) ON DELETE CASCADE,
    receipt_date DATE NOT NULL,
    vendor_id INT REFERENCES entity.entity(id),
    ship_from_information TEXT,
    shipping_personnel_information TEXT,
    receiving_location_information TEXT,
    receiving_personnel_information TEXT,
    goods_received_condition TEXT
);

CREATE TABLE purchase.goods_receipt_note_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    goods_receipt_note_id INT NOT NULL REFERENCES purchase.goods_receipt_note(id) ON DELETE CASCADE
);

CREATE TABLE purchase.request_for_quotation (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    request_for_quotation_number VARCHAR(50) UNIQUE NOT NULL,
    valid_until_date DATE,
    vendor_id INT REFERENCES entity.entity(id),
    customer_id INT REFERENCES entity.entity(id),
    ship_to_information TEXT,
    requested_by TEXT
);

CREATE TABLE purchase.request_for_quotation_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    request_for_quotation_id INT NOT NULL REFERENCES purchase.request_for_quotation(id) ON DELETE CASCADE
);

CREATE VIEW purchase.purchase_order_item_view AS
SELECT
    poi.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = poi.base_document_item_id
    ) AS base_document_item
FROM
    purchase.purchase_order_item poi;

CREATE VIEW purchase.purchase_order_view AS
SELECT
    po.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = po.base_document_id
    ) AS base_document,
    (
        SELECT json_agg(row_to_json(poiv))
        FROM purchase.purchase_order_item_view poiv
        WHERE poiv.purchase_order_id = po.id
    ) AS purchase_order_items
FROM
    purchase.purchase_order po;

CREATE VIEW purchase.goods_receipt_note_item_view AS
SELECT
    grni.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = grni.base_document_item_id
    ) AS base_document_item
FROM
    purchase.goods_receipt_note_item grni;

CREATE VIEW purchase.goods_receipt_note_view AS
SELECT
    grn.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = grn.base_document_id
    ) AS base_document,
    (
        SELECT json_agg(row_to_json(grniv))
        FROM purchase.goods_receipt_note_item_view grniv
        WHERE grniv.goods_receipt_note_id = grn.id
    ) AS goods_receipt_note_items
FROM
    purchase.goods_receipt_note grn;

CREATE VIEW purchase.request_for_quotation_item_view AS
SELECT
    rfqi.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = rfqi.base_document_item_id
    ) AS base_document_item
FROM
    purchase.request_for_quotation_item rfqi;

CREATE VIEW purchase.request_for_quotation_view AS
SELECT
    rfq.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = rfq.base_document_id
    ) AS base_document,
    (
        SELECT json_agg(row_to_json(rfqiv))
        FROM purchase.request_for_quotation_item_view rfqiv
        WHERE rfqiv.request_for_quotation_id = rfq.id
    ) AS request_for_quotation_items
FROM
    purchase.request_for_quotation rfq;

-- +migrate Down
DROP VIEW IF EXISTS purchase.request_for_quotation_view;
DROP VIEW IF EXISTS purchase.request_for_quotation_item_view;

DROP VIEW IF EXISTS purchase.goods_receipt_note_view;
DROP VIEW IF EXISTS purchase.goods_receipt_note_item_view;

DROP VIEW IF EXISTS purchase.purchase_order_view;
DROP VIEW IF EXISTS purchase.purchase_order_item_view;

DROP TABLE IF EXISTS purchase.request_for_quotation_item;
DROP TABLE IF EXISTS purchase.request_for_quotation;

DROP TABLE IF EXISTS purchase.goods_receipt_note_item;
DROP TABLE IF EXISTS purchase.goods_receipt_note;

DROP TABLE IF EXISTS purchase.purchase_order_item;
DROP TABLE IF EXISTS purchase.purchase_order;

DROP TYPE IF EXISTS purchase.purchase_order_status;