# A simple HRMS in golang

## env files
```
POSTGRES_CONN_STRING = "user=username dbname=databaseName password=password sslmode=disable"

DB_NAME = "usename"
DB_USERNAME = "password"
DB_PASSWORD = "databaseName"


```

## Run Instruction
1. Install docker for database 
2. Install go
3. Run the following commands
```
$ make dockerBuild
$ make goInstallMod
$ make goBuildR
```