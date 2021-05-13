CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS booking_bases
(
    uuid            text            NOT NULL  DEFAULT uuid_generate_v4(),
    name            text            NOT NULL,
    owner_uuid      text            NOT NULL,
    location        jsonb           NOT NULL,
    pool            boolean,
    play_area       boolean,
    dance_floor     boolean,
    photos          jsonb,    
    open_at         text,
    close_at        text,
    blocked         boolean
);

CREATE TABLE IF NOT EXISTS booking_owners
(
    uuid        text        NOT NULL                 DEFAULT uuid_generate_v4(),
    name        text        NOT NULL,
    phone       text        NOT NULL,
    requisites  text,
    blocked     boolean,
    created_at  timestamp WITH TIME ZONE    NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS booking_bookings
(
    owner_uuid      text                        NOT NULL,
    client_uuid     text                        NOT NULL,
    house_uuid      text,
    gazebo_uuid     text,
    price           integer                     NOT NULL,
    created_at      timestamp WITH TIME ZONE    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    cancelled_at    timestamp WITH TIME ZONE,
    check_in_time   timestamp WITH TIME ZONE    NOT NULL,
    move_out_time   timestamp WITH TIME ZONE    NOT NULL
);

CREATE TABLE IF NOT EXISTS booking_gazebos
(
    uuid            text        NOT NULL  DEFAULT uuid_generate_v4(),
    base_uuid       text        NOT NULL,
    base_name       text        NOT NULL,
    numder          text,
    price           integer     NOT NULL,
    number_seats    integer     NOT NULL,
    closed          boolean     NOT NULL,
    photos          jsonb
);

CREATE TABLE IF NOT EXISTS booking_houses
(
    uuid                        text        NOT NULL  DEFAULT uuid_generate_v4(),
    owner_uuid                  text        NOT NULL,
    lacation                    jsonb       NOT NULL,
    price                       integer     NOT NULL,
    photos                      text[],        
    sleeping_places             text        NOT NULL,
    allowed_number_people       text        ,
    check_in_time               text        NOT NULL,
    move_out_time               text        NOT NULL,
    assessment                  float,
    prepaymen                   integer,   
    pool                        boolean,
    description                 text,
    blocked                     boolean        
);

CREATE TABLE IF NOT EXISTS booking_clients
(
    uuid                        text        NOT NULL  DEFAULT uuid_generate_v4(),
    name                        text        NOT NULL,
    age                         integer     NOT NULL,
    bookings                    text[],
    requisites                  text,
    assessment                  float,
    blocked                     boolean
);
