# Registration Endpoint

## POST /accounts/signup

### Request

```json
{
    "email": "user1@example.com",
    "password": "password123",
    "first_name": "User",
    "last_name": "Test"
}
```

### Responses
#### 200 OK: The user has been successfully registered. The response contains the data of the new user.
#### 400 Bad Request: The input data is incorrect. The response body contains an error message.
```json
{
"error": "Invalid request data"
}
```

#### 409 Conflict: A user with the provided email already exists.
```json
{
"error": "Email already exists"
}
```

#### 500 Internal Server Error: An error occurred on the server while processing the request.
```json
{
"error": "Internal server error"
}
```