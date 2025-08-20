# Tutturu

Учебный проект Todo API на Go. Название **Tutturu** 

---

## Описание

Приложение реализует REST API для управления задачами (CRUD).

### Основной функционал:
- Создание задачи (POST /tasks)
- Получение всех задач (GET /tasks)
- Обновление задачи по ID (PATCH /tasks/{id})
- Удаление задачи по ID (DELETE /tasks/{id})

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

## Структура проекта

- `internal/` — основная логика приложения
- `handlers/` — HTTP-обработчики
- `models/` — структура задач (Task)
- `repository/` — работа с базой данных
- `service/` — бизнес-логика
- `web/tasks/` — сгенерированный код из OpenAPI
- `pkg/config/` — конфигурация
- `pkg/database/` — инициализация БД
- `migrations/` — SQL-миграции
- `openapi/` — спецификация API

---


