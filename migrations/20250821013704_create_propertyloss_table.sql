-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    propertylosses (
        id SERIAL PRIMARY KEY,
        propertyLossName VARCHAR(255) NOT NULL,
        propertyLossPrice INTEGER NOT NULL,
        category VARCHAR(255) NOT NULL,
        description VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE propertylosses;
-- +goose StatementEnd