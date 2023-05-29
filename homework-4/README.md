# Домашнее задание
 
## Задача: 
обработать N-заказов M-воркерами с применением паттернов Pipeline и WorkerPool.
Параметры N >= 5 и M >= 2 задаются как константы в глобальной области видимости.

Структура заказа состоит из следующих полей:
- ID товара
- ID склада
- ID пункта выдачи заказа
- ID воркера, который обработал заказ
- Массив состояний, в которых находился заказ в рамках пайплайна

Структура состояния заказа состоит из следующих полей:
- Наименование: [Создан, Обработан, Выполнен]
- Время перехода в данное состояние

Каждое состояние является шагом пайплайна.
Входные данные пайплайна:
- Канал, содержащий структуры заказов с инициализированным полем ID товара

Шаг "Создан" пайплайна:
- В массив состояний добавляется новое состояние "Создан"

Шаг "Обработан" пайплайна:
- Инициализируется склад для заказа - результат взятия ID товара по модулю 2
- В массив состояний добавляется новое состояние "Обработан"

Шаг "Завершен" пайплайна:
- Инициализируется пункт выдачи заказа - результат суммы ID товара и ID склада
- В массив состояний добавляется новое состояние "Завершен"

Обработка канала с заказами осуществляется M-воркерами.
Все шаги пайплайна конкретного заказа обрабатываются в рамках одного воркера.
Перед началом выполнения пайплайна каждому заказу присваивается ID воркера, который его обрабатывает.
После того, как заказ прошел все стадии обработки, необходимо передать заказ в канал с обработанными заказами.

Основной поток приложения передает в стандартный поток вывода обработанные заказы в JSON формате структуры заказа.

💎 Использовать паттерны FanIn/FanOut для шага с обработкой заказа 

## Решение:

Проект имеет следующую структуру:

```
homework-4
├───cmd
│   └───gateway
│       ├───base
│       ├───worker_pool_pipeline
│       └───worker_pool_pipeline_fan
└───internal
    ├───config
    ├───model
    └───pkg
        ├───gateway
        │   └───steps
        │       ├───complete
        │       ├───create
        │       └───process
        └───producer
```

В директории ```cmd\gateway``` расположены:
- ```base\main.go``` — демонстрирует обработку заказов в одном потоке одним worker'ов. 

    *Для запуска используйте команду:* 
    
    ```go run .\cmd\gateway\base\main.go ```
- ```worker_pool_pipeline\worker_pool_pipeline.go``` — демонстрирует обработку заказов с использованием паттернов WorkerPool и Pipeline. 

    *Для запуска используйте команду:*
    
    ```go run .\cmd\gateway\worker_pool_pipeline\worker_pool_pipeline.go ``` 
- ```worker_pool_pipeline_fan\worker_pool_pipeline_fan.go``` — демонстрирует обработку заказов с использованием паттернов WorkerPool, Pipeline и FanIn/FanOut. 

    *Для запуска используйте команду:*
    
    ```go run .\cmd\gateway\worker_pool_pipeline_fan\worker_pool_pipeline_fan.go ``` 

В директории ```internal``` расположены:
- ```config\config.go``` — файл, в котором задается количество заказов, количество worker'ов и время задержки этапа обработки заказа в секундах. 
    + Если вы хотите изменить количество worker'ов, присвойте желаемое значение константе ```NumberOfWorkers```, по умолчанию их количество равно 3. 
    + Если вы хотите изменить количество заказов, присвойте желаемое значение константе ```NumberOfOrders```, по умолчанию их количество равно 5. 
    + Если вы хотите изменить время задержки этапа обработки заказа, присвойте желаемое значение константе ```ProcessStepDuration```, по умолчанию задержка составляет 2 секунды.

- ```model\order.go``` — отвечает за структуру заказа ```Order```, а так же содержит конвертор структуры ```Order``` в строку в JSON-формате и функцию для вывода завершенных заказов из канала ```ordersCompleted```.  
- ```pkg``` — содержит реализацию основной логики работы 
    + базового случая (метод ```Process``` структуры ```Implementation``` внутри файла ```gateway.go```)
    + WorkerPool и Pipeline паттернов (метод ```Pipeline``` структуры ```Implementation``` внутри файла ```gateway.go```)
    + WorkerPool, Pipeline и FanIn/FanOut паттернов (метод ```PipelineFan``` структуры ```Implementation``` внутри файла ```gateway.go```)
    + ```steps``` — содержит логику создания, обработки и завершения заказа

Время работы каждого из случаев при ```NumberOfOrders=5```,  ```NumberOfWorkers=3``` и ```ProcessStepDuration=2```
- base case c 1 worker'ом: 10.026794s
- WorkerPool и Pipeline: 4.010462s
- WorkerPool, Pipeline и FanIn/FanOut: 4.019380s

Время работы каждого из случаев при ```NumberOfOrders=100```, ```NumberOfWorkers=5``` и ```ProcessStepDuration=2```
- base case c 1 worker'ом: 200.747155s
- WorkerPool и Pipeline: 40.169833s
- WorkerPool, Pipeline и FanIn/FanOut: 14.098333s

При большем количестве заказов реализациая с использованием WorkerPool, Pipeline и FanIn/FanOut оказалась наиболее эффективной в сравнении с WorkerPool и Pipeline. 

Решение запускалось на компьютере со следующими характеристиками:
- Операционная система Windows 11 Pro 21H2
- Процессор Intel(R) Core(TM) i5-8300H CPU @ 2.30GHz 
- Оперативная память 16,0 ГБ
Однако это далеко не все параметры, от которых зависит время исполнения. 