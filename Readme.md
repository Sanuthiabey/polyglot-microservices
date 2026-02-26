# Polyglot Microservices Architecture

This project demonstrates a polyglot microservices architecture using multiple programming languages integrated through an API Gateway and containerized with Docker.

## 🏗 Architecture Overview

The system consists of four services:

- **API Gateway** – Spring Boot (Port 8080)
- **Item Service** – Go (Port 8081)
- **Order Service** – Python FastAPI (Port 8082)
- **Payment Service** – Spring Boot (Port 8083)

All services communicate through an internal Docker bridge network.


---

## 🚀 Technologies Used

- Java (Spring Boot)
- Go
- Python (FastAPI)
- Docker & Docker Compose
- RESTful APIs

---

## ▶ How to Run

1. Build all services
2. Build Docker images
3. Start the system


---

## 🔗 API Endpoints (via Gateway)

Examples:

- `GET /items`
- `GET /orders`
- `GET /payments`
- `POST /items`
- `POST /orders`
- `POST /payments/process`

---

## 📦 Dockerized Deployment

Each service runs in its own container and communicates over a shared Docker bridge network.

---

## 🎯 Purpose

This project demonstrates:

- Polyglot microservices architecture
- API Gateway routing
- Containerized deployment
- Multi-language service integration

---
