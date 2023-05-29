-- students
CREATE TABLE public.students (
    id         serial      PRIMARY KEY,
    first_name varchar(63) NOT NULL,
    last_name  varchar(63) NOT NULL,
    age        int2        NOT NULL CHECK (age > 0)
);

INSERT INTO public.students (first_name, last_name, age)
VALUES
       ('Bob', 'Brown', 20),
       ('Will', 'Williams', 21),
       ('Harry', 'Bell', 19)
;

-- groups
CREATE TABLE IF NOT EXISTS public.groups (
    id         serial      PRIMARY KEY,
    name       varchar(63)
);

INSERT INTO public.groups (name)
VALUES
       ('group-1'),
       ('group-2')
;

-- students_groups
CREATE TABLE IF NOT EXISTS public.students_groups (
    student_id int REFERENCES students(id),
    group_id   int REFERENCES groups(id),
    UNIQUE(student_id)
);

INSERT INTO public.students_groups (student_id, group_id)
VALUES
       (1, 1),
       (2, 2),
       (3, 2)
;