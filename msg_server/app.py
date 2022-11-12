from flask import Flask, request, jsonify
from message import send_twilio_message, \
get_total_message_length, \
get_message_balance, api_request
from db import get_twilio_number, write_total_message, get_account_id, get_account_token
from waitress import serve
import argparse
import json
from dotenv import load_dotenv
import os
from twilio.rest import Client
from datetime import datetime, timedelta
from markovitz import MarkovitzModel

app = Flask(__name__)
parser = argparse.ArgumentParser(description="Just an example",
                                 formatter_class=argparse.ArgumentDefaultsHelpFormatter)
parser.add_argument("-prod", "--prod", action="store_true", help="server status")
args = vars(parser.parse_args())

@app.route("/api/v1/fmb/send_message", methods=['POST'])
def index():
    req = request.get_json()
    
    msg = req['message']
    sender_name = req['senderName']
    sender_phone = req['sender']
    file_path = req['fileName']
    twilio_phone = req['twilioPhone']
    try:
        error, _ = send_twilio_message(msg, twilio_phone, file_path)
    except Exception as e:
        print(e)
        return e, 500 
    return f'Created, total failed requests {error}!', 201

@app.route("/api/v1/fmb/get_history", methods=["GET"])
def get_history():
    try:
        length = get_total_message_length()
        resp = jsonify(length = length)
    except Exception as e:
        print(e)
        resp = jsonify(length = 0)
    return resp

@app.route("/api/v1/fmb/app_balance", methods=["GET"])
def get_balance_fmb():
    try:
        balance = get_message_balance()
        resp = jsonify(balance = balance)
    except Exception as e:
        print(e)
        resp = jsonify(balance = 0)
    return resp


@app.route("/api/v1/sms/balance", methods=["POST"])
def get_balance():
    req = request.get_json()
    email = req['email']
    try:
        account = get_account_id(email)
        token = get_account_token(email)

        client = Client(account, token)

        balance = get_message_balance(client)

        resp = jsonify(Data = round(balance, 3))

    except Exception as e:
        print(e)
        resp = jsonify(Data = 0)

    return resp

@app.route("/api/v1/sms/total_cost", methods=["POST"])
def get_total_cost():
    req = request.get_json()
    email = req['email']
    try:
        account = get_account_id(email)
        token = get_account_token(email)
        client = Client(account, token)

        total_cost = 0
        for msg in client.messages.list():
            if msg.price != None:
                total_cost += float(msg.price)*-1

        resp = jsonify(Data = round(total_cost, 3))

    except Exception as e:
        print(e)
        resp = jsonify(Data = 0)
    
    return resp, 201

@app.route("/api/v1/sms/single_sms", methods=["POST"])
def send_single_sms():
    req = request.get_json()
    msg = req['message']
    sender_name = req['senderName']
    sender_phone = req['sender']
    email = req['email']

    try:
        account = get_account_id(email)
        token = get_account_token(email)

        client = Client(account, token)

        twilio_phone = get_twilio_number(email)
        
        api_request(client, msg, twilio_phone, sender_phone)
        write_total_message(email, 1)
    
    except Exception as e:
        print(e)
        return 'ok', 500

    return 'ok', 201

@app.route("/api/v1/sms/bulk_sms", methods=["POST"])
def send_bulk_sms():
    req = request.get_json()
    msg = req['message']
    sender_name = req['senderName']
    sender_phone = req['sender']
    email = req['email']
    twilio_phone = get_twilio_number(email)
    file_path = req['fileName']

    try:
        account = get_account_id(email)
        token = get_account_token(email)

        client = Client(account, token)

        error, total = send_twilio_message(client, msg, twilio_phone, file_path)

    except Exception as e:
        print(e)
        return e, 500 

    write_total_message(email, total-error)

    return 'ok', 201

@app.route("/api/v1/sms/total_messages_today", methods=["POST"])
def get_total_messages_today():
    req = request.get_json()
    email = req['email']

    try:
        account = get_account_id(email)
        token = get_account_token(email)
        client = Client(account, token)
        twilio_phone = get_twilio_number(email)
        messages = client.messages.list(date_sent=datetime.today(), 
        from_= twilio_phone)
        resp = jsonify(Data = len(messages))

    
    except Exception as e:
        print(e)
        resp = jsonify(Data = 0)
    
    return resp, 201


@app.route("/api/v1/sms/total_messages", methods=["POST"])
def get_total_messages():
    req = request.get_json()
    email = req['email']

    try:
        account = get_account_id(email)
        token = get_account_token(email)
        client = Client(account, token)
        twilio_phone = get_twilio_number(email)
        messages = client.messages.list(from_= twilio_phone)
        resp = jsonify(Data = len(messages))
    
    except Exception as e:
        print(e)
        resp = jsonify(Data = 0)
    
    return resp, 201

@app.route("/api/v1/optimize/crunch", methods=["POST"])
def crunchMarkovitz():
    req = request.get_json()

    try:
        tickers = req["Tickers"].upper()
        amount = req["Amount"]
        time = req["TimeHorizon"]
        print(time)
        print(amount)
        print(tickers.split(" "))
        markovitz = MarkovitzModel(time, "1d", tickers.split(" "))
        stats, stock = markovitz.run()
        print(stats, stock)
        resp = jsonify(returns = stats[0], volatility=stats[1],
        sharpe=stats[2],
        stock = stock)
    except Exception as e:
        print(e)
        resp = jsonify(stats=0, stock=0)

    return resp, 201


if __name__ == "__main__":
    if args['prod']: 
        print("Running on port 3001 in production mode ..")
        serve(app, host='0.0.0.0', port=3001)
    else: 
        # load_dotenv()
        app.run(debug=True, port=3001)
    