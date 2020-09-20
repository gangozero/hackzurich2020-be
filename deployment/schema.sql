CREATE EXTENSION pgcrypto;

CREATE TABLE public.users
(
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    migrosID TEXT NOT NULL,
    address TEXT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT false
);

CREATE UNIQUE INDEX users_nondel_idx ON users (email) WHERE is_deleted = false;


CREATE TABLE public.products
(
    productID TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    userID UUID NOT NULL REFERENCES public.users(id)
);


CREATE TABLE public.matching
(
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    state TEXT NOT NULL,
    is_full BOOLEAN NOT NULL,
    is_colleague BOOLEAN NOT NULL,
    distance INTEGER NOT NULL,
    details JSONB NOT NULL DEFAULT '{}'::jsonb,
    messages JSONB NOT NULL DEFAULT '[]'::jsonb
);