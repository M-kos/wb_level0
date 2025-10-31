INSERT INTO delivery (customer_id, address_id)
VALUES ($1, $2)
ON CONFLICT (customer_id, address_id) DO NOTHING
RETURNING id;
