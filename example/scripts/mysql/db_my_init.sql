DROP TABLE IF EXISTS my_vendors;
DROP TABLE IF EXISTS my_auth;
DROP TABLE IF EXISTS my_customers;
DROP TABLE IF EXISTS my_contacts;
CREATE TABLE my_vendors (  
    vendor_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    tradename VARCHAR(255) NOT NULL,
    current_type ENUM('personal', 'business'),
    deleted_at DATE,
    created_at DATE,
    updated_at DATE
);

CREATE TABLE my_auth (  
    auth_id BINARY(16) default (UUID_TO_BIN(UUID())),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT,
    vendor_id INT,
    current_provider ENUM('email', 'google', 'apple', 'others'),
    deleted_at DATE,
    created_at DATE,
    updated_at DATE,
    PRIMARY KEY (auth_id),
    CONSTRAINT fk_vendor
        FOREIGN KEY(vendor_id) 
            REFERENCES my_vendors(vendor_id)

);

CREATE TABLE my_customers(
    customer_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    deleted_at DATE,
    created_at DATE,
    updated_at DATE
);

CREATE TABLE my_contacts(
    contact_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT,
    contact_name VARCHAR(255) NOT NULL,
    phone VARCHAR(15),
    email VARCHAR(100),
    deleted_at DATE,
    created_at DATE,
    updated_at DATE,
   CONSTRAINT fk_customer
      FOREIGN KEY(customer_id) 
        REFERENCES my_customers(customer_id)
);