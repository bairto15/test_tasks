Инструкция:
----------

Логин: admin
Пароль: admin

Запуск postgres
```console
docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
```

Скачивание заивисимостей и запуск front-end
```console
npm i
npm start
```
