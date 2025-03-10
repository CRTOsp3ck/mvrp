-- +migrate Up
CREATE OR REPLACE VIEW invoice.invoice_item_view AS
SELECT
    ii.*,
    (
        SELECT row_to_json(bdi)
        FROM base.base_document_item bdi
        WHERE bdi.id = ii.base_document_item_id
    ) AS base_document_item,
    (
        SELECT row_to_json(iv)
        FROM inventory.inventory_view iv
        WHERE iv.id = (
            SELECT bdi.inventory_id
            FROM base.base_document_item bdi
            WHERE bdi.id = ii.base_document_item_id
        )
    ) AS inventory_info
FROM
    invoice.invoice_item ii;

CREATE OR REPLACE VIEW invoice.invoice_view AS
SELECT
    i.*,
    (
        SELECT row_to_json(bdv)
        FROM base.base_document_view bdv
        WHERE bdv.id = i.base_document_id
    ) AS base_document,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = i.vendor_id
    ) AS vendor_info,
    (
        SELECT row_to_json(e)
        FROM entity.entity e
        WHERE e.id = i.customer_id
    ) AS customer_info,
    (
        SELECT json_agg(row_to_json(iiv))
        FROM invoice.invoice_item_view iiv
        WHERE iiv.invoice_id = i.id
    ) AS invoice_items
FROM
    invoice.invoice i;

-- +migrate Down
DROP VIEW IF EXISTS invoice.invoice_view CASCADE;
DROP VIEW IF EXISTS invoice.invoice_item_view CASCADE;