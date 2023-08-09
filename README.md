
### README - Book Management API

---

#### Introduction
This is a RESTful API for managing books using Go and PostgreSQL. 

#### Project Structure
```
library/
|-- go.mod
|-- main.go
|-- database/
|   |-- connection.go
|-- handlers/
|   |-- create_book.go
|   |-- get_books.go
|   |-- get_book.go
|   |-- update_book.go
|   |-- delete_book.go
|-- middlewares/
|   |-- logging.go
|-- tests/
|   |-- logging_test.go
|-- models/
|   |-- book.go
```

#### Setup and Run

1. **Database Setup**: Ensure you have PostgreSQL installed and running. Modify the `connection.go` file inside the `database` directory to configure the correct connection parameters for your PostgreSQL instance. For simplicity to use current implementation, set the DATABASE_URL with ```export DATABASE_URL="postgres://user:password@localhost:5432/library?sslmode=disable"``` on unix/git bash, or the corresponding command for macOS.

2. **Dependencies**: Navigate to the root directory (`library/`) and run:
    ```bash
    go mod tidy
    ```

3. **Run the API**: From the root directory (`library/`), run:
    ```bash
    go run main.go
    ```

4. The API should now be running on `http://localhost:8080`.

5. **Run logging_test**: Alternatively, From the root directory (`library/`), run:
    ```bash
    go test ./tests
    ```

---

#### API Endpoints

1. **Create a Book**
    - **Endpoint**: `POST /books`
    - **Payload**:
      ```json
      {
          "title": "Sample Title",
          "author": "Sample Author",
          "publishedyear": 2023,
          "genre": "Sample Genre",
          "summary": "Sample Summary"
      }
      ```
    - **Response**:
    - `400 Bad Request` status for invalid request payload
    - `500 Internal Server Error` status for failed book insertion
    - `200 OK` status with the created book's details, including its ID.

2. **Get All Books**
    - **Endpoint**: `GET /books`
    - **Response**:
    - `500 Internal Server Error` status for failed book retrieval
    - `200 OK` status with a list of all books' details.

3. **Get a Specific Book**
    - **Endpoint**: `GET /books/{id}`
    - **Parameters**: `id` (path parameter) - ID of the book to be retrieved.
    - **Response**:
    - `404 Not Found` status for book not found
    - `200 OK` status with the details of the requested book.

4. **Update a Book**
    - **Endpoint**: `PUT /books/{id}`
    - **Parameters**: `id` (path parameter) - ID of the book to be updated.
    - **Payload**:
      ```json
      {
          "title": "Updated Title",
          "author": "Updated Author",
          "publishedyear": 2023,
          "genre": "Updated Genre",
          "summary": "Updated Summary"
      }
      ```
    - **Response**:
    - `400 Bad Request` status for invalid request payload
    - `500 Internal Server Error` status for failed book update
    - `200 OK` status with the updated book's details.

5. **Delete a Book**
    - **Endpoint**: `DELETE /books/{id}`
    - **Parameters**: `id` (path parameter) - ID of the book to be deleted.
    - **Response**: `204 No Content` status on successful deletion.

---

#### Middlewares & Utilities

- **Logging Middleware**: This middleware logs details of every API request. Check `middlewares/logging.go` for implementation.
---
#### Improvement 1
For the sake of simplicity, the handlers just connect to the database setup in `database/connection.go`. However, for scalable solutions and better testing, this can be adjusted to have the handlers accept a database as an argument. This allows easier swapping between various real and test/mock databases.
---
#### Improvement 2
Another improvement is to deploy the entire environment in Docker and/or use an orchestration tool like Kubernetes (k8s).
