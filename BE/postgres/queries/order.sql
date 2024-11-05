-- name: GetOrder :one
SELECT * FROM order_groups
WHERE id = $1 LIMIT 1;

-- name: GetSellerOrders :many
SELECT DISTINCT
    og.id AS order_id,
    og.created_at AS order_created_at,
    og.updated_at AS order_updated_at,
    oi.status_name AS order_status,
    og.profile_id AS buyer_profile_id,
    oi.thumbnail_url
FROM 
    order_groups og
JOIN 
    order_items oi ON og.id = oi.order_id
JOIN 
    products p ON oi.product_id = p.id
WHERE 
    p.profile_id = $1  -- This is the seller's profile_id
ORDER BY 
    og.id DESC;

-- name: GetUserOrders :many
SELECT 
    og.*, 
    oi.thumbnail_url
FROM 
    order_groups og
JOIN 
    order_items oi ON og.id = oi.order_id
WHERE 
    og.profile_id = $1
GROUP BY 
    og.id, oi.thumbnail_url
ORDER BY 
    og.id DESC;

-- name: CreateOrderWithStatus :one
SELECT oi.*, os.status_name
FROM order_items oi
JOIN order_statuses os ON oi.status_id = os.id
WHERE oi.id = $1;

-- name: CreateOrder :one
INSERT INTO order_groups (profile_id)
VALUES ($1) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id, price, quantity, status_name, thumbnail_url)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateOrderItemStatus :one
UPDATE order_items
SET status_name = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1 RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM order_groups
WHERE id = $1;

-- name: ListOrders :many
SELECT 
    og.*, 
    oi.status_name AS order_status
FROM 
    order_groups og
JOIN 
    order_items oi ON og.id = oi.order_id
ORDER BY 
    og.id DESC;

-- name: GetOrders :many
SELECT * FROM order_groups
ORDER BY id;

-- name: GetOrderItems :many
SELECT 
    oi.id, 
    oi.product_id, 
    oi.price, 
    oi.quantity, 
    oi.status_name, 
    oi.tracking_url, 
    oi.tracking_number, 
    oi.tracking_service, 
    oi.thumbnail_url AS order_item_thumbnail,
    p.title AS product_title,
    pi.image_url AS product_thumbnail_url
FROM
    order_items oi
LEFT JOIN
    products p ON oi.product_id = p.id
LEFT JOIN
    product_image_thumbnail pit ON p.id = pit.product_id
LEFT JOIN
    product_images pi ON pit.product_image_id = pi.id
WHERE 
    oi.order_id = $1;

-- name: GetSellerOrderItems :many
SELECT 
    oi.id,
    oi.product_id,
    oi.price,
    oi.quantity,
    oi.status_name,
    oi.tracking_url,
    oi.tracking_number,
    oi.tracking_service,
    oi.thumbnail_url
FROM
    order_items oi
JOIN
    products p ON oi.product_id = p.id
WHERE 
    oi.order_id = $1
    AND p.profile_id = $2;

-- name: GetSellerProductsSold :one
SELECT COUNT(DISTINCT oi.id) as total_sold
FROM order_items oi
JOIN products p ON oi.product_id = p.id
WHERE p.profile_id = $1;

-- name: GetSellerTotalRevenue :one
SELECT COALESCE(SUM(oi.price * oi.quantity), 0) as total_revenue
FROM order_items oi
JOIN products p ON oi.product_id = p.id
JOIN order_groups og ON oi.order_id = og.id
WHERE p.profile_id = $1 AND og.created_at >= $2 AND og.created_at < $3;

-- name: GetSellerStatsComprehensive :one
WITH seller_stats AS (
    SELECT 
        COUNT(DISTINCT oi.id) as total_sold,
        COALESCE(SUM(oi.price * oi.quantity), 0) as total_revenue,
        COUNT(DISTINCT og.id) as total_orders
    FROM order_items oi
    JOIN products p ON oi.product_id = p.id
    JOIN order_groups og ON oi.order_id = og.id
    WHERE p.profile_id = $1 AND og.created_at >= $2 AND og.created_at < $3
),
top_products AS (
    SELECT p.id, p.title, COUNT(oi.id) as total_sold, SUM(oi.price * oi.quantity) as revenue
    FROM products p
    JOIN order_items oi ON p.id = oi.product_id
    JOIN order_groups og ON oi.order_id = og.id
    WHERE p.profile_id = $1 AND og.created_at >= $2 AND og.created_at < $3
    GROUP BY p.id, p.title
    ORDER BY total_sold DESC
    LIMIT 5
)
SELECT 
    (SELECT total_sold FROM seller_stats) as total_sold,
    (SELECT total_revenue FROM seller_stats) as total_revenue,
    (SELECT total_orders FROM seller_stats) as total_orders,
    json_agg(json_build_object(
        'id', tp.id,
        'title', tp.title,
        'totalSold', tp.total_sold,
        'revenue', tp.revenue
    )) as top_products
FROM top_products tp;

-- name: GetSellerTopProducts :many
SELECT p.id, p.title, COUNT(oi.id) as total_sold, SUM(oi.price * oi.quantity) as revenue
FROM products p
JOIN order_items oi ON p.id = oi.product_id
JOIN order_groups og ON oi.order_id = og.id
WHERE p.profile_id = $1 AND og.created_at >= $2 AND og.created_at < $3
GROUP BY p.id, p.title
ORDER BY total_sold DESC
LIMIT $4;

-- name: GetSellerOrdersCount :one
SELECT COUNT(DISTINCT og.id) as total_orders
FROM order_groups og
JOIN order_items oi ON og.id = oi.order_id
JOIN products p ON oi.product_id = p.id
WHERE p.profile_id = $1 AND og.created_at >= $2 AND og.created_at < $3;

-- name: LogUserEvent :one
INSERT INTO user_events (profile_id, event_type, product_id, order_item_id, duration, additional_data)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserEvents :many
SELECT * FROM user_events
WHERE profile_id = $1
ORDER BY occurred_at DESC
LIMIT $2 OFFSET $3;


-- name: GetOrderItemStatuses :many
SELECT oi.id, oi.status_name AS item_status, os.status_name AS order_status
FROM order_items oi
LEFT JOIN order_statuses os ON oi.order_id = os.order_id
WHERE oi.order_id = $1;