SELECT id,
       chrt_id,
       track_number,
       price,
       rid,
       name,
       sale,
       size,
       total_price,
       nm_id,
       brand_id,
       status_id
FROM item
WHERE id = $1
