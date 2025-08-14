INSERT INTO provider (name)
VALUES ($1)
RETURNING id;
