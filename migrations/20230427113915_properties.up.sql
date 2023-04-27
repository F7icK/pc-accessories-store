CREATE TABLE properties (
    id      UUID         DEFAULT uuid_generate_v4() PRIMARY KEY,
    name    VARCHAR(255) NOT NULL
);