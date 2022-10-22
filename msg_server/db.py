import sqlite3

def get_twilio_number(email):
    con = sqlite3.connect("../sms.mmuhammad.net/db/sms_mmuhammad.db")
    cursor = con.execute(f"SELECT twilio_phone FROM users WHERE email = (?)", (email,))
    res = cursor.fetchone()
    if res == None:
        return None
    return res[0]