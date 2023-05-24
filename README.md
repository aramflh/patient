# Projet BD

## Features

- [x] Ajouter un médecin
- [x] Ajouter un pharmacien
- [ ] Créer un compte patient
- [ ] Se connecter en tant que patient
- [ ] Modifier son pharmacien ou médecin de référence (Connecté)
- [ ] Consulter SES informations médicales (Connectés)
- [ ] Consulter SES traitements (Connecté)


## Initialization
0. Change [.env](.env) file with Postgres credentials

1. [Install Go](https://go.dev/doc/install)

2. Creates a new module
```Shell
go mod init <module_path>
```
3. Install dependacy
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
```

4. add module requirements and sums
```Shell
go mod tidy
```

5. Compiles and run the code
```Shell
go run <file.go>
```

6. Connect to the server (see port in [.env](.env):
   http://localhost:3000/


## Request

- [X] La liste des noms commerciaux de médicaments correspondant à un nom en DCI, classés par ordre alphabétique et taille de conditionnement.

```SQL
SELECT nom_medic FROM "Medicament" ORDER BY conditionnement, nom_medic;
```

- [ ] La liste des pathologies qui peuvent être prise en charge par un seul type de spécialistes.

```SQL
TODO
```

- [ ] La spécialité de médecins pour laquelle les médecins prescrivent le plus de médicaments.

```SQL
TODO
```

- [ ] Tous les utilisateurs ayant consommé un médicament spécifique (sous son nom commercial) après une date donnée, par exemple en cas de rappel de produit pour lot contaminé.

```SQL
TODO
```

- [ ] Tous les patients ayant été traités par un médicament (sous sa DCI) à une date antérieure mais qui ne le sont plus, pour vérifier qu’un patients suive bien un traitement chronique.

```SQL
TODO
```

- [ ] La liste des médecins ayant prescrit des médicaments ne relevant pas de leur spécialité.

```SQL
TODO
```

- [ ] Pour chaque décennie entre 1950 et 2020, (1950 − 59, 1960 − 69, ...), le médicament le plus consommé par des patients nés durant cette décennie.

```SQL
TODO
```

- [ ] Quelle est la pathologie la plus diagnostiquée ?

```SQL
TODO
```

- [ ] Pour chaque patient, le nombre de médecin lui ayant prescrit un médicament.

```SQL
TODO
```

- [ ] La liste de médicament n’étant plus prescrit depuis une date spécifique.

```SQL
TODO
```

