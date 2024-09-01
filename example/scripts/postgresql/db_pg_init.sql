-- Active: 1721852222098@@127.0.0.1@5432@protomagic@public
DROP TABLE IF EXISTS pg_vendors;
DROP TABLE IF EXISTS pg_auth;
DROP TABLE IF EXISTS pg_customers;
DROP TABLE IF EXISTS pg_contacts;

CREATE TYPE type_pg_vendor AS ENUM ('personal', 'business');
CREATE TYPE type_pg_provider AS ENUM ('email', 'google', 'apple', 'others');
CREATE TABLE pg_vendors (  
    vendor_id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    tradename VARCHAR(255) NOT NULL,
    current_type type_pg_vendor,
    deleted_at DATE,
    created_at DATE,
    updated_at DATE
);
COMMENT ON TABLE pg_vendors IS 'Vendor table';
COMMENT ON COLUMN pg_vendors.name IS 'Vendor Name';

CREATE TABLE pg_auth (  
    auth_id uuid DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT,
    vendor_id INT,
    current_provider type_pg_provider,
    deleted_at DATE,
    created_at DATE,
    updated_at DATE,
    PRIMARY KEY (auth_id),
    CONSTRAINT fk_vendor
        FOREIGN KEY(vendor_id) 
            REFERENCES pg_vendors(vendor_id)

);
COMMENT ON TABLE pg_auth IS 'pg_auth table';
COMMENT ON COLUMN pg_auth.username IS 'Username of vendor test';

CREATE TABLE pg_customers(
    customer_id INT GENERATED ALWAYS AS IDENTITY,
    customer_name VARCHAR(255) NOT NULL,
    deleted_at DATE,
    created_at DATE,
    updated_at DATE,
   PRIMARY KEY(customer_id)
);

CREATE TABLE pg_contacts(
    contact_id INT GENERATED ALWAYS AS IDENTITY,
    customer_id INT,
    contact_name VARCHAR(255) NOT NULL,
    phone VARCHAR(15),
    email VARCHAR(100),
    deleted_at DATE,
    created_at DATE,
    updated_at DATE,
   PRIMARY KEY(contact_id),
   CONSTRAINT fk_customer
      FOREIGN KEY(customer_id) 
        REFERENCES pg_customers(customer_id)
);