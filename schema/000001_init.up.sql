create table order_delivery
(
    name    varchar not null
        primary key,
    phone   varchar not null,
    zip     varchar not null,
    city    varchar not null,
    address varchar,
    region  varchar,
    email   varchar
);

alter table order_delivery
    owner to postgres;

create table order_payment
(
    transaction   varchar not null
        primary key,
    request_id    varchar,
    currency      varchar not null,
    provider      varchar not null,
    amount        integer not null,
    payment_dt    bigint  not null,
    bank          varchar not null,
    delivery_cost integer not null,
    goods_total   integer not null,
    custom_fee    integer not null
);

alter table order_payment
    owner to postgres;

create table orders
(
    order_uid          varchar(255) not null
        primary key,
    track_number       varchar(255) not null,
    entry              varchar(255) not null,
    delivery           varchar      not null
        constraint delivery
            references order_delivery,
    payment            varchar(255) not null
        constraint payment
            references order_payment,
    locale             varchar(5)   not null,
    internal_signature varchar,
    customer_id        varchar      not null,
    delivery_service   varchar      not null,
    shardkey           char         not null,
    sm_id              integer      not null,
    data_created       varchar      not null,
    oof_shard          varchar      not null
);

alter table orders
    owner to postgres;

create table order_items
(
    chrt_id      bigint   not null
        primary key,
    track_number varchar  not null,
    price        integer  not null,
    rid          varchar  not null,
    name         varchar  not null,
    sale         integer  not null,
    size         varchar  not null,
    total_price  integer  not null,
    nm_id        bigint   not null,
    brand        varchar  not null,
    status       smallint not null,
    order_uid    varchar  not null
        references orders
);

alter table order_items
    owner to postgres;

