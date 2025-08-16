SELECT id,
       transaction,
       request_id,
       currency_id,
       provider_id,
       amount,
       payment_dt,
       bank_id,
       delivery_cost,
       goods_total,
       custom_fee
FROM payment
WHERE id = $1
