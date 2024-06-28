CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE productList (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    os VARCHAR(255),
    ram INTEGER(16),
    cpu INTEGER(16),
    storage INTEGER(16),
    firewall BOOLEAN,
    selinux VARCHAR(255),
    location VARCHAR(255),
    owner VARCHAR(255),
    timelimit TIMESTAMP,
    username VARCHAR(64),
    password VARCHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
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
