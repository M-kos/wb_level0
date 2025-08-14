INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand_id, status_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id;
