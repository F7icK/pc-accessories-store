CREATE TABLE products
(
    id          UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR(255)                                       NOT NULL,
    price       BIGINT                                             NOT NULL,
    category_id UUID REFERENCES categories (id) ON DELETE CASCADE,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE
);