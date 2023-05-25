from bs4 import BeautifulSoup
import psycopg2

DB_NAME = "dbtest"
DB_USER = "ayoub"
DB_PASS = "root"
DB_HOST = "localhost"
DB_PORT = "5432"


conn = psycopg2.connect(database=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)

print("Database connected successfully")

def sanitize(data):
    return data.strip().replace("'","''").lower()

with open('data/medecins.xml', 'r', encoding="utf8") as f:
    data = f.read()

soup = BeautifulSoup(data, 'lxml')
medecins = []

for medecin in soup.find_all("medecin"):
    row = {
            "nom":sanitize(medecin.nom.string),
            "inami":sanitize(medecin.inami.string),
            "mail":sanitize(medecin.mail.string if medecin.mail.string != None else ''),
            "specialite":sanitize(medecin.specialite.string),
            "telephone":sanitize(medecin.telephone.string)
        }
    if row not in medecins:
        medecins.append(row)

cur = conn.cursor()
for medecin in medecins:
    cur.execute(f"""
    INSERT INTO "Medecin" (nom, a_mail, n_telephone, inami, specialite)
    VALUES ('{medecin['nom']}','{medecin['mail']}','{medecin['telephone']}','{medecin['inami']}','{medecin['specialite']}');
    """)



conn.commit()

print('Data added successfully')
conn.close()