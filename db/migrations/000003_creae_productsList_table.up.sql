CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE productList (
    id VARCHAR(45) PRIMARY KEY,
    id_products VARCHAR(45) REFERENCES products(id),
    owner UUID REFERENCES users(id),
    timelimit TIME,
    username VARCHAR(64),
    pass VARCHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION add_prefix_to_productlist_uuid() RETURNS TRIGGER AS $$
BEGIN
    NEW.id := 'productList-' || uuid_generate_v4();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_productlist
BEFORE INSERT ON productList
FOR EACH ROW
EXECUTE FUNCTION add_prefix_to_productlist_uuid();
