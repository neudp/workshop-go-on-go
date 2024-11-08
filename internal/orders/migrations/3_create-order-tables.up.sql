CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    number BIGINT UNSIGNED NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    customer_id BIGINT UNSIGNED NOT NULL,
    shipping_address_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers (id),
    FOREIGN KEY (shipping_address_id) REFERENCES addresses (id)
);

ALTER TABLE orders ADD UNIQUE INDEX order_number_idx (number);

CREATE TABLE IF NOT EXISTS order_status_history (
    order_id BIGINT UNSIGNED NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    PRIMARY KEY (order_id, status)
);
