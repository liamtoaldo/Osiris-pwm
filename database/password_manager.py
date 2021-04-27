import database


def create_new_user(username, mail, password):
    query = database.query("pws.db", "users", "username", username)

    if len(query) == 0:
        database.insert("pws.db", "users", {'username':  username, 'mail': mail, 'password': password})
        database.create_table("pws.db", username, {'website': 'text', 'mail': 'text', 'password': 'text'})
        print("You're now registered!")
        main(username)
    else:
        print("You're already registered!")


def login(username, password):
    query = database.query("pws.db", "users", "username", username)

    if len(query) > 0:
        print("You're now logged in!")
        main(username)
    else:
        print("You're not registered!")


def create_new_password(username, website, mail, password):
    database.insert("pws.db", username, {'website': website, 'mail': mail, 'password': password})
    print("Succesfuly saved data!\n")


def show_all(username):
    database.visualize("pws.db", username)


def query(username, search, value):
    query = database.query("pws.db", username, search, value)
    for item in query:
        print(item)

