INSERT INTO customer (first_name, last_name, phone, email)
VALUES ($1, $2, $3, $4)
RETURNING id;
