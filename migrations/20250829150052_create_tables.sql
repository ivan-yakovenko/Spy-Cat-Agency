-- +goose Up
-- +goose StatementBegin
CREATE TABLE spy_cats
(
    id               uuid primary key not null,
    name             varchar(255)     not null,
    experience_years integer          not null,
    breed            varchar(255)     not null,
    salary           numeric(9, 2),
    created_at       timestamp default now(),
    updated_at       timestamp default now()
);

CREATE TYPE complete_state_enum AS ENUM ('in progress', 'completed');

CREATE TABLE missions
(
    id             uuid primary key    not null,
    spycat_id      uuid unique         references spy_cats (id) ON DELETE SET NULL,
    complete_state complete_state_enum not null,
    created_at     timestamp default now(),
    updated_at     timestamp default now()
);

CREATE TABLE targets
(
    id             uuid primary key    not null,
    mission_id     uuid                not null references missions (id) on delete cascade,
    name           varchar(255)        not null,
    country        varchar(255)        not null,
    notes          text                not null,
    complete_state complete_state_enum not null,
    created_at     timestamp default now(),
    updated_at     timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS spy_cats;
DROP TABLE IF EXISTS missions;
DROP TABLE IF EXISTS targets;
DROP TYPE IF EXISTS complete_state_enum;
-- +goose StatementEnd
