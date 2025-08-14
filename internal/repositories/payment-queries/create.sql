INSERT INTO payment (transaction, request_id, currency_id, provider_id, amount, payment_dt, bank_id, delivery_cost,
                     goods_total, custom_fee)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id;
