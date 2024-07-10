# Login Endpoint
### Description
The POST /auth/login endpoint is responsible for user authentication and issuing a JWT token.

## POST /accounts/login

## Request
The request must contain a JSON object with the following fields:

```json
{
"email": "string",
"password": "string"
}
```

## Responses
In the case of successful authentication, a JSON object with the JWT token is returned:

```json
{
"token": "string"
}
```
#### 200 OK: The user has been successfully authenticated. A JWT token is returned.
```json
{
"token": "jwt_token_here"
}
```

#### 400 Bad Request: The input data is incorrect. The response body contains an error message.
```json
{
"error": "Invalid request data"
}
```

#### 401 Unauthorized: The user's credentials are invalid.

```json
{
"error": "Invalid credentials"
}
```

#### 500 Internal Server Error: An error occurred on the server while processing the request.
```json
{
"error": "Could not login"
}
```