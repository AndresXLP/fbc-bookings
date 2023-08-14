CREATE TABLE IF NOT EXISTS public.users
(
    id            SERIAL         NOT NULL PRIMARY KEY,
    document_type VARCHAR(50)    NOT NULL,
    document_id   varchar UNIQUE NOT NULL,
    names         VARCHAR(100)   NOT NULL,
    surnames      VARCHAR(100)   NOT NULL,
    birthdate     TIMESTAMP      NOT NULL,
    age           INT            NOT NULL,
    gender        VARCHAR(1)     NULL,
    phone         VARCHAR        NULL,
    email         VARCHAR(100)   NULL,
    created_at    TIMESTAMP      NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP      NOT NULL DEFAULT now(),
    deleted_at    TIMESTAMP      NULL     DEFAULT NULL
);