-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    parent_id INTEGER REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    banner_url TEXT,
    background_url TEXT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE profiles (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    tag_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    profile_id VARCHAR(36) REFERENCES profiles(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    short_description VARCHAR(155),
    price NUMERIC(10, 2) NOT NULL,
    quantity INTEGER NOT NULL,
    discount_price NUMERIC(10, 2) NOT NULL,
    regular_price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE category_products (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id),
    category_id INTEGER NOT NULL
);
CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id),
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_image_thumbnail (
    product_id INTEGER NOT NULL REFERENCES products(id),
    product_image_id INTEGER NOT NULL REFERENCES product_images(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_id, product_image_id)
);

CREATE TABLE attribute_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

INSERT INTO attribute_types (name) VALUES 
    ('string'), ('float'), ('integer'), ('boolean'), ('enum');

CREATE TABLE category_attributes (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    attribute_type_id INTEGER REFERENCES attribute_types(id),
    is_required BOOLEAN DEFAULT false,
    enum_values TEXT[] -- Only used for enum type, stores possible values
);

CREATE TABLE product_attribute_values (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id),
    category_attribute_id INTEGER REFERENCES category_attributes(id),
    value_string TEXT,
    value_float FLOAT,
    value_integer INTEGER,
    value_boolean BOOLEAN
);

CREATE TABLE product_tags (
    tag_id INTEGER NOT NULL REFERENCES tags(id),
    product_id INTEGER NOT NULL REFERENCES products(id),
    PRIMARY KEY (tag_id, product_id)
);

CREATE TABLE profile_addresses (
    id SERIAL PRIMARY KEY,
    profile_id VARCHAR(36) NOT NULL REFERENCES profiles(id),
    address_line1 VARCHAR(255) NOT NULL,
    address_line2 VARCHAR(255),
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20)
);

CREATE TABLE profile_product_favorites (
    product_id INTEGER NOT NULL REFERENCES products(id),
    profile_id VARCHAR(36) NOT NULL REFERENCES profiles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (product_id, profile_id)
);


CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    profile_id VARCHAR(36) NOT NULL REFERENCES profiles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id),
    order_id INTEGER NOT NULL REFERENCES orders(id),
    price NUMERIC(10, 2) NOT NULL,
    status_name VARCHAR(255),
    quantity INTEGER NOT NULL
);

CREATE TABLE order_statuses (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(255) NOT NULL,
    order_id INTEGER NOT NULL UNIQUE REFERENCES orders(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_statuses;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS profile_product_favorites;
DROP TABLE IF EXISTS profile_addresses;
DROP TABLE IF EXISTS product_tags;
DROP TABLE IF EXISTS product_attribute_values;
DROP TABLE IF EXISTS category_attributes;
DROP TABLE IF EXISTS attribute_types;
DROP TABLE IF EXISTS category_products;
-- DROP TABLE IF EXISTS attribute_values;
DROP TABLE IF EXISTS product_image_thumbnail;
DROP TABLE IF EXISTS product_images;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS orders;
-- DROP TABLE IF EXISTS attributes;
DROP TABLE IF EXISTS product_types;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd