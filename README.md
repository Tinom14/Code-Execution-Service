# Сервис для выполнения программного кода 

# Описание
Сервис предоставляет REST API для выполнения программного кода на языках C, C++ и Python в безопасной изолированной среде. Каждое выполнение происходит в отдельном Docker-контейнере.

# Возможности
+ Выполнение кода на C, C++ и Python

+ Изолированная среда выполнения (Docker)

+ Асинхронная обработка задач

+ Аутентификация пользователей

+ Мониторинг выполнения задач

+ Логирование выполнения

# Архитектура

Сервис состоит из нескольких компонентов:

1. HTTP_service: Обрабатывает запросы пользователей и управление задачами

2. Processor: Выполняет задачи в изолированных Docker-контейнерах

3. RabbitMQ: Брокер сообщений для распределения задач

4. PostgreSQL: Хранит данные пользователей и задач

5. Redis: Хранит информацию о сессиях

6. Prometheus + Grafana: Мониторинг и визуализация

# API

+ POST /register - Регистрация нового пользователя

Request example:
Content-Type: application/json
```json
{"username": "username",
"password": "password"}
```

+ POST /login - Вход и получение токена

Request example:
Content-Type: application/json
```json
{"username": "username",
"password": "password"}
```

Response example:
```json
{"token":"svIgNDdaFoXEEbjzCIsKOVQdqyoV3IyqCHZKeiVWVHA="}
```


+ POST /task - Отправка кода на выполнение

Request example:
Authorization: Bearer {token}
Content-Type: application/json
```json
{"translator": "python3",
"code": "print('Hello, stdout world!')"}
```

Response example:
```json
{"task_id":"2fc1e17f-575a-4afa-ba30-9a1781f5fdb4"}
```

+ GET /status/2fc1e17f-575a-4afa-ba30-9a1781f5fdb4 - Получение статуса задачи
Authorization: Bearer {token}
Content-Type: application/json

Response example:
```json
{"status":"ready"}
```

+ GET /result/2fc1e17f-575a-4afa-ba30-9a1781f5fdb4  - Получение результатов выполнения
Authorization: Bearer {token}
Content-Type: application/json

Response example:
```json
{"result":"Hello, stdout world!"}
```





