# Polyglot Microservices Architecture

This project demonstrates a polyglot microservices architecture built using multiple programming languages, integrated through a Spring Boot API Gateway and containerized using Docker.

---

## 🏗 Architecture Overview

The system consists of four independent services:

- **API Gateway** – Spring Boot (Port 8080)
- **Item Service** – Go (Port 8081)
- **Order Service** – Python FastAPI (Port 8082)
- **Payment Service** – Node.js (Port 8083)

All services communicate over an internal Docker bridge network.

Architecture Flow:

Browser → API Gateway → Microservices (Go / Python / Node)

---

## 🚀 Technologies Used

- Java (Spring Boot)
- Go
- Python (FastAPI)
- Node.js (Express)
- Docker & Docker Compose
- RESTful APIs

---

---

## 📦 Dockerized Deployment

Each microservice runs in its own Docker container and communicates through a shared Docker bridge network defined in `docker-compose.yml`.

---

## 🎯 Purpose

This project demonstrates:

- Polyglot microservices architecture
- API Gateway routing
- Docker-based deployment
- Multi-language backend integration
- Service-to-service communication
