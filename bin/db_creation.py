import psycopg2
from ddl import ddl
from dml.dml_diagnostiques import dml_diag
from dml.dml_dossier import dml_dossier
from dml.dml_medecin import dml_medecin
from dml.dml_medicament import dml_medic
from dml.dml_patho import dml_patho
from dml.dml_patient import dml_patient
from dml.dml_pharmaciens import dml_pha
from dml.dmlspe import dml_spe

def insert_data(requests, table):
    for request in requests:
        cursor.execute(request)
    print(f"Data {table} imported")


# connection establishment
conn = psycopg2.connect(
   database="postgres",
    user='ayoub',
    password='root',
    host='localhost',
    port= '5432'
)
 
conn.autocommit = True
 
cursor = conn.cursor()
 
sql_creation = ''' CREATE database dbtest ; '''
sql_alter_date = ''' ALTER DATABASE dbtest SET datestyle TO "ISO, MDY";'''
cursor.execute(sql_creation)
cursor.execute(sql_alter_date)

print("Database created")

conn.close()

conn = psycopg2.connect(
   database="dbtest",
    user='ayoub',
    password='root',
    host='localhost',
    port= '5432'
)

conn.autocommit = True
cursor = conn.cursor()


# Création des tables
cursor.execute(ddl())
print("Tables created")

# Ajout des données des tables spécialité et systèmes anatomique

sql_spe, sql_sys, sql_spe_sys = dml_spe()

insert_data(sql_spe, "Specialite")
insert_data(sql_sys, "Systeme_ana")
insert_data(sql_spe_sys, "Specialite_systeme_ana")

# Ajout des données de la table Pathologie

sql_patho, sql_patho_spe = dml_patho()

insert_data(sql_patho, "Pathologie")
insert_data(sql_patho_spe, "Pathologie_specialite")

# Ajout des données de la table Medecin

sql_medecin = dml_medecin()

insert_data(sql_medecin, "Medecin")

# Ajout des données de la table Pharmacien
sql_pharmacien = dml_pha()

insert_data(sql_pharmacien, "Pharmacien")

# Ajout des données de la table Patient

sql_patient = dml_patient()

insert_data(sql_patient, "Patient")

# Ajout des données des tables Medicament et Conditionnement

sql_medic, sql_quantite, sql_medic_quantite = dml_medic()

insert_data(sql_medic, "Medicament")
insert_data(sql_quantite, "Conditionnement")
insert_data(sql_medic_quantite, "Medicament_conditionnement")

# Ajout des données de la table Diagnostic

sql_diagnostic = dml_diag()

insert_data(sql_diagnostic, "Diagnostic")

# Ajout des données des tables Prescription et Traitement

sql_prescription, sql_traitement = dml_dossier()

insert_data(sql_prescription, "Prescription")
insert_data(sql_traitement, "Traitement")

print('EVERYTHING DONE ....')