# Byfood Assignment Backend Section


## Prerequisite
- PostgreSQL
- `make` command on terminal
- Go version 1.21.6 (or using Go Version Manager)


## How to Run
- Create a PostgresSQL database
- Duplicate `./config/.env.example` named `./config/.env.local`, fill the value of the variable accordingly
- Install the project simply by executing :
```bash
make install
```
- Run DB migration by executing :
```bash
# Replace the placeholders with your DB credential
DB_NAME=DB_NAME_PLACEHOLDER DB_USER=DB_USERNAME_PLACEHOLDER DB_PASS=DB_PASSWORD_PLACEHOLDER make migrate
```
- Alternatively, you could run a seed script to populate the DB :
```bash
# Replace the placeholders with your DB credential
DB_NAME=DB_NAME_PLACEHOLDER DB_USER=DB_USERNAME_PLACEHOLDER DB_PASS=DB_PASSWORD_PLACEHOLDER make seed
```
- You could either directly run the app using development environment or build it into a binary first using following command :
```bash
# Run as development environment
make run

# Or you could build if first than run as a service
make build
./bin/byfood
```


## Project Structure
- `bin` is where the binary file is generated and ready to execute natively
- `cmd` is a service(s) entry point
- `config` is where the key-value variable and other configurable file located
- `db` is a place for writing database related scripts inside
  - `pg` all PostgreSQL related script were hosted
    - `factory` a directory for writing table Schema
    - `migration` all migration file went here (including initial data)
    - `seed` only use seed on development, contain dummy data
- `docs` REST API documentation (Swagger/OpenAPI Spec)
- `internal` is a place where all services/modules located. The services/modules will and should be encapsulated
- `pkg` almost identical with `internal` but all the module/packages should not be encapsulated
- `scripts` all the shell script for building, deploying, or other related to ops


## Sample Working App
![](https://raw.githubusercontent.com/Novando/byfood/master/images/url-modifier.png)