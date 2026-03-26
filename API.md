# API.md

# IkonKutz Backend API Documentation

Base URL:


http://localhost:3000


API Base:


http://localhost:3000/api


Health Check
http://localhost:3000/api/health


---

# Authentication

JWT is used for protected routes.

For protected routes, send this header:

```txt
Authorization: Bearer your-jwt-token


---

# 1. Health

## GET
```txt
http://localhost:3000/api/health


### Purpose
Check if the API is running.

### How to send
No body required.

### Success response
```json
{
  "message": "API is healthy"
}


---

# 2. Authentication

## 2.1 Register User

## POST
```txt
http://localhost:3000/api/auth/register


### Purpose
Create a new customer account.

### How to send
Send JSON in the request body:

```json
{
  "name": "Peter",
  "email": "peter@example.com",
  "password": "strongpassword123"
}


### Success response
```json
{
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "name": "Peter",
      "email": "peter@example.com",
      "role": "customer"
    },
    "token": "jwt-token-here"
  }
}


### Error responses
```json
{
  "error": "Invalid request body"
}


```json
{
  "error": "Name, email and password are required"
}


```json
{
  "error": "Email is already in use"
}


```json
{
  "error": "Failed to hash password"
}


```json
{
  "error": "Failed to create user"
}


```json
{
  "error": "Failed to generate token"
}


---

## 2.2 Login User

## POST
```txt
http://localhost:3000/api/auth/login


### Purpose
Log in an existing user.

### How to send
Send JSON in the request body:

```json
{
  "email": "peter@example.com",
  "password": "strongpassword123"
}


### Success response
```json
{
  "message": "Login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "Peter",
      "email": "peter@example.com",
      "role": "customer"
    },
    "token": "jwt-token-here"
  }
}


### Error responses
```json
{
  "error": "Invalid request body"
}


```json
{
  "error": "Email and password are required"
}


```json
{
  "error": "Invalid email or password"
}


```json
{
  "error": "Failed to generate token"
}


---

## 2.3 Get Current User

## GET
```txt
http://localhost:3000/api/auth/me


### Purpose
Return the currently authenticated user.

### How to send
No body required.

Header:
```txt
Authorization: Bearer your-jwt-token


### Success response
```json
{
  "message": "Current user fetched successfully",
  "data": {
    "id": 1,
    "name": "Peter",
    "email": "peter@example.com",
    "role": "customer"
  }
}


### Error responses
```json
{
  "error": "Authorization header is required"
}


```json
{
  "error": "Authorization header must be in the format: Bearer <token>"
}


```json
{
  "error": "Invalid or expired token"
}


```json
{
  "error": "Unauthorized"
}


```json
{
  "error": "User not found"
}


---

## 2.4 Logout User

## POST
```txt
http://localhost:3000/api/auth/logout


### Purpose
Logout in this v1 means the client removes the token.

### How to send
No body required.

Header:
```txt
Authorization: Bearer your-jwt-token


### Success response
```json
{
  "message": "Logout successful",
  "data": {
    "message": "Remove the token on the client side"
  }
}


### Error responses
```json
{
  "error": "Authorization header is required"
}


```json
{
  "error": "Authorization header must be in the format: Bearer <token>"
}


```json
{
  "error": "Invalid or expired token"
}


---

# 3. Services

## 3.1 Get All Services

## GET
```txt
http://localhost:3000/api/services


### Purpose
Return all services.

### How to send
No body required.

### Success response
```json
{
  "message": "Services fetched successfully",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2026-03-26T10:00:00Z",
      "UpdatedAt": "2026-03-26T10:00:00Z",
      "DeletedAt": null,
      "name": "Haircut",
      "price": 120,
      "durationMinutes": 50,
      "description": "Standard haircut"
    }
  ]
}
```

### Error responses
```json
{
  "error": "Failed to fetch services"
}
```

---

## 3.2 Get One Service

## GET
```txt
http://localhost:3000/api/services/1
```

### Purpose
Return one service by ID.

### How to send
No body required.

### Success response
```json
{
  "message": "Service fetched successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "name": "Haircut",
    "price": 120,
    "durationMinutes": 50,
    "description": "Standard haircut"
  }
}
```

### Error responses
```json
{
  "error": "Invalid service ID"
}
```

```json
{
  "error": "Service not found"
}
```

---

## 3.3 Create Service (Admin Only)

## POST
```txt
http://localhost:3000/api/services
```

### Purpose
Create a new service.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "name": "Haircut",
  "price": 120,
  "durationMinutes": 50,
  "description": "Standard haircut"
}
```

### Success response
```json
{
  "message": "Service created successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "name": "Haircut",
    "price": 120,
    "durationMinutes": 50,
    "description": "Standard haircut"
  }
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

```json
{
  "error": "Admin access required"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Name, price and durationMinutes are required"
}
```

```json
{
  "error": "Failed to create service"
}
```

---

## 3.4 Update Service (Admin Only)

## PUT
```txt
http://localhost:3000/api/services/1
```

### Purpose
Update an existing service.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "name": "Premium Haircut",
  "price": 150,
  "durationMinutes": 60,
  "description": "Premium haircut service"
}
```

### Success response
```json
{
  "message": "Service updated successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:30:00Z",
    "DeletedAt": null,
    "name": "Premium Haircut",
    "price": 150,
    "durationMinutes": 60,
    "description": "Premium haircut service"
  }
}
```

### Error responses
```json
{
  "error": "Invalid service ID"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Name, price and durationMinutes are required"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Failed to update service"
}
```

```json
{
  "error": "Admin access required"
}
```

---

## 3.5 Delete Service (Admin Only)

## DELETE
```txt
http://localhost:3000/api/services/1
```

### Purpose
Delete a service.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

### Success response
```json
{
  "message": "Service deleted successfully",
  "data": null
}
```

### Error responses
```json
{
  "error": "Invalid service ID"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Failed to delete service"
}
```

```json
{
  "error": "Admin access required"
}
```

---

# 4. Slots

## 4.1 Get All Slots

## GET
```txt
http://localhost:3000/api/slots
```

### Purpose
Return all slots.

### How to send
No body required.

Optional query examples:
```txt
http://localhost:3000/api/slots?date=2026-03-30
http://localhost:3000/api/slots?booked=false
http://localhost:3000/api/slots?locked=false
http://localhost:3000/api/slots?date=2026-03-30&booked=false&locked=false
```

### Success response
```json
{
  "message": "Slots fetched successfully",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2026-03-26T10:00:00Z",
      "UpdatedAt": "2026-03-26T10:00:00Z",
      "DeletedAt": null,
      "date": "2026-03-30",
      "time": "10:00",
      "isBooked": false,
      "isLocked": false
    }
  ]
}
```

### Error responses
```json
{
  "error": "Failed to fetch slots"
}
```

---

## 4.2 Get One Slot

## GET
```txt
http://localhost:3000/api/slots/1
```

### Purpose
Return one slot by ID.

### How to send
No body required.

### Success response
```json
{
  "message": "Slot fetched successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

---

## 4.3 Create Slot (Admin Only)

## POST
```txt
http://localhost:3000/api/slots
```

### Purpose
Create a new slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "date": "2026-03-30",
  "time": "10:00",
  "isLocked": false
}
```

### Success response
```json
{
  "message": "Slot created successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Date and time are required"
}
```

```json
{
  "error": "A slot with the same date and time already exists"
}
```

```json
{
  "error": "Failed to create slot"
}
```

```json
{
  "error": "Admin access required"
}
```

---

## 4.4 Update Slot (Admin Only)

## PUT
```txt
http://localhost:3000/api/slots/1
```

### Purpose
Update a slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "date": "2026-03-30",
  "time": "11:00",
  "isLocked": false
}
```

### Success response
```json
{
  "message": "Slot updated successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:30:00Z",
    "DeletedAt": null,
    "date": "2026-03-30",
    "time": "11:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Date and time are required"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Booked slots cannot be moved to a different date or time"
}
```

```json
{
  "error": "Another slot with the same date and time already exists"
}
```

```json
{
  "error": "Failed to update slot"
}
```

---

## 4.5 Lock Slot (Admin Only)

## PATCH
```txt
http://localhost:3000/api/slots/1/lock
```

### Purpose
Lock a slot so it cannot be booked.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Slot locked successfully",
  "data": {
    "ID": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": true
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Failed to lock slot"
}
```

---

## 4.6 Unlock Slot (Admin Only)

## PATCH
```txt
http://localhost:3000/api/slots/1/unlock
```

### Purpose
Unlock a slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Slot unlocked successfully",
  "data": {
    "ID": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Failed to unlock slot"
}
```

---

## 4.7 Delete Slot (Admin Only)

## DELETE
```txt
http://localhost:3000/api/slots/1
```

### Purpose
Delete a slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Slot deleted successfully",
  "data": null
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Booked slots cannot be deleted"
}
```

```json
{
  "error": "Failed to delete slot"
}
```

---

# 5. Appointments

## 5.1 Get My Appointments

## GET
```txt
http://localhost:3000/api/appointments/me
```

### Purpose
Return only the logged-in user's appointments.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Your appointments fetched successfully",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2026-03-26T10:00:00Z",
      "UpdatedAt": "2026-03-26T10:00:00Z",
      "DeletedAt": null,
      "userId": 1,
      "customerName": "Peter",
      "serviceId": 1,
      "serviceName": "Haircut",
      "slotId": 1,
      "date": "2026-03-30",
      "time": "10:00",
      "price": 120,
      "status": "confirmed"
    }
  ]
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

```json
{
  "error": "Failed to fetch your appointments"
}
```

---

## 5.2 Get One Appointment

## GET
```txt
http://localhost:3000/api/appointments/1
```

### Purpose
Return one appointment.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointment fetched successfully",
  "data": {
    "ID": 1,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "price": 120,
    "status": "confirmed"
  }
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only view your own appointments"
}
```

---

## 5.3 Create Appointment

## POST
```txt
http://localhost:3000/api/appointments
```

### Purpose
Create a new appointment for the logged-in user.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

JSON body:
```json
{
  "serviceId": 1,
  "date": "2026-03-30",
  "time": "10:00"
}
```

### Success response
```json
{
  "message": "Appointment created successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "price": 120,
    "status": "confirmed"
  }
}
```

### Error responses
```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "ServiceId, date and time are required"
}
```

```json
{
  "error": "User not found"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Slot not found for the given date and time"
}
```

```json
{
  "error": "Slot is locked"
}
```

```json
{
  "error": "Slot is already booked"
}
```

```json
{
  "error": "Failed to create appointment"
}
```

```json
{
  "error": "Failed to reserve slot"
}
```

---

## 5.4 Update Appointment

## PUT
```txt
http://localhost:3000/api/appointments/1
```

### Purpose
Update an existing appointment.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

JSON body:
```json
{
  "serviceId": 1,
  "date": "2026-03-30",
  "time": "11:00",
  "status": "confirmed"
}
```

### Success response
```json
{
  "message": "Appointment updated successfully",
  "data": {
    "ID": 1,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 2,
    "date": "2026-03-30",
    "time": "11:00",
    "price": 120,
    "status": "confirmed"
  }
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "ServiceId, date and time are required"
}
```

```json
{
  "error": "Status must be either confirmed or cancelled"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only update your own appointments"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Current slot linked to appointment was not found"
}
```

```json
{
  "error": "Target slot not found for the given date and time"
}
```

```json
{
  "error": "Target slot is locked"
}
```

```json
{
  "error": "Target slot is already booked"
}
```

```json
{
  "error": "Failed to release old slot"
}
```

```json
{
  "error": "Failed to reserve target slot"
}
```

```json
{
  "error": "Failed to update appointment"
}
```

---

## 5.5 Cancel Appointment

## PATCH
```txt
http://localhost:3000/api/appointments/1/cancel
```

### Purpose
Cancel an appointment and free the linked slot.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointment cancelled successfully",
  "data": {
    "ID": 1,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "price": 120,
    "status": "cancelled"
  }
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only cancel your own appointments"
}
```

```json
{
  "error": "Linked slot not found"
}
```

```json
{
  "error": "Failed to cancel appointment"
}
```

```json
{
  "error": "Failed to free slot"
}
```

---

## 5.6 Delete Appointment

## DELETE
```txt
http://localhost:3000/api/appointments/1
```

### Purpose
Delete an appointment and free the slot first.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointment deleted successfully",
  "data": null
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only delete your own appointments"
}
```

```json
{
  "error": "Failed to free slot"
}
```

```json
{
  "error": "Failed to delete appointment"
}
```

---

## 5.7 Get All Appointments (Admin Only)

## GET
```txt
http://localhost:3000/api/appointments
```

### Purpose
Return all appointments in the system.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointments fetched successfully",
  "data": [
    {
      "ID": 1,
      "userId": 1,
      "customerName": "Peter",
      "serviceId": 1,
      "serviceName": "Haircut",
      "slotId": 1,
      "date": "2026-03-30",
      "time": "10:00",
      "price": 120,
      "status": "confirmed"
    }
  ]
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

```json
{
  "error": "Admin access required"
}
```

```json
{
  "error": "Failed to fetch appointments"
}
```

---

# Common Errors

## Unauthorized
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Authorization header must be in the format: Bearer <token>"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

## Forbidden
```json
{
  "error": "Admin access required"
}
```

```json
{
  "error": "You can only view your own appointments"
}
```

```json
{
  "error": "You can only update your own appointments"
}
```

```json
{
  "error": "You can only cancel your own appointments"
}
```

```json
{
  "error": "You can only delete your own appointments"
}
```

## Validation / Booking errors
```json
{
  "error": "Slot is locked"
}
```

```json
{
  "error": "Slot is already booked"
}
```

```json
{
  "error": "A slot with the same date and time already exists"
}
```

```json
{
  "error": "Booked slots cannot be deleted"
}
```

```json
{
  "error": "Booked slots cannot be moved to a different date or time"
}
```

---

# Quick Testing Order

## 1. Check health
```txt
GET http://localhost:3000/api/health
```

## 2. Register a customer
```txt
POST http://localhost:3000/api/auth/register
```

## 3. Login
```txt
POST http://localhost:3000/api/auth/login
```

## 4. Promote one user to admin in the database

## 5. Login as admin

## 6. Create a service
```txt
POST http://localhost:3000/api/services
```

## 7. Create a slot
```txt
POST http://localhost:3000/api/slots
```

## 8. Login as customer

## 9. Create an appointment
```txt
POST http://localhost:3000/api/appointments
```

## 10. View my appointments
```txt
GET http://localhost:3000/api/appointments/me
```

---

# Endpoints

## Public
- `GET http://localhost:3000/api/health`
- `POST http://localhost:3000/api/auth/register`
- `POST http://localhost:3000/api/auth/login`
- `GET http://localhost:3000/api/services`
- `GET http://localhost:3000/api/services/:id`
- `GET http://localhost:3000/api/slots`
- `GET http://localhost:3000/api/slots/:id`

## Authenticated
- `GET http://localhost:3000/api/auth/me`
- `POST http://localhost:3000/api/auth/logout`
- `GET http://localhost:3000/api/appointments/me`
- `GET http://localhost:3000/api/appointments/:id`
- `POST http://localhost:3000/api/appointments`
- `PUT http://localhost:3000/api/appointments/:id`
- `PATCH http://localhost:3000/api/appointments/:id/cancel`
- `DELETE http://localhost:3000/api/appointments/:id`

## Admin Only
- `POST http://localhost:3000/api/services`
- `PUT http://localhost:3000/api/services/:id`
- `DELETE http://localhost:3000/api/services/:id`
- `POST http://localhost:3000/api/slots`
- `PUT http://localhost:3000/api/slots/:id`
- `PATCH http://localhost:3000/api/slots/:id/lock`
- `PATCH http://localhost:3000/api/slots/:id/unlock`
- `DELETE http://localhost:3000/api/slots/:id`
- `GET http://localhost:3000/api/appointments`
# API.md

# IkonKutz Backend API Documentation

Base URL:

```txt
http://localhost:3000
```

API Base:

```txt
http://localhost:3000/api
```

Health Check:

```txt
http://localhost:3000/api/health
```

---

# Authentication

JWT is used for protected routes.

For protected routes, send this header:

```txt
Authorization: Bearer your-jwt-token
```

---

# 1. Health

## GET
```txt
http://localhost:3000/api/health
```

### Purpose
Check if the API is running.

### How to send
No body required.

### Success response
```json
{
  "message": "API is healthy"
}
```

---

# 2. Authentication

## 2.1 Register User

## POST
```txt
http://localhost:3000/api/auth/register
```

### Purpose
Create a new customer account.

### How to send
Send JSON in the request body:

```json
{
  "name": "Peter",
  "email": "peter@example.com",
  "password": "strongpassword123"
}
```

### Success response
```json
{
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "name": "Peter",
      "email": "peter@example.com",
      "role": "customer"
    },
    "token": "jwt-token-here"
  }
}
```

### Error responses
```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Name, email and password are required"
}
```

```json
{
  "error": "Email is already in use"
}
```

```json
{
  "error": "Failed to hash password"
}
```

```json
{
  "error": "Failed to create user"
}
```

```json
{
  "error": "Failed to generate token"
}
```

---

## 2.2 Login User

## POST
```txt
http://localhost:3000/api/auth/login
```

### Purpose
Log in an existing user.

### How to send
Send JSON in the request body:

```json
{
  "email": "peter@example.com",
  "password": "strongpassword123"
}
```

### Success response
```json
{
  "message": "Login successful",
  "data": {
    "user": {
      "id": 1,
      "name": "Peter",
      "email": "peter@example.com",
      "role": "customer"
    },
    "token": "jwt-token-here"
  }
}
```

### Error responses
```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Email and password are required"
}
```

```json
{
  "error": "Invalid email or password"
}
```

```json
{
  "error": "Failed to generate token"
}
```

---

## 2.3 Get Current User

## GET
```txt
http://localhost:3000/api/auth/me
```

### Purpose
Return the currently authenticated user.

### How to send
No body required.

Header:
```txt
Authorization: Bearer your-jwt-token
```

### Success response
```json
{
  "message": "Current user fetched successfully",
  "data": {
    "id": 1,
    "name": "Peter",
    "email": "peter@example.com",
    "role": "customer"
  }
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Authorization header must be in the format: Bearer <token>"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

```json
{
  "error": "Unauthorized"
}
```

```json
{
  "error": "User not found"
}
```

---

## 2.4 Logout User

## POST
```txt
http://localhost:3000/api/auth/logout
```

### Purpose
Logout in this v1 means the client removes the token.

### How to send
No body required.

Header:
```txt
Authorization: Bearer your-jwt-token
```

### Success response
```json
{
  "message": "Logout successful",
  "data": {
    "message": "Remove the token on the client side"
  }
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Authorization header must be in the format: Bearer <token>"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

---

# 3. Services

## 3.1 Get All Services

## GET
```txt
http://localhost:3000/api/services
```

### Purpose
Return all services.

### How to send
No body required.

### Success response
```json
{
  "message": "Services fetched successfully",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2026-03-26T10:00:00Z",
      "UpdatedAt": "2026-03-26T10:00:00Z",
      "DeletedAt": null,
      "name": "Haircut",
      "price": 120,
      "durationMinutes": 50,
      "description": "Standard haircut"
    }
  ]
}
```

### Error responses
```json
{
  "error": "Failed to fetch services"
}
```

---

## 3.2 Get One Service

## GET
```txt
http://localhost:3000/api/services/1
```

### Purpose
Return one service by ID.

### How to send
No body required.

### Success response
```json
{
  "message": "Service fetched successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "name": "Haircut",
    "price": 120,
    "durationMinutes": 50,
    "description": "Standard haircut"
  }
}
```

### Error responses
```json
{
  "error": "Invalid service ID"
}
```

```json
{
  "error": "Service not found"
}
```

---

## 3.3 Create Service (Admin Only)

## POST
```txt
http://localhost:3000/api/services
```

### Purpose
Create a new service.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "name": "Haircut",
  "price": 120,
  "durationMinutes": 50,
  "description": "Standard haircut"
}
```

### Success response
```json
{
  "message": "Service created successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "name": "Haircut",
    "price": 120,
    "durationMinutes": 50,
    "description": "Standard haircut"
  }
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

```json
{
  "error": "Admin access required"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Name, price and durationMinutes are required"
}
```

```json
{
  "error": "Failed to create service"
}
```

---

## 3.4 Update Service (Admin Only)

## PUT
```txt
http://localhost:3000/api/services/1
```

### Purpose
Update an existing service.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "name": "Premium Haircut",
  "price": 150,
  "durationMinutes": 60,
  "description": "Premium haircut service"
}
```

### Success response
```json
{
  "message": "Service updated successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:30:00Z",
    "DeletedAt": null,
    "name": "Premium Haircut",
    "price": 150,
    "durationMinutes": 60,
    "description": "Premium haircut service"
  }
}
```

### Error responses
```json
{
  "error": "Invalid service ID"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Name, price and durationMinutes are required"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Failed to update service"
}
```

```json
{
  "error": "Admin access required"
}
```

---

## 3.5 Delete Service (Admin Only)

## DELETE
```txt
http://localhost:3000/api/services/1
```

### Purpose
Delete a service.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

### Success response
```json
{
  "message": "Service deleted successfully",
  "data": null
}
```

### Error responses
```json
{
  "error": "Invalid service ID"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Failed to delete service"
}
```

```json
{
  "error": "Admin access required"
}
```

---

# 4. Slots

## 4.1 Get All Slots

## GET
```txt
http://localhost:3000/api/slots
```

### Purpose
Return all slots.

### How to send
No body required.

Optional query examples:
```txt
http://localhost:3000/api/slots?date=2026-03-30
http://localhost:3000/api/slots?booked=false
http://localhost:3000/api/slots?locked=false
http://localhost:3000/api/slots?date=2026-03-30&booked=false&locked=false
```

### Success response
```json
{
  "message": "Slots fetched successfully",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2026-03-26T10:00:00Z",
      "UpdatedAt": "2026-03-26T10:00:00Z",
      "DeletedAt": null,
      "date": "2026-03-30",
      "time": "10:00",
      "isBooked": false,
      "isLocked": false
    }
  ]
}
```

### Error responses
```json
{
  "error": "Failed to fetch slots"
}
```

---

## 4.2 Get One Slot

## GET
```txt
http://localhost:3000/api/slots/1
```

### Purpose
Return one slot by ID.

### How to send
No body required.

### Success response
```json
{
  "message": "Slot fetched successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

---

## 4.3 Create Slot (Admin Only)

## POST
```txt
http://localhost:3000/api/slots
```

### Purpose
Create a new slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "date": "2026-03-30",
  "time": "10:00",
  "isLocked": false
}
```

### Success response
```json
{
  "message": "Slot created successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Date and time are required"
}
```

```json
{
  "error": "A slot with the same date and time already exists"
}
```

```json
{
  "error": "Failed to create slot"
}
```

```json
{
  "error": "Admin access required"
}
```

---

## 4.4 Update Slot (Admin Only)

## PUT
```txt
http://localhost:3000/api/slots/1
```

### Purpose
Update a slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

JSON body:
```json
{
  "date": "2026-03-30",
  "time": "11:00",
  "isLocked": false
}
```

### Success response
```json
{
  "message": "Slot updated successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:30:00Z",
    "DeletedAt": null,
    "date": "2026-03-30",
    "time": "11:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "Date and time are required"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Booked slots cannot be moved to a different date or time"
}
```

```json
{
  "error": "Another slot with the same date and time already exists"
}
```

```json
{
  "error": "Failed to update slot"
}
```

---

## 4.5 Lock Slot (Admin Only)

## PATCH
```txt
http://localhost:3000/api/slots/1/lock
```

### Purpose
Lock a slot so it cannot be booked.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Slot locked successfully",
  "data": {
    "ID": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": true
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Failed to lock slot"
}
```

---

## 4.6 Unlock Slot (Admin Only)

## PATCH
```txt
http://localhost:3000/api/slots/1/unlock
```

### Purpose
Unlock a slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Slot unlocked successfully",
  "data": {
    "ID": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "isBooked": false,
    "isLocked": false
  }
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Failed to unlock slot"
}
```

---

## 4.7 Delete Slot (Admin Only)

## DELETE
```txt
http://localhost:3000/api/slots/1
```

### Purpose
Delete a slot.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Slot deleted successfully",
  "data": null
}
```

### Error responses
```json
{
  "error": "Invalid slot ID"
}
```

```json
{
  "error": "Slot not found"
}
```

```json
{
  "error": "Booked slots cannot be deleted"
}
```

```json
{
  "error": "Failed to delete slot"
}
```

---

# 5. Appointments

## 5.1 Get My Appointments

## GET
```txt
http://localhost:3000/api/appointments/me
```

### Purpose
Return only the logged-in user's appointments.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Your appointments fetched successfully",
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2026-03-26T10:00:00Z",
      "UpdatedAt": "2026-03-26T10:00:00Z",
      "DeletedAt": null,
      "userId": 1,
      "customerName": "Peter",
      "serviceId": 1,
      "serviceName": "Haircut",
      "slotId": 1,
      "date": "2026-03-30",
      "time": "10:00",
      "price": 120,
      "status": "confirmed"
    }
  ]
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

```json
{
  "error": "Failed to fetch your appointments"
}
```

---

## 5.2 Get One Appointment

## GET
```txt
http://localhost:3000/api/appointments/1
```

### Purpose
Return one appointment.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointment fetched successfully",
  "data": {
    "ID": 1,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "price": 120,
    "status": "confirmed"
  }
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only view your own appointments"
}
```

---

## 5.3 Create Appointment

## POST
```txt
http://localhost:3000/api/appointments
```

### Purpose
Create a new appointment for the logged-in user.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

JSON body:
```json
{
  "serviceId": 1,
  "date": "2026-03-30",
  "time": "10:00"
}
```

### Success response
```json
{
  "message": "Appointment created successfully",
  "data": {
    "ID": 1,
    "CreatedAt": "2026-03-26T10:00:00Z",
    "UpdatedAt": "2026-03-26T10:00:00Z",
    "DeletedAt": null,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "price": 120,
    "status": "confirmed"
  }
}
```

### Error responses
```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "ServiceId, date and time are required"
}
```

```json
{
  "error": "User not found"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Slot not found for the given date and time"
}
```

```json
{
  "error": "Slot is locked"
}
```

```json
{
  "error": "Slot is already booked"
}
```

```json
{
  "error": "Failed to create appointment"
}
```

```json
{
  "error": "Failed to reserve slot"
}
```

---

## 5.4 Update Appointment

## PUT
```txt
http://localhost:3000/api/appointments/1
```

### Purpose
Update an existing appointment.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

JSON body:
```json
{
  "serviceId": 1,
  "date": "2026-03-30",
  "time": "11:00",
  "status": "confirmed"
}
```

### Success response
```json
{
  "message": "Appointment updated successfully",
  "data": {
    "ID": 1,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 2,
    "date": "2026-03-30",
    "time": "11:00",
    "price": 120,
    "status": "confirmed"
  }
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Invalid request body"
}
```

```json
{
  "error": "ServiceId, date and time are required"
}
```

```json
{
  "error": "Status must be either confirmed or cancelled"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only update your own appointments"
}
```

```json
{
  "error": "Service not found"
}
```

```json
{
  "error": "Current slot linked to appointment was not found"
}
```

```json
{
  "error": "Target slot not found for the given date and time"
}
```

```json
{
  "error": "Target slot is locked"
}
```

```json
{
  "error": "Target slot is already booked"
}
```

```json
{
  "error": "Failed to release old slot"
}
```

```json
{
  "error": "Failed to reserve target slot"
}
```

```json
{
  "error": "Failed to update appointment"
}
```

---

## 5.5 Cancel Appointment

## PATCH
```txt
http://localhost:3000/api/appointments/1/cancel
```

### Purpose
Cancel an appointment and free the linked slot.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointment cancelled successfully",
  "data": {
    "ID": 1,
    "userId": 1,
    "customerName": "Peter",
    "serviceId": 1,
    "serviceName": "Haircut",
    "slotId": 1,
    "date": "2026-03-30",
    "time": "10:00",
    "price": 120,
    "status": "cancelled"
  }
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only cancel your own appointments"
}
```

```json
{
  "error": "Linked slot not found"
}
```

```json
{
  "error": "Failed to cancel appointment"
}
```

```json
{
  "error": "Failed to free slot"
}
```

---

## 5.6 Delete Appointment

## DELETE
```txt
http://localhost:3000/api/appointments/1
```

### Purpose
Delete an appointment and free the slot first.

### How to send
Header:
```txt
Authorization: Bearer your-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointment deleted successfully",
  "data": null
}
```

### Error responses
```json
{
  "error": "Invalid appointment ID"
}
```

```json
{
  "error": "Appointment not found"
}
```

```json
{
  "error": "You can only delete your own appointments"
}
```

```json
{
  "error": "Failed to free slot"
}
```

```json
{
  "error": "Failed to delete appointment"
}
```

---

## 5.7 Get All Appointments (Admin Only)

## GET
```txt
http://localhost:3000/api/appointments
```

### Purpose
Return all appointments in the system.

### How to send
Header:
```txt
Authorization: Bearer your-admin-jwt-token
```

No body required.

### Success response
```json
{
  "message": "Appointments fetched successfully",
  "data": [
    {
      "ID": 1,
      "userId": 1,
      "customerName": "Peter",
      "serviceId": 1,
      "serviceName": "Haircut",
      "slotId": 1,
      "date": "2026-03-30",
      "time": "10:00",
      "price": 120,
      "status": "confirmed"
    }
  ]
}
```

### Error responses
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

```json
{
  "error": "Admin access required"
}
```

```json
{
  "error": "Failed to fetch appointments"
}
```

---

# Common Errors

## Unauthorized
```json
{
  "error": "Authorization header is required"
}
```

```json
{
  "error": "Authorization header must be in the format: Bearer <token>"
}
```

```json
{
  "error": "Invalid or expired token"
}
```

## Forbidden
```json
{
  "error": "Admin access required"
}
```

```json
{
  "error": "You can only view your own appointments"
}
```

```json
{
  "error": "You can only update your own appointments"
}
```

```json
{
  "error": "You can only cancel your own appointments"
}
```

```json
{
  "error": "You can only delete your own appointments"
}
```

## Validation / Booking errors
```json
{
  "error": "Slot is locked"
}
```

```json
{
  "error": "Slot is already booked"
}
```

```json
{
  "error": "A slot with the same date and time already exists"
}
```

```json
{
  "error": "Booked slots cannot be deleted"
}

```json
{
  "error": "Booked slots cannot be moved to a different date or time"
}

---

# Quick Testing Order

## 1. Check health

GET http://localhost:3000/api/health
```

## 2. Register a customer

POST http://localhost:3000/api/auth/register

## 3. Login

POST http://localhost:3000/api/auth/login

## 4. Promote one user to admin in the database

## 5. Login as admin

## 6. Create a service

POST http://localhost:3000/api/services

## 7. Create a slot

POST http://localhost:3000/api/slots

## 8. Login as customer

## 9. Create an appointment

POST http://localhost:3000/api/appointments


## 10. View my appointments

GET http://localhost:3000/api/appointments/me

---

# Endpoints

## Public
- `GET http://localhost:3000/api/health`
- `POST http://localhost:3000/api/auth/register`
- `POST http://localhost:3000/api/auth/login`
- `GET http://localhost:3000/api/services`
- `GET http://localhost:3000/api/services/:id`
- `GET http://localhost:3000/api/slots`
- `GET http://localhost:3000/api/slots/:id`

## Authenticated
- `GET http://localhost:3000/api/auth/me`
- `POST http://localhost:3000/api/auth/logout`
- `GET http://localhost:3000/api/appointments/me`
- `GET http://localhost:3000/api/appointments/:id`
- `POST http://localhost:3000/api/appointments`
- `PUT http://localhost:3000/api/appointments/:id`
- `PATCH http://localhost:3000/api/appointments/:id/cancel`
- `DELETE http://localhost:3000/api/appointments/:id`

## Admin Only
- `POST http://localhost:3000/api/services`
- `PUT http://localhost:3000/api/services/:id`
- `DELETE http://localhost:3000/api/services/:id`
- `POST http://localhost:3000/api/slots`
- `PUT http://localhost:3000/api/slots/:id`
- `PATCH http://localhost:3000/api/slots/:id/lock`
- `PATCH http://localhost:3000/api/slots/:id/unlock`
- `DELETE http://localhost:3000/api/slots/:id`
- `GET http://localhost:3000/api/appointmets`

3485