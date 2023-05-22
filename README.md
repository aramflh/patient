# Projet BD

## Dependancy

```Shell
# ORM
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
# Web framework
go get -u github.com/gin-gonic/gin
# .env management
go get github.com/joho/godotenv
# auto compile go project
go get github.com/githubnemo/CompileDaemon 
go install github.com/githubnemo/CompileDaemon   # Use in CLI
# Generates models from DB
go get -u gorm.io/gen
```

## Initialization commands

- Creates a new module
```Shell
go mod init <module path>
```

- add module requirements and sums
```Shell
go mod tidy
```

## Run command

- Compiles and run the code
```Shell
go run <file.go>
```

- Activate automatic compilation
```Shell
CompileDaemon -command="./patient"
```

## Project structure

- dal/ :
- model/ :
- initializers/ :
- controllers/ : 
