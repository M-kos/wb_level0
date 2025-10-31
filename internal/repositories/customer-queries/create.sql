INSERT INTO customer (first_name, last_name, phone, email)
VALUES ($1, $2, $3, $4)
ON CONFLICT (phone) DO UPDATE SET first_name = EXCLUDED.first_name,
                                  last_name  = EXCLUDED.last_name,
                                  email      = EXCLUDED.email
RETURNING id;
