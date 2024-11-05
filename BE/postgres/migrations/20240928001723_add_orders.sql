-- +goose Up
-- +goose StatementBegin

-- Rename orders table to order_groups
ALTER TABLE orders RENAME TO order_groups;

-- Add columns to order_items table
ALTER TABLE order_items
ADD COLUMN tracking_url TEXT,
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN tracking_number VARCHAR(255),
ADD COLUMN tracking_service VARCHAR(255),
ADD COLUMN thumbnail_url TEXT;

-- Add columns to profiles table
ALTER TABLE profiles
ADD COLUMN thumbnail_url TEXT,
ADD COLUMN email VARCHAR(255),
ADD COLUMN username VARCHAR(255);

-- Create new table for event tracking
CREATE TABLE user_events (
    id SERIAL PRIMARY KEY,
    profile_id VARCHAR(36) REFERENCES profiles(id),
    event_type VARCHAR(255) NOT NULL,
    product_id INTEGER REFERENCES products(id),
    order_item_id INTEGER REFERENCES order_items(id),
    duration INTEGER, -- in seconds, for product views
    occurred_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    additional_data JSONB
);

-- Create index on event_type for faster queries
CREATE INDEX idx_user_events_event_type ON user_events(event_type);

-- Create index on profile_id for faster queries
CREATE INDEX idx_user_events_profile_id ON user_events(profile_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Revert the changes in case of rollback

-- Rename order_groups table back to orders
ALTER TABLE order_groups RENAME TO orders;

-- Remove added columns from order_items table
ALTER TABLE order_items
DROP COLUMN status,
DROP COLUMN tracking_url,
DROP COLUMN tracking_number,
DROP COLUMN tracking_service,
DROP COLUMN thumbnail_url;

-- Remove added columns from profiles table
ALTER TABLE profiles
DROP COLUMN thumbnail_url,
DROP COLUMN email,
DROP COLUMN username;

-- Drop the user_events table
DROP TABLE IF EXISTS user_events;

-- +goose StatementEnd