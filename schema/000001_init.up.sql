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

CREATE TABLE operations
(
    id                serial       not null unique,
    operation_name    varchar(255) not null unique, /* возможно не unique ! */
    materials_list_id serial       not null,
    positions_list_id serial       not null
);

CREATE TABLE materials
(
    id            serial       not null unique,
    material_name varchar(255) not null unique
);

CREATE TABLE machines
(
    id           serial       not null unique,
    machine_name varchar(255) not null
);

CREATE TABLE orders
(
    id                serial       not null unique,
    order_title       varchar(255) not null,
    status            varchar(255) not null,
    customer_name     varchar(255) not null,
    end_date          varchar(255) not null,
    phone_number      varchar(255) not null,
    manager_id        serial       not null,
    operation_list_id serial       not null
    /* ФПП */
);

CREATE TABLE operations_materials
(
    operation_id serial not null,
    material_id  serial not null
);

CREATE TABLE orders_managers
(
    order_id   serial not null,
    manager_id serial not null
);

CREATE TABLE orders_operations
(
    order_id     serial not null,
    operation_id serial not null
);
