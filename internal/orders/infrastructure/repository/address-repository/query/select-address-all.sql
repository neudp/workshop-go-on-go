SELECT
    addr.id,
    addr.street,
    addr.postal_code,
    addr.address_line,
    cnt.id AS country_id,
    cnt.name AS country_name,
    prv.id AS province_id,
    prv.name AS province_name,
    cty.id AS city_id,
    cty.name AS city_name
FROM addresses addr
JOIN countries cnt ON addr.country_id = cnt.id
JOIN provinces prv ON addr.province_id = prv.id
JOIN cities cty ON addr.city_id = cty.id
