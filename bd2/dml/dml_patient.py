from bs4 import BeautifulSoup
import psycopg2

DB_NAME = "dbtest"
DB_USER = "ayoub"
DB_PASS = "root"
DB_HOST = "localhost"
DB_PORT = "5432"


conn = psycopg2.connect(database=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)

def sanitize(data):
    return data.strip().replace("'","''").lower()

print("Database connected successfully")
with open('data/patients_corrige.xml', 'r', encoding='utf8') as f:
    data = f.read()

soup = BeautifulSoup(data, 'lxml')
patients = []

for patient in soup.find_all("patient"):
    row = {
        "nom":sanitize(patient.nom.string),
        "prenom":sanitize(patient.prenom.string),
        "niss":sanitize(patient.niss.string),
        "genre":sanitize(patient.genre.string),
        "date_naissance":sanitize(patient.date_de_naissance.string),
        "inami_med":sanitize(patient.inami_medecin.string),
        "inami_pha":sanitize(patient.inami_pharmacien.string),
        "mail":sanitize(patient.mail.string if patient.mail.string != "None"  else ''),
        "telephone":sanitize(patient.telephone.string if patient.telephone.string != "None" else '')
        }
    if row not in patients:
        patients.append(row)

cur = conn.cursor()
for patient in patients:
    cur.execute(f"""
    INSERT INTO "Patient" (n_niss, nom, prenom, genre, date_naissance, a_mail, n_telephone, inami, n_inami_pha)
    VALUES ('{patient['niss']}','{patient['nom']}','{patient['prenom']}','{patient['genre']}','{patient['date_naissance']}','{patient['mail']}','{patient['telephone']}','{patient['inami_med']}','{patient['inami_pha']}');
    """)


conn.commit()

print('Data added successfully')
conn.close()