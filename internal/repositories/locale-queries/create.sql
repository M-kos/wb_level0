INSERT INTO locale (name)
VALUES ($1)
RETURNING id;
