DROP TABLE IF EXISTS public.author CASCADE ;
DROP TABLE IF EXISTS public.book CASCADE;
DROP TABLE IF EXISTS public.book_authors CASCADE;

CREATE TABLE public.author (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age int,
    is_alive bool,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE INDEX idx_author_created_at_pagination ON public.author (created_at, id);
CREATE INDEX idx_author_age_pagination ON public.author (age, id);

CREATE TABLE public.book (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    age INT,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);
CREATE INDEX idx_book_created_at_pagination ON public.book (created_at, id);
CREATE INDEX idx_book_age_pagination ON public.book (age, id);

CREATE TABLE public.book_authors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id UUID NOT NULL,
    author_id UUID NOT NULL,
    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id),
    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id),
    CONSTRAINT book_authors_unique unique (book_id, author_id)
);

INSERT INTO public.book (id, name, age) VALUES ('c6254b6e-d18e-489d-8859-93b4d9b1219b','Властелин колец', 2006);
INSERT INTO public.book (id, name, age) VALUES ('6ad58474-9a89-4abb-8f49-bdddb33746af', 'Гордость и предубеждение', 2001);
INSERT INTO public.book (id, name) VALUES ('cae7c9f4-6478-49c5-b19c-cc861d5e4c1f', 'Тёмные начала');

INSERT INTO public.book (name, age) VALUES ('A1', 1);
INSERT INTO public.book (name, age) VALUES ('A2', 2);
INSERT INTO public.book (name, age) VALUES ('A3', 3);
INSERT INTO public.book (name, age) VALUES ('A4', 4);
INSERT INTO public.book (name, age) VALUES ('A5', 5);
INSERT INTO public.book (name, age) VALUES ('A6', 6);
INSERT INTO public.book (name, age) VALUES ('A7', 7);

INSERT INTO public.author (id, name, age, is_alive) VALUES ('4cf14ef9-9c0e-4480-8c86-709acc03114e', 'Джон Р. Р. Толкин', 72, true);
INSERT INTO public.author (id, name, age, is_alive) VALUES ('0f739c77-59bc-437f-9c18-38f2b4fa93bc', 'Филип Пулман', 46, true);
INSERT INTO public.author (id, name, age, is_alive) VALUES ('40452dec-04e8-4379-9f61-07b0016e428f', 'Джейн Остин', 32, false);
INSERT INTO public.author (name, age, is_alive) VALUES ('Робби Уильямс', 65, false);
INSERT INTO public.author (name, age, is_alive) VALUES ('Питтер Паркер', 23, true);
INSERT INTO public.author (name, age, is_alive) VALUES ('Иван Ургант', 64, true);
INSERT INTO public.author (name, age, is_alive) VALUES ('Дон Алькапоне', 87, false);
INSERT INTO public.author (name, age, is_alive) VALUES ('Лютер Баумен', 21, true);
INSERT INTO public.author (name, age, is_alive) VALUES ('Ирина Ламба', 48, true);
INSERT INTO public.author (name, age, is_alive) VALUES ('Автор 1', 48, true);
INSERT INTO public.author (name, age, is_alive) VALUES ('Автор 2', 12, false);
INSERT INTO public.author (name, age, is_alive) VALUES ('Автор 3', 43, true);
INSERT INTO public.author (name, age, is_alive) VALUES ('Автор 4', 23, false);
INSERT INTO public.author (name, age, is_alive) VALUES ('Автор 5', 67, true);


INSERT INTO public.book_authors(book_id, author_id) VALUES ('c6254b6e-d18e-489d-8859-93b4d9b1219b', '4cf14ef9-9c0e-4480-8c86-709acc03114e');
INSERT INTO public.book_authors(book_id, author_id) VALUES ('6ad58474-9a89-4abb-8f49-bdddb33746af', '40452dec-04e8-4379-9f61-07b0016e428f');
INSERT INTO public.book_authors(book_id, author_id) VALUES ('6ad58474-9a89-4abb-8f49-bdddb33746af', '0f739c77-59bc-437f-9c18-38f2b4fa93bc');
INSERT INTO public.book_authors(book_id, author_id) VALUES ('cae7c9f4-6478-49c5-b19c-cc861d5e4c1f', '0f739c77-59bc-437f-9c18-38f2b4fa93bc');


SELECT *
FROM book
ORDER BY created_at DESC, id DESC
LIMIT 2;

SELECT
    b.id,
    b.name,
    b.created_at,
    b.age
FROM book b
WHERE (created_at, id) < ('2022-08-12 11:21:28.236088' :: timestamp, 'fa93df85-c052-424a-b198-eaec81048879')
ORDER BY created_at DESC, id DESC
-- OFFSET 3
LIMIT 2