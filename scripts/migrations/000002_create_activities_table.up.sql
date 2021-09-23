create table if not exists activities
(
    id         bigint not null,
    name       text   not null,
    created_at bigint not null,
    constraint activities_pk
        primary key (id)
);

create table if not exists activities_users_map
(
    id          int auto_increment,
    activity_id bigint not null,
    user_id     bigint not null,
    constraint activities_users_map_pk
        primary key (id)
);
