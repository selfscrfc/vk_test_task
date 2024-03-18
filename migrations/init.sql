-- we don't know how to generate root <with-no-name> (class Root) :(

comment on database postgres is 'default administrative connection database';

create table actor
(
    name  varchar(128),
    sex   integer,
    birth date,
    id    varchar(64) not null
        primary key
);

alter table actor
    owner to postgres;

create table film
(
    name          varchar(256),
    description   varchar(1024),
    date_released date,
    rate          integer,
    id            varchar(64) not null
        primary key
);

alter table film
    owner to postgres;

create table "user"
(
    id       serial
        primary key,
    login    varchar(128),
    password varchar(256) not null,
    is_admin boolean      not null,
    user_id  varchar(64)
);

alter table "user"
    owner to postgres;

create table film_actor
(
    film_id  varchar(64) not null
        references film
        constraint film_actor_film_id_fkey1
            references film,
    actor_id varchar(64) not null
        references actor
        constraint film_actor_actor_id_fkey1
            references actor
);

alter table film_actor
    owner to postgres;

