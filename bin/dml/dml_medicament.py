import csv


def dml_medic():
    sql_medic = []
    sql_quantite = []
    sql_medic_quantite = []
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

    for unique_medicament in medicaments:
        sql_medic.append(f"""
        INSERT INTO "Medicament" (nom_commercial, dci, nom_sys_ana)
        VALUES ('{unique_medicament["nom Commercial"].strip().replace("'","''").lower()}', '{unique_medicament["dci"].strip().replace("'","''").lower()}', '{unique_medicament["système anatomique"].strip().replace("'","''").lower()}');
        """) 

    for unique_quantite in quantites_unique:
        sql_quantite.append(f"""
        INSERT INTO "Conditionnement" (quantite)
        VALUES ('{unique_quantite}');
        """) 

    for element in data:
        sql_medic_quantite.append(f"""
        INSERT INTO "Medicament_conditionnement" (quantite, nom_commercial)
        VALUES ('{element['conditionnement'].strip().replace("'","''").lower()}', '{element['nom Commercial'].strip().replace("'","''").lower()}');
        """)
    
    return sql_medic, sql_quantite, sql_medic_quantite

