create table if not exists spend_history
(
    id          bigint not null,
    activity_id bigint not null,
    payer_id    bigint not null,
    remark      text   not null,
    created_at  bigint not null,
    constraint spend_history_pk
        primary key (id)
);

create index spend_history_activity_id_index
    on spend_history (activity_id);

create index spend_history_payer_id_index
    on spend_history (payer_id);

create table if not exists spend_details
(
    id       int auto_increment,
    spend_id bigint not null,
    user_id  bigint not null,
    value    int    not null,
    constraint spend_details_pk
        primary key (id)
);

create index spend_details_spend_id_index
    on spend_details (spend_id);

create index spend_details_user_id_index
    on spend_details (user_id);
