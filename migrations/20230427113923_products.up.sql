CREATE TABLE products (
    id          UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    price       BIGINT NOT NULL,
    category_id UUID REFERENCES categories(id) ON DELETE CASCADE
);