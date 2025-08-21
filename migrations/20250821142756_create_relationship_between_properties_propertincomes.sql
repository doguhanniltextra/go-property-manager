-- +goose Up
-- +goose StatementBegin
ALTER TABLE propertyincomes
ADD COLUMN properties_id INTEGER NOT NULL;

ALTER TABLE propertyincomes
ADD CONSTRAINT fk_propertyincomes_properties
FOREIGN KEY (properties_id) REFERENCES properties(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE propertyincomes DROP CONSTRAINT fk_propertyincomes_properties;
ALTER TABLE propertyincomes DROP COLUMN properties_id;
-- +goose StatementEnd
