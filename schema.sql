DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS users_to_subscribers CASCADE;

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    password VARCHAR(10),
    username VARCHAR(10) UNIQUE,
    date_of_birth DATE
);

CREATE TABLE users_to_subscribers (
    user_id INT REFERENCES users(user_id),
    friend_id INT REFERENCES users(user_id)
);