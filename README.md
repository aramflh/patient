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
```

## Commands

go mod init <module path>
- Creates a new module

go mod tidy
- add module requirements and sums

go run <file.go>
- Compiles and run the code



