# Gopherbank

Gopherbank is a simple bank application witt the backend written in Go and frontend in React.
It is a learning project for me to learn Go and to learn how to write a web application in Go.

## How to Run

There is a docker compose file in the project, in order to use it you gotta follow the steps below
Obs.: make sure you have 

1. Create a .env file, there's a .env.example file inside the project that you can use in order to create you own
2. Run the command bellow:
```sh
docker-compose up --build -d
```
3. After the image is built and running, run the commands below to start the db migrations and seed the databse with a few accounts and users:
```sh
# this will handle the db migration
docker exec gopherbank-backend-1 go run cmd/migration/main.go

# this will insert a few accounts and users, including an admin account for testing
docker exec gopherbank-backend-1 go run cmd/seed/main.go 
```

4. Go to your favorite browser and the application should be good and running on *localhost:5173/signin*
