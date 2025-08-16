SELECT id, customer_id, address_id
FROM delivery
WHERE customer_id = $1
  AND address_id = $2
