
# JWT Authentication

A GO project to demonstrate API endpoint protection using JWT token.

## Steps to follow

- Do Sign up
- Login and copy generated token
- Pass that token as Authorization Header to protected APIs
- While signing up, note down the password


## API Reference

#### Get all events

```http
  GET /events
```

#### Get event by id

```http
  GET /events/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of event to fetch |

#### Create event

```http
  POST /events
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `event`      | `json` | **Refer `./api` folder for request**. |
| `token` | `string` | **Required:** pass as Authorization Header |

#### Update event

```http
  PUT /events/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `int` | **Required**. Id of event to update |
| `event`      | `json` | **Refer `./api` folder for request**. |
| `token` | `string` | **Required:** pass as Authorization Header |

#### Delete event

```http
  DELETE /events/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `int` | **Required**. Id of event to delete |
| `token` | `string` | **Required:** pass as Authorization Header |

#### Register for an event

```http
  POST /events/{id}/register
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `int` | **Required**. Id of event to register |
| `token` | `string` | **Required:** pass as Authorization Header |

#### Unregister for an event

```http
  DELETE /events/{id}/unregister
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `int` | **Required**. Id of event to unregister |
| `token` | `string` | **Required:** pass as Authorization Header |

#### User Sign Up

```http
  POST /signup
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `user`      | `json` | **Refer `./api` folder for request**.|

#### User Login

```http
  POST /login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `user`      | `json` | **Refer `./api` folder for request**.|
