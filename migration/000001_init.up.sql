CREATE SCHEMA IF NOT EXISTS main;

CREATE TABLE IF NOT EXISTS main.seekers (
    chat_id     BIGINT      NOT NULL UNIQUE PRIMARY KEY,
    nickname    VARCHAR(32) NOT NULL   FOREIGN KEY,
    f_name      VARCHAR(64) NOT NULL,
    s_name      VARCHAR(64) NOT NULL,
    resume      TEXT        NOT NULL
);

CREATE TABLE IF NOT EXISTS main.employers (
    chat_id     BIGINT      NOT NULL UNIQUE,
    nickname    VARCHAR(32) NOT NULL,
    company     VARCHAR(64) NOT NULL UNIQUE  PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS main.vacancies (
    vacancy_id          SERIAL      NOT NULL UNIQUE  PRIMARY KEY,
    company             VARCHAR(64) NOT NULL,
    title               TEXT        NOT NULL,
    description         TEXT        NOT NULL,
    chat_id_employer    BIGINT      NOT NULL
);

CREATE TABLE IF NOT EXISTS main.filters (
    vacancy_id      BIGINT      NOT NULL,
    tags            VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS main.responses (
    chat_id             BIGINT      NOT NULL,
    vacancy_id          BIGINT      NOT NULL,
    chat_id_employer    BIGINT      NOT NULL,
    status              VARCHAR(64) NOT NULL
);