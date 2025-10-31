INSERT INTO city (name, region_id)
VALUES ($1, $2)
ON CONFLICT (name, region_id) DO UPDATE SET name=EXCLUDED.name
RETURNING id;
