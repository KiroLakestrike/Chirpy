I'll create a comprehensive API documentation file for you!

```markdown
# Chirpy API Documentation

Chirpy is a social media API that allows users to post short messages (chirps), manage their accounts, and interact with the platform.

## Base URL
```

http://localhost:8080
```
## Table of Contents

- [Authentication](#authentication)
- [Users](#users)
  - [Create User](#create-user)
  - [Login](#login)
  - [Update User](#update-user)
  - [Refresh Token](#refresh-token)
  - [Revoke Token](#revoke-token)
  - [Upgrade to Chirpy Red](#upgrade-to-chirpy-red)
- [Chirps](#chirps)
  - [Create Chirp](#create-chirp)
  - [Get All Chirps](#get-all-chirps)
  - [Get Single Chirp](#get-single-chirp)
  - [Delete Chirp](#delete-chirp)

---

## Authentication

The Chirpy API uses JWT (JSON Web Tokens) for authentication. There are two types of tokens:

- **Access Token**: Short-lived token used for authenticating API requests (1 hour expiry)
- **Refresh Token**: Long-lived token used to obtain new access tokens (60 days expiry)

### How to Authenticate

Include the access token in the `Authorization` header:
```

Authorization: Bearer <your_access_token>
```
For refresh token endpoints, use:
```

Authorization: Bearer <your_refresh_token>
```
---

## Users

### Create User

Create a new user account.

**Endpoint:** `POST /api/users`

**Authentication:** Not required

**Request Body:**
```
json
{
"email": "user@example.com",
"password": "securepassword123"
}
```
**Response:** `201 Created`
```
json
{
"id": "123e4567-e89b-12d3-a456-426614174000",
"created_at": "2024-01-15T10:30:00Z",
"updated_at": "2024-01-15T10:30:00Z",
"email": "user@example.com",
"is_chirpy_red": false
}
```
**Error Responses:**

- `400 Bad Request`: Invalid email or password format
- `500 Internal Server Error`: Server error

---

### Login

Authenticate a user and receive access and refresh tokens.

**Endpoint:** `POST /api/login`

**Authentication:** Not required

**Request Body:**
```
json
{
"email": "user@example.com",
"password": "securepassword123"
}
```
**Response:** `200 OK`

```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z",
  "email": "user@example.com",
  "is_chirpy_red": false,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0..."
}
```
```


**Error Responses:**

- `401 Unauthorized`: Invalid email or password
- `500 Internal Server Error`: Server error

---

### Update User

Update the authenticated user's email and/or password.

**Endpoint:** `PUT /api/users`

**Authentication:** Required (Access Token)

**Request Headers:**

```
Authorization: Bearer <access_token>
```


**Request Body:**

```json
{
  "email": "newemail@example.com",
  "password": "newsecurepassword123"
}
```


**Response:** `200 OK`

```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T11:00:00Z",
  "email": "newemail@example.com",
  "is_chirpy_red": false
}
```


**Error Responses:**

- `401 Unauthorized`: Missing or invalid access token
- `400 Bad Request`: Invalid request format
- `500 Internal Server Error`: Server error

---

### Refresh Token

Get a new access token using a refresh token.

**Endpoint:** `POST /api/refresh`

**Authentication:** Required (Refresh Token)

**Request Headers:**

```
Authorization: Bearer <refresh_token>
```


**Response:** `200 OK`

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```


**Error Responses:**

- `401 Unauthorized`: Missing, invalid, expired, or revoked refresh token

---

### Revoke Token

Revoke a refresh token (logout).

**Endpoint:** `POST /api/revoke`

**Authentication:** Required (Refresh Token)

**Request Headers:**

```
Authorization: Bearer <refresh_token>
```


**Response:** `204 No Content`

**Error Responses:**

- `401 Unauthorized`: Missing or invalid refresh token
- `500 Internal Server Error`: Server error

---

### Upgrade to Chirpy Red

Webhook endpoint for upgrading users to Chirpy Red (premium membership).

**Endpoint:** `POST /api/polka/webhooks`

**Authentication:** Required (API Key)

**Request Headers:**

```
Authorization: ApiKey <polka_api_key>
```


**Request Body:**

```json
{
  "event": "user.upgraded",
  "data": {
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  }
}
```


**Response:** `204 No Content`

**Error Responses:**

- `401 Unauthorized`: Missing or invalid API key
- `400 Bad Request`: Invalid request format

---

## Chirps

### Create Chirp

Post a new chirp (message).

**Endpoint:** `POST /api/chirps`

**Authentication:** Required (Access Token)

**Request Headers:**

```
Authorization: Bearer <access_token>
```


**Request Body:**

```json
{
  "body": "This is my first chirp!"
}
```


**Validation Rules:**

- Body must not be empty
- Body must be 140 characters or less
- Profane words ("kerfuffle", "sharbert", "fornax") are automatically replaced with "****"

**Response:** `201 Created`

```json
{
  "id": "456e7890-e12b-34d5-a678-901234567890",
  "created_at": "2024-01-15T12:00:00Z",
  "updated_at": "2024-01-15T12:00:00Z",
  "body": "This is my first chirp!",
  "user_id": "123e4567-e89b-12d3-a456-426614174000"
}
```


**Error Responses:**

- `401 Unauthorized`: Missing or invalid access token
- `400 Bad Request`: Empty body, body too long, or invalid format
- `500 Internal Server Error`: Server error

---

### Get All Chirps

Retrieve all chirps or filter by author and sort order.

**Endpoint:** `GET /api/chirps`

**Authentication:** Not required

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `author_id` | UUID | No | - | Filter chirps by author (user ID) |
| `sort` | string | No | `asc` | Sort order: `asc` (oldest first) or `desc` (newest first) |

**Examples:**

```
GET /api/chirps
GET /api/chirps?sort=desc
GET /api/chirps?author_id=123e4567-e89b-12d3-a456-426614174000
GET /api/chirps?author_id=123e4567-e89b-12d3-a456-426614174000&sort=desc
```


**Response:** `200 OK`

```json
[
  {
    "id": "456e7890-e12b-34d5-a678-901234567890",
    "created_at": "2024-01-15T12:00:00Z",
    "updated_at": "2024-01-15T12:00:00Z",
    "body": "This is my first chirp!",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  },
  {
    "id": "789e0123-e45b-67d8-a901-234567890123",
    "created_at": "2024-01-15T13:00:00Z",
    "updated_at": "2024-01-15T13:00:00Z",
    "body": "Another chirp here!",
    "user_id": "123e4567-e89b-12d3-a456-426614174000"
  }
]
```


**Error Responses:**

- `400 Bad Request`: Invalid `author_id` format
- `500 Internal Server Error`: Server error

---

### Get Single Chirp

Retrieve a specific chirp by ID.

**Endpoint:** `GET /api/chirps/{chirpID}`

**Authentication:** Not required

**URL Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `chirpID` | UUID | Yes | The ID of the chirp to retrieve |

**Example:**

```
GET /api/chirps/456e7890-e12b-34d5-a678-901234567890
```


**Response:** `200 OK`

```json
{
  "id": "456e7890-e12b-34d5-a678-901234567890",
  "created_at": "2024-01-15T12:00:00Z",
  "updated_at": "2024-01-15T12:00:00Z",
  "body": "This is my first chirp!",
  "user_id": "123e4567-e89b-12d3-a456-426614174000"
}
```


**Error Responses:**

- `404 Not Found`: Chirp does not exist
- `500 Internal Server Error`: Server error

---

### Delete Chirp

Delete a chirp. Users can only delete their own chirps.

**Endpoint:** `DELETE /api/chirps/{chirpID}`

**Authentication:** Required (Access Token)

**Request Headers:**

```
Authorization: Bearer <access_token>
```


**URL Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `chirpID` | UUID | Yes | The ID of the chirp to delete |

**Example:**

```
DELETE /api/chirps/456e7890-e12b-34d5-a678-901234567890
```


**Response:** `204 No Content`

**Error Responses:**

- `401 Unauthorized`: Missing or invalid access token
- `403 Forbidden`: User doesn't own this chirp
- `404 Not Found`: Chirp does not exist
- `500 Internal Server Error`: Server error

---

## Error Response Format

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```


## Common Status Codes

- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `204 No Content`: Request successful, no content to return
- `400 Bad Request`: Invalid request format or parameters
- `401 Unauthorized`: Missing or invalid authentication
- `403 Forbidden`: Authenticated but not authorized for this action
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

---

## Rate Limiting

Currently, there are no rate limits implemented. This may change in future versions.

## Support

For issues or questions, please contact the development team.

---

**API Version:** 1.0  
**Last Updated:** 2024-01-15
```
This documentation covers all the main endpoints in your API. You can save this as `API_DOCUMENTATION.md` in your project root! 📚
```
