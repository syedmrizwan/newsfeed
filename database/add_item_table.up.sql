CREATE TABLE if not exists item
(
    title character varying COLLATE pg_catalog."default" NOT NULL DEFAULT 'up'::character varying,
    post character varying COLLATE pg_catalog."default" NOT NULL DEFAULT 'up'::character varying,
    stats json NOT NULL
)


insert into item (title, post, stats) Values ('Item Two', 'Sports Post', '{"views": 250, "likes": 85}')