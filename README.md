# bct-go-assessment
Simple Go application that provides a REST API to maintain a table of users.

## Usage

### Get All Users
GET /users

### Create User
POST /users

### Update User
PUT /users/:id

### Delete User
DELETE /users/:id

### Get User
NOT IMPLEMENTED

### Examples
`curl localhost:8080/users`

`curl -X POST -v -H 'Content-Type: application/json' -d '{"user_name": "gwashington", "first_name": "George", "last_name": "Washington", "user_status": "A", "email": "gwashington@example.com", "department": "CEO"}' localhost:8080/users`

`curl -X PUT -v -H 'Content-Type: application/json' -d '{"user_name": "gwashington", "first_name": "George", "last_name": "Washington", "user_status": "I", "email": "gwashington@example.com", "department": "President"}' localhost:8080/users/10017`

`curl -X DELETE -v localhost:8080/users/10017`

## Not Completed
* Swagger documentation.  While I've worked on projects that have used Swagger, I haven't done any work with it myself.

## Notes
* Configuration information is hardcoded.  For a production application sensitive information would typically be passed into the application via environment variables.

* user_status is maintained in a separate table.

* Once created, the user_id and user_name cannot be changed.

* Unique user_name is enforced by a unique constraint on the users table.
  * Attempts to create a duplicate user_name will result in a `Username already exists` error.

* Error handling is corse grained with the exception of duplicate user name

* Controller unit tests mock the model and primarily test how request are handled and if the requests are returning the correct response status codes.  The controller tests do not examine the repsonse body at this time.  This can be added later.
  * To run the controller tests execute `go test -v ./controller` from the project's root directory.

* I did not add unit tests to the model at this time.  To do so I would use a package like [sqlmock](https://pkg.go.dev/github.com/data-dog/go-sqlmock) to mock out the database operations.

## Setting up the PostgreSQL Database
These commands have been tested on macOS.

### Prerequisites
Docker is installed
psql client is installed

### Actions

#### Running PostgreSQL
1. Pull Docker Image
`docker pull postgres`

2. Build data directory
`mkdir -p ~/srv/postgres`

3. Run docker image
`docker run --rm --name postgres-db -e POSTGRES_PASSWORD=password -d -v $HOME/srv/postgres:/var/lib/postgresql/data -p 5432:5432 postgres`

#### Stopping PostgreSQL
`docker stop postgres-db`

#### Logging into Database
* `psql -h localhost -U postgres -d userdb`

#### Creating starter data
1. `psql -h localhost -U postgres -f database.sql`
2. `psql -h localhost -U postgres -d userdb -f users.sql`
