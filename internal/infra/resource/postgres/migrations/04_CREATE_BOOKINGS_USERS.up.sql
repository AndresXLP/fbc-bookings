CREATE TABLE IF NOT EXISTS public.booking_users
(
    booking_id INT REFERENCES public.bookings (id) NOT NULL,
    user_id    INT REFERENCES public.users (id)    NOT NULL,
    PRIMARY KEY (booking_id, user_id)
);