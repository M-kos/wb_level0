INSERT INTO "order" (order_uid, track_number, entry, delivery_id, payment_id, locale_id, internal_signature,
                     customer_id,
                     delivery_service_id, shardkey, sm_id, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id;
