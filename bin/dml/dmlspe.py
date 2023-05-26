from bs4 import BeautifulSoup

def dml_spe():

    sql_spe = []
    sql_sys = []
    sql_spe_sys = []

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

    for spe in specialites:
        sql_spe.append(f"""
        INSERT INTO "Specialite" (specialite)
        VALUES ('{spe}');
        """)

    for sys in unique_systemes:
        sql_sys.append(f"""
        INSERT INTO "Systeme_ana" (nom_sys_ana)
        VALUES ('{sys}');
        """)

    for s in soup.find_all("specialite"):
        name = s.find("name")
        sys_ana = s.find_all("medicament")
        for a in sys_ana:
            sql_spe_sys.append(f"""
            INSERT INTO "Specialite_systeme_ana" (nom_sys_ana, specialite)
            VALUES ('{a.string.strip().lower()}', '{name.string.strip().lower()}');
            """)
    return sql_spe, sql_sys, sql_spe_sys
