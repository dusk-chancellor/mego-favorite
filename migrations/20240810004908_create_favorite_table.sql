-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS favorites
(
    post_id VARCHAR(255) NOT NULL,
    count INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS favorites
-- +goose StatementEnd
