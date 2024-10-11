-- +goose Up
create table if not exists material
(
    id      serial primary key,
    name    text not null default '',
    country text not null default ''
);

create table if not exists brand
(
    id   serial primary key,
    name text not null default ''
);

create table if not exists factory
(
    id      serial primary key,
    name    text not null default '',
    country text not null default '',
    city    text not null default '',
    address text not null default ''
);

create table if not exists product
(
    id          serial primary key,
    brand_id    int references brand (id),
    factory_id  int references factory (id),
    name        text    not null default '',
    description text    not null default '',
    price       decimal not null default 0.0,
    is_approved bool not null default false, -- только на курсач
    is_deleted bool not null default false -- только на курсач
);

create table if not exists product_item
(
    id          serial primary key,
    product_id  int references product (id),
    stock_count int     not null default 0,
    size        int     not null default 0,
    weight      decimal not null default 0.0,
    color       text    not null default ''
);

create table if not exists product_material
(
    id          serial primary key,
    product_id  int references product (id),
    material_id int references material (id)
);

create table if not exists images
(
    id         serial primary key,
    url        text,
    product_id int references product (id)
);

-- +goose Down
drop table if exists images;
drop table if exists product_material;
drop table if exists product_item;
drop table if exists product;
drop table if exists factory;
drop table if exists brand;
drop table if exists material;
