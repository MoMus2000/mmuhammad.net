import pandas as pd
import requests
import phonenumbers
import time
import os
from twilio.rest import Client
from datetime import date



# account = os.environ.get("TWILIO_ACCOUNT")
# token = os.environ.get("TWILIO_TOKEN")
# twilio_phone = os.environ.get("TWILIO_PHONE")

def api_request(client, msg, sender, reciever):
    try:
        message = client.messages.create(
        body=msg, from_=sender, to=reciever
        )
        print(message.body, message.price, reciever, message.error_code, message.error_message)
    except Exception as e:
        raise Exception(e)
    time.sleep(1)

def send_twilio_message(client, msg, sender, file_path):
    df = pd.read_excel(f"../temp/{file_path}")
    error = 0
    for i in range(0, len(df['Phone Number'])):
        try:
            ph_num = format_number(df['Phone Number'].iloc[i])
            api_request(client, str(msg), str(sender), ph_num)
        except Exception as e:
            print(e)
            error += 1
        if(error >= len(df)):
            raise Exception("Something seems to be going wrong, all requests failed")
    return error, len(df)

def get_total_message_length(client, twilio_phone):
    messages = client.messages.list(date_sent=date.today(), 
    from_= twilio_phone)
    return len(messages)


def format_number(phonenumber):
    return phonenumbers.format_number(phonenumbers.parse(str(phonenumber), 'CA'),
    phonenumbers.PhoneNumberFormat.NATIONAL)

def get_message_balance(client):
    return float(client.api.v2010.balance.fetch().balance)


if __name__ == "__main__":
    pass