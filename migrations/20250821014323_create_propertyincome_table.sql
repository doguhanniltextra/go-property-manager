-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    propertyincomes (
        id SERIAL PRIMARY KEY,
        propertyIncomeName VARCHAR(255) NOT NULL,
        propertyIncomePrice INTEGER NOT NULL,
        category VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE propertyincomes;

-- +goose StatementEnd