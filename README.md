# Projet BD

## Prepare the DB

- Create a DB instance with credential
```shell
docker run --name dbtest -e POSTGRES_PASSWORD=root -e POSTGRES_USER=ayoub -e POSTGRES_DB=dbtest -p 5432:5432 -d postgres
```

- run the [script](bin) in order

## Features

- [x] Ajouter un médecin
- [x] Ajouter un pharmacien
- [X] Créer un compte patient
- [X] Se connecter en tant que patient
- [X] Modifier son pharmacien ou médecin de référence (Connecté)
- [ ] Consulter ses informations médicales (Connectés)  ??
- [X] Consulter ses traitements (Connecté)


## Initialization
0. Change [.env](.env) file with Postgres credentials if different

1. [Install Go](https://go.dev/doc/install)

2. Creates a new module if there is no [go.mod](go.mod) file

```Shell
go mod init patient
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

### 1
- [X] La liste des noms commerciaux de médicaments correspondant à un nom en DCI, classés par ordre alphabétique et taille de conditionnement.
```SQL
SELECT nom_medic
FROM "Medicament"
WHERE dci = 'Nom_en_DCI'
ORDER BY nom_medic ASC, conditionnement ASC;

```

### 2
- [X] La liste des pathologies qui peuvent être prise en charge par un seul type de spécialistes.

```SQL
SELECT p.nom_pathologie, m.specialite
FROM Pathologie p
INNER JOIN Medecin m ON p.nom_sys_ana = m.nom_sys_ana
GROUP BY p.nom_pathologie, m.specialite
HAVING COUNT(DISTINCT m.specialite) = 1;

```

### 3
- [X] La spécialité de médecins pour laquelle les médecins prescrivent le plus de médicaments.

```SQL
SELECT m.specialite, COUNT(*) AS total_prescriptions
FROM "Medecin" m
INNER JOIN "Prescription" p ON m.n_inami_med = p.n_inami_med
GROUP BY m.specialite
ORDER BY total_prescriptions DESC
LIMIT 1;
```

### 4
- [X] Tous les utilisateurs ayant consommé un médicament spécifique (sous son nom commercial) après une date donnée, par exemple en cas de rappel de produit pour lot contaminé.
```SQL
SELECT DISTINCT p.nom, p.prenom
FROM "Patient" p
INNER JOIN "Prescription" pr ON p.n_niss = pr.n_niss
INNER JOIN "Medicament" m ON pr.nom_medic = m.nom_medic
WHERE m.nom_medic = 'Nom_du_medicament'
AND pr.date_emission > 'Date_donnee';
```

### 5
- [X] Tous les patients ayant été traités par un médicament (sous sa DCI) à une date antérieure mais qui ne le sont plus, pour vérifier qu’un patients suive bien un traitement chronique.

```SQL
SELECT DISTINCT p.nom, p.prenom
FROM "Patient" p
INNER JOIN "Traitement" t ON p.n_niss = t.n_niss
INNER JOIN "Medicament" m ON t.nom_medic = m.nom_medic
WHERE m.dci = 'Nom_en_DCI'
AND t.date_debut < CURDATE()
AND t.date_fin IS NOT NULL;
```

### 6
- [X] La liste des médecins ayant prescrit des médicaments ne relevant pas de leur spécialité.

```SQL
SELECT DISTINCT m.n_inami_med, m.specialite
FROM "Medecin" m
INNER JOIN "Prescription" p ON m.n_inami_med = p.n_inami_med
INNER JOIN "Medicament" med ON p.nom_medic = med.nom_medic
WHERE med.nom_pathologie NOT IN (
SELECT nom_pathologie
FROM "Pathologie"
WHERE nom_sys_ana = m.nom_sys_ana
);
```

### 7
- [X] Pour chaque décennie entre 1950 et 2020, (1950 − 59, 1960 − 69, ...), le médicament le plus consommé par des patients nés durant cette décennie.

```SQL
SELECT SUBSTRING(YEAR(date_naissance), 1, 3) || '0s' AS decade, nom_medic
FROM "Patient" p
INNER JOIN "Prescription" pr ON p.n_niss = pr.n_niss
INNER JOIN "Medicament" m ON pr.nom_medic = m.nom_medic
WHERE YEAR(date_naissance) BETWEEN 1950 AND 2020
GROUP BY decade
HAVING COUNT(*) > 0
ORDER BY decade;
```

### 8
- [X] Quelle est la pathologie la plus diagnostiquée ?

```SQL
SELECT nom_pathologie, COUNT(*) AS total_diagnosis
FROM "Dossier_med"
GROUP BY nom_pathologie
ORDER BY total_diagnosis DESC
LIMIT 1;
```

### 9
- [X] Pour chaque patient, le nombre de médecin lui ayant prescrit un médicament.

```SQL
SELECT p.n_niss, p.nom, p.prenom, COUNT(DISTINCT pr.n_inami_med) AS nombre_medecins
FROM "Patient" p
INNER JOIN "Prescription" pr ON p.n_niss = pr.n_niss
GROUP BY p.n_niss, p.nom, p.prenom;
```
### 10
- [X] La liste de médicament n’étant plus prescrit depuis une date spécifique.

```SQL
SELECT DISTINCT nom_medic
FROM "Medicament"
WHERE nom_medic NOT IN (
SELECT DISTINCT nom_medic
FROM "Prescription"
WHERE date_emission > 'Date_specifique'
);
```

