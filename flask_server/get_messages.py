from datetime import datetime
import os
from twilio.rest import Client


account_sid = 'AC7c1d4068211dfa361cfc6be3a3af78a8'
auth_token = 'ebd2cd21905321b9fd1612d6b4682338'
client = Client(account_sid, auth_token)

messages = client.messages.list(
                               date_sent=datetime.now(),
                               from_='+16476916189',
                           )

if __name__ == "__main__":
    print(messages)