## Job Gopportunities __[Project in Progress]__

- Platform built using Golang to manage job postings efficiently. This API leverages cutting-edge technologies and follows best practices to ensure reliability, scalability, and maintainability.

## Technologies Used: 
[__Golang:__](https://golang.org/) The core programming languaged used to develop the API, renowned for its performance and concurrency support.

[__Go-Gin:__](https://github.com/gin-gonic/gin) A lightweight and flexible HTTP web framework used as the router, prividing powerful routing capabilities and middleware support.

[__SQLite:__](https://www.sqlite.org/index.html) A reliable and lightweight relational database management system used as the backend database, ensuring data persistence and integrity

[__GoORM:__](https://gorm.io/) A simples and powerful ORM library for Go, facilitating seamless communication between the application and the SQLite database.

[__Swagger:__](https://swagger.io/) Integration with Swagger for comprehensive API documentation and interactive testing, enabling developers to explore endpoints and make requests effortlessly

## Key Features: 
__Structured Package Layout:__ The API codebase follows a well-organized package stucture, enhancing readability and maintainability.

__Efficient Routing with Go-gin:__ Utilizing Go-Gin routing capabilities for handling endpoints and middleware integration

__Secure Database Communication:__ GoORM ensures secure and efficient communication with the SQLite Database, safeguarding data integrity

__Comprehensive Documentation with Swagger:__ Swagger integration provides detailed API documentation, facilitating exploration and testing. 

## Important Note
- __*That this is a project developed for educational purposes, aimed at learning and practicing Golang and related technologies.*__


### Getting Started
Follow these steps:

1. Clone the repository: `https://github.com/reishenrique/job-gopportunities`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

### Makefile Commands
- Makefile to help you manage common tasks more easily

    - `make` - Run the application, generating documentation by default

    - `make run` - Just run the application locally

    - `make test` - Script would handle running the tests for the application

    - `make run-with-docs` - Generate the API documentation using Swagger, then run the application

    - `make build` - Build the application and create an executable file named `job-gopportunities`

Type `make` followed by the desired command in your terminal
```sh
make run
```

## Usage

After launching the API, you can leverage the Swagger UI to interact with endpoints for searching, creating, editing, and deleting job opportunities. Access the API at http://localhost:$PORT/swagger/index.html


## Implemented and future improvements
- [X] Loggers
- [X] SQLite configuration enabling scalability
- [X] Error handling
- [ ] Implementation of request field validation (go-playground/validator).
- [ ] Implementation of authentication
- [ ] Unit tests
- [ ] Docker file and docker-compose
