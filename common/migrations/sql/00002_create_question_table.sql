-- +goose Up
-- +goose StatementBegin
CREATE TABLE question (
  id int PRIMARY KEY,
  question text,
  created_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS question;
-- +goose StatementEnd
