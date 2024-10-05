# Studio Classes API

## Setup

1. Install dependencies:

   ```sh
   go mod tidy
   ```

2. Run the application:

   ```sh
   run the command
   first:

   make build

   next:

   make run
   ```

## Endpoints

- `POST /classes`

  - Creates a new class.
  - Request body:

    ```json
    {
      "className": "string",
      "startDate": "YYYY-MM-DD",
      "endDate": "YYYY-MM-DD",
      "capacity": "int"
    }

    Response:
    {
    "data": {
        "class": "className",
        "class_id": "uuid"
    },
    "message": "Class created successfully",
    "success": true
    }

    ```

- `POST /bookings`

  - Books a class for a member.
  - Request body:

    ```json
    {
      "class_id": "uuid",
      "member_name": "name",
      "date": "YYYY-MM-DD"
    }

    Response:
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

- `GET /classes`

  - Gives the array of list of classes created.

  - Response will be like:

  ```
      {
      "data": [
      {
      "id": "uuid",
      "class": "className",
      "start_date": "2024-12-01T00:00:00Z",
      "end_date": "2024-12-20T00:00:00Z",
      "capacity": 10
      },
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

  ```

  ```

## Approach Taken

- Developed RESTful API'S using Go and Gin.
- Organized the project into handlers and models for clarity.
- Used in-memory arrays for simplicity, ready for database integration.

## Testing

- Use tools like Postman to test the API endpoints.
