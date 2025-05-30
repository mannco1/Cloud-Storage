# ЛОКАЛЬНЫЙ ЗАПУСК(БЕЗ ДОКЕРА) Auth Service

## 📌 Описание
Этот сервис отвечает за аутентификацию и управление пользователями, используя PostgreSQL для хранения данных и JWT для аутентификации.

---

## Запуск проекта

### 1️⃣ Настройка переменных окружения
Создайте файл `.env` в корневой директории и добавьте следующие параметры, пример:

```ini
POSTGRES_USER=postgres
POSTGRES_PASSWORD=MYpassword
POSTGRES_DB=auth-service
STORAGE_PATH=postgres://postgres:MYpassword@db:5432/auth-service?sslmode=disable
HTTP_SERVER_ADDRESS=0.0.0.0:8080
ACCESS_TOKEN_TTL=15m
REFRESH_TOKEN_TTL=720h
PUBLIC_KEY_PATH=public_key.pem
PRIVATE_KEY_PATH=private_key.pem
```

> **Важно:** Замените `MYpassword` на ваш реальный пароль от PostgreSQL.

---

### 2️⃣ Создание базы данных

Создайте пустую базу данных `auth-service` в PostgreSQL:
```sh
psql -U postgres -c "CREATE DATABASE \"auth-service\";"
```
> Если база уже существует, этот шаг можно пропустить.

---

### 3️⃣ Запуск миграций

Выполните миграции для создания необходимых таблиц:
```sh
make pg-migrate
```

> Убедитесь, что у вас установлен `make`, иначе выполните миграции вручную.

---

### 4️⃣ Генерация ключей для подписи JWT

Генерируем **приватный** и **публичный** ключи с помощью OpenSSL:

```sh
# Генерация приватного ключа
openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
```
```sh
# Генерация публичного ключа
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

> Эти ключи необходимы для подписи и верификации JWT-токенов.

---

### 5️⃣ Запуск сервера

Запустите сервер командой:
```sh
make server-run
```
Или в случаи "make: *** No rule to make target 'server-run'.  Stop."
```sh
make -f Makefile server-run
```
---

> После запуска сервер будет доступен по адресу `http://localhost:8080`

---

## 🛠 Технологии
- **Go** – версия 1.23.1
- **PostgreSQL** – база данных
- **JWT** – аутентификация
- **Makefile** – автоматизация задач
- **OpenSSL** – генерация ключей

---
