from flask import Flask, request
from send_message import send_twilio_message
from waitress import serve
import argparse

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

    send_twilio_message(msg, sender_phone, file_path)
    return 'Hello', 201
    

if __name__ == "__main__":
    if args['prod']: 
        print("Running on port 3001 in production mode ..")
        serve(app, host='0.0.0.0', port=3001)
    else: 
        app.run(debug=True, port=3001)
    