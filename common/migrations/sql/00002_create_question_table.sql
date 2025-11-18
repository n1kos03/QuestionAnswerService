-- +goose Up
-- +goose StatementBegin
CREATE TABLE questions (
  id SERIAL PRIMARY KEY,
  text text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS questions;
-- +goose StatementEnd
