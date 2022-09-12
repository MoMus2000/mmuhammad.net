import requests
import os
from datetime import datetime, timedelta
import sqlite3
import json

FOREX_API_TOKEN = os.environ.get("FOREX_API_TOKEN")
METALS_API_TOKEN = os.environ.get("METAL_API_TOKEN")
COMMODITY_API_TOKEN = os.environ.get("COMMODITY_API_TOKEN")

con = sqlite3.connect("../db/lenslocked_dev.db")

today_date = datetime.today()
today_date = today_date.strftime('%Y-%m-%d')

url = f"https://api.polygon.io/v2/aggs/ticker/C:USDPKR/range/1/day/{today_date}/\
{today_date}?adjusted=true&sort=asc&limit=120&apiKey={FOREX_API_TOKEN}"

data = requests.get(url)

if data.status_code == 200:
    response = json.loads(data.content)
    open_price = float(response['results'][-1]["o"])
    close_price = float(response['results'][-1]["c"])
    cur = con.cursor()
    data = (
        [today_date, "OPEN_USD", open_price],
        [today_date, "CLOSE_USD", close_price]
    )
    try:
        cur.executemany("""INSERT INTO monitors 
        (date, metric, value) VALUES (?, ?, ?)""", data)
        con.commit()
    except sqlite3.IntegrityError:
        print("DUPLICATE insertion ", data)

else:
    print(data.status_code)

url = f"https://metals-api.com/api/latest?access_key={METAL_API_TOKEN}\
&base=PKR&symbols=XAU%2CXAG%2CSTEEL-SC%2CSTEEL-RE%2CSTEEL-HR%09"

data = requests.get(url)

if data.status_code == 200:
    response = json.loads(data.content)
    
    china_hot_roll_steel = float(response['rates']["STEEL-HR"])
    rebar_steel_turkey = float(response['rates']["STEEL-RE"])
    steel_sc = float(response['rates']["STEEL-SC"])
    gold = float(response['rates']["XAU"])
    silver = float(response['rates']["XAG"])

    cur = con.cursor()
    
    data = (
        [today_date, "CHINA_HOT_ROLL", china_hot_roll_steel],
        [today_date, "TURKEY_REBAR", rebar_steel_turkey],
        [today_date, "TURKEY_SC", steel_sc],
        [today_date, "GOLD", gold],
        [today_date, "SILVER", silver],
    )
    try:
        cur.executemany("""INSERT INTO monitors 
        (date, metric, value) VALUES (?, ?, ?)""", data)
        con.commit()
    except sqlite3.IntegrityError:
        print("DUPLICATE insertion ", data)

else:
    print(data.status_code)


url = f"https://www.commodities-api.com/api/latest?access_key={COMMODITY_API_TOKEN}&base=USD&symbols=WTIOIL%2CBRENTOIL"
response = requests.get(url)
if response.status_code == 200:
    rates = json.loads(response.content)['data']['rates']
    BRENTOIL = 1/float(rates['BRENTOIL'])
    WTIOIL = 1/float(rates['WTIOIL'])

    cur = con.cursor()

    data = (
        [today_date, "BRENTOIL", BRENTOIL],
        [today_date, "WTIOIL", WTIOIL],
    )
    try:
        cur.executemany("""INSERT INTO monitors 
        (date, metric, value) VALUES (?, ?, ?)""", data)
        con.commit()
    except sqlite3.IntegrityError:
        print("DUPLICATE insertion ", data)
else:
    print(response.status_code)

