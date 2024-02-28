# Lightweight Prototype of a RESTful API
## Go, Echo and PostgreSQL. 

This prototype of a RESTful API is designed for demonstration purposes in a bachelor's thesis. 

### Manual

1. `cd`into project root
2. Create Module: `go mod init project-name`
3. Add Dependencies: `go mod tidy` or manually with `go get ...`
4. Start Service locally: `go run project-name.go`
5. Send Requests on `localhost:1323` - the routing is defined in `main.go`!

> For proper functionality, the service requires a connection to a PostgreSQL database with the following table structure. The database credentials must be provided in line 18 of the main.go file.

```
CREATE TABLE offerings (
    id SERIAL PRIMARY KEY,
    item VARCHAR(255) NOT NULL,
    preis VARCHAR(255) NOT NULL,
    kontakt VARCHAR(255) NOT NULL
);
```

### API Endpoints

The service provides a suite of endpoints to manage offerings:

#### 1. Get All Offerings

- **Method:** GET
- **URL:** `http://localhost:1323/offerings`
- Retrieves all offerings in a JSON array format.

#### 2. Get a Specific Offering

- **Method:** GET
- **URL:** `http://localhost:1323/offerings/:id`
- Fetches a specific offering identified by its `id`.

#### 3. Add an Offering

- **Method:** POST
- **URL:** `http://localhost:1323/offerings`
- Creates a new offering with the provided JSON object.

#### 4. Update an Offering

- **Method:** PUT
- **URL:** `http://localhost:1323/offerings/:id`
- Updates or creates an offering identified by its `id`.

#### 5. Delete an Offering

- **Method:** DELETE
- **URL:** `http://localhost:1323/offerings/:id`
- Removes a specific offering by its `id`.

### Testing the API

Use Postman or `curl` commands to interact with the API endpoints:

- **Get All Offerings**
  ```bash
  curl -X GET http://localhost:1323/offerings
  ```

- **Get a Specific Offering**
  ```bash
  curl -X GET http://localhost:1323/offerings/42
  ```

- **Add an Offering**
  ```bash
  curl -X POST http://localhost:1323/offerings -H "Content-Type: application/json" -d '{"item":"Waschmaschine","preis":"9000","kontakt":"waschmaschine@schleuderpreis.com"}'
  ```

- **Update an Offering**
  ```bash
  curl -X PUT http://localhost:1323/offerings/42 -H "Content-Type: application/json" -d '{"preis":"799"}'
  ```

- **Delete an Offering**
  ```bash
  curl -X DELETE http://localhost:1323/offerings/42
  ```
