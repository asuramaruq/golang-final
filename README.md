# golang article backend application

## Olzhas Imangeldin SE-2201
## Alibek Zhampeisov SE-2201

## installation

1. `git clone https://github.com/asuramaruq/golang-final.git`
2. change values in the `.env` file to you local postgresql server
3. navigate to local repository using cd in terminal
4. `go mod download`
5. `go build`
6. `go run .`

## requests to test the api

### Post Register:

`curl --location 'http://localhost:8080/api/auth/register' \
-H "Content-Type: application/json" \
--data-raw '{
    "name":"testuser",
    "email":"testuser@gmail.com",
    "password":"12345678"
}'`

### Post Login:

`curl --location 'http://localhost:8080/api/auth/login' \
-H "Content-Type: application/json" \
--data-raw '{
    "email":"testuser@gmail.com",
    "password":"12345678"
}'`

### Get All Users:

`curl --location 'http://localhost:8080/api/users' \
-H "Authorization: Bearer auth_token"`

### Get User Profile:

`curl --location 'http://localhost:8080/api/users/profile' \
-H "Authorization: Bearer auth_token"`

### Put User Update:

`curl --location --request PUT 'http://localhost:8080/api/users/update' \
-H "Content-Type: application/json" \
-H "Authorization: Bearer auth_token" \
--data-raw '{
    "name":"update test",
    "email":"updatetest@gmail.com",
    "password":"12345678"
}'
`
### All Articles:

`curl --location 'http://localhost:8080/api/articles' \
-H "Authorization: Bearer auth_token"`

### Add Article:

`curl --location 'http://localhost:8080/api/articles/create' \
-H "Content-Type: application/json" \
-H "Authorization: Bearer auth_token" \
--data-raw '{
    "title":"article test",
    "description":"test article o_O"
}'`

### Update Article:

`curl --location --request PUT 'http://localhost:8080/api/articles/update' \
-H "Content-Type: application/json" \
-H "Authorization: Bearer auth_token" \
--data-raw '{
    "id":1,
    "title":"article edit",
    "description":"test edit O_o"
}'`

### Delete Article:

`curl --location --request DELETE 'http://localhost:8080/api/articles/{id}' \
-H "Authorization: Bearer auth_token"`

### Get Specific Article:

`curl --location 'http://localhost:8080/api/articles/{id}' \
-H "Authorization: Bearer auth_token"`
