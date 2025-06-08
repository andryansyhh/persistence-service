# persistence-service

This service listens to messages from Kafka (`message-topic`) and persists them to PostgreSQL.

---

## ⚙️ Tech Stack

- Golang
- Kafka-Go
- GORM + PostgreSQL

---

## 📂 Environment Variables

Create `.env` file:

```env
BROKERS=localhost:9092

PORT=9094

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=messages_db


## how to run 
go run main.go
