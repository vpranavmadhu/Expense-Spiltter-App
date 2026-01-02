# Expense Splitter – Group Expense Management System

A full-stack web application built to simplify shared expense management.  
Expense Splitter helps users create groups, track expenses, calculate balances, and generate settlement suggestions in a clean and efficient way.

---

## Project Overview

Expense Splitter is designed for scenarios such as trips, roommates, and shared workspaces where expenses need to be tracked and split accurately.  
The application follows a layered backend architecture and is deployed using containerized services for scalability and maintainability.

---

## Key Features

### Group & User Management
- User registration and authentication
- Create and manage expense groups
- Add members to groups securely

### Expense Tracking
- Add expenses with flexible split options
- Track who paid and who owes
- Maintain expense history per group

### Balance Calculation
- Automatic calculation of balances
- Clear debtor and creditor breakdown
- Group-level financial summaries

### Settlement Suggestions
- Optimized settlement recommendations
- Reduces unnecessary transactions
- Improves clarity during settlements

---

## Security & Design

- JWT-based authentication
- Password hashing using BCrypt
- Clean separation of layers (Handler → Service → Repository)
- DTO-based request validation
- Secure cookie-based authentication flow

---

## Tech Stack

### Backend
- Golang (Gin Framework)
- GORM
- PostgreSQL

### Frontend
- React (Vite)
- Tailwind CSS
- Axios

### DevOps & Deployment
- Docker
- Docker Compose
- Nginx (Reverse Proxy)

---

## Getting Started

### Prerequisites
- Docker
- Docker Compose
- PostgreSQL (for local, non-docker runs)

---

## How to Run

### 1.Clone the Project
```bash
git clone https://github.com/vpranavmadhu/Expense-Spiltter-App.git
cd Expense-Spiltter-App
```

## 2.Set Environment Configuration

The application relies on environment variables for database access and authentication.

### Create `.env` File

Create a `.env` file in the **root directory** and add the following:

```env
DB_USER=user
DB_PASS=password
DB_NAME=name
DB_PORT=5432

JWT_SECRET=random_string
```
---

## 3.Running the Project with Docker compose

This project is fully containerized and can be started using **Docker Compose**.

```bash
docker compose up --build -d
```