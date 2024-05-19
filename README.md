# Кейс для відбору до Software Engineering School 4.0

Сервіс написаний на Go з використанням Docker для контейнеризації. Використовується база даних PostgreSQL для зберігання підписаних електроних адрес і SMTP для надсилання листів. 
Для реалізації періодичного надсилання листів раз на день використовується бібліотека [cron](https://github.com/robfig/cron).

### Можливості

- GET /api/rate для отримання поточного курсу USD у UAH
- POST /api/subscribe для додавання нової електронної адреси до розсилки
- Автоматична розсилка інформації про курс раз на день для підписників

### Інструкція із запуску

1. Створити .env файл з конфігурацією SMTP провайдера і бази даних (ввести власні дані замість лапок `<>`).

```
DATABASE_URL=<host>://<login>:<password>@db:<port>/goapi?sslmode=disable
SMTP_HOST=<example.com>
SMTP_PORT=<port>
SMTP_USERNAME=<email@example.com>
SMTP_PASSWORD=<password>
```

2. Розгорнути Docker-контейнер наступною командою:

```bash
docker-compose up --build
```

Після цього, сервіс має бути запущений на `http://localhost:8080`

### Згортання

```bash
docker-compose down --rmi all
```

### Тестування

Запуск unit-тестів здійснюється за допомогою команди:

```bash
go test ./...
```
