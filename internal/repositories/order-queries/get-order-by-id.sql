SELECT o.id,
       o.order_uid,
       o.track_number,
       o.entry,

       d.id      as delivery_id,
       cr.id     as customer_id,
       cr.first_name,
       cr.last_name,
       cr.phone,
       a.id      as address_id,
       a.zip,
       c.id      AS city_id,
       c.name    AS delivery_city,
       a.address,
       r.id      AS delivery_region_id,
       r.name    AS delivery_region,
       cr.email,

       p.id      as payment_id,
       p.transaction,
       p.request_id,
       cur.id    AS payment_currency_id,
       cur.name  AS payment_currency,
       prov.id   AS payment_provider_id,
       prov.name AS payment_provider,
       p.amount,
       p.payment_dt,
       b.id      AS payment_bank_id,
       b.name    AS payment_bank,
       p.delivery_cost,
       p.goods_total,
       p.custom_fee,

       l.id      AS locale_id,
       l.name    AS locale,

       o.internal_signature,
       o.customer_id,
       ds.id     AS delivery_service_id,
       ds.name   AS delivery_service,
       o.shardkey,
       o.sm_id,
       o.date_created,
       o.oof_shard
FROM "order" as o
	     LEFT JOIN
     delivery AS d ON d.id = o.delivery_id
	     LEFT JOIN
     customer cr ON d.customer_id = cr.id
	     LEFT JOIN
     address a ON d.address_id = a.id
	     LEFT JOIN
     city c ON a.city_id = c.id
	     LEFT JOIN
     region r ON c.region_id = r.id
	     LEFT JOIN
     payment AS p ON p.id = o.payment_id
	     LEFT JOIN
     currency cur ON p.currency_id = cur.id
	     LEFT JOIN
     provider prov ON p.provider_id = prov.id
	     LEFT JOIN
     bank b ON p.bank_id = b.id
	     LEFT JOIN
     locale AS l ON l.id = o.locale_id
	     LEFT JOIN
     delivery_service AS ds ON ds.id = o.delivery_service_id
WHERE o.order_uid = $1;
