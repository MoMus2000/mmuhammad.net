import pandas as pd
import requests
import phonenumbers
import time
import os


account = os.environ.get("TWILIO_ACCOUNT")
token = os.environ.get("TWILIO_TOKEN")

def api_request(msg, sender, reciever):
    headers = {
        'Content-Type': 'application/x-www-form-urlencoded',
    }
    data = f'Body={msg}&From={sender}&To={reciever}'
    response = requests.post\
    (f'https://api.twilio.com/2010-04-01/Accounts/{account}/Messages.json',
     headers=headers, data=data, auth=(account, token)
    )
    if response.status_code != 201:
        print(response.text)
    time.sleep(2)

def send_twilio_message(msg, sender, file_path):
    df = pd.read_excel(f"../temp/{file_path}")
    for i in range(0, len(df['Phone Number'])):
        ph_num = format_number(df['Phone Number'].iloc[i])
        api_request(str(msg), str(sender), ph_num)

def format_number(phonenumber):
    return phonenumbers.format_number(phonenumbers.parse(str(phonenumber), 'CA'),
    phonenumbers.PhoneNumberFormat.NATIONAL)


if __name__ == "__main__":
    send_twilio_message("Kevin Bacon", "+13862515211")