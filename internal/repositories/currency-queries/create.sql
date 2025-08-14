INSERT INTO currency (name)
VALUES ($1)
RETURNING id;
