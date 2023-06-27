# **API Documentation**

## Get All User

Endpoint: GET /users

Response Success

```json
{
  "message": "success",
  "data": [
      {
          "id": 1,
          "name": "Dedit Hery Suprastyo",
          "email": "deditherys@gmail.com",
          "password": "sangat rahasia"
      },
      {
          "id": 2,
          "name": "Dedit Baru",
          "email": "deditherys_baru@gmail.com",
          "password": "rahasia_baru"
      }
  ]
}
```

## Get User By ID

Endpoint: GET /users/1

Response Success

```json
{
  "message": "success",
  "data": [
      {
          "id": 1,
          "name": "Dedit Hery Suprastyo",
          "email": "deditherys@gmail.com",
          "password": "sangat rahasia"
      }
  ]
}
```

## Add New User

Endpoint: POST /users

Request Body:

```json
{
    "name": "Dedit",
    "email": "deditherys@gmail.com",
    "password": "rahasia"
}
```

Response Success:

```json
{
  "message": "success",
  "data": [
      {
          "id": 1,
          "name": "Dedit",
          "email": "deditherys@gmail.com",
          "password": "rahasia"
      }
  ]
}
```

## Edit User

Endpoint: PUT /users/1

Request Body:

```json
{
    "name": "Dedit hery Suprastyo",
    "email": "deditherys@gmail.com",
    "password": "sangat rahasia"
}
```

Response Success:

```json
{
  "message": "success",
  "data": [
      {
          "id": 1,
          "name": "Dedit Hery Suprastyo",
          "email": "deditherys@gmail.com",
          "password": "sangat rahasia"
      }
  ]
}
```

## Delete User

Endpoint: DELETE /users/1

Response Success

```json
{
  "message": "success",
  "data": null
}
```
