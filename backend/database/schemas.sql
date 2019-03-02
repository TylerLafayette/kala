CREATE TABLE users_kala (
    user_id SERIAL NOT NULL UNIQUE PRIMARY KEY,
    username text,
    current_balance numeric
);

CREATE TABLE transactions (
    transaction_id SERIAL NOT NULL UNIQUE PRIMARY KEY,
    description text,
    price numeric,
    final_balance numeric
);

CREATE TABLE subscriptions (
    subscription_id SERIAL NOT NULL UNIQUE PRIMARY KEY,
    amount numeric,
    reason text
);