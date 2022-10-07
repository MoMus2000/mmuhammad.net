from flask import Flask, request, jsonify
from message import send_twilio_message, \
get_total_message_length, \
get_message_balance
from waitress import serve
import argparse
import json


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
        error = send_twilio_message(msg, twilio_phone, file_path)
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
def get_balance():
    try:
        balance = get_message_balance()
        resp = jsonify(balance = balance)
    except Exception as e:
        print(e)
        resp = jsonify(balance = 0)
    return resp
    

if __name__ == "__main__":
    if args['prod']: 
        print("Running on port 3001 in production mode ..")
        serve(app, host='0.0.0.0', port=3001)
    else: 
        app.run(debug=True, port=3001)
    