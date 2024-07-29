# 📅 Telegram Reminder Service

A simple CRUD service to send delayed messages (reminders) to Telegram. Built with Go, Gin framework, PostgreSQL, and Swagger API documentation. Deployed using Docker Compose. Logs are formatted in JSON using Zap.

## 🌟 Features

- **Create** reminders
- **Read** reminders
- **Update** reminders (if not already sent)
- **Delete** reminders (if not already sent)
- **JSON Logging** for all events
- **Swagger API Documentation**

## 🛠️ Tech Stack

- **Go**: Programming language
- **Gin**: Web framework
- **PostgreSQL**: Database
- **Swagger**: API documentation
- **Docker Compose**: Container orchestration
- **Zap**: Structured logging

## 🚀 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/) (if running locally)
- [PostgreSQL](https://www.postgresql.org/) (if running locally)

### Setup
1. **Clone the repository**:
   ```sh
   git clone https://github.com/yourusername/telegram-reminder-service.git
   cd telegram-reminder-service
   ```
2. **Create a .env file in the root directory.**
   
4. **Run with Docker Compose:**
   ```sh
   docker-compose up --build
   ```
   
5. **Access Swagger Documentation:**
   Open your browser and go to http://localhost:8080/swagger/index.html.
   
## 📝 Logging
All requests and responses are logged in JSON format using Zap for better traceability and debugging. Each log entry includes the request method, path, status code, latency, and more.

## 📚 Swagger Documentation
Swagger is used to describe the API. Once the service is running, you can access the documentation at:
```sh
http://localhost:8080/swagger/index.html
```
