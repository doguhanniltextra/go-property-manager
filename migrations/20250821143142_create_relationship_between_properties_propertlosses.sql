-- +goose Up
-- +goose StatementBegin
ALTER TABLE propertylosses
ADD COLUMN properties_id INTEGER NOT NULL;

ALTER TABLE propertylosses
ADD CONSTRAINT fk_propertylosses_properties
FOREIGN KEY (properties_id) REFERENCES properties (id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE propertylosses DROP CONSTRAINT fk_propertylosses_properties
ALTER TABLE propertylosses DROP COLUMN properties_id;
-- +goose StatementEnd
