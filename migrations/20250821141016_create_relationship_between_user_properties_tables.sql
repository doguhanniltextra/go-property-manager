-- +goose Up
-- +goose StatementBegin
ALTER TABLE properties
ADD user_id INTEGER NOT NULL;

ALTER TABLE properties
ADD CONSTRAINT fk_properties_user
FOREIGN KEY (user_id)
REFERENCES users (id)
ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE properties DROP CONSTRAINT fk_properties_user;
ALTER TABLE properties DROP COLUMN user_id;
-- +goose StatementEnd
