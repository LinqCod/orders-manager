CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        title VARCHAR NOT NULL,
                                        type VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS racks (
                                        id SERIAL PRIMARY KEY,
                                        title VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS products_racks (
                                     product_type VARCHAR NOT NULL,
                                     rack_id BIGINT NOT NULL REFERENCES racks(id),
                                     rack_type VARCHAR NOT NULL,
                                     CONSTRAINT products_racks_pkey PRIMARY KEY (product_type, rack_id)
);

CREATE TABLE IF NOT EXISTS orders (
                                     id BIGINT NOT NULL,
                                     product_id BIGINT NOT NULL REFERENCES products(id),
    count BIGINT NOT NULL,
                                     CONSTRAINT orders_pkey PRIMARY KEY (id, product_id)
);

INSERT INTO racks (title) VALUES ('А');
INSERT INTO racks (title) VALUES ('Б');
INSERT INTO racks (title) VALUES ('Ж');
INSERT INTO racks (title) VALUES ('З');
INSERT INTO racks (title) VALUES ('В');

INSERT INTO products (title, type) VALUES ('Samsung 1', 'Ноутбук');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Ноутбук', 1, 'Main');
INSERT INTO products (title, type) VALUES ('Samsung Tel 1', 'Телевизор');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Телевизор', 1, 'Main');
INSERT INTO products (title, type) VALUES ('Samsung Galaxy S1', 'Телефон');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Телефон', 2, 'Main');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Телефон', 4, 'Secondary');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Телефон', 5, 'Secondary');
INSERT INTO products (title, type) VALUES ('PC 1', 'Системный блок');
INSERT INTO products (title, type) VALUES ('SmartWatch 1', 'Часы');
INSERT INTO products (title, type) VALUES ('Microphone 1', 'Микрофон');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Системный блок', 3, 'Main');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Часы', 3, 'Main');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Микрофон', 3, 'Main');
INSERT INTO products_racks (product_type, rack_id, rack_type) VALUES ('Часы', 1, 'Secondary');

INSERT INTO orders (id, product_id, count) VALUES (10, 1, 2),(11, 2, 3),(14, 1, 3),(10, 3, 1),(14, 4, 4),(15, 5, 1),(10, 6, 1);
