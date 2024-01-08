# Gopherbank

Gopherbank is a simple bank application with the backend written in Go and frontend in React.
It is a learning project for me to learn Go and to learn how to write a web application in Go.

Showcase video(in Portuguese): https://www.linkedin.com/feed/update/urn:li:activity:7149169538620723200/

## How to Run

There is a docker compose file in the project, in order to use it you gotta follow the steps below

1. Create a .env file, there's a .env.example file inside the project that you can use in order to create your own
2. Run the command bellow:
```sh
docker-compose up --build -d
```
3. After the image is built and running, run the commands below to start the db migrations and seed the database with a few accounts and users:
```sh
# this will handle the db migration
docker exec gopherbank-backend-1 go run cmd/migration/main.go

# this will insert a few accounts and users, including an admin account for testing
docker exec gopherbank-backend-1 go run cmd/seed/main.go 
```

4. Go to your favorite browser and the application should be good and running on *localhost:5173/signin*

Obs.: So far this is a mobile only app, future additions will be made to support other layouts
