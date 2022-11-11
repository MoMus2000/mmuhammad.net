import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
import statsmodels.api as sm
from statsmodels.tsa.stattools import adfuller
import json
import os
import requests
import math
import sqlite3
from datetime import datetime

ALPHA_API_TOKEN = os.environ.get("ALPHA_API_TOKEN")
con = sqlite3.connect("../db/lenslocked_dev.db")

def calculate_return(ticker):
    url = f'https://www.alphavantage.co/query?function=TIME_SERIES_WEEKLY&symbol={ticker}&apikey={ALPHA_API_TOKEN}'
    r = requests.get(url)
    df = pd.DataFrame()
    
    date = []
    close_prices = []
    
    if r.status_code == 200:
        data = r.json()
        keys = sorted(list(data['Weekly Time Series'].keys()))
        for key in keys:
            close_price = data['Weekly Time Series'][key]['4. close']
            close_prices.append(float(close_price))
            date.append(key)
    
    df['Close'] = np.array(close_prices)
    df['Date'] = np.array(date)

    df = df.set_index("Date")
    df.index = pd.to_datetime(df.index)

    df_ret = df['Close'].resample('W').last().pct_change().dropna() #Get weekly returns
    return df_ret

def run_adfuller(returns):
    if math.isclose(adfuller(returns.dropna())[1], 0.0):
        print("valid P test")
    return

def train(returns):
    #Fit the model
    mod_kns = sm.tsa.MarkovRegression(returns.dropna(), k_regimes=3, trend='n', switching_variance=True)
    res_kns = mod_kns.fit()
    print(res_kns.summary())
    return res_kns

def final_probability_for_what_we_have(trained_model):
    df = trained_model.smoothed_marginal_probabilities
    df = df.tail(1)
    low_volatility = df[0][0]
    medium_volatility = df[1][0]
    high_volatility = df[2][0]
    return low_volatility, medium_volatility, high_volatility

def write_to_db(date, data, ticker):
    cur = con.cursor()
    
    data = (
        [date, f"LOW_VOL_PROB_{ticker}", data[0]],
        [date, f"MED_VOL_PROB_{ticker}", data[1]],
        [date, f"HIGH_VOL_PROB_{ticker}", data[2]],
    )
    try:
        cur.executemany("""INSERT INTO monitors 
        (date, metric, value) VALUES (?, ?, ?)""", data)
        con.commit()
        print("Wrote regime change probabilities")

    except sqlite3.IntegrityError:
        print("DUPLICATE insertion ", data)

def pipeline(ticker):
    print("Fetching weekly data")
    returns = calculate_return(ticker)
    print("Running Adfuller test")
    run_adfuller(returns)
    date = str(returns.tail(1).index[0]).split(" ")[0]
    print("Training MarkovRegressive model")
    model = train(returns)
    print("Generating probs")
    l, m, h = final_probability_for_what_we_have(model)
    print("Write to db")
    write_to_db(date, (l, m, h), ticker)

if __name__ == "__main__":
    print("RUN for SPY")
    pipeline("SPY")
    print("RUN for Canadian Housing REIT")
    pipeline("XRE.TO")