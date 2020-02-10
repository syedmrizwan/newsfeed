CREATE TABLE if not exists item
(
    title character varying COLLATE pg_catalog."default" NOT NULL DEFAULT 'up'::character varying,
    post character varying COLLATE pg_catalog."default" NOT NULL DEFAULT 'up'::character varying
)
