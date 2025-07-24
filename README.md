# Aggregation Service — REST API для агрегации онлайн-подписок

## Описание

Это приложение реализует REST API для сбора и управления данными о подписках пользователей на онлайн-сервисы. Используется Go, PostgreSQL, всё упаковано в Docker-контейнеры.

**Возможности:**
- CRUD-операции с подписками
- Подсчёт суммарной стоимости подписок за период
- Фильтрация по пользователю и названию сервиса
- Логирование
- Документация Swagger/OpenAPI

API реализован в формате **REST + JSON**.

---

## Как запустить

### 1. Через Docker Compose (рекомендуется)

```sh
docker-compose up --build
```
- Приложение будет доступно на [http://localhost:9000](http://localhost:9000)
- Swagger-документация: [http://localhost:9000/swagger/index.html](http://localhost:9000/swagger/index.html)
- База данных PostgreSQL поднимается автоматически (порт 5432, пользователь `admin`, пароль `123`, база `service`)

### 2. Локально (Go 1.20+)

```sh

cd app
make build
make run
```

> Не забудьте поднять PostgreSQL и прописать параметры подключения в `cfg.yaml` (или используйте docker-compose).

### 3. Генерация Swagger

```sh

cd app
make swag
```

### 4. Тесты

```sh

cd app
make test
```

---

## Структура проекта

```
app/internal/handlers/      — HTTP-обработчики (ручки подписок)
app/internal/usecases/      — бизнес-логика (usecase-ы по подпискам)
app/internal/repo/          — репозитории для доступа к данным (PostgreSQL)
app/internal/model/         — структуры данных (модели)
app/internal/dto/           — DTO для запросов/ответов
app/internal/db/            — инициализация БД
app/internal/config/        — конфигурация
app/internal/logger/        — логирование
app/internal/apperr/        — централизованная обработка ошибок
app/internal/router/        — маршрутизация
app/docs/                   — Swagger/OpenAPI спецификация
```

---

## Реализовано

- CRUD-операции с подписками (создание, получение, обновление, удаление, полная замена)
- Подсчёт суммарной стоимости подписок с фильтрацией по пользователю и названию
- Логирование всех операций
- Конфигурация через YAML (`cfg.yaml`)
- Миграция и инициализация БД через docker-compose и init.sql
- Swagger-документация (автоматически генерируется)
- Пример Postman-коллекции для ручного тестирования API (`service.postman_collection.json`)
- Запуск через docker-compose

---

## Документация

- Swagger UI: [http://localhost:9000/swagger/index.html](http://localhost:9000/swagger/index.html)
- OpenAPI спецификация: `app/docs/swagger.yaml`
- Примеры запросов — в Postman-коллекции `service.postman_collection.json`

---

## Контейнеризация

- Всё приложение и БД запускаются одной командой через Docker Compose
```sh

docker compose up --build
```
- Конфиг и миграции БД монтируются автоматически

---

## Требования

- Go 1.20+
- Docker, Docker Compose

---

## Покрытие тестами

- Покрыты unit-тестами usecase-слоя:
  - Создание подписки
  - Получение всех подписок и по id
  - Подсчёт суммарной стоимости
  - Обновление и полная замена подписки
  - Обработка ошибок и валидация входных данных
- Интеграционных тестов нет (можно добавить для проверки работы с реальной БД и ручек)

---

## Что можно улучшить

1. **Добавить интеграционные тесты**  
   Проверка работы с реальной БД и HTTP-ручками (например, через Testcontainers).
2. **Внедрить CI/CD**  
   Автоматизация тестирования и деплоя (Github Actions, Gitlab CI и др.).
3. **Покрытие кода**  
   Проверка покрытия (make test-cover).

---

## Ссылки

- [Swagger UI (локально)](http://localhost:9000/swagger/index.html)
- [Postman-коллекция](service.postman_collection.json)

---

