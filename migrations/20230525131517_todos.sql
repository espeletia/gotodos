-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
    id SERIAL NOT NULL,
    user_id integer NOT NULL,
    title VARCHAR(255) NOT NULL,
    completed boolean NOT NULL,

    CONSTRAINT "todo_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "todo_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd
