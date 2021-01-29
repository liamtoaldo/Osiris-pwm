# UMSR Copyright (c) Alphabet Software, Inc
# Tutti i diritti riservati


import sqlite3


def get_connection(filename: str, debug=False):
    """
    Creates a connection to the .db file if it exists
    Creates the .db file if it doesn't exist

    :param filename: Name of the file you'd like to use to store the DB

    :return: connection, cursor
    """

    conn = sqlite3.connect(str(filename))
    c    = conn.cursor()
    return conn, c


def create_table(filename: str, table_name: str, fields: dict, debug=False) -> bool:
    """
    Creates an SQL table inside a given database

    :param filename: .db filename (String)
    :param table_name: Name of the table you'd like to create (String)
    :param fields: List of variable fields and types (Dict {'var_name': 'type'})

    :return: Bool
    """

    conn, cursor = get_connection(filename)
    command      = f"""CREATE TABLE {table_name} (\n"""
    fields       = str(fields)
    fields      = eval(fields)
    var_list     = fields.copy()

    for field in fields:
        if len(var_list) > 1:
            command += f"{field} {fields[field]},\n"
        else:
            command += f"{field} {fields[field]}\n)"

        var_list.pop(field)

    try:
        cursor.execute(command)
        if debug: print("Successfuly created table!")
        return True
    except Exception as e:
        if debug: print(f"Error while creating database table '{table_name}' for database '{filename}'\n\tException: {e}")
        return False


def query(filename: str, table: str, field_name: str, filed_value, debug=False) -> list:
    """
    Returns a list of tuples of values that contains the specified value

    :param filename: .db filename (String)
    :param table: Name of the table (String)
    :param field_name: Name of the field in the table (String)
    :param field_value: Value that the query should look for in the field

    :return: List
    """

    conn, cursor = get_connection(filename)
    table        = str(table)
    field_name   = str(field_name)

    if type(filed_value) == str:
        cursor.execute(f"SELECT * FROM {table} WHERE {field_name}='{filed_value}'")
    elif type(filed_value) == (int or float):
        cursor.execute(f"SELECT * FROM {table} WHERE {field_name}={filed_value}")

    return cursor.fetchall()


def insert(filename: str, table: str, values: dict, debug=False) -> bool:
    """
    Adds an entry to a given table

    :param filename: .db file name (String)
    :param table: Name of the table (String)
    :param values: Dictionary of values (Dict)

    :return: Bool
    """

    conn, cursor = get_connection(filename)
    table        = str(table)
    values       = str(values)
    values       = eval(values)
    _values      = []

    for value in values:
        _values.append(f":{value}")

    _values      = tuple(_values)
    command      = f"INSERT INTO {table} VALUES {_values}"
    
    for value in values:
        command = command.replace(f"':{value}'", f":{value}")

    if command[len(command) - 2] == ",":
        _command = list(command)
        _command.reverse()
        _command.pop(1)
        _command.reverse()
        command  = ""

        for item in _command:
            command += item

    try:
        cursor.execute(command, values)
        conn.commit()
        if debug: print(f"values {values} succesfuly added to {table}!")
        return True
    except Exception as e:
        if debug: print(f"Something went wrong while trying to insert {values} in {table}!\n\tException: {e}")
        return False


def visualize(filename: str, table: str, debug=False) -> bool:
    """
    Formats the contents of a db table using the texttable package

    :param filename: .db file name (String)
    :param table: Name of the table to plot (String)

    :return: Bool
    """

    from texttable import Texttable

    conn, cursor   = get_connection(filename)
    
    try:
        cursor.execute(f"SELECT * FROM {table}")
    except Exception as e:
        if debug: print(f"Something went wrong while trying to access table '{table}'\n\tException: {e}")

    table_elements = cursor.fetchall()

    arr            =   [[]]

    if len(table_elements) > 0:
        for item in table_elements[0]:
            arr[0].append("")
    else:
        print("This is an empty table")
        return False

    t              = Texttable()
    allign         = []
    vallign        = []

    for item in arr[0]:
        allign.append("l")
        vallign.append("m")

    t.set_cols_align(allign)
    t.set_cols_valign(vallign)

    for item in table_elements:
        arr.append(list(item))

    t.add_rows(arr)
    print(t.draw())
    return True


def update(filename: str, table: str, condition: str, attribute: str, new_value, debug=False) -> bool:
    """
    Edits the value at the given table, condition and attribute

    :param filename: .db file name (String)
    :param table: Name of the table (String)
    :param condition: SQL 'WHERE' condition (String: "colum = value")

    :return: Bool
    """

    conn, cursor = get_connection(filename)
    table        = str(table)
    condition        = str(condition)
    attribute    = str(attribute)
    
    try:
        command = f"""
UPDATE {table}
SET {attribute} = ?
WHERE {condition};
"""

        cursor.execute(command, (new_value,))
        conn.commit()
        if debug: print("Value updated successfuly!")
        return True
    except Exception as e:
        if debug: print(f"Unable to update value {attribute} in table {table}\n\tException: {e}")
        return False


def delete_entry(filename: str, table: str, condition: str, debug=False) -> bool:
    """
    Deletes an entry from a given table with a given condition

    :param filename: .db file name (String)
    :param table: Name of the table (String)
    :param condition: SQL 'WHERE' condition (String: "colum = value")

    :return: Bool
    """

    conn, cursor = get_connection(filename)

    try:
        cursor.execute(f"DELETE FROM {table} WHERE {condition}")
        conn.commit()
        if debug: print("Value deleted succesfuly!")
        return True
    except Exception as e:
        if debug: print(f"Something went wrong while trying to delete a value from {table} with condition {condition}\n\tException: {e}")
        return False


def delete_table(filename: str, table: str, debug=False) -> bool:
    """
    Deletes the specified table

    :param filename: .db file name (String)
    :param table: Name of the table you'd like to delete (String)

    :return: Bool
    """

    conn, cursor = get_connection(filename)
    table = str(table)

    try:
        cursor.execute(F"DROP TABLE {table}")
        if debug: print("Table delete succesfuly!")
        return True
    except Exception as e:
        if debug: print(f"Something went wrong while trying to delete table {table} from DB {filename}\n\tException: {e}")
        return False
