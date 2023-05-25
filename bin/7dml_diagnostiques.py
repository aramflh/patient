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


with open('data/diagnostiques.xml', 'r', encoding='utf8') as f:
    data = f.read()

soup = BeautifulSoup(data, 'lxml')
diags = []

for diag in soup.find_all("diagnostique"):
    diags.append(
        {
        "niss":sanitize(diag.niss.string),
        "date_diagnostic":sanitize(diag.date_diagnostic.string),
        #"naissance":diag.naissance.string,
        "pathology":sanitize(diag.pathology.string),
        #"specialite":diag.specialite.string
        }
        )

cur = conn.cursor()
for diag in diags:
    cur.execute(f"""
    INSERT INTO "Diagnostic" (date_diagnostic, n_niss, nom_pathologie)
    VALUES ('{diag['date_diagnostic']}','{diag['niss']}','{diag['pathology']}');
    """)

    
conn.commit()

print('Data added successfully')
conn.close()