create table a(
    a_id serial primary key
);

create table b(
    b_id serial primary key,
    t text
);

insert into a
select gen.id
    from generate_series(1,2000) as gen(id)
order by gen.id;

insert into b
select gen.id,t.t
    from generate_series(1,2000) as gen(id),(values(('a'),('b'),('c'))) as t(t)
order by gen.id;

--  считаем статистику и сохраняем ее в пг
analyse a;
analyse b;

explain select * from a join b on( a.a_id = b.b_id) where a_id between 1 and 2;