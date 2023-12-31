# Orders Manager

Решение тестового задания

## Старт сервиса

1. Клонируем репозиторий проекта:
```
git clone https://github.com/LinqCod/orders-manager.git
```

Дальше необходимо запустить само приложение. Бд сервиса обернута в докер контейнер. Для развертывания использовался `docker-compose`.

2. Для начала необходимо запустить контейнер с postgres. Для этого достаточно использовать команду `make db-up`, которая создаст и запустит контейнер с базой данных (создание и заполнение таблиц происходит автоматически при сборке контейнера):

```
make db-up
```

Для завершения работы бд следует воспользоваться командой:

```
make db-down
```

3. Далее необходимо запустить само приложение. Сделать это можно командой `make app-run`, в которой по умолчанию задается массив индексов заказов из задания (10,11,14,15). Также можно запустить приложение вручную указав этот массив в качестве аргумента `go run ./cmd/main.go 10,11,14,15`
