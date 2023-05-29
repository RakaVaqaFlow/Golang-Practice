CREATE INDEX assessor_task_task_id_assessor_id_index ON assessor_task (task_id, assessor_id);

CREATE INDEX tasks_group_id_index ON tasks (group_id);