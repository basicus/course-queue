create schema  if not exists  remains;
create schema if not exists  orders;

create table if not exists  orders.transactions
(
    id        uuid    not null
        constraint order_transactions_pk
            primary key,
    good_id   integer not null,
    operation varchar not null,
    pcs       integer,
    created_at timestamp with time zone default NOW() not null
);

create index if not exists order_transactions_good_id_index
    on orders.transactions (good_id);


create table if not exists remains.remains (
  good_id integer primary key not null,
  cnt integer not null
);

create table if not exists  remains.transactions_log
(
    transaction_id uuid not null
        constraint transactions_log_pk
            unique
);