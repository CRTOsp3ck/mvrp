-- +migrate Up
CREATE SCHEMA IF NOT EXISTS entity;

CREATE TYPE entity.entity_type AS ENUM (
    'Customer', 
    'Supplier',
    'Manufacturer', 
    'Employee', 
    'Other'
);

CREATE TABLE entity.entity (
    id INT PRIMARY KEY,
    code VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT NOT NULL,
    address TEXT,
    phone VARCHAR(32),
    email VARCHAR(128),
    website VARCHAR(128),
    type entity.entity_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS entity.entity;

DROP TYPE IF EXISTS entity.entity_type;