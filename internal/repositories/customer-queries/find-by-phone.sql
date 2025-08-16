SELECT id, first_name, last_name, phone, email
FROM customer
WHERE phone = $1
