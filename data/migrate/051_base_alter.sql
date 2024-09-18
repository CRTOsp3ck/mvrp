-- +migrate Up
CREATE OR REPLACE VIEW base.base_document_item_view AS
SELECT
    bdi.*,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = bdi.inventory_id
    ) AS inventory_info
FROM
    base.base_document_item bdi;

CREATE OR REPLACE VIEW base.base_document_view AS
SELECT
    bd.*,
    (
        SELECT json_agg(row_to_json(bdiv))
        FROM base.base_document_item_view bdiv
        WHERE bdiv.base_document_id = bd.id
    ) AS base_document_items
FROM
    base.base_document bd;

-- +migrate Down
DROP VIEW IF EXISTS base.base_document_view;
DROP VIEW IF EXISTS base.base_document_item_view;

