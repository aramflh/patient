from bs4 import BeautifulSoup

def sanitize(data):
    return data.strip().replace("'","''").lower()

def dml_diag():
    sql_diag = []
    with open('data/diagnostiques.xml', 'r', encoding='utf8') as f:
        data = f.read()

    soup = BeautifulSoup(data, 'lxml')
    diags = []

    for diag in soup.find_all("diagnostique"):
        diags.append(
            {
            "niss":sanitize(diag.niss.string),
            "date_diagnostic":sanitize(diag.date_diagnostic.string),
            #"naissance":diag.naissance.string,
            "pathology":sanitize(diag.pathology.string),
            #"specialite":diag.specialite.string
            }
            )


    for diag in diags:
        sql_diag.append(f"""
        INSERT INTO "Diagnostic" (date_diagnostic, n_niss, nom_pathologie)
        VALUES ('{diag['date_diagnostic']}','{diag['niss']}','{diag['pathology']}');
        """)
    return sql_diag
