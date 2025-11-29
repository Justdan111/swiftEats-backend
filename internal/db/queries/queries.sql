-- name: GetMenuItemsByIDs :many
SELECT id, restaurant_id, category_id, name, description, price_cents, is_available, created_at
FROM menu_items
WHERE id = ANY($1);

-- name: GetMenuItemByID :one
SELECT id, restaurant_id, category_id, name, description, price_cents, is_available, created_at
FROM menu_items
WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO orders (user_id, restaurant_id, total_cents, status, idempotency_key)
VALUES ($1,$2,$3,$4,$5)
RETURNING id, user_id, restaurant_id, total_cents, status, idempotency_key, payment_provider, payment_reference, created_at;

-- name: CreateOrderItem :exec
INSERT INTO order_items (order_id, menu_item_id, quantity, price_cents)
VALUES ($1,$2,$3,$4);

-- name: FindOrderByIdempotency :one
SELECT id, user_id, restaurant_id, total_cents, status, idempotency_key, payment_provider, payment_reference, created_at
FROM orders WHERE idempotency_key = $1 LIMIT 1;

-- name: GetOrderByID :one
SELECT id, user_id, restaurant_id, total_cents, status, idempotency_key, payment_provider, payment_reference, created_at
FROM orders WHERE id = $1;

-- name: InsertPayment :one
INSERT INTO payments (order_id, provider, provider_reference, amount_cents, status, raw_payload)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, order_id, provider, provider_reference, amount_cents, status, raw_payload, created_at;

-- name: UpdateOrderPaymentAndStatus :exec
UPDATE orders SET payment_reference = $2, payment_provider = $3, status = $4 WHERE id = $1;
