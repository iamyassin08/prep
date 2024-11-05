-- name: CreateCategoryAttribute :one
INSERT INTO category_attributes (
    category_id, name, attribute_type_id, is_required, enum_values
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCategoryAttributes :many
SELECT * FROM category_attributes 
WHERE category_id = $1;


-- name: GetAttribute :one
SELECT * FROM category_attributes 
WHERE id = $1
LIMIT 1;

-- name: UpdateAttribute :exec
UPDATE category_attributes
SET  category_id = $1, name = $2, attribute_type_id = $3, is_required = $4, enum_values = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $6 RETURNING *;

-- name: UpdateProductAttributeValue :exec
UPDATE product_attribute_values
SET product_id = $1, category_attribute_id = $2, value_string = $3, value_float = $4, value_integer = $5, value_boolean = $6
Where id = $7;

-- name: CreateProductAttributeValue :one
INSERT INTO product_attribute_values (
    product_id, category_attribute_id, value_string, value_float, value_integer, value_boolean
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetProductAttributeValues :many
SELECT pav.*, ca.name as attribute_name, at.name as attribute_type
FROM product_attribute_values pav
JOIN category_attributes ca ON pav.category_attribute_id = ca.id
JOIN attribute_types at ON ca.attribute_type_id = at.id
WHERE pav.product_id = $1;

-- name: ListAttributes :many
SELECT id, name
FROM attribute_types
ORDER BY name;




