CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS duser (
id uuid primary key default uuid_generate_v1(),
name varchar(100) not null
);

CREATE TABLE IF NOT EXISTS portfolio (
    user_id uuid primary key,
    cash_bal numeric(19,2) default 100000,
    net_worth numeric(19,2) default 100000,
    rank int,
    no_trans numeric(19,2) default 0,
    margin numeric(19,2) default 0
);

CREATE TABLE IF NOT EXISTS transaction_buy (
    user_id uuid not null,
    symbol varchar(10) not null,
    quantity numeric(19,2) default 0,
    value numeric(19,2) not null,
    time timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transaction_short_sell (
    user_id uuid not null,
    symbol varchar(10) not null,
    quantity numeric(19,2) default 0,
    value numeric(19,2) not null,
    time timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS stocks_data (
    symbol varchar(30) primary key,
    name varchar(100),
    current_price numeric(19,2),
    high numeric(19,2),
    low numeric(19,2),
    open_price numeric(19,2),
    change numeric(19,2),
    change_per numeric(19,2),
    trade_qty numeric(19,2),
    trade_value numeric(19,2)
);

CREATE TABLE IF NOT EXISTS stock_data_history (
    symbol varchar(30) not null,
    current_price numeric(19,2),
    time timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS history (
    uid uuid not null,
    symbol varchar(10) not null,
    buy_ss int not null,
    quantity numeric(19,2) default 0,
    price numeric(19,2) not null
    time timestamp default CURRENT_TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS pending (
    uid uuid not null,
    symbol varchar(10) not null,
    buy_ss varchar(30) not null,
    quantity numeric(19,2) default 0,
    value numeric(19,2) not null,
    time timestamp default CURRENT_TIMESTAMP
)