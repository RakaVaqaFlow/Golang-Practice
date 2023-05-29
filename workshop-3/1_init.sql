create table tasks
(
    id          bigserial
        constraint tasks_pk
            primary key,
    title       varchar(255)        not null,
    description text                not null,
    started_at  timestamp,
    finished_at timestamp,
    answer      varchar(255),
    group_id    bigint              not null,
    overlap     smallint  default 1 not null,
    created_at  timestamp default now(),
    updated_at  timestamp default now()
);

create table assessors
(
    id bigserial
        constraint assessors_pk
            primary key,
    name varchar(255) not null,
    surname varchar(255) not null,
    patronymic varchar(255),
    age integer not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table task_groups
(
    id bigserial
        constraint task_groups_pk
            primary key,
    name varchar(255) not null,
    description text,
    price integer,
    seconds_to_decide integer,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table rating
(
    id bigserial
        constraint rating_pk
            primary key,
    assessor_id bigint not null,
    value smallint not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table assessor_task
(
    assessor_id bigint not null,
    task_id bigint not null,
    answer varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()
);


