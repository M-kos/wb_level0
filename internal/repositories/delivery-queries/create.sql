INSERT INTO delivery (name, phone, email, address_id)
VALUES ($1, $2, $3, $4)
RETURNING id;
