INSERT INTO delivery (customer_id, address_id)
VALUES ($1, $2)
RETURNING id;
