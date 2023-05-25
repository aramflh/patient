from bs4 import BeautifulSoup
import psycopg2

DB_NAME = "dbtest"
DB_USER = "ayoub"
DB_PASS = "root"
DB_HOST = "localhost"
DB_PORT = "5432"


conn = psycopg2.connect(database=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)

print("Database connected successfully")

with open('data/specialites.xml', 'r', encoding='utf8') as f:
    data = f.read()

soup = BeautifulSoup(data, 'lxml')

systemes = []
specialites = []
for sys_ana in soup.find_all("medicament"):
    systemes.append(sys_ana.string.strip().lower())

for specialite in soup.find_all("name"):
    specialites.append(specialite.string.strip().lower())

unique_systemes = set(systemes)

cur = conn.cursor()

for spe in specialites:
    cur.execute(f"""
    INSERT INTO "Specialite" (specialite)
    VALUES ('{spe}');
    """)

for sys in unique_systemes:
    cur.execute(f"""
    INSERT INTO "Systeme_ana" (nom_sys_ana)
    VALUES ('{sys}');
    """)

for s in soup.find_all("specialite"):
    name = s.find("name")
    sys_ana = s.find_all("medicament")
    for a in sys_ana:
        print(a.string)
        cur.execute(f"""
        INSERT INTO "Specialite_systeme_ana" (nom_sys_ana, specialite)
        VALUES ('{a.string.strip().lower()}', '{name.string.strip().lower()}');
        """)

conn.commit()

print('Data added successfully')
conn.close()