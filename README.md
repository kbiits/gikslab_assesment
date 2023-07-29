# Gikslab Golang Assesment

## Pre requisite
- Golang ^1.18
- PostgreSQL ^14.2
- Redis
- Make (for Makefile)

## Config
You need to fill out all configurations and adjust it with your environment in file ${project directory}/config.json
especially for database DSN and Redis URL

## Postman
There's 2 file for postman that you need to import before test the app, which are Gikslab.postman_collection.json and Gikslab.postman_environment.json

## Database
The filename of database dump is database.sql. Please find it in the project root directory.
By default, the script will create a database with name `gikslab`.

***Notes: I dump the database using pg_dump, so you should run the script inside psql (CLI for postgres), if not the script can't successfully run without error***

## Credentials
In the database dump, I already provide users for easily test the API.
- Board user:
  - username: nabiel, password: password
- Expert user:
  - username: expert_1, password: password
  - username: expert_2, password: password 
- Trainer user:
  - username: trainer_1, password: password
  - username: trainer_2, password: password 
- Competitor user:
  - username: competitor_1, password: password
  - username: competitor_2, password: password 

## How to Run
To run the app, you just need to execute below commands
```bash
go mod download
make run
```
By default, the API will listen to port 3000