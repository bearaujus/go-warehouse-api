CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255) UNIQUE,
    phone         VARCHAR(20) UNIQUE,
    password_hash VARCHAR(255),
    role          VARCHAR(6) NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version       BIGINT
);

CREATE TABLE shops
(
    id          SERIAL PRIMARY KEY,
    user_id     INT          REFERENCES users (id) ON DELETE SET NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version     BIGINT
);

CREATE TABLE warehouses
(
    id         SERIAL PRIMARY KEY,
    shop_id    INT REFERENCES shops (id) ON DELETE SET NULL,
    name       VARCHAR(255) NOT NULL,
    location   VARCHAR(255),
    status     VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version    BIGINT
);

CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    user_id     INT REFERENCES users (id) ON DELETE SET NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    price       DECIMAL(10,2) NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version     BIGINT
);

CREATE TABLE product_stocks
(
    product_id   INT REFERENCES products (id) ON DELETE CASCADE,
    warehouse_id INT REFERENCES warehouses (id) ON DELETE CASCADE,
    quantity     INT NOT NULL,
    version      BIGINT,
    PRIMARY KEY (product_id, warehouse_id)
);

CREATE TABLE orders
(
    id          SERIAL PRIMARY KEY,
    user_id     INT            REFERENCES users (id) ON DELETE SET NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    status      VARCHAR(50)    NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version     BIGINT
);

CREATE TABLE order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders (id) ON DELETE CASCADE,
    product_id INT            REFERENCES products (id) ON DELETE SET NULL,
    quantity   INT            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    version    BIGINT
);

CREATE TABLE stock_reservations
(
    id                SERIAL PRIMARY KEY,
    product_id        INT REFERENCES products (id) ON DELETE CASCADE,
    warehouse_id      INT REFERENCES warehouses (id) ON DELETE CASCADE,
    order_id          INT REFERENCES orders (id) ON DELETE CASCADE,
    reserved_quantity INT                      NOT NULL,
    reserved_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at        TIMESTAMP WITH TIME ZONE NOT NULL,
    version           BIGINT
);

CREATE TABLE warehouse_transfers
(
    id                SERIAL PRIMARY KEY,
    product_id        INT REFERENCES products (id) ON DELETE CASCADE,
    from_warehouse_id INT REFERENCES warehouses (id) ON DELETE SET NULL,
    to_warehouse_id   INT REFERENCES warehouses (id) ON DELETE SET NULL,
    quantity          INT NOT NULL,
    transferred_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version           BIGINT
);
