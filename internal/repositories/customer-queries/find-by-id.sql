SELECT id, first_name, last_name, phone, email
FROM customer
WHERE id = $1
