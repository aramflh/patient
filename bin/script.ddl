-- Database: BD

CREATE TABLE "Systeme_ana"(
                              nom_sys_ana VARCHAR PRIMARY KEY,
                              specialite VARCHAR
);


CREATE TABLE "Pathologie" (
                              nom_pathologie VARCHAR PRIMARY KEY,
                              nom_sys_ana VARCHAR,
                              CONSTRAINT fk_sys_ana FOREIGN KEY (nom_sys_ana) REFERENCES "Systeme_ana"(nom_sys_ana)
);

CREATE TABLE "Pro_sante" (
                             nom VARCHAR  NOT NULL,
                             prenom VARCHAR NOT NULL,
                             a_mail VARCHAR,
                             n_telephone VARCHAR
);

CREATE TABLE "Medecin" (
                           inami VARCHAR PRIMARY KEY ,
                           specialite VARCHAR NOT NULL
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
                           a_mail VARCHAR UNIQUE ,
                           pwd VARCHAR,
                           n_telephone VARCHAR,
                           n_inami_med VARCHAR,
                           n_inami_pha VARCHAR,
                           CONSTRAINT fk_n_inami_med FOREIGN KEY (n_inami_med) REFERENCES "Medecin"(inami),
                           CONSTRAINT fk_n_inami_pha FOREIGN KEY (n_inami_pha) REFERENCES "Pharmacien"(inami)
);

CREATE TABLE "Medicament" (
                              nom_medic VARCHAR PRIMARY KEY,
                              dci VARCHAR,
                              conditionnement VARCHAR,
                              nom_pathologie VARCHAR,
                              CONSTRAINT fk_nom_pathologie FOREIGN KEY (nom_pathologie) REFERENCES "Pathologie"(nom_pathologie)
);

CREATE TABLE "Dossier_med"(
                              date_diagnostic date,
                              n_niss VARCHAR,
                              nom_pathologie VARCHAR,
                              CONSTRAINT CompKey_diag_niss_path PRIMARY KEY (date_diagnostic,n_niss,nom_pathologie),
                              CONSTRAINT fk_n_niss FOREIGN KEY (n_niss) REFERENCES "Patient"(n_niss),
                              CONSTRAINT fk_nom_pathologie FOREIGN KEY (nom_pathologie) REFERENCES "Pathologie"(nom_pathologie)
);

CREATE TABLE "Prescription"(
                               date_emission date,
                               duree_traitement VARCHAR,
                               n_niss VARCHAR,
                               nom_medic VARCHAR,
                               n_inami_med VARCHAR,
                               n_inami_pha VARCHAR,
                               CONSTRAINT CompKey_date_inamim_med_niss PRIMARY KEY (date_emission,n_niss,nom_medic,n_inami_med),
                               CONSTRAINT fk_n_niss FOREIGN KEY (n_niss) REFERENCES "Patient"(n_niss),
                               CONSTRAINT fk_nom_medic FOREIGN KEY (nom_medic) REFERENCES "Medicament"(nom_medic),
                               CONSTRAINT fk_n_inami_med FOREIGN KEY (n_inami_med) REFERENCES "Medecin"(inami),
                               CONSTRAINT fk_n_inami_pha FOREIGN KEY (n_inami_pha) REFERENCES "Pharmacien"(inami)
);

CREATE TABLE "Traitement"(
                             date_debut date,
                             duree_traitement VARCHAR,
                             n_niss VARCHAR,
                             nom_medic VARCHAR,
                             n_inami_med VARCHAR,
                             n_inami_pha VARCHAR,
                             CONSTRAINT CompKey_date_inamip_inamim_med_niss PRIMARY KEY (date_debut,n_niss,nom_medic,n_inami_pha,n_inami_med),
                             CONSTRAINT fk_n_niss FOREIGN KEY (n_niss) REFERENCES "Patient"(n_niss),
                             CONSTRAINT fk_nom_medic FOREIGN KEY (nom_medic) REFERENCES "Medicament"(nom_medic),
                             CONSTRAINT fk_n_inami_med FOREIGN KEY (n_inami_med) REFERENCES "Medecin"(inami),
                             CONSTRAINT fk_n_inami_pha FOREIGN KEY (n_inami_pha) REFERENCES "Pharmacien"(inami)
);