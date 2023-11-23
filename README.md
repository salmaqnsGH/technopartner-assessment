## Table of Contents

- [Requirements](#requirements)
- [Instalation](#instalation)
- [Documentation](#documentation)



## Requirements

- Go version 1.15 or higher
- PostgreSQL (13 or higher)


## Installation
1. Create database 
2. Import collection
3. Install the dependencies:
    ```bash
    go get
    go mod tidy
    ```

4. Configure your database connection APP_DB_USERNAME, APP_DB_PASSWORD, APP_DB_NAME in env.example file then rename file to .env

5. Run the application:
    ```bash
    go run main.go
    ```
6. Run in postman with method based on port you created
  ``` http://localhost:3000 ```


## Documentation
### Postman Documentation
[https://documenter.getpostman.com/view/14858801/2s9YeBetpY](https://documenter.getpostman.com/view/14858801/2s9YeBetpY)