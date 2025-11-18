-- +goose Up
-- +goose StatementBegin
INSERT INTO users (user_name) VALUES ('Jhon'), ('Fill'), ('Anna');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id = 1, id = 2, id = 3;
-- +goose StatementEnd
