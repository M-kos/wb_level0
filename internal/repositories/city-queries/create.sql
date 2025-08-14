INSERT INTO city (name, region_id)
VALUES ($1, $2)
RETURNING id;
