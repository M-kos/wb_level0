INSERT INTO address (zip, address, city_id)
VALUES ($1, $2, $3)
RETURNING id;
