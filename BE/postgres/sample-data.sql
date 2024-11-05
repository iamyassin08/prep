-- Function to generate a random date
CREATE OR REPLACE FUNCTION random_date(start_date DATE, end_date DATE) 
RETURNS DATE AS $$
BEGIN
    RETURN start_date + (random() * (end_date - start_date))::INT;
END;
$$ LANGUAGE plpgsql;

-- Function to generate a random user name
CREATE OR REPLACE FUNCTION generate_user_name(category TEXT) 
RETURNS TEXT AS $$
DECLARE
    adjectives TEXT[] := ARRAY['Premium', 'Deluxe', 'Professional', 'Classic', 'Modern', 'Ultra', 'Elite', 'Essential', 'Signature', 'Advanced'];
    instruments TEXT[] := ARRAY['Electric Guitar', 'Acoustic Guitar', 'Bass Guitar', 'Drum Kit', 'Digital Piano', 'Synthesizer', 'Violin', 'Trumpet', 'Saxophone', 'Flute'];
    music_brands TEXT[] := ARRAY['Fender', 'Gibson', 'Yamaha', 'Roland', 'Ibanez', 'Pearl', 'Steinway', 'Selmer', 'Zildjian', 'Martin'];
    shoe_types TEXT[] := ARRAY['Running Shoes', 'Sneakers', 'Boots', 'Sandals', 'Loafers', 'Oxfords', 'Slip-ons', 'Heels', 'Flats', 'Athletic Shoes'];
    shoe_brands TEXT[] := ARRAY['Nike', 'Adidas', 'New Balance', 'Asics', 'Skechers', 'Clarks', 'Timberland', 'Vans', 'Converse', 'Puma'];
    clothing_items TEXT[] := ARRAY['T-shirt', 'Jeans', 'Jacket', 'Sweater', 'Dress', 'Skirt', 'Pants', 'Shorts', 'Hoodie', 'Blouse'];
    clothing_brands TEXT[] := ARRAY['Levi''s', 'Gap', 'H&M', 'Zara', 'Uniqlo', 'Calvin Klein', 'Ralph Lauren', 'Tommy Hilfiger', 'Gucci', 'North Face'];
    game_types TEXT[] := ARRAY['Action', 'Adventure', 'RPG', 'Strategy', 'Sports', 'Puzzle', 'Simulation', 'Racing', 'Fighting', 'Shooter'];
    game_titles TEXT[] := ARRAY['Legends', 'Chronicles', 'Odyssey', 'Quest', 'Warfare', 'Empire', 'Horizon', 'Saga', 'Realms', 'Legacy'];
    tools TEXT[] := ARRAY['Power Drill', 'Circular Saw', 'Hammer', 'Screwdriver Set', 'Wrench Set', 'Pliers', 'Measuring Tape', 'Level', 'Chisel Set', 'Sander'];
    tool_brands TEXT[] := ARRAY['DeWalt', 'Makita', 'Bosch', 'Milwaukee', 'Craftsman', 'Stanley', 'Ryobi', 'Black & Decker', 'Ridgid', 'Hilti'];
    audio_items TEXT[] := ARRAY['Headphones', 'Bluetooth Speaker', 'Microphone', 'Amplifier', 'Turntable', 'Soundbar', 'Subwoofer', 'Wireless Earbuds', 'DAC', 'AV Receiver'];
    audio_brands TEXT[] := ARRAY['Bose', 'Sony', 'Sennheiser', 'Audio-Technica', 'JBL', 'Harman Kardon', 'Beyerdynamic', 'Shure', 'Marshall', 'Focal'];
    sports_items TEXT[] := ARRAY['Basketball', 'Soccer Ball', 'Tennis Racket', 'Golf Clubs Set', 'Mountain Bike', 'Skateboard', 'Yoga Mat', 'Dumbbell Set', 'Baseball Bat', 'Swimming Goggles'];
    sports_brands TEXT[] := ARRAY['Wilson', 'Spalding', 'Nike', 'Adidas', 'Under Armour', 'Callaway', 'Speedo', 'Trek', 'Everlast', 'Coleman'];
BEGIN
    IF category = 'Musical Instruments' THEN
        RETURN (SELECT music_brands[floor(random() * array_length(music_brands, 1) + 1)] || ' ' ||
                       adjectives[floor(random() * array_length(adjectives, 1) + 1)] || ' ' ||
                       instruments[floor(random() * array_length(instruments, 1) + 1)]);
    ELSIF category LIKE '%Shoes%' THEN
        RETURN (SELECT shoe_brands[floor(random() * array_length(shoe_brands, 1) + 1)] || ' ' ||
                       adjectives[floor(random() * array_length(adjectives, 1) + 1)] || ' ' ||
                       shoe_types[floor(random() * array_length(shoe_types, 1) + 1)]);
    ELSIF category LIKE '%Clothes%' THEN
        RETURN (SELECT clothing_brands[floor(random() * array_length(clothing_brands, 1) + 1)] || ' ' ||
                       adjectives[floor(random() * array_length(adjectives, 1) + 1)] || ' ' ||
                       clothing_items[floor(random() * array_length(clothing_items, 1) + 1)]);
    ELSIF category LIKE '%Games%' THEN
        RETURN (SELECT game_types[floor(random() * array_length(game_types, 1) + 1)] || ' ' ||
                       game_titles[floor(random() * array_length(game_titles, 1) + 1)] || ' ' ||
                       (floor(random() * 5) + 1)::TEXT);
    ELSIF category = 'Hardware Tools' THEN
        RETURN (SELECT tool_brands[floor(random() * array_length(tool_brands, 1) + 1)] || ' ' ||
                       adjectives[floor(random() * array_length(adjectives, 1) + 1)] || ' ' ||
                       tools[floor(random() * array_length(tools, 1) + 1)]);
    ELSIF category = 'Audio Equipment' THEN
        RETURN (SELECT audio_brands[floor(random() * array_length(audio_brands, 1) + 1)] || ' ' ||
                       adjectives[floor(random() * array_length(adjectives, 1) + 1)] || ' ' ||
                       audio_items[floor(random() * array_length(audio_items, 1) + 1)]);
    ELSIF category = 'Sporting Goods' THEN
        RETURN (SELECT sports_brands[floor(random() * array_length(sports_brands, 1) + 1)] || ' ' ||
                       adjectives[floor(random() * array_length(adjectives, 1) + 1)] || ' ' ||
                       sports_items[floor(random() * array_length(sports_items, 1) + 1)]);
    ELSE
        RETURN 'Generic User';
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Function to generate a user description
CREATE OR REPLACE FUNCTION generate_user_description(name TEXT, category TEXT) 
RETURNS TEXT AS $$
DECLARE
    features TEXT[] := ARRAY['durable', 'exciting','lightweight', 'compact', 'versatile', 'high-performance', 'ergonomic', 'energy-efficient', 'customizable', 'portable', 'user-friendly'];
    benefits TEXT[] := ARRAY['enhances your experience', 'improves userivity', 'provides superior comfort', 'delivers exceptional quality', 'offers great value', 'meets professional standards', 'perfect for beginners and experts alike', 'stands out from the competition', 'built to last', 'elevates your style'];
BEGIN
    RETURN 'Introducing the ' || name || ' - a ' || features[floor(random() * array_length(features, 1) + 1)] || ' ' || category || ' user that ' || benefits[floor(random() * array_length(benefits, 1) + 1)] || '. ' ||
           'Crafted with premium materials, this ' || features[floor(random() * array_length(features, 1) + 1)] || ' item ' || benefits[floor(random() * array_length(benefits, 1) + 1)] || '. ' ||
           'Whether you''re a professional or enthusiast, the ' || name || ' is designed to ' || benefits[floor(random() * array_length(benefits, 1) + 1)] || '. ' ||
           'Experience the difference with our ' || features[floor(random() * array_length(features, 1) + 1)] || ' design and ' || features[floor(random() * array_length(features, 1) + 1)] || ' functionality.';
END;
$$ LANGUAGE plpgsql;


-- Function to generate a user short description
CREATE OR REPLACE FUNCTION generate_short_description(name TEXT, category TEXT) 
RETURNS VARCHAR(155) AS $$
DECLARE
    features TEXT[] := ARRAY['durable', 'exciting', 'lightweight', 'compact', 'versatile', 'high-performance', 'ergonomic', 'efficient', 'customizable', 'portable'];
    benefits TEXT[] := ARRAY['enhances experience', 'boosts userivity', 'ensures comfort', 'delivers quality', 'offers value', 'meets pro standards', 'suits all skill levels', 'stands out', 'built to last'];
    description VARCHAR(155);
BEGIN
    description := 'The ' || name || ': A ' || features[floor(random() * array_length(features, 1) + 1)] || ' ' || category || ' that ' || 
                   benefits[floor(random() * array_length(benefits, 1) + 1)] || '. ' ||
                   features[floor(random() * array_length(features, 1) + 1)] || ' design ' || 
                   benefits[floor(random() * array_length(benefits, 1) + 1)] || '.';
    
    -- Ensure the description doesn't exceed 155 characters
    RETURN substring(description, 1, 155);
END;
$$ LANGUAGE plpgsql;

-- Insert categories
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

INSERT INTO category_attributes (attribute_id, category_id)
VALUES
    (1, 1),   -- Color for Musical Instruments
    (2, 1),   -- Size for Musical Instruments
    (4, 1),   -- Brand for Musical Instruments
    (1, 2),   -- Color for Shoes
    (2, 2),   -- Size for Shoes
    (4, 2),   -- Brand for Shoes
    (1, 3),   -- Color for Men's Shoes
    (2, 3),   -- Size for Men's Shoes
    (4, 3),   -- Brand for Men's Shoes
    (2, 4),   -- Size for Men's Shoes
    (4, 4),   -- Brand for Men's Shoes
    (1, 4),   -- Color for Men's Shoes
    (2, 5),   -- Size for Men's Shoes
    (4, 5),   -- Brand for Men's Shoes
    (1, 5),   -- Color for Men's Shoes
    (2, 6),   -- Size for Men's Shoes
    (4, 6),   -- Brand for Men's Shoes
    (1, 6),   -- Color for Men's Shoes
    (2, 7),   -- Size for Men's Shoes
    (4, 7),   -- Brand for Men's Shoes
    (1, 7),   -- Color for Men's Shoes
    (2, 8),   -- Size for Men's Shoes
    (4, 8),   -- Brand for Men's Shoes
    (1, 8),   -- Color for Men's Shoes
    (2, 9),   -- Size for Men's Shoes
    (4, 9),   -- Brand for Men's Shoes
    (1, 9),   -- Color for Men's Shoes
    (2, 10),   -- Size for Men's Shoes
    (4, 10),   -- Brand for Men's Shoes
    (1, 10),   -- Brand for Men's Shoes
    (2, 11),   -- Brand for Men's Shoes
    (4, 11),   -- Brand for Men's Shoes
    (1, 11),   -- Brand for Men's Shoes
    (2, 12),   -- Brand for Men's Shoes
    (4, 12),   -- Brand for Men's Shoes
    (1, 12),   -- Brand for Men's Shoes
    (2, 13),   -- Brand for Men's Shoes
    (4, 13),   -- Brand for Men's Shoes
    (1, 13),   -- Brand for Men's Shoes
    (2, 14),   -- Brand for Men's Shoes
    (4, 14),   -- Brand for Men's Shoes
    (1, 14),   -- Brand for Men's Shoes
    (2, 15),   -- Brand for Men's Shoes
    (4, 15),   -- Brand for Men's Shoes
    (1, 15),   -- Brand for Men's Shoes
    (4, 16),   -- Brand for Men's Shoes
    (1, 16),   -- Brand for Men's Shoes
    (2, 16);
-- Insert attributes
INSERT INTO attributes (name, description)
VALUES 
    ('Color', 'Description for Color'),
    ('Size', 'Description for Size'),
    ('Material', 'Description for Material'),
    ('Brand', 'Description for Brand'),
    ('Weight', 'Description for Weight'),
    ('Dimensions', 'Description for Dimensions');

-- Insert attribute values
INSERT INTO attribute_values (attribute_id, name, description)
SELECT a.id, a.name, 'Description for ' || a.name
FROM attributes a
CROSS JOIN (
    VALUES 
        ('Color', ARRAY['Red', 'Blue', 'Green', 'Black', 'White']),
        ('Size', ARRAY['Small', 'Medium', 'Large', 'X-Large']),
        ('Material', ARRAY['Cotton', 'Leather', 'Plastic', 'Metal', 'Wood']),
        ('Brand', ARRAY['BrandA', 'BrandB', 'BrandC', 'BrandD']),
        ('Weight', ARRAY['Light', 'Medium', 'Heavy']),
        ('Dimensions', ARRAY['Small', 'Medium', 'Large'])
) AS v (attr_name, values)
CROSS JOIN unnest(v.values) AS name
WHERE a.name = v.attr_name;

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



-- Insert users and related data
DO $$
DECLARE
    category_rec RECORD;
    profile_id UUID;
    user_id INT;
    user_name TEXT;
    user_description TEXT;
    user_short_description TEXT;
    price NUMERIC(10, 2);
    discount_price NUMERIC(10, 2);
    attribute_id INT;
    attribute_value_id INT;
BEGIN
    FOR category_rec IN SELECT id, name FROM categories LOOP
        FOR i IN 1..10 LOOP
            -- -- Select a random profile
            SELECT id INTO profile_id FROM profiles WHERE id = 'caaa600c-ef66-4fc5-a341-fe54c164961a' LIMIT 1;
            
            -- Generate user name and description
            user_name := generate_user_name(category_rec.name);
            user_description := generate_user_description(user_name, category_rec.name);
            user_short_description := generate_short_description(user_name, category_rec.name);
            
            -- Generate prices
            price := round(cast(random() * 990 + 10 as numeric), 2);
            discount_price := round(price * cast(random() * 0.25 + 0.7 as numeric), 2);
            
            -- Insert user
            INSERT INTO users (profile_id, title, description, short_description, price, quantity, discount_price, regular_price)
            VALUES (profile_id, user_name, user_description, user_short_description, price, floor(random() * 100 + 1), discount_price, price)
            RETURNING id INTO user_id;
            
            -- Insert category_users
            INSERT INTO category_users (user_id, category_id)
            VALUES (user_id, category_rec.id);

            -- Insert user_attribute_values
            INSERT INTO user_attribute_values (attribute_value_id, user_id)
            SELECT id, user_id
            FROM attribute_values
            ORDER BY random()
            LIMIT floor(random() * 3 + 1);
            
            -- Insert user_tags
            INSERT INTO user_tags (tag_id, user_id)
            SELECT id, user_id
            FROM tags
            ORDER BY random()
            LIMIT floor(random() * 3 + 1);
        END LOOP;
    END LOOP;
END $$;

-- Insert profile_addresses
INSERT INTO profile_addresses (profile_id, address_line1, postal_code, country, city, phone_number)
SELECT 
    id,
    (floor(random() * 999) + 1)::TEXT || ' Main St',
    lpad((floor(random() * 90000) + 10000)::TEXT, 5, '0'),
    'United States',
    'Sample City',
    lpad((floor(random() * 9000000000) + 1000000000)::TEXT, 10, '0')
FROM profiles;

-- Insert possible status names without associating them with orders yet




DO $$
DECLARE
    profile_id VARCHAR(36);
    sn VARCHAR(255)[] = ARRAY['Pending', 'Processing', 'Shipped', 'Delivegreen', 'Cancelled'];
    new_order_id INTEGER;
BEGIN
    -- Example profile_id (replace with your actual profile_id)
    profile_id := '5319834e-1b74-4e70-a37c-5d2679f05c01';

    FOR i IN 1..10 LOOP -- Example: Inserting 10 orders
        -- Insert an order
        INSERT INTO orders (profile_id)
        VALUES (profile_id) -- Assuming 'Pending' has status_name_id = 1
        RETURNING id INTO new_order_id;

        INSERT INTO order_statuses (order_id, status_name)
        VALUES (new_order_id, sn[floor(random() * array_length(sn, 1) + 1)] );
        -- -- Update the corresponding order status to associate it with the new order
        -- UPDATE order_statuses
        -- SET order_id = new_order_id, status_name = 
        -- WHERE status_name = 'Pending' AND order_id IS NULL; -- Update only if not associated already

        -- Insert order items (example: inserting 1 to 5 items per order)
        FOR j IN 1..floor(random() * 5 + 1) LOOP
            INSERT INTO order_items (user_id, order_id, price, quantity)
            SELECT id, new_order_id, price, floor(random() * 5 + 1)
            FROM users
            ORDER BY random()
            LIMIT 1;
        END LOOP;
    END LOOP;
END $$;


