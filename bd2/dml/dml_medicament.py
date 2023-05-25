import csv
import psycopg2


DB_NAME = "dbtest"
DB_USER = "ayoub"
DB_PASS = "root"
DB_HOST = "localhost"
DB_PORT = "5432"


conn = psycopg2.connect(database=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)


print("Database connected successfully")

nom_medic = []
medicaments = []
quantites = []
data = []

with open("data/medicaments.csv", "r", encoding='utf8') as medicament:
    reader = csv.DictReader(medicament)
    for row in reader:
        if row['nom Commercial'] not in nom_medic:
            medicaments.append({'nom Commercial': row['nom Commercial'], 'système anatomique': row['système anatomique'],'dci' :row['dci'] })
            nom_medic.append(row['nom Commercial'] )
        quantites.append(row['conditionnement'])
        data.append(row)

quantites_unique = set(quantites)

cur = conn.cursor()

for unique_medicament in medicaments:
    cur.execute(f"""
    INSERT INTO "Medicament" (nom_commercial, dci, nom_sys_ana)
    VALUES ('{unique_medicament["nom Commercial"].strip().replace("'","''").lower()}', '{unique_medicament["dci"].strip().replace("'","''").lower()}', '{unique_medicament["système anatomique"].strip().replace("'","''").lower()}');
    """) 

for unique_quantite in quantites_unique:
    cur.execute(f"""
    INSERT INTO "Conditionnement" (quantite)
    VALUES ('{unique_quantite}');
    """) 

for element in data:
    cur.execute(f"""
    INSERT INTO "Medicament_conditionnement" (quantite, nom_commercial)
    VALUES ('{element['conditionnement'].strip().replace("'","''").lower()}', '{element['nom Commercial'].strip().replace("'","''").lower()}');
    """)


conn.commit()

print('Data added successfully')
conn.close()