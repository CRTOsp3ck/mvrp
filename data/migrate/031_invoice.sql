-- +migrate Up
CREATE SCHEMA IF NOT EXISTS invoice;

CREATE TABLE invoice.invoice (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    vendor_id INT DEFAULT 0,
    customer_id INT DEFAULT 0,
    payment_due_date DATE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE invoice.invoice_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    invoice_id INT NOT NULL REFERENCES invoice.invoice(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE invoice.payment_receipt (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    payment_receipt_number VARCHAR(50) UNIQUE NOT NULL,
    invoice_id INT NOT NULL REFERENCES invoice.invoice(id) ON DELETE CASCADE,
    date_of_payment DATE,
    payer_id INT REFERENCES entity.entity(id),
    payee_id INT REFERENCES entity.entity(id),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE invoice.payment_receipt_item (
    id INT PRIMARY KEY,
    base_document_item_id INT NOT NULL REFERENCES base.base_document_item(id) ON DELETE CASCADE,
    payment_receipt_id INT NOT NULL REFERENCES invoice.payment_receipt(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE invoice.credit_note (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    credit_note_number VARCHAR(50) UNIQUE NOT NULL,
    reference_invoice_id INT REFERENCES invoice.invoice(id),
    reason_for_issuance TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE invoice.debit_note (
    id INT PRIMARY KEY,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    debit_note_number VARCHAR(50) UNIQUE NOT NULL,
    reference_invoice_id INT REFERENCES invoice.invoice(id),
    additional_charges NUMERIC(12, 2) DEFAULT 0,
    reason_for_issuance TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE VIEW invoice.credit_note_view AS
SELECT
    cn.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = cn.base_document_id
    ) AS base_document
FROM
    invoice.credit_note cn;

CREATE VIEW invoice.debit_note_view AS
SELECT
    dn.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = dn.base_document_id
    ) AS base_document
FROM
    invoice.debit_note dn;

CREATE VIEW invoice.payment_receipt_item_view AS
SELECT
    pri.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = pri.base_document_item_id
    ) AS base_document_item
FROM
    invoice.payment_receipt_item pri;

CREATE VIEW invoice.payment_receipt_view AS
SELECT
    pr.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = pr.base_document_id
    ) AS base_document,
    (
        SELECT json_agg(row_to_json(priv))
        FROM invoice.payment_receipt_item_view priv
        WHERE priv.payment_receipt_id = pr.id
    ) AS payment_receipt_items
FROM
    invoice.payment_receipt pr;

CREATE VIEW invoice.invoice_item_view AS
SELECT
    ii.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = ii.base_document_item_id
    ) AS base_document_item
FROM
    invoice.invoice_item ii;

CREATE VIEW invoice.invoice_view AS
SELECT
    i.*,
    (
        SELECT row_to_json(bd)
        FROM base.base_document bd
        WHERE bd.id = i.base_document_id
    ) AS base_document,
    (
        SELECT json_agg(row_to_json(iiv))
        FROM invoice.invoice_item_view iiv
        WHERE iiv.invoice_id = i.id
    ) AS invoice_items
FROM
    invoice.invoice i;

-- +migrate Down
DROP VIEW IF EXISTS invoice.invoice_view;
DROP VIEW IF EXISTS invoice.invoice_item_view;

DROP VIEW IF EXISTS invoice.payment_receipt_view;
DROP VIEW IF EXISTS invoice.payment_receipt_item_view;

DROP VIEW IF EXISTS invoice.debit_note_view;
DROP VIEW IF EXISTS invoice.credit_note_view;

DROP TABLE IF EXISTS invoice.debit_note;
DROP TABLE IF EXISTS invoice.credit_note;

DROP TABLE IF EXISTS invoice.payment_receipt_item;
DROP TABLE IF EXISTS invoice.payment_receipt;

DROP TABLE IF EXISTS invoice.invoice_item;
DROP TABLE IF EXISTS invoice.invoice;