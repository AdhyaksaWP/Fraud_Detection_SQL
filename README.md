# Fraud Detection + Credit Activity API

A RESTful API built with [Fiber](https://gofiber.io/) and Go, designed for managing **users**, **cards**, and **transactions**, with built-in **fraud detection** powered by a Python script. This is a inspired by the project https://github.com/maitree7/Fraud_Detection_SQL 

---

## 📚 API Endpoints

Base URL:

```
/api
```

---

### 🚹 Users

| Method | Endpoint     | Description         |
| :----: | ------------ | ------------------- |
|   GET  | `/users/`    | Retrieve all users  |
|  POST  | `/users/`    | Create a new user   |
| DELETE | `/users/:id` | Delete a user by ID |

#### 🔹 Create a User (POST `/users/`)

**Request Body**:

```json
{
  "name": "John Doe"
}
```

**Response**:

```json
{
  "id": 1,
  "name": "John Doe"
}
```

---

### 💳 Cards

| Method | Endpoint              | Description             |
| :----: | --------------------- | ----------------------- |
|   GET  | `/cards/`             | Retrieve all cards      |
|  POST  | `/cards/`             | Add a new card          |
| DELETE | `/cards/:card_number` | Delete a card by number |

#### 🔹 Add a Card (POST `/cards/`)

**Request Body**:

```json
{
  "id": 1
}
```

**Response**:

```json
{
  "id": 1,
  "card_number": "123456789876"
}
```

---

### 💸 Transactions

| Method | Endpoint            | Description                             |
| :----: | ------------------- | --------------------------------------- |
|   GET  | `/transactions/`    | Retrieve all transactions               |
|   GET  | `/transactions/:id` | Get transactions by merchant ID         |
|  POST  | `/transactions/`    | Create a transaction (with fraud check) |

#### 🔹 Create a Transaction (POST `/transactions/`)

**Request Body**:

```json
{
  "amount": 150.75,
  "card_number": "123456789876",
  "id_merchant": 5
}
```

**Response (Success)**:

```json
{
  "id": 1001,
  "amount": 150.75,
  "card_number": "123456789876",
  "id_merchant": 5
}
```

**Response (Fraud Detected)**:

```json
{
  "error": "Fraudulent transaction detected!"
}
```

---

## 🔒 Fraud Detection

When a transaction is created, the API triggers a **Python-based machine learning model** to predict whether the transaction is fraudulent:

* **Process**:

  * The amount is passed to the `predictor.py` script.
  * If the Python script returns `1`, the transaction is flagged as **fraudulent** and rejected.
  * Otherwise, the transaction is saved into the database.

---

## ⚙️ Technologies Used

* [Go Fiber](https://gofiber.io/)
* [Go SQL Package](https://pkg.go.dev/database/sql)
* Python 3.10 for fraud detection model
* Machine Learning model (details in `predictor.py`)

---

## 🛠️ Running Locally

1. Install Go, Fiber, and Python 3.
2. Make sure `predictor.py` is present in the `/controllers/` folder.
3. Configure your database connection.
4. Start the Go Fiber app:

```bash
cd src
go run main.go
```

---
