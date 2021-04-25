# API Documentation

## Authentication

All API must use this authentication

Request :

- Header :
  - Authorization : "Bearer | token"

Except :

- Search Products
- List Products
- Get Products
- List Categories
- Get Categories
- Create User
- Login User

---

## Login User

<details>
  <summary>Example</summary>

Request :

- Method : POST
- Endpoint : `/api/login`
- Header :

  - Accept: application/json
  - Content-Type: application/json

- Body :

```json
{
  "email": "string",
  "password": "string"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "token": "string"
    }
  ]
}
```

</details>

## List Users

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/users`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "string, unique",
      "full_name": "string",
      "phone_number": "string",
      "email": "string",
      "password": "string",
      "gender": "string",
      "date_of_birth": "date",
      "district": "string",
      "sub_district": "string",
      "address": "string",
      "created_at": "date",
      "updated_at": "date"
    },
    {
      "id": "string, unique",
      "full_name": "string",
      "phone_number": "string",
      "email": "string",
      "password": "string",
      "gender": "string",
      "date_of_birth": "date",
      "district": "string",
      "sub_district": "string",
      "address": "string",
      "created_at": "date",
      "updated_at": "date"
    }
  ]
}
```

</details>

## Create User

<details>
  <summary>Example</summary>

Request :

- Method : POST
- Endpoint : `/api/users`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "full_name": "string",
  "phone_number": "string",
  "email": "string",
  "password": "string",
  "gender": "enum",
  "date_of_birth": "date",
  "district": "string",
  "sub_district": "string",
  "address": "string"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "full_name": "string",
    "phone_number": "string",
    "email": "string",
    "password": "string",
    "gender": "string",
    "date_of_birth": "date",
    "district": "string",
    "sub_district": "string",
    "address": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Get User

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/users/{id_user}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "full_name": "string",
    "phone_number": "string",
    "email": "string",
    "password": "string",
    "gender": "string",
    "date_of_birth": "date",
    "district": "string",
    "sub_district": "string",
    "address": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Update User

<details>
  <summary>Example</summary>

Request :

- Method : PUT
- Endpoint : `/api/users/{id_user}`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "full_name": "string",
  "phone_number": "string",
  "email": "string",
  "password": "string",
  "gender": "enum",
  "date_of_birth": "date",
  "district": "string",
  "sub_district": "string",
  "address": "string"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "full_name": "string",
    "phone_number": "string",
    "email": "string",
    "password": "string",
    "gender": "string",
    "date_of_birth": "date",
    "district": "string",
    "sub_district": "string",
    "address": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Delete User

<details>
  <summary>Example</summary>

Request :

- Method : DELETE
- Endpoint : `/api/users/{id_user}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string"
}
```

</details>

---

## List Products

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/products`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "string, unique",
      "name": "string",
      "description": "string",
      "stock": "integer",
      "price": "integer",
      "category": "string",
      "created_at": "date",
      "updated_at": "date"
    },
    {
      "id": "string, unique",
      "name": "string",
      "description": "string",
      "stock": "integer",
      "price": "integer",
      "category": "string",
      "created_at": "date",
      "updated_at": "date"
    }
  ]
}
```

</details>

## Search Products

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/products?q=`
- Header :
  - Accept: application/json
- Query Param :
  - category : string

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "string, unique",
      "name": "string",
      "description": "string",
      "stock": "integer",
      "price": "integer",
      "category": "string",
      "created_at": "date",
      "updated_at": "date"
    },
    {
      "id": "string, unique",
      "name": "string",
      "description": "string",
      "stock": "integer",
      "price": "integer",
      "category": "string",
      "created_at": "date",
      "updated_at": "date"
    }
  ]
}
```

</details>

## Create Product

<details>
  <summary>Example</summary>

Request :

- Method : POST
- Endpoint : `/api/products`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "name": "string",
  "description": "string",
  "stock": "integer",
  "price": "integer",
  "category": "string"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "name": "string",
    "description": "string",
    "stock": "integer",
    "price": "integer",
    "category": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Get Product

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/products/{id_product}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "name": "string",
    "description": "string",
    "stock": "integer",
    "price": "integer",
    "category": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Update Product

<details>
  <summary>Example</summary>

Request :

- Method : PUT
- Endpoint : `/api/products/{id_product}`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "name": "string",
  "description": "string",
  "stock": "integer",
  "price": "integer",
  "category": "string"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "name": "string",
    "description": "string",
    "stock": "integer",
    "price": "integer",
    "category": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Delete Product

<details>
  <summary>Example</summary>

Request :

- Method : DELETE
- Endpoint : `/api/products/{id_product}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string"
}
```

</details>

---

## List Categories

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/categories`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "string, unique",
      "name": "string",
      "description": "string",
      "created_at": "date",
      "updated_at": "date"
    },
    {
      "id": "string, unique",
      "name": "string",
      "description": "string",
      "created_at": "date",
      "updated_at": "date"
    }
  ]
}
```

</details>

## Create Category

<details>
  <summary>Example</summary>

Request :

- Method : POST
- Endpoint : `/api/categories`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "name": "string",
  "description": "string"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "name": "string",
    "description": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Get Category

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/categories/{id_category}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "name": "string",
    "description": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Update Category

<details>
  <summary>Example</summary>

Request :

- Method : PUT
- Endpoint : `/api/categories/{id_category}`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "name": "string",
  "description": "string"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "name": "string",
    "description": "string",
    "created_at": "date",
    "updated_at": "date"
  }
}
```

</details>

## Delete Category

<details>
  <summary>Example</summary>

Request :

- Method : DELETE
- Endpoint : `/api/categories/{id_category}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string"
}
```

</details>

---

## List Transactions

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/transactions`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "string, unique",
      "user": "string",
      "date": "date",
      "total": "integer",
      "shipping": "integer",
      "status": "string",
      "items": [
        {
          "product": "string",
          "quantity": "integer",
          "price": "integer"
        },
        {
          "product": "string",
          "quantity": "integer",
          "price": "integer"
        }
      ]
    },
    {
      "id": "string, unique",
      "user": "string",
      "date": "date",
      "total": "integer",
      "shipping": "integer",
      "status": "string",
      "items": [
        {
          "product": "string",
          "quantity": "integer",
          "price": "integer"
        },
        {
          "product": "string",
          "quantity": "integer",
          "price": "integer"
        }
      ]
    }
  ]
}
```

</details>

## Create Transaction

<details>
  <summary>Example</summary>

Request :

- Method : POST
- Endpoint : `/api/transactions`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "product": "string",
  "quantity": "integer"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "user": "string",
    "date": "date",
    "total": "integer",
    "shipping": "integer",
    "status": "string",
    "items": [
      {
        "product": "string",
        "quantity": "integer",
        "price": "integer"
      },
      {
        "product": "string",
        "quantity": "integer",
        "price": "integer"
      }
    ]
  }
}
```

</details>

## Get Transaction

<details>
  <summary>Example</summary>

Request :

- Method : GET
- Endpoint : `/api/transactions/{id_transaction}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "user": "string",
    "date": "date",
    "total": "integer",
    "shipping": "integer",
    "status": "string",
    "items": [
      {
        "product": "string",
        "quantity": "integer",
        "price": "integer"
      },
      {
        "product": "string",
        "quantity": "integer",
        "price": "integer"
      }
    ]
  }
}
```

</details>

## Update Transaction

<details>
  <summary>Example</summary>

Request :

- Method : PUT
- Endpoint : `/api/transactions/{id_transaction}`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "product": "string",
  "quantity": "integer"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "string, unique",
    "user": "string",
    "date": "date",
    "total": "integer",
    "shipping": "integer",
    "status": "string",
    "items": [
      {
        "product": "string",
        "quantity": "integer",
        "price": "integer"
      },
      {
        "product": "string",
        "quantity": "integer",
        "price": "integer"
      }
    ]
  }
}
```

</details>

## Delete Transaction

<details>
  <summary>Example</summary>

Request :

- Method : DELETE
- Endpoint : `/api/transactions/{id_transaction}`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string"
}
```

</details>

---
