CREATE TABLE product_properties (
    product_id  UUID REFERENCES products(id) ON DELETE CASCADE,
    property_id UUID REFERENCES properties(id) ON DELETE CASCADE,
    value       VARCHAR(255) NOT NULL,
    PRIMARY KEY (product_id, property_id)
);