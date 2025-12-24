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
-- ============ RESTAURANT QUERIES ============

-- name: GetAllRestaurants :many
SELECT id, name, description, address, created_at
FROM restaurants
ORDER BY created_at DESC;

-- name: GetRestaurantByID :one
SELECT id, name, description, address, created_at
FROM restaurants
WHERE id = $1;

-- name: GetMenuItemsByRestaurantID :many
SELECT id, restaurant_id, category_id, name, description, price_cents, is_available, created_at
FROM menu_items
WHERE restaurant_id = $1 AND is_available = true
ORDER BY name ASC;

-- ============ CART QUERIES ============

-- name: GetCartByUserID :one
SELECT id, user_id, created_at, updated_at
FROM carts
WHERE user_id = $1;

-- name: CreateCart :one
INSERT INTO carts (user_id)
VALUES ($1)
ON CONFLICT (user_id) DO NOTHING
RETURNING id, user_id, created_at, updated_at;

-- name: GetCartItemsByCartID :many
SELECT id, cart_id, menu_item_id, quantity, price_cents, created_at
FROM cart_items
WHERE cart_id = $1
ORDER BY created_at DESC;

-- name: AddCartItem :one
INSERT INTO cart_items (cart_id, menu_item_id, quantity, price_cents)
VALUES ($1, $2, $3, $4)
RETURNING id, cart_id, menu_item_id, quantity, price_cents, created_at;

-- name: RemoveCartItem :exec
DELETE FROM cart_items
WHERE id = $1 AND cart_id = $2;

-- name: ClearCart :exec
DELETE FROM cart_items
WHERE cart_id = $1;

-- name: UpdateCartItemQuantity :one
UPDATE cart_items
SET quantity = $2
WHERE id = $1 AND cart_id = $3
RETURNING id, cart_id, menu_item_id, quantity, price_cents, created_at;

-- ============ ORDER QUERIES ============

-- name: GetUserOrders :many
SELECT id, user_id, restaurant_id, total_cents, status, idempotency_key, payment_provider, payment_reference, created_at
FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetOrderItems :many
SELECT id, order_id, menu_item_id, quantity, price_cents, created_at
FROM order_items
WHERE order_id = $1
ORDER BY created_at ASC;

-- name: GetOrderCount :one
SELECT COUNT(*) as count
FROM orders
WHERE user_id = $1;

-- ============ PAYMENT QUERIES ============

-- name: GetPaymentByOrderID :one
SELECT id, order_id, provider, provider_reference, amount_cents, status, raw_payload, created_at
FROM payments
WHERE order_id = $1;

-- name: UpdatePaymentStatus :exec
UPDATE payments
SET status = $2
WHERE id = $1;

-- name: GetPaymentByProviderReference :one
SELECT id, order_id, provider, provider_reference, amount_cents, status, raw_payload, created_at
FROM payments
WHERE provider_reference = $1;

-- ============ ADMIN RESTAURANT QUERIES ============

-- name: CreateRestaurant :one
INSERT INTO restaurants (name, description, address)
VALUES ($1, $2, $3)
RETURNING id, name, description, address, created_at;

-- name: UpdateRestaurant :one
UPDATE restaurants
SET name = $2, description = $3, address = $4
WHERE id = $1
RETURNING id, name, description, address, created_at;

-- name: DeleteRestaurant :exec
DELETE FROM restaurants
WHERE id = $1;

-- ============ ADMIN MENU ITEM QUERIES ============

-- name: CreateMenuItem :one
INSERT INTO menu_items (restaurant_id, category_id, name, description, price_cents, is_available)
VALUES ($1, $2, $3, $4, $5, true)
RETURNING id, restaurant_id, category_id, name, description, price_cents, is_available, created_at;

-- name: UpdateMenuItem :one
UPDATE menu_items
SET name = $2, description = $3, price_cents = $4, is_available = $5
WHERE id = $1
RETURNING id, restaurant_id, category_id, name, description, price_cents, is_available, created_at;

-- name: DeleteMenuItem :exec
DELETE FROM menu_items
WHERE id = $1;

-- name: UpdateMenuItemAvailability :exec
UPDATE menu_items
SET is_available = $2
WHERE id = $1;