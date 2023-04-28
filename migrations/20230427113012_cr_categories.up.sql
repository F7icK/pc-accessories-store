CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE categories
(
    id         UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    name       VARCHAR(255)                                       NOT NULL,
    parent_id  UUID,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);





