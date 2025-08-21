-- +goose Up
-- +goose StatementBegin
CREATE TABLE properties(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    purchasePrice FLOAT NOT NULL,
    purchaseDate VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    propertyType VARCHAR(255) NOT NULL,
    areaSqm VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE properties;
-- +goose StatementEnd
