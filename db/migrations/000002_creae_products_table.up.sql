CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    id VARCHAR(45) PRIMARY KEY,
    name VARCHAR(255),
    os VARCHAR(255),
    ram VARCHAR(255),
    cpu VARCHAR(255),
    storage VARCHAR(255),
    firewall TEXT[],
    selinux BOOLEAN,
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION add_prefix_to_products() RETURNS TRIGGER AS $$
BEGIN
    NEW.id := 'products-' || uuid_generate_v4();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_products
BEFORE INSERT ON products
FOR EACH ROW
EXECUTE FUNCTION add_prefix_to_products();
