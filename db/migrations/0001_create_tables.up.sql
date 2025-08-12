CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255),
    phone VARCHAR(20),
    address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            description TEXT,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price INTEGER NOT NULL,
    image_url TEXT,
    category_id INTEGER REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cart (
                      id SERIAL PRIMARY KEY,
                      user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                      product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                      quantity VARCHAR(255),
                      size VARCHAR(255),
                      color VARCHAR(255),
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE inventory (
                           id SERIAL PRIMARY KEY,
                           product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                           quantity INTEGER NOT NULL,
                           size VARCHAR(255),
                           color VARCHAR(255),
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                        order_status VARCHAR(255),
                        total_amount DECIMAL(10, 2) NOT NULL,
                        shipping_address TEXT,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
                             product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
                             quantity INTEGER NOT NULL,
                             size VARCHAR(255),
                             color VARCHAR(255),
                             price_at_order DECIMAL(10, 2) NOT NULL
);

CREATE TABLE payments (
                          id SERIAL PRIMARY KEY,
                          order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
                          payment_method VARCHAR(255),
                          payment_status VARCHAR(255),
                          transaction_id VARCHAR(255) UNIQUE,
                          paid_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE promocodes (
                            id SERIAL PRIMARY KEY,
                            code VARCHAR(255) UNIQUE NOT NULL,
                            discount_percent INTEGER,
                            discount_amount DECIMAL(10, 2),
                            valid_from TIMESTAMP,
                            valid_until TIMESTAMP,
                            usage_limit INTEGER,
                            times_used INTEGER DEFAULT 0
);