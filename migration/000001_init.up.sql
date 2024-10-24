CREATE SCHEMA IF NOT EXISTS main;

CREATE TABLE IF NOT EXISTS main.seekers (
    id SERIAL,
    username        VARCHAR(32) NOT NULL UNIQUE PRIMARY KEY,
    f_name          VARCHAR(64) NOT NULL,
    password_hash   TEXT        NOT NULL,
    resume          TEXT        NOT NULL
);

CREATE TABLE IF NOT EXISTS main.employers (
    employer_id     SERIAL      NOT NULL UNIQUE,
    username        VARCHAR(32) NOT NULL UNIQUE,
    f_name          VARCHAR(64) NOT NULL,
    password_hash   TEXT        NOT NULL,
    company         VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS main.vacancies (
    vacancy_id          SERIAL      NOT NULL UNIQUE PRIMARY KEY,
    company             TEXT        NOT NULL,
    title               TEXT        NOT NULL,
    description         TEXT        NOT NULL,
    employer_id         BIGINT      NOT NULL,
    status              TEXT        NOT NULL
);

CREATE TABLE IF NOT EXISTS main.filters (
    vacancy_id      BIGINT      NOT NULL,
    tag             VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS main.responses (
    response_id SERIAL      NOT NULL    UNIQUE,
    vacancy_id  BIGINT      NOT NULL,
    employer_id BIGINT      NOT NULL,
    username    VARCHAR(32) NOT NULL, 
    status      VARCHAR(64) NOT NULL
);

CREATE TABLE main.messages (
    id SERIAL PRIMARY KEY,
    response_id INT,
    sender_type VARCHAR(10),
    sender_id INT,
    receiver_type VARCHAR(10), -- "seeker" или "employer"
    receiver_id INT,
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);