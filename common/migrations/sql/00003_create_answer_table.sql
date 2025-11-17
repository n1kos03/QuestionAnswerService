-- +goose Up
-- +goose StatementBegin
CREATE TABLE answer (
  id int PRIMARY KEY,
  question_id int,
  user_id int,
  answer TEXT,
  created_at TIMESTAMP,
  FOREIGN KEY (question_id) REFERENCES question(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
