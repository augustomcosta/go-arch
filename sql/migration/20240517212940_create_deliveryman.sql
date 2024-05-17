-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS deliverymen (
                                     id uuid PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
                                     age INT NOT NULL CHECK (age BETWEEN 0 AND 120),
                                     address_state VARCHAR(255) NOT NULL,
                                     address_city VARCHAR(255) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS deliverymen;
-- +goose StatementEnd