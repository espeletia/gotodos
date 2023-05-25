-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,

    CONSTRAINT "user_pkey" PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
