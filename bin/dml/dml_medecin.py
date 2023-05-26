from bs4 import BeautifulSoup


def sanitize(data):
    return data.strip().replace("'","''").lower()


def dml_medecin():
    sql_medecin = []
    with open('data/medecins.xml', 'r', encoding="utf8") as f:
        data = f.read()

    soup = BeautifulSoup(data, 'lxml')
    medecins = []

    for medecin in soup.find_all("medecin"):
        row = {
                "nom":sanitize(medecin.nom.string),
                "inami":sanitize(medecin.inami.string),
                "mail":sanitize(medecin.mail.string if medecin.mail.string != None else ''),
                "specialite":sanitize(medecin.specialite.string),
                "telephone":sanitize(medecin.telephone.string)
            }
        if row not in medecins:
            medecins.append(row)

    for medecin in medecins:
        sql_medecin.append(f"""
        INSERT INTO "Medecin" (nom, a_mail, n_telephone, inami, specialite)
        VALUES ('{medecin['nom']}','{medecin['mail']}','{medecin['telephone']}','{medecin['inami']}','{medecin['specialite']}');
        """)
    return sql_medecin


