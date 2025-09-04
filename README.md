# Tutturu

Учебный проект Todo на Go. Название **Tutturu** 

---

## Описание

Приложение реализует REST API для управления задачами (CRUD).

### Основной функционал:
- Создание задачи (POST /tasks)
- Получение всех задач (GET /tasks)
- Обновление задачи по ID (PATCH /tasks/{id})
- Удаление задачи по ID (DELETE /tasks/{id})
- Регистрация нового пользователя (POST /users)
- Получение всех пользователей (GET /users)
- Обновление пользователя по ID (PATCH /users/{id})
- Удаление пользователя по ID (DELETE /users/{id})
---

##  Технологии

- [Go 1.21+](https://go.dev/)
- [Echo](https://echo.labstack.com/) — HTTP фреймворк
- [GORM](https://gorm.io/) — ORM
- [PostgreSQL](https://www.postgresql.org/)
- [OpenAPI 3.0](https://swagger.io/specification/) + [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [Docker](https://www.docker.com/)
- [golang-migrate](https://github.com/golang-migrate/migrate) — миграции базы данных
- [GolangCI-Lint](https://golangci-lint.run/) — линтинг кода


---

## 🏗️ Структура проекта

```
Tutturu/
├── cmd/server/
│   └── main.go                 # Entry point
│
├── internal/
│   ├── handlers/               # HTTP handlers
│   │   ├── tasks.go           # Task handlers
│   │   └── userHandlers.go    # User handlers
│   ├── models/                # Data models
│   │   ├── task.go           # Task model
│   │   └── user.go           # User model
│   ├── repository/            # Database layer (Tasks)
│   │   └── task_repository.go
│   ├── service/               # Business logic (Tasks)
│   │   └── task_service.go
│   ├── user/                  # User module
│   │   ├── repository/        # User repository
│   │   │   └── user_repository.go
│   │   └── service/          # User service
│   │       └── user_service.go
│   └── web/                  # Generated code
│       ├── tasks/            # Tasks API
│       │   └── api.gen.go
│       └── users/           # Users API
│           └── api.gen.go
│
├── pkg/                      # Utility packages
│   ├── config/              # Configuration
│   │   └── config.go
│   └── database/            # Database
│       └── db.go
│
├── migrations/              # Database migrations
│   ├── 000001_create_tasks_table.up.sql
│   ├── 000001_create_tasks_table.down.sql
│   ├── 000002_create_users_table.up.sql
│   └── 000002_create_users_table.down.sql
│
├── openapi/                # OpenAPI specs
│   ├── .openapi           # Codegen config
│   └── openapi.yaml       # API documentation
│
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
---
# Команды Makefile

# Генерация кода для задач
- make gen

# Генерация кода для пользователей 
- make gen-users

# Запуск линтера
- make lint

# Выполнение миграций
- make migrate

# Создание новой миграции
- make migrate-new NAME=create_table_name

# Запуск всех генераторов 
- make gen-all

---

# Статус проекта
- CRUD операции для задач

- RESTful API архитектура

- Работа с PostgreSQL через GORM

- Автоматическая генерация кода из OpenAPI

- Миграции базы данных

- Модульная структура проекта

- CRUD операции для пользователей 

- Разделение на tasks и users модули 

