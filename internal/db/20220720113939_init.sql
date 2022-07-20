-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    Id SERIAL PRIMARY KEY not null unique ,
    Name varchar(255) not null,
    PhoneNumber varchar(255) not null,
    Email varchar(255) not null ,
    Password varchar(255) not null,
    CreatedAt timestamp not null,
    UpdatedAt timestamp not null
);
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
SELECT 'down SQL query';
-- +goose StatementEnd
