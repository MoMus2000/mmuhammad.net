from datetime import datetime
import os
from twilio.rest import Client




account_sid = ""
auth_token = ""
client = Client(account_sid, auth_token)

def get_total_message_length():
    messages = client.messages.list(
                                date_sent=datetime.now(),
                                from_='+16476916189',
                            )
    return messages


if __name__ == "__main__":
    print(messages)