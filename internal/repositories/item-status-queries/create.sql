INSERT INTO item_status (value)
VALUES ($1)
RETURNING id;
