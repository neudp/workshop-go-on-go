CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    billing_address_id BIGINT UNSIGNED NOT NULL,
    shipping_address_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (billing_address_id) REFERENCES addresses (id),
    FOREIGN KEY (shipping_address_id) REFERENCES addresses (id)
);
