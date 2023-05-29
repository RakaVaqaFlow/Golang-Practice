
create table some_table(
    id serial primary key,
    some_jsonb_field jsonb
);

-- создание
-- Из текстового представления
SELECT '{}'::jsonb;
-- конструктор
SELECT jsonb_build_object();

UPDATE some_table SET some_jsonb_field = jsonb_set(some_jsonb_field,'{some_flag}','true',true);

-- Из объектов в БД
SELECT row_to_json(some_field) from some_table;


-- извлечение из jsonb
-- если -> то будет json
-- если ->> то будет text
SELECT to_jsonb(some_jsonb_field)->>'some_key' from some_table;
SELECT to_jsonb(some_jsonb_field)->'some_key' from some_table;

-- операторы
SELECT json #> text[];
SELECT json #>> text[];

-- функции
SELECT * from json_each_text('{"a":"foo","b":"bar"}')

--поиск по jsonb
SELECT * from some_table where some_jsonb_field ->>'key' = 'value'; -- можно использовать B tree
SELECT * from some_table where exists( select * from json_each_text('{"a":"foo","b":"bar"}') where key='a' and value = 'foo')
SELECT * from some_table where exists( select * from json_each_text(some_jsonb_field) where key='a' and value = 'foo')
    -- можно формулировать сложные запросы, но индексы вам не помогут

-- индекс
CREATE INDEX some_jsonb on some_table USING GIN (some_jsonb_field)