-- Table: public.sequences

-- DROP TABLE IF EXISTS public.sequences;

CREATE TABLE IF NOT EXISTS public.sequences
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    letters text COLLATE pg_catalog."default" NOT NULL,
    is_valid boolean NOT NULL,
    CONSTRAINT sequences_pkey PRIMARY KEY (id),
    CONSTRAINT un_letters UNIQUE (letters)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.sequences
    OWNER to root;