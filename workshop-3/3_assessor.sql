-- сессия решения задачи

-- выдаем актуальные группы заданий
SELECT
       tg.id,
       tg.name,
       tg.description,
       price,
       seconds_to_decide,
       COUNT(t.id)
FROM task_groups tg
         INNER JOIN tasks t
                    ON tg.id = t.group_id
                        AND t.overlap > 0
                        AND finished_at isnull
GROUP BY tg.id;

-- получаем задание из группы
SELECT id, title, description
FROM tasks
WHERE group_id = 1
  AND overlap > 0
  AND finished_at isnull
LIMIT 1;

-- связываем задание и исполнителя
INSERT INTO assessor_task
    (task_id, assessor_id)
VALUES (1, 1);

-- резервируем перекрытие и стартуем задание, если оно еще не стартануло
UPDATE tasks
SET overlap = overlap - 1 AND started_at = now() WHERE id = 1;

-- даем ответ пользователя
UPDATE assessor_task
SET answer = 'сыр1'
WHERE task_id = 1
  and assessor_id = 1;

-- проверяем перекрытие у задачи
SELECT overlap FROM tasks WHERE id = 1;

-- если оно 0, то сначала получим все ответы, чтобы принять окончательный
SELECT answer FROM assessor_task WHERE task_id = 1;

-- если оно 0 и есть ответы, завершаем задание
UPDATE tasks
SET answer = 'сыр2', finished_at = now()
WHERE id = 1;




