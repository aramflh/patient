import csv
import psycopg2


DB_NAME = "dbtest"
DB_USER = "ayoub"
DB_PASS = "root"
DB_HOST = "localhost"
DB_PORT = "5432"


conn = psycopg2.connect(database=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)


print("Database connected successfully")

prescriptions = []
traitements = []

with open("data\dossiers_patients.csv", "r", encoding='utf8') as dossier:
    reader = csv.DictReader(dossier)
    for row in reader:
        prescription = {"n_niss":row["NISS_patient"], "date_prescription":row["date_prescription"],"nom_commercial":row["medicament_nom_commercial"],"inami_med":row["inami_medecin"]}
        traitement = {"date_vente":row["date_vente"],"duree_traitement":row["duree_traitement"]}
        prescriptions.append(prescription)
        traitements.append(traitement)

cur = conn.cursor()

for i in range(len(prescriptions)):
    cur.execute(f"""
    INSERT INTO "Prescription" (id, date_prescription,n_niss,nom_commercial,inami_med)
    VALUES ('{i}', '{prescriptions[i]['date_prescription'].strip().lower()}', '{prescriptions[i]['n_niss'].strip().lower()}', '{prescriptions[i]['nom_commercial'].strip().lower()}', '{prescriptions[i]['inami_med'].strip().lower()}');
    """)  

    cur.execute(f"""
    INSERT INTO "Traitement" (id, date_vente,duree_traitement, id_prescription)
    VALUES ('{i}', '{traitements[i]['date_vente'].strip().lower()}', '{traitements[i]['duree_traitement'].strip().lower()}', '{i}');
    """)       
     

conn.commit()

print('Data added successfully')
conn.close()