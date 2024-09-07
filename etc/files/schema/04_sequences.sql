SELECT setval(pg_get_serial_sequence('users', 'id'), COALESCE(MAX(id), 1), false)
FROM users;
SELECT setval(pg_get_serial_sequence('warehouses', 'id'), COALESCE(MAX(id), 1), false)
FROM warehouses;
SELECT setval(pg_get_serial_sequence('products', 'id'), COALESCE(MAX(id), 1), false)
FROM products;
SELECT setval(pg_get_serial_sequence('warehouse_product_stocks', 'id'), COALESCE(MAX(id), 1), false)
FROM warehouse_product_stocks;
SELECT setval(pg_get_serial_sequence('orders', 'id'), COALESCE(MAX(id), 1), false)
FROM orders;
SELECT setval(pg_get_serial_sequence('order_items', 'id'), COALESCE(MAX(id), 1), false)
FROM order_items;
SELECT setval(pg_get_serial_sequence('order_item_reservations', 'id'), COALESCE(MAX(id), 1), false)
FROM order_item_reservations;
SELECT setval(pg_get_serial_sequence('warehouse_product_transfers', 'id'), COALESCE(MAX(id), 1), false)
FROM warehouse_product_transfers;
