CREATE EXTENSION pgcrypto;

CREATE TABLE public.users
(
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    migrosID TEXT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT false
);

CREATE UNIQUE INDEX users_nondel_idx ON users (email) WHERE is_deleted = false;