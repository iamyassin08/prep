-- +goose Up
-- +goose StatementBegin

CREATE TABLE profile_product_cart_items (
    product_id INTEGER NOT NULL REFERENCES products(id),
    profile_id VARCHAR(36) NOT NULL REFERENCES profiles(id),
    product_quantity INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_id, profile_id)
);

ALTER TABLE product_image_thumbnail
ADD CONSTRAINT product_image_thumbnail_product_id_key UNIQUE (product_id);

ALTER TABLE products
ADD COLUMN type VARCHAR(20) DEFAULT 'Buy It Now' CHECK (type IN ('Buy It Now', 'Reservation', 'External', 'Delisted'));


CREATE TABLE product_external_brands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    logo_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_external_urls (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id),
    brand_id  INTEGER NOT NULL REFERENCES product_external_brands(id),
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO categories (name, parent_id, description)
VALUES 
    ('Musical Instruments', NULL, 'Description for Musical Instruments'),
    ('Shoes', NULL, 'Description for Shoes'),
    ('Men''s Shoes', 2, 'Description for Men''s Shoes'),
    ('Women''s Shoes', 2, 'Description for Women''s Shoes'),
    ('Children''s Shoes', 2, 'Description for Children''s Shoes'),
    ('Clothes', NULL, 'Description for Clothes'),
    ('Men''s Clothes', 6, 'Description for Men''s Clothes'),
    ('Women''s Clothes', 6, 'Description for Women''s Clothes'),
    ('Children''s Clothes', 6, 'Description for Children''s Clothes'),
    ('Video Games', NULL, 'Description for Video Games'),
    ('Xbox Games', 10, 'Description for Xbox Games'),
    ('Playstation Games', 10, 'Description for Playstation Games'),
    ('Nintendo Games', 10, 'Description for Nintendo Games'),
    ('Hardware Tools', NULL, 'Description for Hardware Tools'),
    ('Audio Equipment', NULL, 'Description for Audio Equipment'),
    ('Sporting Goods', NULL, 'Description for Sporting Goods');


-- Insert tags
INSERT INTO tags (tag_name)
VALUES 
    ('New'),
    ('Sale'),
    ('Bestseller'),
    ('Limited Edition'),
    ('Eco-friendly');

-- SELLER
INSERT INTO Profiles (id)
VALUES 
    ('caaa600c-ef66-4fc5-a341-fe54c164961a'), -- SELLER
    ('5319834e-1b74-4e70-a37c-5d2679f05c01'), -- BUYER
    ('53fc6e40-ed4b-403f-b120-e6db0377e05f'); -- ADMIN

-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
ALTER TABLE products
DROP COLUMN IF EXISTS type;
ALTER TABLE product_image_thumbnail
DROP CONSTRAINT product_image_thumbnail_product_id_key;
-- DROP TABLE IF EXISTS category_backgrounds;
-- DROP TABLE IF EXISTS category_banners;
DROP TABLE IF EXISTS product_external_urls;
DROP TABLE IF EXISTS product_external_brands;
DROP TABLE IF EXISTS profile_product_cart_items;
-- +goose StatementEnd




