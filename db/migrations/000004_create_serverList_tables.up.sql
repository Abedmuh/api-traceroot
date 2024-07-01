CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE serverList (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    os VARCHAR(255),
    ram INTEGER,
    cpu INTEGER,
    storage INTEGER,
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

CREATE OR REPLACE FUNCTION add_prefix_to_serverlist_uuid() RETURNS TRIGGER AS $$
BEGIN
    NEW.id := 'serverList-' || uuid_generate_v4();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_serverlist
BEFORE INSERT ON serverList
FOR EACH ROW
EXECUTE FUNCTION add_prefix_to_serverlist_uuid();