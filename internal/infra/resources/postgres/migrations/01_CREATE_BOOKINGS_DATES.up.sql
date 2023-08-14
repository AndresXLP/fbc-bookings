set time zone 'America/Bogota';

CREATE TABLE IF NOT EXISTS public.booking_dates
(
    id         SERIAL           NOT NULL PRIMARY KEY,
    check_in   TIMESTAMP UNIQUE NOT NULL,
    check_out  TIMESTAMP UNIQUE NOT NULL,
    places     INT              NOT NULL,
    created_at TIMESTAMP        NOT NULL DEFAULT now(),
    updated_at TIMESTAMP        NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP        NULL     DEFAULT NULL
);