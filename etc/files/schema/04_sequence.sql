-- Adjust the sequence for the users table
SELECT setval(pg_get_serial_sequence('users', 'id'), COALESCE(MAX(id), 1)) FROM users;

-- Adjust the sequence for the shops table (user_id is a serial field here)
SELECT setval(pg_get_serial_sequence('shops', 'user_id'), COALESCE(MAX(user_id), 1)) FROM shops;

-- Adjust the sequence for the warehouses table
SELECT setval(pg_get_serial_sequence('warehouses', 'id'), COALESCE(MAX(id), 1)) FROM warehouses;

-- Adjust the sequence for the products table
SELECT setval(pg_get_serial_sequence('products', 'id'), COALESCE(MAX(id), 1)) FROM products;

-- Adjust the sequence for the warehouse_product_stocks table
SELECT setval(pg_get_serial_sequence('warehouse_product_stocks', 'id'), COALESCE(MAX(id), 1)) FROM warehouse_product_stocks;

-- Adjust the sequence for the orders table
SELECT setval(pg_get_serial_sequence('orders', 'id'), COALESCE(MAX(id), 1)) FROM orders;

-- Adjust the sequence for the order_items table
SELECT setval(pg_get_serial_sequence('order_items', 'id'), COALESCE(MAX(id), 1)) FROM order_items;

-- Adjust the sequence for the order_item_reservations table
SELECT setval(pg_get_serial_sequence('order_item_reservations', 'id'), COALESCE(MAX(id), 1)) FROM order_item_reservations;

-- Adjust the sequence for the warehouse_product_transfers table
SELECT setval(pg_get_serial_sequence('warehouse_product_transfers', 'id'), COALESCE(MAX(id), 1)) FROM warehouse_product_transfers;