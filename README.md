# Projet BD

## API Documentation

TODO

## Initialization

- Creates a new module
```Shell
go mod init <module_path>
```
- Install dependacy
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
- add module requirements and sums
```Shell
go mod tidy
```

## Run command

- Compiles and run the code
```Shell
go run <file.go>
```

**_OR_**

- Activate automatic compilation
```Shell
CompileDaemon -command="./patient"
```

## Project structure

TODO

## Request

- La liste des noms commerciaux de ḿedicaments correspondant à un nom en DCI, class ́es par ordre alphab ́etique et taille de conditionnement.

```SQL
SELECT nom_medic FROM "Medicament" ORDER BY  conditionnement, nom_medic;
```

- La liste des pathologies qui peuvent ˆetre prise en charge par un seul type de sp ́ecialistes.


