
# Go_Postgres_Gorm

This is a Golang project that uses a PostgreSQL Database and Go Fiber package for the HTTP server to create a REST API for a Book Database and perform CRUD operations on it using the GORM package


## Roadmap

### Project Structure
![ProjectStructure](https://github.com/user-attachments/assets/4ebb764c-711b-47d6-907c-61d779f69af8)

### Routes
![Routes](https://github.com/user-attachments/assets/63ddf52f-704b-497d-b3ce-46b085297aff)



## Testing Demo

Refer to the provided Google Doc for the [Testing Demo](https://docs.google.com/document/d/e/2PACX-1vTv7V2kIw7SE28KGtRioY4o0keBATvRVnYYITpXH9FzY8INQztkKaZJT7zuYLDM7kd6pVb7He6Z-dHZ/pub)




## Run Locally

#### Clone the project

```bash
  git clone https://github.com/Debsnil24/Go_Postgres_GORM.git
```

#### Go to the project directory

```bash
  cd {Directory}/Go_Postgres_GORM
```

#### Install dependencies

```go
  go get "github.com/gofiber/fiber/v2"
```
```go
  go get "github.com/joho/godotenv"
```
```go
  go get "gorm.io/gorm"
```
```go
  go get "gorm.io/driver/postgres"
```

#### Pre-Requisites

- Ensure to have PostgreSQL installed on your Local Machine
*(Note: Don't use the postgres user profile for this project)* 
- Login to Admin Profile
```bash
    psql -U postgres
```
*Enter the Password if prompted*

- Create a New User 
```sql
    CREATE DATABASE username;
    CREATE USER username WITH PASSWORD 'yourpassword';
    ALTER USER username WITH SUPERUSER;
```
*Quit using \q & Relogin using the new user*
- Create a New Database called go_gorm_fiber
```sql
    CREATE DATABASE go_gorm_fiber;
```
- The Table will be automatically created using the AutoMigrate Function in the code based on the Books Struct defined.
- Update the .env file with username and yourpassword

#### Starting the server
- Go into the main folder 
```bash
  cd {Directory}/Go_Postgres_GORM/main
```
- Build the project
```go
  go build main.go
```
- Run the main.go file
```go
  go run main.go
```


## Acknowledgements

#### Check Out Akhil Sharma's Youtube Video for Project Tutorial
*Note - This project has different file & code structure than the Video Tutorial & contains an extra Update Function to update data in the Database*

[Postgres With Go Using GORM & Go-Fiber](https://youtu.be/1XPktts9USg?si=G388AnFbQZ0EE82c)



