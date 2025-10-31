INSERT INTO item_status (value)
VALUES ($1)
ON CONFLICT (value) DO UPDATE SET value=EXCLUDED.value
RETURNING id;
