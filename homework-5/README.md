# Домашнее задание

## Задача
- Разработать схему базы данных, состояющую минимум из 2 связанных таблиц(минимум 2 поля кроме primary key)
- Написать миграции при помощи goose, подготовить скрипт для их выполнения(MakeFile)
- При помощи паттерна репозиторий написать crud методы, работающие с таблицами из п.1, каждый метод должен быть доступен при помощи консольной команды
- Составить план по подключению кэшей, предложить тип кэша(инструмент), обосновать свой выбор.(текстовое обоснование вашего видения как должны заполнятся и инвалидироваться кэши)
- 💎 Подготовить sql скрипты, которые заполняют таблицы тестовыми данными

## Решение 

### 1. Схема базы данных

Решение будет основываться на базе данных со следующей структурой:

**Users**                                                        
| Поле       | Тип                      |
| :--------: | :----------------------: |
| id         | BIGSERIAL                |
| name       | VARCHAR(50)              |
| email      | VARCHAR(100)             |
| password   | VARCHAR(100)             |
| created_at | TIMESTAMP WITH TIME ZONE |
| updated_at | TIMESTAMP WITH TIME ZONE |


**Tasks**

| Поле        | Тип                      |
| :---------: | :----------------------: |
| id          | BIGSERIAL                |
| user_id     | INTEGER                  |
| title       | VARCHAR(50)              |
| description | TEXT                     |
| created_at  | TIMESTAMP WITH TIME ZONE |
| updated_at  | TIMESTAMP WITH TIME ZONE |

### 2. Миграции

C помощью goose можно выполнить миграции на создание и удаление таблиц ```Users``` и ```Tasks```

- В первую очередь, для корректной работы необходимо изменить следующие поля для взаимодействия с СУБД в Makefile:
```
PG_HOST=<host>
PG_PORT=<port>
PG_USER=<user name> 
PG_PASSWORD=<password>
PG_DATABASE=<database name>
```
*Для тестирования своего решения мной была запущена локальная сессия PostgreSQL*


- Для создания таблиц необходимо выполнить:
```
make migration-up
``` 

- При успешном выполнении команды вы увидете подобное сообщение:

```
goose -dir "/migrations" postgres "user=postgres password=test dbname=test host=localhost port=5432 sslmode=disable" up
2023/04/09 01:21:27 OK   20230408163009_.sql (55.8ms)
2023/04/09 01:21:27 goose: no migrations to run. current version: 20230408163009
```
- Для их удаления необходимо выполнить:

```
make migration-down
```
- При успешном выполнении команды вы увидете подобное сообщение:
```
goose -dir "/migrations" postgres "user=postgres password=test dbname=test host=localhost port=5432 sslmode=disable" down
2023/04/09 01:21:21 OK   20230408163009_.sql (57.39ms)
```

### 3. CRUD и работа с СУБД

Проект имеет следующую структуру 

```
homework-5
├───cmd
├───migrations
└───internal
    └───pkg
        ├───db   
        └───repository
            └───postgres
```

- cmd ― содержит main.go, который выполняет роль интерфейса для взаимодействия с СУБД
- internal/pkg/db содержит:
    - сlient.go ― структуру клиента для взаимодействия с СУБД
    - database.go ― содержит методы для выполнения базовых операций с базой данных используя библиотеку pgx, 
- internal/pkg/repository ― содержит реализацию паттерна репозиторий:
    - struct.go ― содержит структуры таблиц БД
    - repository.go ― содержит интерфейсы для взаимодействия с таблицами
    - postgres ― содержит реализацию интерфейса для взаимодействия с таблицами PostgreSQL
- migrations ― директория с миграциями

Для корректной работы необходимо изменить следующие поля для взаимодействия с СУБД в internal/pkg/db/сlient.go : 
```
  host     = <host>
  port     = <port>
  user     = <user name>
  password = <password>
  dbname   = <database name>
```

Далее, можно выполнить запуск проекта с помощью команды
```
go run cmd/main.go
```

После успешного запуска вы увидете набор доступных команд (CRUD):
```
Available commands:
        'help' to print list of commands
        'create-user' to create new user
        'create-task' to create new task
        'get-user-by-id' to get user by id
        'get-task-by-id' to get task by id
        'get-all-users' to get all users
        'get-all-tasks' to get all tasks for specific user
        'update-user' to update user
        'update-task' to update task for specific user
        'delete-user' to delete user by id
        'delete-task' to delete task by id
        'exit'
```

### 4. Кэширование

На занятиях мы познакомились со следующими инструментами для кэширования:
- Redis ― мощный инструмент для кэширования данных, надежный, поддерживает множество типов для хранения.
- Memcache ― более простой инструмент, но в тоже время более производительный. 

Поскольку в проекте реализованы простые методы для работы с данными, то операций чтения/записи и хранения данных в виде ключ-значение Memcache'a будет вполне достаточно, а использование Redis'a было бы излишне.

План по подключению к кэшам:

    1. Для работы можно использовать библиотеку gomemcache
    2. Создать новый интерфейс в internal/pkg/repository/repository.go для взаимодействия с memcache
    3. Создать пакет memcache внутри internal/pkg/repository, в котором будет находится реализация интерфейса из internal/pkg/repository/repository.go, а именно логика кэширования

Кэши должны заполняться следующим образом:
- при создании нового пользователя или задачи, записать значение в соответсвующий кэш с помощью Set

Кэши должны инвалидироваться следующим образом:
- при удалении записи о пользователе или задаче, удалить их из кэша
- после изменения данных в базе данных, обновить записи в кэше

### 5. 💎 SQL запрос для заполнения таблиц данными

В папке sql находится файл script.sql с запросом на заполнение БД данными
