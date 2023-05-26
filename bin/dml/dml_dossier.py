import csv

def dml_dossier():
    sql_prescription = []
    sql_traitement = []
    prescriptions = []
    traitements = []

    with open("data/dossiers_patients.csv", "r", encoding='utf8') as dossier:
        reader = csv.DictReader(dossier)
        for row in reader:
            prescription = {"n_niss":row["NISS_patient"], "date_prescription":row["date_prescription"],"nom_commercial":row["medicament_nom_commercial"],"inami_med":row["inami_medecin"]}
            traitement = {"date_vente":row["date_vente"],"duree_traitement":row["duree_traitement"], "inami_pha":row["inami_pharmacien"] }
            prescriptions.append(prescription)
            traitements.append(traitement)



    for i in range(len(prescriptions)):
        sql_prescription.append(f"""
        INSERT INTO "Prescription" (id, date_prescription,n_niss,nom_commercial,inami_med)
        VALUES ('{i}', '{prescriptions[i]['date_prescription'].strip().lower()}', '{prescriptions[i]['n_niss'].strip().lower()}', '{prescriptions[i]['nom_commercial'].strip().lower()}', '{prescriptions[i]['inami_med'].strip().lower()}');
        """)  

        sql_traitement.append(f"""
        INSERT INTO "Traitement" (id, date_vente,duree_traitement, id_prescription, inami_pha)
        VALUES ('{i}', '{traitements[i]['date_vente'].strip().lower()}', '{traitements[i]['duree_traitement'].strip().lower()}', '{i}', '{traitements[i]['inami_pha'].strip().lower()}');
        """)       
    
    return [sql_prescription, sql_traitement]
