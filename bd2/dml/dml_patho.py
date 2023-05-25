import csv
import psycopg2


DB_NAME = "dbtest"
DB_USER = "ayoub"
DB_PASS = "root"
DB_HOST = "localhost"
DB_PORT = "5432"


conn = psycopg2.connect(database=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)


print("Database connected successfully")
pathologies = []
specialites = []
content = []
cur = conn.cursor()
headers = ['pathologie','specialite']


with open('data/pathologies.csv', "r", encoding='utf8') as data:
    reader = csv.DictReader(data, fieldnames=headers)
    for row in reader:
        pathologies.append(row['pathologie'].strip().replace("'","''").lower())
        specialites.append(row['specialite'].strip().replace("'","''").lower())
        if row not in content:
            content.append(row)

pathologies_unique = set(pathologies)

for unique_pathologie in pathologies_unique:
    cur.execute(f"""
    INSERT INTO "Pathologie" (nom_pathologie)
    VALUES ('{unique_pathologie}');
    """) 


for element in content:
    cur.execute(f"""
    INSERT INTO "Pathologie_specialite" (nom_pathologie, specialite)
    VALUES ('{element["pathologie"].strip().replace("'","''").lower()}', '{element['specialite'].strip().lower()}');
    """)    
conn.commit()

print('Data added successfully')
conn.close()