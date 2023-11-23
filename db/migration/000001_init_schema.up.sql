CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);

INSERT INTO categories (name, description) VALUES
('Pemasukan', 'Description for Category 1'),
('Pengeluaran', 'Description for Category 2');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255)
);