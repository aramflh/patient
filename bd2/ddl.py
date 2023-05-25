import psycopg2

DB_NAME = "dbtest"
DB_USER = "ayoub"
DB_PASS = "root"
DB_HOST = "localhost"
DB_PORT = "5432"


conn = psycopg2.connect(database=DB_NAME, user=DB_USER, password=DB_PASS, host=DB_HOST, port=DB_PORT)

print("Database connected successfully")


cur = conn.cursor()
cur.execute("""


CREATE TABLE "Specialite"(
    specialite VARCHAR PRIMARY KEY
);

CREATE TABLE "Systeme_ana"(
	nom_sys_ana VARCHAR PRIMARY KEY
);

CREATE TABLE "Specialite_systeme_ana"(
	nom_sys_ana VARCHAR,
    specialite VARCHAR,
    CONSTRAINT fk_specialite FOREIGN KEY (specialite) REFERENCES "Specialite"(specialite),
    CONSTRAINT fk_nom_sys_ana FOREIGN KEY (nom_sys_ana) REFERENCES "Systeme_ana"(nom_sys_ana),
    CONSTRAINT CompKey_sys_spe PRIMARY KEY (nom_sys_ana,specialite)
);

CREATE TABLE "Pathologie" (
	nom_pathologie VARCHAR PRIMARY KEY
);

CREATE TABLE "Pathologie_specialite" (
	nom_pathologie VARCHAR NOT NULL,
	specialite VARCHAR NOT NULL,
	CONSTRAINT fk_specialite FOREIGN KEY (specialite) REFERENCES "Specialite"(specialite),
    CONSTRAINT fk_nom_pathologie FOREIGN KEY (nom_pathologie) REFERENCES "Pathologie"(nom_pathologie),
    CONSTRAINT CompKey_path_spe PRIMARY KEY (nom_pathologie,specialite)
);


CREATE TABLE "Pro_sante" (
	nom VARCHAR  NOT NULL,
	a_mail VARCHAR,
	n_telephone VARCHAR
);

CREATE TABLE "Medecin" (
	inami VARCHAR PRIMARY KEY ,
	specialite VARCHAR NOT NULL,
    CONSTRAINT fk_specialite FOREIGN KEY (specialite) REFERENCES "Specialite"(specialite)
) INHERITS ("Pro_sante");

CREATE TABLE "Pharmacien" (
	inami VARCHAR PRIMARY KEY
) INHERITS ("Pro_sante");

CREATE TABLE "Patient" (
	n_niss CHAR (16) PRIMARY KEY,
	nom VARCHAR  NOT NULL,
	prenom VARCHAR NOT NULL,
	genre CHAR (1) NOT NULL,
	date_naissance DATE,
	a_mail VARCHAR,
	pwd VARCHAR,
	n_telephone VARCHAR,
	inami VARCHAR,
	n_inami_pha VARCHAR,
	CONSTRAINT fk_inami FOREIGN KEY (inami) REFERENCES "Medecin"(inami),
	CONSTRAINT fk_n_inami_pha FOREIGN KEY (n_inami_pha) REFERENCES "Pharmacien"(inami)
);

CREATE TABLE "Medicament" (
	nom_commercial VARCHAR PRIMARY KEY,
	dci VARCHAR,
	nom_sys_ana VARCHAR,
	CONSTRAINT fk_nom_sys_ana FOREIGN KEY (nom_sys_ana) REFERENCES "Systeme_ana"(nom_sys_ana)
);

CREATE TABLE "Conditionnement" (
	quantite INTEGER PRIMARY KEY
);

CREATE TABLE "Medicament_conditionnement" (
	quantite INTEGER,
    nom_commercial VARCHAR,
	CONSTRAINT fk_nom_commercial FOREIGN KEY (nom_commercial) REFERENCES "Medicament"(nom_commercial),
    CONSTRAINT fk_quantite FOREIGN KEY (quantite) REFERENCES "Conditionnement"(quantite),
    CONSTRAINT CompKey_cond_nom_quantite PRIMARY KEY (quantite,nom_commercial)
);

CREATE TABLE "Diagnostic"(
	date_diagnostic date,
	n_niss VARCHAR,
	nom_pathologie VARCHAR,
	CONSTRAINT CompKey_diag_niss_path PRIMARY KEY (date_diagnostic,n_niss,nom_pathologie),
	CONSTRAINT fk_n_niss FOREIGN KEY (n_niss) REFERENCES "Patient"(n_niss),
	CONSTRAINT fk_nom_pathologie FOREIGN KEY (nom_pathologie) REFERENCES "Pathologie"(nom_pathologie)
);

CREATE TABLE "Prescription"(
	id SERIAL PRIMARY KEY,
	date_prescription date,
	n_niss VARCHAR,
	nom_commercial VARCHAR,
	inami_med VARCHAR,
	CONSTRAINT fk_n_niss FOREIGN KEY (n_niss) REFERENCES "Patient"(n_niss),
	CONSTRAINT fk_nom_commercial FOREIGN KEY (nom_commercial) REFERENCES "Medicament"(nom_commercial),
	CONSTRAINT fk_inami FOREIGN KEY (inami_med) REFERENCES "Medecin"(inami)
);

CREATE TABLE "Traitement"(
	id SERIAL PRIMARY KEY,
	date_vente date,
	duree_traitement VARCHAR,
	id_prescription SERIAL,
	CONSTRAINT fk_id_prescription FOREIGN KEY (id_prescription) REFERENCES "Prescription"(id)
);

""")

conn.commit()

print('Tables created successfully')
conn.close()