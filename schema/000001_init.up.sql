CREATE TABLE users
(
    id            serial       not null unique,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    first_name    varchar(255) not null,
    second_name   varchar(255) not null,
    third_name    varchar(255) not null,
    user_role     varchar(255) not null,
    status        varchar(255) not null
);
