CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255)
);

INSERT INTO categories (name, description) VALUES
('Pemasukan', 'Description for Category 1'),
('Pengeluaran', 'Description for Category 2');

CREATE TABLE transactions (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255),
    description VARCHAR(255),
    nominal DECIMAL,
    date DATE,
	category_id INT REFERENCES categories(id)
);

INSERT INTO transactions (category_id, name, description, nominal, date) VALUES
(1, 'Gaji', 'Description for Transactions 1', 1000, now()),
(1, 'Tunjangan', 'Description for Transactions 2', 2000, now()),
(1, 'Bonus', 'Description for Transactions 3', 3000, now()),
(2, 'Sewa Kost', 'Description for Transactions 1', 1000, now()),
(2, 'Makan', 'Description for Transactions 2', 2000, now()),
(2, 'Pakaian', 'Description for Transactions 3', 3000, now()),
(2, 'Nonton Bioskop', 'Description for Transactions 4', 4000, now());
