# API Specification
GET /api/services
GET /api/services/:id
POST /api/services
PUT /api/services/:id
DELETE /api/services/:id

## Endpoints
1) GET /api/services

Purpose: Get all services

Method: GET
Endpoint: /api/services `localhost:3000/api/services`

How to send it:
No body required.

Success response

Status: 200 OK

[
  {
    "ID": 1,
    "CreatedAt": "2026-03-24T10:00:00Z",
    "UpdatedAt": "2026-03-24T10:00:00Z",
    "DeletedAt": null,
    "name": "Haircut",
    "price": 80,
    "durationMinutes": 50,
    "description": "Standard haircut service"
  }
]
Error response

Status: 500 Internal Server Error

{
  "error": "Failed to fetch services"
}
2) GET /api/services/:id  

Purpose: Get one service by ID

Method: GET
Endpoint: /api/services/:id `localhost:3000/api/services/:id`

How to send it:
Example:

GET /api/services/1
Success response

Status: 200 OK

{
  "ID": 1,
  "CreatedAt": "2026-03-24T10:00:00Z",
  "UpdatedAt": "2026-03-24T10:00:00Z",
  "DeletedAt": null,
  "name": "Haircut",
  "price": 80,
  "durationMinutes": 50,
  "description": "Standard haircut service"
}
Error responses

Status: 404 Not Found

{
  "error": "Service not found"
}

Status: 400 Bad Request

{
  "error": "Invalid service ID"
}
3) POST /api/services

Purpose: Create a new service

Method: POST
Endpoint: /api/services `localhost:3000/api/services/`

How to send it:
JSON body:

{
  "name": "Beard Trim",
  "price": 50,
  "durationMinutes": 30,
  "description": "Clean beard shaping and lining"
}
Success response

Status: 201 Created

{
  "message": "Service created successfully",
  "service": {
    "ID": 2,
    "CreatedAt": "2026-03-24T10:05:00Z",
    "UpdatedAt": "2026-03-24T10:05:00Z",
    "DeletedAt": null,
    "name": "Beard Trim",
    "price": 50,
    "durationMinutes": 30,
    "description": "Clean beard shaping and lining"
  }
}
Error responses

Status: 400 Bad Request

{
  "error": "Name, price and durationMinutes are required"
}

Status: 500 Internal Server Error

{
  "error": "Failed to create service"
}
4) PUT /api/services/:id

Purpose: Update a service

Method: PUT
Endpoint: /api/services/:id `localhost:3000/api/services/:id`

How to send it:
Example:

PUT /api/services/2

JSON body:

{
  "name": "Premium Beard Trim",
  "price": 60,
  "durationMinutes": 35,
  "description": "Detailed beard grooming"
}
Success response

Status: 200 OK

{
  "message": "Service updated successfully",
  "service": {
    "ID": 2,
    "CreatedAt": "2026-03-24T10:05:00Z",
    "UpdatedAt": "2026-03-24T10:10:00Z",
    "DeletedAt": null,
    "name": "Premium Beard Trim",
    "price": 60,
    "durationMinutes": 35,
    "description": "Detailed beard grooming"
  }
}
Error responses

Status: 400 Bad Request

{
  "error": "Invalid service ID"
}
{
  "error": "Name, price and durationMinutes are required"
}

Status: 404 Not Found

{
  "error": "Service not found"
}

Status: 500 Internal Server Error

{
  "error": "Failed to update service"
}
5) DELETE /api/services/:id

Purpose: Delete a service

Method: DELETE
Endpoint: /api/services/:id `localhost:3000/api/services/:id`

How to send it:
Example:

DELETE /api/services/2
Success response

Status: 200 OK

{
  "message": "Service deleted successfully"
}
Error responses

Status: 400 Bad Request

{
  "error": "Invalid service ID"
}

Status: 404 Not Found

{
  "error": "Service not found"
}

Status: 500 Internal Server Error

{
  "error": "Failed to delete service"
}
### Service Code Plan

These are the files needed for the Services:

models/service.go
controllers/service_controller.go
routes/service_routes.go
routes/index.go
initializers/database.go
initializers/load_env.go
initializers/sync_database.go
main.go

### Final Thought

Use DTO/request structs

Do not bind raw DB models directly for create/update requests if you can avoid it.

Validate required fields

At least check:

name not empty
price > 0
durationMinutes > 0
Return consistent JSON

Prefer:

{
  "message": "...",
  "data": service
}

#### Licence Part
