# WB Level 0 — Order Service

## Установка и запуск
- Клонировать репозиторий:
```
  git clone https://github.com/M-kos/wb_level0.git
  cd wb_level0
```

- Настроить переменные окружения:
  - Создай файл .env (или .env.local для локального запуска). Пример есть в файле .env.template

- Собрать и запустить сервисы:
  - Полный запуск (всё в Docker)
    ```
    make compose-up
    ```
    После старта:

    Backend: http://localhost:8083

    Frontend: http://localhost:3001

    Kafka UI: http://localhost:8091    

  - Локальная разработка (Go — локально, всё остальное в Docker)
    ```
    make local-dev
    ```
    Эта команда:

    Поднимет контейнеры postgres, kafka, kafka-ui, migrations-up

    Запустит order_service локально через go run ./cmd/order-app/main.go

- Остановка всех контейнеров
  ```
  make compose-down
  ```

## Основные команды Makefile

| Команда | Описание |
|----------|-----------|
| `make build` | Сборка Go-приложения (`order_service`) |
| `make build-producer` | Сборка Kafka producer |
| `make tidy` | Очистка и обновление зависимостей (`go mod tidy`) |
| `make lint` | Запуск линтера `golangci-lint` |
| `make compose-up` | Поднять **все контейнеры** из `docker-compose.yaml` |
| `make compose-up-local` | Поднять **только инфраструктуру** (Postgres, Kafka, UI) без `order_service` |
| `make local-dev` | Локальная разработка: поднимает Docker-инфраструктуру и запускает Go-сервис локально |
| `make compose-down` | Остановка и очистка всех контейнеров и томов |

## Проверка Kafka

- **Kafka брокер:** `localhost:29092`
- **Kafka UI:** [http://localhost:8091](http://localhost:8091)
