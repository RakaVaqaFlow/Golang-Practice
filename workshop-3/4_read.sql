-- чтение заказчиком
SELECT t.id,
       t.title,
       t.description,
       t.answer,
       tg.id,
       tg.name,
       tg.description,
       tg.price,
       tg.seconds_to_decide
FROM tasks t
         INNER JOIN task_groups tg
                    ON tg.id = t.group_id
                        AND t.overlap = 0
                        AND t.finished_at is not null;

-- задачи, решение исполнителями от 25 до 35 лет с рейтингом выше 50
SELECT t.id,
       t.title,
       t.description,
       t.answer,
       tg.id,
       tg.name,
       tg.description,
       tg.price,
       tg.seconds_to_decide,
       at.answer
FROM tasks t
         INNER JOIN task_groups tg
                    ON tg.id = t.group_id
         INNER JOIN assessor_task at on t.id = at.task_id
         INNER JOIN rating r on at.assessor_id = r.assessor_id
         INNER JOIN assessors a on at.assessor_id = a.id
WHERE t.overlap = 0
  AND t.finished_at is not null
  AND r.value > 50
  AND a.age BETWEEN 25 AND 35;

-- задачи за прошлые сутки, которые не были решены
SELECT t.id,
       t.title,
       t.description,
       t.answer,
       tg.id,
       tg.name,
       tg.description,
       tg.price,
       tg.seconds_to_decide
FROM tasks t
         INNER JOIN task_groups tg
                    ON tg.id = t.group_id
                        AND t.finished_at is null
                        AND t.created_at
                           BETWEEN current_date - interval '1 day' AND current_date;

-- развернутый отчет по всем задачам заказчика с исполнителями
SELECT t.id,
       t.title,
       t.description,
       t.answer,
       tg.id,
       tg.name,
       tg.description,
       tg.price,
       tg.seconds_to_decide,
       at.answer,
       a.id,
       a.name,
       a.surname,
       a.patronymic,
       a.age
FROM tasks t
        INNER JOIN task_groups tg
                    ON tg.id = t.group_id
        INNER JOIN assessor_task at on t.id = at.task_id
        INNER JOIN assessors a on at.assessor_id = a.id;