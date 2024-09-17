CREATE TABLE IF NOT EXISTS authentication (
    guid CHARACTER VARYING, 
    id SERIAL,
    employee_guid CHARACTER VARYING UNIQUE,
    username CHARACTER VARYING NOT NULL UNIQUE,
    password CHARACTER VARYING NOT NULL,
    forgot_password_token CHARACTER VARYING,
    forgot_password_expiry TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    last_login TIMESTAMP WITH TIME ZONE,
    
    status CHARACTER VARYING NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by CHARACTER VARYING NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by CHARACTER VARYING,

    CONSTRAINT authentication_pkey PRIMARY KEY(guid)
);