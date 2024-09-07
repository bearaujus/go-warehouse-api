-- INSERT INTO users table
INSERT INTO users (id, email, phone, password_hash, role, created_at, version)
VALUES
    (1, 'dmp.seller@domain.com', '080000000001', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'seller', '2024-09-07 18:00:00.000000 +00:00', 1),
    (2, 'dmp.seller2@domain.com', '080000000002', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'seller', '2024-09-07 18:00:00.000000 +00:00', 1),
    (3, 'dmp.buyer@domain.com', '080000000003', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'buyer', '2024-09-07 18:00:00.000000 +00:00', 1);

-- INSERT INTO shops table
INSERT INTO shops (user_id, name, description, version)
VALUES
    (1, 'Tech Gadgets', 'A shop specializing in the latest tech gadgets and accessories', 1),
    (2, 'Home Essentials', 'A shop offering a variety of home goods and essentials', 1);

-- INSERT INTO warehouses table
INSERT INTO warehouses (id, shop_user_id, name, location, status, created_at, version)
VALUES
    (1, 1, 'Main Warehouse - Tech Gadgets', 'Jakarta', 'active', '2024-09-07 18:00:00.000000 +00:00', 1),
    (2, 1, 'Backup Warehouse - Tech Gadgets', 'Bandung', 'active', '2024-09-07 18:00:00.000000 +00:00', 1),
    (3, 1, 'Inactive Warehouse - Tech Gadgets', 'Surabaya', 'inactive', '2024-09-07 18:00:00.000000 +00:00', 1),
    (4, 2, 'Main Warehouse - Home Essentials', 'Jakarta', 'active', '2024-09-07 18:00:00.000000 +00:00', 1),
    (5, 2, 'Backup Warehouse - Home Essentials', 'Bandung', 'active', '2024-09-07 18:00:00.000000 +00:00', 1),
    (6, 2, 'Inactive Warehouse - Home Essentials', 'Surabaya', 'inactive', '2024-09-07 18:00:00.000000 +00:00', 1);

-- INSERT INTO products table
INSERT INTO products (id, shop_user_id, name, description, price, created_at, version)
VALUES
    (1, 1, 'Smartphone X', 'The latest Smartphone X with cutting-edge features', 999.99, '2024-09-07 18:00:00.000000 +00:00', 1),
    (2, 1, 'Wireless Headphones', 'High-quality wireless headphones with noise cancellation', 199.99, '2024-09-07 18:00:00.000000 +00:00', 1),
    (3, 1, 'Smartwatch Pro', 'Advanced smartwatch with fitness tracking and heart rate monitoring', 299.99, '2024-09-07 18:00:00.000000 +00:00', 1),
    (4, 2, 'Vacuum Cleaner', 'Powerful vacuum cleaner for home cleaning', 149.99, '2024-09-07 18:00:00.000000 +00:00', 1),
    (5, 2, 'Air Purifier', 'Air purifier with advanced filtration technology', 249.99, '2024-09-07 18:00:00.000000 +00:00', 1),
    (6, 2, 'Kitchen Blender', 'High-performance kitchen blender for smoothies and food preparation', 99.99, '2024-09-07 18:00:00.000000 +00:00', 1);

-- INSERT INTO warehouse_product_stocks
INSERT INTO warehouse_product_stocks (product_id, warehouse_id, quantity, version)
VALUES
    -- Stocks for product_id = 1 (Smartphone X)
    (1, 1, 100, 1), -- Main Warehouse - Tech Gadgets
    (1, 2, 50, 1),  -- Backup Warehouse - Tech Gadgets
    (1, 3, 70, 1),  -- Inactive Warehouse - Tech Gadgets

    -- Stocks for product_id = 2 (Wireless Headphones)
    (2, 1, 200, 1), -- Main Warehouse - Tech Gadgets
    (2, 2, 75, 1),  -- Backup Warehouse - Tech Gadgets
    (2, 3, 66, 1),  -- Inactive Warehouse - Tech Gadgets

    -- Stocks for product_id = 3 (Smartwatch Pro)
    (3, 1, 150, 1), -- Main Warehouse - Tech Gadgets
    (3, 2, 100, 1), -- Backup Warehouse - Tech Gadgets
    (3, 3, 23, 1),  -- Inactive Warehouse - Tech Gadgets

    -- Stocks for product_id = 4 (Vacuum Cleaner)
    (4, 4, 80, 1),  -- Main Warehouse - Home Essentials
    (4, 5, 40, 1),  -- Backup Warehouse - Home Essentials
    (4, 6, 15, 1),  -- Inactive Warehouse - Home Essentials

    -- Stocks for product_id = 5 (Air Purifier)
    (5, 4, 60, 1),  -- Main Warehouse - Home Essentials
    (5, 5, 30, 1),  -- Backup Warehouse - Home Essentials
    (5, 6, 29, 1),  -- Inactive Warehouse - Home Essentials

    -- Stocks for product_id = 6 (Kitchen Blender)
    (6, 4, 120, 1), -- Main Warehouse - Home Essentials
    (6, 5, 50, 1),  -- Backup Warehouse - Home Essentials
    (6, 6, 95, 1); -- Inactive Warehouse - Home Essentials