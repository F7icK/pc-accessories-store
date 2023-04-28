CREATE TABLE properties
(
    id         UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    name       VARCHAR(255)                                       NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);