INSERT INTO region (name)
VALUES ($1)
RETURNING id;
