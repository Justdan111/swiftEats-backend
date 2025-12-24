# SwiftEats API — Postman Testing Guide

A practical guide to test the SwiftEats backend endpoints with Postman.

## Prerequisites
- Backend running locally at `http://localhost:8080`.
- Environment variables set for the server (see `CHECKLIST.md`).
- Postman installed.

## Recommended Postman Setup
1. Create a new Postman Environment:
   - `baseUrl` = `http://localhost:8080`
   - `jwt` = (leave empty; will be set after login)
   - `adminJwt` = (optional; you can reuse `jwt` for now)
   - `restaurantId` = (optional; set after you create/list)
   - `menuItemId` = (optional; set after you create/list)
2. In requests, set `Authorization` to `Bearer {{jwt}}` for protected admin endpoints.
3. Use `{{baseUrl}}` in request URLs (e.g., `{{baseUrl}}/api/restaurants`).

## Collection Structure (suggested)
- Auth
  - Register
  - Login
- User — Restaurants
  - List Restaurants
  - Get Restaurant by ID
  - Get Restaurant Menu
- Admin — Restaurants
  - Create Restaurant
  - Update Restaurant
  - Delete Restaurant
- Admin — Menu
  - Create Menu Item
  - Update Menu Item
  - Delete Menu Item
  - Update Availability

## Auth Flow

### Register
- Method: POST
- URL: `{{baseUrl}}/api/auth/register`
- Body (JSON):
```json
{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "Passw0rd!"
}
```
- Expect: `200` or `201` with user info.

### Login
- Method: POST
- URL: `{{baseUrl}}/api/auth/login`
- Body (JSON):
```json
{
  "email": "jane@example.com",
  "password": "Passw0rd!"
}
```
- Tests (Postman → Tests tab) to capture JWT:
```javascript
pm.test("Login OK", function () {
  pm.response.to.have.status(200);
});
const data = pm.response.json();
if (data && data.token) {
  pm.environment.set("jwt", data.token);
}
```
- Expect: `200` with `{ "token": "..." }`.

### Me (verify token)
- Method: GET
- URL: `{{baseUrl}}/api/me`
- Headers: `Authorization: Bearer {{jwt}}`
- Expect: `200` with current user details.

## User — Restaurants (Public Read)

### List Restaurants
- Method: GET
- URL: `{{baseUrl}}/api/restaurants`
- Expect: `200` with array of restaurants.

### Get Restaurant by ID
- Method: GET
- URL: `{{baseUrl}}/api/restaurants/{{restaurantId}}`
- Path Variable: `restaurantId` (uuid)
- Expect: `200` with single restaurant.

### Get Menu for a Restaurant
- Method: GET
- URL: `{{baseUrl}}/api/restaurants/{{restaurantId}}/menu`
- Expect: `200` with array of menu items.

## Admin — Restaurants & Menu (Protected)
Note: For now, any valid JWT may suffice. Later, role checks will require admin.

### Create Restaurant
- Method: POST
- URL: `{{baseUrl}}/api/admin/restaurants`
- Headers: `Authorization: Bearer {{jwt}}`
- Body (JSON):
```json
{
  "name": "SwiftEats Deli",
  "description": "Fresh bowls and sandwiches",
  "address": "123 Main Street"
}
```
- Tests (optional): capture `restaurantId`
```javascript
const x = pm.response.json();
if (x && x.id) pm.environment.set("restaurantId", x.id);
```
- Expect: `201` with created restaurant.

### Update Restaurant
- Method: PUT
- URL: `{{baseUrl}}/api/admin/restaurants/{{restaurantId}}`
- Headers: `Authorization: Bearer {{jwt}}`
- Body (JSON):
```json
{
  "name": "SwiftEats Deli Updated",
  "description": "Updated menu and hours",
  "address": "456 Elm Avenue"
}
```
- Expect: `200` with updated restaurant.

### Delete Restaurant
- Method: DELETE
- URL: `{{baseUrl}}/api/admin/restaurants/{{restaurantId}}`
- Headers: `Authorization: Bearer {{jwt}}`
- Expect: `204` or `200`.

### Create Menu Item
- Method: POST
- URL: `{{baseUrl}}/api/admin/restaurants/{{restaurantId}}/menu`
- Headers: `Authorization: Bearer {{jwt}}`
- Body (JSON):
```json
{
  "name": "Chicken Bowl",
  "description": "Grilled chicken, rice, veggies",
  "price_cents": 1299,
  "is_available": true,
  "category_id": null
}
```
- Tests (optional): capture `menuItemId`
```javascript
const m = pm.response.json();
if (m && m.id) pm.environment.set("menuItemId", m.id);
```
- Expect: `201` with created menu item.

### Update Menu Item
- Method: PUT
- URL: `{{baseUrl}}/api/admin/restaurants/{{restaurantId}}/menu/{{menuItemId}}`
- Headers: `Authorization: Bearer {{jwt}}`
- Body (JSON):
```json
{
  "name": "Chicken Bowl (Large)",
  "description": "Extra serving",
  "price_cents": 1599,
  "is_available": true,
  "category_id": null
}
```
- Expect: `200` with updated item.

### Delete Menu Item
- Method: DELETE
- URL: `{{baseUrl}}/api/admin/restaurants/{{restaurantId}}/menu/{{menuItemId}}`
- Headers: `Authorization: Bearer {{jwt}}`
- Expect: `204` or `200`.

### Update Availability (PATCH)
- Method: PATCH
- URL: `{{baseUrl}}/api/admin/menu/{{menuItemId}}/availability`
- Headers: `Authorization: Bearer {{jwt}}`
- Body (JSON):
```json
{ "is_available": false }
```
- Expect: `200` with updated item.

## Common Tips
- Always set `Authorization: Bearer {{jwt}}` for admin routes.
- Ensure UUIDs for `restaurantId` / `menuItemId` are valid.
- Prices are integer cents (`price_cents`) — avoid floats.
- Use Postman `Tests` tab to chain variables from responses.

## Troubleshooting
- `401 Unauthorized`: Missing/invalid token — re-login and refresh `{{jwt}}`.
- `404 Not Found`: Check IDs exist.
- `400 Bad Request`: Validate JSON body types and required fields.

## Next
- Cart, Orders, and Payments will follow similar patterns and be added as modules mature.
