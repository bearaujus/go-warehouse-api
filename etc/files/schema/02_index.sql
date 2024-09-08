-- Index for users table
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);

-- Index for shops table
CREATE INDEX idx_shops_user_id ON shops(user_id);

-- Index for warehouses table
CREATE INDEX idx_warehouses_shop_user_id ON warehouses(shop_user_id);
CREATE INDEX idx_warehouses_status ON warehouses(status);

-- Index for products table
CREATE INDEX idx_products_shop_user_id ON products(shop_user_id);

-- Index for warehouse_product_stocks table
CREATE INDEX idx_warehouse_product_stocks_product_id ON warehouse_product_stocks(product_id);
CREATE INDEX idx_warehouse_product_stocks_warehouse_id ON warehouse_product_stocks(warehouse_id);

-- Index for orders table
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_expires_at ON orders(expires_at);
CREATE INDEX idx_orders_completed_at ON orders(completed_at);

-- Index for order_items table
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_product_id ON order_items(product_id);

-- Index for order_item_reservations table
CREATE INDEX idx_order_item_reservations_order_id ON order_item_reservations(order_id);
CREATE INDEX idx_order_item_reservations_warehouse_product_stock_id ON order_item_reservations(warehouse_product_stock_id);

-- Index for warehouse_product_transfers table
CREATE INDEX idx_warehouse_product_transfers_product_id ON warehouse_product_transfers(product_id);
CREATE INDEX idx_warehouse_product_transfers_from_warehouse_id ON warehouse_product_transfers(from_warehouse_id);
CREATE INDEX idx_warehouse_product_transfers_to_warehouse_id ON warehouse_product_transfers(to_warehouse_id);