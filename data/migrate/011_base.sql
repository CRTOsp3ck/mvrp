-- +migrate Up
CREATE SCHEMA IF NOT EXISTS base;

CREATE TYPE base.payment_terms AS ENUM (
    'full', 
    'partial_before_and_after_delivery', 
    'net_30',
    'net_60',
    'net_90'
);

CREATE TYPE base.payment_status AS ENUM (
    'pending', 
    'paid', 
    'partially_paid'
);

CREATE TYPE base.shipping_terms AS ENUM (
    'free_on_board',  -- Free on Board
    'cost_insurance_and_freight',  -- Cost, Insurance, and Freight
    'ex_works',  -- Ex Works
    'delivered_duty_paid',  -- Delivered Duty Paid
    'delivered_at_place',  -- Delivered at Place
    'free_carrier',  -- Free Carrier
    'carriage_paid_to'   -- Carriage Paid To
);

CREATE TYPE base.shipping_method AS ENUM (
    'air_freight',   -- Air Freight
    'sea_freight',   -- Sea Freight
    'ground_shipping', -- Ground Shipping (Truck or Rail)
    'courier_service', -- Courier (DHL, FedEx, UPS)
    'rail_freight',  -- Rail Freight
    'multimodal',    -- Combination of different shipping modes
    'drop_shipping', -- Supplier ships directly to customer
    '3pl',           -- Third-Party Logistics
    'expedited',     -- Expedited shipping for faster delivery
    'standard',      -- Standard shipping option
    'same_day',      -- Same-day delivery
    'next_day',      -- Next-day delivery
    'economy'        -- Economy shipping, slower but cheaper
);

-- tax_amount - to be inserted by the user according to items
-- shipping_fees - to be inserted by the user according to items
CREATE TABLE base.base_document (
    id INT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    issue_date TIMESTAMPTZ DEFAULT (now() at TIME ZONE 'UTC'),
    gross_amount_gen NUMERIC(12, 2),
    discount_amount_gen NUMERIC(12, 2) DEFAULT 0,
    discount_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS (discount_amount_gen / gross_amount_gen * 100) STORED,
    additional_discount_amount NUMERIC(12, 2) DEFAULT 0,
    additional_discount_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS (additional_discount_amount / gross_amount_gen * 100) STORED,
    gross_amount_after_discount_gen NUMERIC(12, 2) GENERATED ALWAYS AS (gross_amount_gen - discount_amount_gen - additional_discount_amount) STORED,
    tax_amount_gen NUMERIC(12, 2),
    tax_rate_gen NUMERIC(5, 2) GENERATED ALWAYS AS (tax_amount_gen / gross_amount_gen * 100) STORED,
    shipping_fees_gen NUMERIC(12, 2) DEFAULT 0,
    other_fees NUMERIC(12, 2) DEFAULT 0,
    custom_adjustment_amount NUMERIC(12, 2) DEFAULT 0,
    net_amount_gen NUMERIC(12, 2) GENERATED ALWAYS AS (gross_amount_gen - discount_amount_gen - additional_discount_amount + tax_amount_gen + shipping_fees_gen + other_fees + custom_adjustment_amount) STORED,
    shipping_terms base.shipping_terms,
    shipping_method base.shipping_method,
    shipping_date TIMESTAMPTZ,
    payment_terms base.payment_terms,
    payment_instructions TEXT,
    payment_status base.payment_status,
    remarks TEXT,
    terms_and_conditions TEXT
);

CREATE TABLE base.base_document_item (
    id INT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    base_document_id INT NOT NULL REFERENCES base.base_document(id) ON DELETE CASCADE,
    -- item_id INT,
    -- item_code VARCHAR(50),
    -- item_name VARCHAR(255),
    -- item_description TEXT,
    inventory_id INT,
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

DROP TYPE IF EXISTS base.shipping_method;
DROP TYPE IF EXISTS base.shipping_terms;
DROP TYPE IF EXISTS base.payment_status;
DROP TYPE IF EXISTS base.payment_terms;