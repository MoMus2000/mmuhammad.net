import sqlite3
from datetime import datetime

def get_twilio_number(email):
    con = sqlite3.connect("../sms.mmuhammad.net/db/sms_mmuhammad.db")
    cursor = con.execute(f"SELECT twilio_phone FROM users WHERE email = (?)", (email,))
    res = cursor.fetchone()
    if res == None:
        return None
    con.close()
    return res[0]

def write_total_message(email, message_count):
    now = datetime.now()
    con = sqlite3.connect("../sms.mmuhammad.net/db/sms_mmuhammad.db")
    c = con.cursor()
    c.execute(f"INSERT INTO sms_metrics (created_at, email, metric_name, metric_value) VALUES (?, ?, ?, ?) ", (now.strftime("%d-%m-%Y %H:%M:%S"), email, "SMS_SENT", message_count,))
    con.commit()
    con.close()

def get_account_id(email):
    con = sqlite3.connect("../sms.mmuhammad.net/db/sms_mmuhammad.db")
    cursor = con.execute("SELECT account_id from users WHERE email = ?", (email,))
    res = cursor.fetchone()
    if res == None:
        return None
    con.close()
    return res[0]

def get_account_token(email):
    con = sqlite3.connect("../sms.mmuhammad.net/db/sms_mmuhammad.db")
    cursor = con.execute("SELECT account_token from users WHERE email = ?", (email,))
    res = cursor.fetchone()
    if res == None:
        return None
    con.close()
    return res[0]


if __name__ == "__main__":
    write_total_message("muhammadmustafa4000@gmail.com", 20)