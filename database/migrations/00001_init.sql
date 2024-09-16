-- +goose Up
create table if not exists clothes_type
(
    id   serial primary key,
    name text not null default ''
);

create table if not exists clothes_categories
(
    id   serial primary key,
    name text not null default ''
);

create table if not exists materials
(
    id      serial primary key,
    name    text    not null default '',
    country text    not null default '',
    year    date,
    price   decimal not null default 0.0
);

create table if not exists products
(
    id          serial primary key,
    name        text    not null default '',
    description text    not null default '',
    price       decimal not null default 0.0,
    discount    decimal not null default 0.0,
    stock_count int     not null default 0,
    size        int     not null default 0,
    color       text    not null default '',

    type_id     int references clothes_type (id),
    category_id int references clothes_categories (id),
    material_id int references materials (id)
);

create table if not exists images
(
    id         serial primary key,
    url        text,
    product_id int references products (id)
);

-- +goose Down
drop table clothes_type cascade;
drop table clothes_categories cascade;
drop table materials cascade;
drop table products cascade;
drop table images cascade;