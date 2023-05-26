import csv


pathologies = []
specialites = []
content = []
headers = ['pathologie','specialite']

def dml_patho():
    sql_patho = []
    sql_patho_spe = []
    with open('data/pathologies.csv', "r", encoding='utf8') as data:
        reader = csv.DictReader(data, fieldnames=headers)
        for row in reader:
            pathologies.append(row['pathologie'].strip().replace("'","''").lower())
            specialites.append(row['specialite'].strip().replace("'","''").lower())
            if row not in content:
                content.append(row)

    pathologies_unique = set(pathologies)

    for unique_pathologie in pathologies_unique:
        sql_patho.append(f"""
        INSERT INTO "Pathologie" (nom_pathologie)
        VALUES ('{unique_pathologie}');
        """) 


    for element in content:
        sql_patho_spe.append(f"""
        INSERT INTO "Pathologie_specialite" (nom_pathologie, specialite)
        VALUES ('{element["pathologie"].strip().replace("'","''").lower()}', '{element['specialite'].strip().lower()}');
        """)
    
    return sql_patho, sql_patho_spe
