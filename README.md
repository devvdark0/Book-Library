# 🚀 Учебный Go Сервис (Book-Library)

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)

Учебный REST API сервер, написанный на Go с использованием только стандартной библиотеки. Проект демонстрирует создание веб-сервиса с подключением к PostgreSQL, контейнеризацию с помощью Docker и оркестрацию через Docker Compose.

## 📋 О проекте

Этот проект был создан в **учебных целях** для освоения следующих технологий и концепций:
- Разработка backend-приложений на Go без внешних фреймворков
- Работа с реляционными базами данных (PostgreSQL)
- Миграция базы данных с помощью библеотеки `pressly/goose`
- Контейнеризация приложений с помощью Docker
- Настройка многоконтейнерных сред с Docker Compose
- Создание RESTful API

## 🛠 Технологический стек

### Backend
- **Go** (чистая стандартная библиотека `net/http`, `database/sql`)
- **PostgreSQL** - реляционная база данных

### Инфраструктура
- **Docker** - контейнеризация приложения
- **Docker Compose** - оркестрация контейнеров

## 🚀 Быстрый старт

### Предварительные требования

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Запуск проекта

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/devvdark0/Book-library
   cd Book-Library
   
   1. Запустите приложение с помощью Docker Compose:
   docker-compose up -d
   
   2. Проверьте статус контейнеров:
   docker-compose ps

   3. Приложение будет доступно по адресу:
   http://localhost:80
   
   Остановка проекта
   docker-compose down🔧 Разработка
