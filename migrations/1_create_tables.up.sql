CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    hashPass VARCHAR(255) NOT NULL
);

CREATE TABLE portfolio(
  id SERIAL PRIMARY KEY,
  user_id int REFERENCES users (id) ON DELETE CASCADE,
  account bigint default 0
);

CREATE TABLE symbol(
    id SERIAL PRIMARY KEY,
    abbr VARCHAR(6),
    full_name VARCHAR(255)
);

CREATE TABLE deal(
    id SERIAL PRIMARY KEY,
    type VARCHAR(10) NOT NULL,
    symbol_id int REFERENCES symbol(id),
    symbol_price DECIMAL NOT NULL,
    number INTEGER NOT NULL,
    amount DECIMAL NOT NULL,
    date DATE,
    portfolio_id int REFERENCES portfolio(id),
    user_id int REFERENCES users (id)
);

CREATE TABLE active_share(
    id SERIAL PRIMARY KEY,
    price DECIMAL NOT NULL ,
    number INT NOT NULL,
    portfolio_id int REFERENCES portfolio(id) ON DELETE CASCADE,
    symbol_id int REFERENCES symbol(id),
    deal_id int REFERENCES deal(id)
);