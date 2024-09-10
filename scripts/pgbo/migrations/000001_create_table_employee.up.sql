CREATE TABLE IF NOT EXISTS employee 
(
    guid CHARACTER VARYING,
    employee_id SERIAL,
    fullname VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255),
    date_of_birth TIMESTAMP,
    status CHARACTER VARYING NOT NULL DEFAULT('active'),    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by CHARACTER VARYING NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by CHARACTER VARYING,
    
    CONSTRAINT employee_pkey PRIMARY KEY(guid)
);