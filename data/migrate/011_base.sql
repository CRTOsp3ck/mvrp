-- +migrate Up
CREATE SCHEMA IF NOT EXISTS base;

CREATE TYPE base.payment_terms AS ENUM (
    'Full', 
    'PartialBeforeAndAfterDelivery', 
    'Net30',
    'Net60',
    'Net90'
);

CREATE TYPE base.payment_status AS ENUM (
    'Pending', 
    'Paid', 
    'PartiallyPaid'
);

-- tax_amount - to be inserted by the user according to items
-- shipping_fees - to be inserted by the user according to items
CREATE TABLE base.base_document (
    id INT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    issue_date TIMESTAMPTZ DEFAULT (now() at TIME ZONE 'UTC'),
    gross_amount NUMERIC(12, 2),
    discount_amount NUMERIC(12, 2) DEFAULT 0,
    discount_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS (discount_amount / gross_amount * 100) STORED,
    additional_discount_amount NUMERIC(12, 2) DEFAULT 0,
    additional_discount_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS (additional_discount_amount / gross_amount * 100) STORED,
    gross_amount_after_discount_gen NUMERIC(12, 2) GENERATED ALWAYS AS (gross_amount - discount_amount - additional_discount_amount) STORED,
    tax_amount NUMERIC(12, 2),
    tax_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS (tax_amount / gross_amount * 100) STORED,
    shipping_fees NUMERIC(12, 2) DEFAULT 0,
    other_fees NUMERIC(12, 2) DEFAULT 0,
    custom_adjustment_amount NUMERIC(12, 2) DEFAULT 0,
    net_amount_gen NUMERIC(12, 2) GENERATED ALWAYS AS (gross_amount - discount_amount - additional_discount_amount + tax_amount + shipping_fees + other_fees + custom_adjustment_amount) STORED,
    shipping_terms TEXT,
    shipping_method VARCHAR(100),
    shipping_date TIMESTAMPTZ,
    payment_terms base.payment_terms NOT NULL,
    payment_instructions TEXT,
    payment_status base.payment_status NOT NULL,
    remarks TEXT,
    terms_and_conditions TEXT
);

CREATE TABLE base.base_document_item (
    id INT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    item_id INT,
    item_code VARCHAR(50),
    item_name VARCHAR(255),
    item_description TEXT,
    quantity NUMERIC(12, 2) DEFAULT 0,
    unit_price NUMERIC(12, 2) DEFAULT 0,
    unit_discount_amount NUMERIC(12, 2) DEFAULT 0,
    total_discount_amount_gen NUMERIC(12, 2) GENERATED ALWAYS AS (unit_discount_amount * quantity) STORED,
    discount_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS (unit_discount_amount / unit_price * 100) STORED,
    unit_tax_amount NUMERIC(12, 2) DEFAULT 0,
    total_tax_amount_gen NUMERIC(12, 2) GENERATED ALWAYS AS (unit_tax_amount * quantity) STORED,
    tax_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS ((unit_tax_amount/unit_price) * 100) STORED,
    unit_shipping_fees NUMERIC(12, 2) DEFAULT 0,
    total_shipping_fees_gen NUMERIC(12, 2) GENERATED ALWAYS AS (unit_shipping_fees * quantity) STORED,
    final_unit_price_gen NUMERIC(12, 2) GENERATED ALWAYS AS (unit_price - unit_discount_amount + unit_tax_amount + unit_shipping_fees) STORED,
    total_sale_price_gen NUMERIC(12, 2) GENERATED ALWAYS AS (quantity * (unit_price - unit_discount_amount + unit_tax_amount + unit_shipping_fees)) STORED
);

-- +migrate Down
DROP TABLE IF EXISTS base.base_document_item;
DROP TABLE IF EXISTS base.base_document;

DROP TYPE IF EXISTS base.payment_status;
DROP TYPE IF EXISTS base.payment_terms;