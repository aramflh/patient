from bs4 import BeautifulSoup


def sanitize(data):
    return data.strip().replace("'","''").lower()

def dml_pha():
    sql_pha = []
    with open('data/pharmaciens.xml', 'r', encoding='utf8') as f:
        data = f.read()

    soup = BeautifulSoup(data, 'lxml')
    pharmaciens = []

    for pharmacien in soup.find_all("pharmacien"):
        row = {
                "nom":sanitize(pharmacien.nom.string),
                "inami":sanitize(pharmacien.inami.string),
                "mail":sanitize(pharmacien.mail.string if pharmacien.mail.string != None else ''),
                "telephone":sanitize(pharmacien.tel.string if pharmacien.tel.string != None else '')
            }
        if row not in pharmaciens:
            pharmaciens.append(row)


    for pharmacien in pharmaciens:
        sql_pha.append(f"""
        INSERT INTO "Pharmacien" (nom, a_mail, n_telephone, inami)
        VALUES ('{pharmacien['nom']}','{pharmacien['mail']}','{pharmacien['telephone']}','{pharmacien['inami']}');
        """)
    
    return sql_pha
