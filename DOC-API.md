# Yuk Ngepet 
### Platform Edukasi Ngepet Digital

Yuk ngepet adalah platform website yang menyediakan kursus edukasi agar bisa di kira ngepet oleh tetangga!

ini adalah dokumentasi API dan Task List
Baik Backend dan Frontend harus membuat berdasarkan Dokumentasi ini
bila ada update Endpoit Detail, ber arti ada penambahan task

&bnsp;

## Daftar Endpoint 
### users
- `GET /users`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:user_id`
- `PUT /users/:user_id`
- `DELETE /users/:user_id`

## userDetails
- `GET /user_details`
- `GET /user_details/:user_id`
- `POST /user_details`
- `PUT /user_details/:detail_id`

## userTransaction
- `GET /users/transactions`
- `GET /users/transactions/:user_id`
- `POST /users/transactions/add`

## Sub Type 
- `GET /sub-types/`
- `POST /sub-types/`
- `PUT /sub-type/`

## categories
- `GET /categories`
- `POST /categories`
- `PUT /categories/:category_id`

## masteries
- `GET /masteries`
- `POST /masteries`
- `PUT /masteries/:mastery_id`

## course 
- `GET /courses/all`
- `GET /courses/basic`
- `GET /courses/premium`
- `GET /courses/:category_id`
- `GET /courses/:mastery_id`
- `GET /courses/:sub_id`
- `GET /courses/`
- `POST /courses/`
- `PUT /courses/courses_id`
- `DELETE /courses/courses_id`

## RESTful endpoint users

### GET /users

> Get All Users

_Request Header_
```json
{
    "Authorization": "<your authorization>"
}
```
_Request Body_
```
not needed
```

_Response (200)_
```json
{
    "meta": {
      "message" : "success get all user",
      "code" : 200,
      "status" : "success"
    },
    "data": [
        {
            "id" : 1,
            "username": "radika90",
            "email": "radikaye@gmail.com"
        },
        {
            "id" : 2,
            "username": "panduwi",
            "email": "panduwii@gmail.com"
        }
    ]
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "input data error",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /users/register

> Create new user

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "username" : "<username to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>",
  "role": "<role to get insert into>"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success create new user",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : 
      {
        "id" : <given id by system>,
        "username" : "<posted username>",
        "email" : "<posted email>",
        "role" : "<poster role>"
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### POST /users/login

> Compare data login on database with data inputed

_Request Header_
```
not needed
```
_Request Body_
```json
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```


_Response (200)_
```json
{
  "meta" : {
      "message" : "success login user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : 
      {
        "token" : "<your access token>"
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "input data error",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### GET /users/:user_id

> Get user by user ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id": 1,
        "username": "radika90",
        "email": "radikaye@gmail.com",
        "role": "admin",
        "user_detail": {
            "id": 1,
            "first_name": "Radika",
            "last_name": "Yudhistira",
            "sub_type": "Premium",
            "no_handphone": 6283213231232,
            "gender": "male",
            "address": "Tangerang",
            "Period": 360,
            "user_id": 1,
            "created_at": "2021-05-06T15:21:02+07:00",
            "updated_at": "2021-05-06T15:21:02+07:00"
        },
        "UserTransaction": [
            {
                "id":1,
                "sub_type": "3",
                "date": "2021-05-013T15:21:02+07:00",
                "status":"succes",
                "transfer_fact":"<link penyimpanan bukti transfer online>",
                "user_id":1
            }
        ]
    }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "error bad request user ID",
      "code" : 400,
      "status" : "error"
  }, 
  "data" : 
      {
        "errors" : "user id <id? not found"
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### PUT /users/:user_id

> Update user by User iD
just user who have same id with param can access endpoint

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "username": "radika90",
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success update user by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 1,
        "username" : "Radika90",
        "email" : "radikaye@mail.com",
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### DELETE /users/:user_id

> Delete user by ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success delete user by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : "",
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## RESTfull API User Detail

### GET /user_details

> Get All User Details
just user with role admin can access endpoint
Authorization will get id, and role

_Request Header_
```json
{
    "Authorization": "<your authorization>"
}
```
 

_Request Body_
```
not needed
```

_Response (200)_
```json
{
    "meta": {
      "message" : "success get all user details",
      "code" : 200,
      "status" : "success"
    },
    "data": [
        {
            "id": 1,
            "first_name": "Radika",
            "last_name": "Yudhistira",
            "sub_type": "Premium",
            "no_handphone": 6283213231232,
            "gender": "male",
            "address": "Tangerang",
            "Period": 360,
            "user_id": 1,
            "created_at": "2021-05-06T15:21:02+07:00",
            "updated_at": "2021-05-06T15:21:02+07:00"
        },
        {
            "id": 2,
            "first_name": "Pandu",
            "last_name": "Wilaaaa",
            "sub_type": "Basic",
            "no_handphone": 6283213231232,
            "gender": "male",
            "address": "unknown",
            "Period": 30,
            "user_id": 2,
            "created_at": "2021-05-06T15:21:02+07:00",
            "updated_at": "2021-05-06T15:21:02+07:00"
        }
    ]
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "input data error",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## GET /user_details/:user_id

> Get User Detail with user id
user with role admin can use this endpoint
user with role others just can acces if have same user_id with the user

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{   "meta": {
        "message" : "success get user details with user id",
        "code" : 200,
        "status" : "success"
    },
    "data":  {
        "id": 1,
        "first_name": "Radika",
        "last_name": "Yudhistira",
        "sub_type": "Premium",            
        "no_handphone": 6283213231232,
        "gender": "male",
        "address": "Tangerang",
        "Period": 360,
        "user_id": 1,
        "created_at": "2021-05-06T15:21:02+07:00",
        "updated_at": "2021-05-06T15:21:02+07:00"
    }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "input data error",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## POST /user_details

> Create user details

_Request Header_
```json
{
    "Authorization": "<your authorization>"
}
```

_Request Body_
```json
{
    "first_name": "Pandu",
    "last_name": "Willaa",
    "sub_type": "Free",
    "no_handphone": 6283213231232,
    "gender": "male",
    "address": "Unknown",
    "Period": 0,
    "user_id": 2,
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "input data error",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### PUT /user_details/:detail_id

> Update user detail by detail iD
just user who have same id with param can access endpoint

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "first_name": "Pandu",
    "last_name": "Willaa",
    "sub_type": "Free",
    "no_handphone": 6283213231232,
    "gender": "male",
    "address": "Unknown",
    "Period": 0,
    "user_id": 2,
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success update user detail by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 2,
        "first_name": "Pandu",
        "last_name": "Willaa",
        "sub_type": "Free",
        "no_handphone": 6283213231232,
        "gender": "male",
        "address": "Unknown",
        "Period": 0,
        "user_id": 2,
        "created_at": "2021-05-06T15:21:02+07:00",
        "updated_at": "2021-05-06T15:21:02+07:00"
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## RESTful endpoint userTransaction

### GET /user-transactions

> Get all user transaction data
just admin can access this endpoint


