# Projet BD

## API Documentation

### Ajouter un médecin

- URL: /medecins 
- Method: POST 
- Request (JSON):
  - INAMI (string)
  - Nom (string)
  - Prenom (string)
  - Email (string)
  - Num (string)
  - Specialite (string)
- Response 
  - HTTP_201_CREATED 
  - HTTP_400_BAD_REQUEST

### Ajouter un pharmacien

- URL: /pharmaciens
- Method: POST
- Request (JSON):
    - INAMI (string)
    - Nom (string)
    - Prenom (string)
    - Email (string)
    - Num (string)
- Response
    - HTTP_201_CREATED
    - HTTP_400_BAD_REQUEST

    
### Créer un compte patient

- URL: /signup
- Method: POST
- Request (JSON):
    - INSS (string)
    - Nom (string)
    - Prenom (string)
    - Genre (1 char max)
    - DateNaissance (string yyyy-mm-dd)
    - Email (string)
    - Password (string)
    - Num (string)
    - INAMIMed (string)
    - INAMIPha (string)
- Response
  - HTTP_201_CREATED
  - HTTP_400_BAD_REQUEST

### Se connecter en tant que patient

- URL:
- Method:
- Request (JSON):
    - A
    - A
- Response
    - A
    - A
  
### Modifier son pharmacien ou médecin de référence (Connecté)
- URL:
- Method:
- Request (JSON):
    - A
    - A
- Response
    - A
    - A
  
### Consulter SES informations médicales (Connectés)
- URL:
- Method:
- Request (JSON):
    - A
    - A
- Response
    - A
    - A
  
### Consulter SES traitements (Connecté)
- URL:
- Method:
- Request (JSON):
    - A
    - A
- Response
    - A
    - A
  
    
## Initialization

- Creates a new module
```Shell
go mod init <module_path>
```
- Install dependacy
```Shell
# ORM - Database management
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
# Web framework API
go get -u github.com/gin-gonic/gin
# Crypto functions
go get -u golang.org/x/crypto/bcrypt
# JWT package
go get -u github.com/golang-jwt/jwt/v5
# .env management
go get github.com/joho/godotenv
# auto compile go project - Not mandatory
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

- [X] La liste des noms commerciaux de médicaments correspondant à un nom en DCI, classées par ordre alphabétique et taille de conditionnement.

```SQL
SELECT nom_medic FROM "Medicament" ORDER BY  conditionnement, nom_medic;
```

- [ ] La liste des pathologies qui peuvent être prise en charge par un seul type de spécialistes.


