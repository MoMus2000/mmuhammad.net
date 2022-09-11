import requests
import os
from datetime import datetime, timedelta
import sqlite3
import json

FOREX_API_TOKEN = os.environ.get("FOREX_API_TOKEN")
METALS_API_TOKEN = os.environ.get("METAL_API_TOKEN")

con = sqlite3.connect("../db/lenslocked_dev.db")

today_date = datetime.today()
today_date = today_date.strftime('%Y-%m-%d')


FOREX_API_TOKEN = "59pfif605nZ54x5c77xrul4l3BUb7Ht2"
METAL_API_TOKEN = "nyqm4mk9y089cvi17m07x19khfixw06nyo7e2d3pmea623m4kjkn9dk1n75a"

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