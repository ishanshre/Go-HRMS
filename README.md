# A simple HRMS in golang

## .env file example
```
POSTGRES_CONN_STRING = "user=username dbname=databaseName password=password sslmode=disable"

DB_NAME = "usename"
DB_USERNAME = "password"
DB_PASSWORD = "databaseName"


```

## Run Instruction
1. Install docker for database 
2. Install go
3. Clone this repo
4. Create .env files and add variable according to above .env example
3. Run the following commands
```
$ git clone https://github.com/ishanshre/Go-HRMS.git
$ cd GO-HRMS
$ touch .env
$ make dockerBuild
$ make goInstallMod
$ make goBuildR
```
