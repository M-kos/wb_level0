INSERT INTO address (zip, address, city_id)
VALUES ($1, $2, $3)
ON CONFLICT (zip, address, city_id) DO UPDATE SET address=EXCLUDED.address
RETURNING id;
