# Studio Classes API

## Overview
The **Studio Classes API** is a RESTful service built using Go and the Gin framework. It allows users to create and manage studio classes, book classes for members, and retrieve class and booking details.

## Setup

### Prerequisites
Ensure you have Go installed on your system.

### Installation
1. Install project dependencies:
   ```sh
   go mod tidy
   ```

2. Build and run the application:
   ```sh
   make build
   make run
   ```

## API Endpoints

### 1. Create a Class
**Endpoint:** `POST /classes`

Creates a new class with specified details.

#### Request Body:
```json
{
  "className": "string",
  "startDate": "YYYY-MM-DD",
  "endDate": "YYYY-MM-DD",
  "capacity": "int"
}
```

#### Response:
```json
{
  "data": {
      "class": "className",
      "class_id": "uuid"
  },
  "message": "Class created successfully",
  "success": true
}
```

### 2. Book a Class
**Endpoint:** `POST /bookings`

Books a class for a member on a specific date.

#### Request Body:
```json
{
  "class_id": "uuid",
  "member_name": "name",
  "date": "YYYY-MM-DD"
}
```

#### Response:
```json
{
  "data": {
      "class": "className",
      "date": "2024-12-02",
      "member_name": "member_name"
  },
  "message": "Booking successful",
  "success": true
}
```

### 3. Get List of Classes
**Endpoint:** `GET /classes`

Retrieves a list of all created classes.

#### Response:
```json
{
  "data": [
    {
      "id": "uuid",
      "class": "className",
      "start_date": "2024-12-01T00:00:00Z",
      "end_date": "2024-12-20T00:00:00Z",
      "capacity": 10
    }
  ],
  "message": "Classes fetched",
  "success": true
}
```

### 4. Get List of Bookings
**Endpoint:** `GET /bookings`

Retrieves a list of all class bookings.

#### Response:
```json
{
  "data": [
    {
      "booking_id": "uuid",
      "class": "className",
      "date": "2024-12-19T00:00:00Z",
      "member_name": "name"
    }
  ],
  "message": "Bookings fetched",
  "success": true
}
```

## Development Approach
- Developed a **RESTful API** using **Go** and **Gin**.
- Implemented a structured architecture with handlers and models for maintainability.
- Used **in-memory storage** for data, with plans for future database integration.

## Testing
You can use **Postman**, **cURL**, or any API testing tool to test the endpoints.

---

### Author
**Santhosh Muthoju**

For any issues or contributions, please submit a pull request or report an issue.

