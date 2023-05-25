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

with open('data/pharmaciens.xml', 'r', encoding='utf8') as f:
    data = f.read()

soup = BeautifulSoup(data, 'lxml')
pharmaciens = []

for pharmacien in soup.find_all("pharmacien"):
    row = {
            "nom":sanitize(pharmacien.nom.string),
            "inami":sanitize(pharmacien.inami.string),
            "mail":sanitize(pharmacien.mail.string if pharmacien.mail.string != None else ''),
            "telephone":sanitize(pharmacien.tel.string if pharmacien.tel.string != None else '')
        }
    if row not in pharmaciens:
        pharmaciens.append(row)

cur = conn.cursor()
for pharmacien in pharmaciens:
    cur.execute(f"""
    INSERT INTO "Pharmacien" (nom, a_mail, n_telephone, inami)
    VALUES ('{pharmacien['nom']}','{pharmacien['mail']}','{pharmacien['telephone']}','{pharmacien['inami']}');
    """)



conn.commit()

print('Data added successfully')
conn.close()