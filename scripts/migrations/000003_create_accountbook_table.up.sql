create table if not exists account_book
(
    id         bigint not null,
    payer_id   bigint not null,
    remark     text   not null,
    created_at bigint not null,
    constraint account_book_pk
        primary key (id)
);

create index account_book_payer_id_index
    on account_book (payer_id);

create table if not exists cost_history
(
    id        int auto_increment,
    record_id bigint not null,
    user_id   bigint not null,
    spend     int    not null,
    constraint cost_history_pk
        primary key (id)
);

create index cost_history_record_id_index
    on cost_history (record_id);

create index cost_history_user_id_index
    on cost_history (user_id);
