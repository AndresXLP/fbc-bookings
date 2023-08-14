CREATE TABLE IF NOT EXISTS public.bookings
(
    id              SERIAL                            NOT NULL PRIMARY KEY,
    booking_date_id INT REFERENCES public.booking_dates (id) NOT NULL,
    reserved_places INT                               NOT NULL,
    booked_by_id    INT REFERENCES public.users (id)         NOT NULL,
    status          VARCHAR(20)                       NOT NULL DEFAULT 'PENDING'
);
