CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    email         TEXT UNIQUE              NOT NULL,
    phone         TEXT UNIQUE              NOT NULL,
    password_hash TEXT                     NOT NULL,
    role          VARCHAR(6)               NOT NULL CHECK (role IN ('seller', 'buyer')),
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version       BIGINT                   NOT NULL DEFAULT 1
);

CREATE TABLE shops
(
    user_id     INT UNIQUE REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    name        TEXT                                               NOT NULL,
    description TEXT                                               NOT NULL,
    version     BIGINT                                             NOT NULL DEFAULT 1
);

CREATE TABLE warehouses
(
    id           SERIAL PRIMARY KEY,
    shop_user_id INT REFERENCES shops (user_id) ON DELETE CASCADE NOT NULL,
    name         TEXT                                             NOT NULL,
    location     TEXT,
    status       VARCHAR(8)                                       NOT NULL CHECK (status IN ('active', 'inactive')),
    created_at   TIMESTAMP WITH TIME ZONE                         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version      BIGINT                                           NOT NULL DEFAULT 1
);

CREATE TABLE products
(
    id           SERIAL PRIMARY KEY,
    shop_user_id INT REFERENCES shops (user_id) ON DELETE CASCADE NOT NULL,
    name         TEXT                                             NOT NULL,
    description  TEXT,
    price        DECIMAL(10, 2)                                   NOT NULL CHECK (price >= 0),
    created_at   TIMESTAMP WITH TIME ZONE                         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    version      BIGINT                                           NOT NULL DEFAULT 1
);

CREATE TABLE warehouse_product_stocks
(
    id           SERIAL PRIMARY KEY,
    product_id   INT REFERENCES products (id) ON DELETE CASCADE,
    warehouse_id INT REFERENCES warehouses (id) ON DELETE CASCADE,
    quantity     INT    NOT NULL CHECK (quantity >= 0),
    version      BIGINT NOT NULL DEFAULT 1
);

CREATE TABLE orders
(
    id           SERIAL PRIMARY KEY,
    user_id      INT REFERENCES users (id) ON DELETE CASCADE,
    total_price  DECIMAL(10, 2)           NOT NULL CHECK (total_price >= 0),
    status       VARCHAR(9)               NOT NULL CHECK (status IN ('pending', 'completed', 'expired')),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at   TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    version      BIGINT                   NOT NULL DEFAULT 1
);

CREATE TABLE order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders (id) ON DELETE CASCADE,
    product_id INT REFERENCES products (id) ON DELETE CASCADE,
    quantity   INT            NOT NULL CHECK (quantity >= 0),
    price      DECIMAL(10, 2) NOT NULL CHECK (price > 0)
);

CREATE TABLE order_item_reservations
(
    id                         SERIAL PRIMARY KEY,
    order_id                   INT REFERENCES orders (id) ON DELETE CASCADE,
    warehouse_product_stock_id INT REFERENCES warehouse_product_stocks (id) ON DELETE CASCADE,
    quantity                   INT    NOT NULL CHECK (quantity >= 0),
    version                    BIGINT NOT NULL DEFAULT 1
);

CREATE TABLE warehouse_product_transfers
(
    id                SERIAL PRIMARY KEY,
    product_id        INT REFERENCES products (id) ON DELETE CASCADE,
    from_warehouse_id INT                      REFERENCES warehouses (id) ON DELETE SET NULL,
    to_warehouse_id   INT                      REFERENCES warehouses (id) ON DELETE SET NULL,
    quantity          INT                      NOT NULL,
    transferred_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
