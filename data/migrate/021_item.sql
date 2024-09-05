-- +migrate Up
CREATE SCHEMA IF NOT EXISTS item;

CREATE TYPE item.item_type AS ENUM (
    'product', 
    'service',
    'other'
);

CREATE TYPE item.item_status AS ENUM (
    'active',
    'inactive',
    'discontinued',
    'testing'
);

CREATE TABLE item.item (
    id INT PRIMARY KEY,
    code VARCHAR NOT NULL,
    sku VARCHAR NOT NULL,
    brand VARCHAR NOT NULL,
    category VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL,
    origin VARCHAR NOT NULL,
    cost DECIMAL(10, 2),
    price DECIMAL(10, 2),
    type item.item_type NOT NULL,
    status item.item_status NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS item.item;

DROP TYPE IF EXISTS item.item_status;
DROP TYPE IF EXISTS item.item_type;