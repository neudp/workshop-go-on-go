CREATE TABLE IF NOT EXISTS countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

ALTER TABLE countries ADD UNIQUE INDEX country_name_unique_idx (name);

CREATE TABLE IF NOT EXISTS provinces (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (country_id) REFERENCES countries (id)
);

ALTER TABLE provinces ADD UNIQUE INDEX province_name_country_id_unique_idx (country_id, name);

CREATE TABLE IF NOT EXISTS cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    province_id BIGINT UNSIGNED NOT NULL,
    country_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (province_id) REFERENCES provinces (id),
    FOREIGN KEY (country_id) REFERENCES countries (id)
);

ALTER TABLE cities ADD UNIQUE INDEX city_name_province_id_country_id_unique_idx (country_id, name, province_id);

CREATE TABLE IF NOT EXISTS addresses (
    id SERIAL PRIMARY KEY,
    street VARCHAR(255) NOT NULL,
    postal_code VARCHAR(255) NOT NULL,
    address_line VARCHAR(255) NOT NULL,
    city_id BIGINT UNSIGNED NOT NULL,
    province_id BIGINT UNSIGNED NOT NULL,
    country_id BIGINT UNSIGNED NOT NULL,
    FOREIGN KEY (city_id) REFERENCES cities (id),
    FOREIGN KEY (province_id) REFERENCES provinces (id),
    FOREIGN KEY (country_id) REFERENCES countries (id)
);

ALTER TABLE addresses ADD INDEX postal_code_idx (postal_code);
ALTER TABLE addresses ADD INDEX street_idx (street);
