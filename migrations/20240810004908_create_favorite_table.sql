-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS favorites
(
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    post_id VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS favorites
-- +goose StatementEnd
