import requests 
from bs4 import BeautifulSoup
import pandas as pd
from tqdm import tqdm
import time
from fuzzywuzzy import fuzz
import re
import numpy as np
import statistics
import sqlite3
import scipy.stats
import matplotlib.pyplot as plt
from distfit import distfit
import math
from datetime import datetime

today_date = datetime.today()
today_date = today_date.strftime('%Y-%m-%d')

def write_to_db(output, prefix):
    con = sqlite3.connect("../db/lenslocked_dev.db")
    cur = con.cursor()
    data = (
        [today_date, f"{prefix}_MEAN", str(output[0])],
        [today_date, f"{prefix}_MIN", str(output[1])],
        [today_date, f"{prefix}_MEDIAN", str(output[2])],
        [today_date, f"{prefix}_MAX", str(output[3])],
        [today_date, f"{prefix}_LIKELY_PRICE", str(output[4])]
    )
    try:
        cur.executemany("""INSERT INTO monitors 
        (date, metric, value) VALUES (?, ?, ?)""", data)
        con.commit()

    except sqlite3.IntegrityError:
        print("DUPLICATE insertion ", data)

def model_distribution(price_arr):
    lin = np.linspace(min(price_arr), max(price_arr))
    dist = distfit()
    dist.fit_transform(price_arr)
    maxVal = float('-inf')
    maxindex = -1
    for idx, prob in enumerate(dist.predict(lin)['y_proba']):
        if maxVal < prob:
            maxVal = prob
            maxindex = idx
    
    print(f"Maximum likelihood of price at which you find a basement: CAD ${int(lin[maxindex])}")
    mean = statistics.mean(price_arr)
    minimum = min(price_arr)
    maximum = max(price_arr)
    median = statistics.median(price_arr)
    
    print("MIN :", minimum)
    print("MAX :", maximum)
    print("MEAN:", mean)
    print("MEDIAN", median)

    return mean, minimum, median, maximum, int(lin[maxindex])

def post_processing(dataset):
    print(len(dataset))
    dataset.drop_duplicates(inplace=True, keep="first") 
    print(len(dataset))   
    price_arr = []
    for i in range(0, len(dataset)):
        city = dataset['Location'].iloc[i]
        title = dataset['Title'].iloc[i]
        price = dataset['Price'].iloc[i]
        if "$" in price:
            price = int(currency_parser(price))
            price_arr.append(price)
    
    print(len(price_arr))
    price_arr = reject_outliers(np.array(price_arr))
    print(len(price_arr))

    return price_arr

def currency_parser(cur_str):
    # Remove any non-numerical characters
    # except for ',' '.' or '-' (e.g. EUR)
    cur_str = re.sub("[^-0-9.,]", '', cur_str)
    # Remove any 000s separators (either , or .)
    cur_str = re.sub("[.,]", '', cur_str[:-3]) + cur_str[-3:]

    if '.' in list(cur_str[-3:]):
        num = float(cur_str)
    elif ',' in list(cur_str[-3:]):
        num = float(cur_str.replace(',', '.'))
    else:
        num = float(cur_str)

    return np.round(num, 2)

def reject_outliers(data, m = 2.):
    d = np.abs(data - np.median(data))
    mdev = np.median(d)
    s = d/mdev if mdev else 0.
    return data[s<m]

def basement():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/markham-york-region/basement+apartment/page-{i}/c37l1700274a29276001?radius=68.0&ad=offering&address=266+Lady+Valentina+Ave%2C+Maple%2C+ON+L6A+0E1%2C+Canada&ll=43.872291,-79.482061"
            x = requests.get(url,timeout=30)
            if x.status_code == 200:
                soup = BeautifulSoup(x.text, 'html.parser')
                info_container = soup.findAll("div", class_="info-container")
                description = soup.findAll("div",class_="description")
                price = soup.findAll("div",class_="price")
                location = soup.findAll("div",class_="location")
                for i in info_container:
                    urls = i.find('a')['href']
                    links.append("https://www.kijiji.ca"+urls)
                    titles = i.find('div',class_="title").text.strip()
                    title.append(titles)
                for p in price:
                    strs = ''.join(p.text.split())
                    prices.append(strs)
                
                for d in description:
                    strs = ' '.join(d.text.split())
                    desc.append(strs)
                for l in location:
                    locs = l.find('span')
                    strs = ' '.join(locs.text.split())
                    locations.append(strs)
            else:
                print(x.status_code)
        except Exception as e:
            print(e)
            break

    output["Title"] = title
    output["Location"] = locations
    output["Price"] = prices
    output["description"] = desc
    output["Links"] = links

    price_arr = post_processing(output)
    output = model_distribution(price_arr)
    write_to_db(output, "BASEMENT")

def bedroom_apartment():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-for-rent/markham-york-region/bedroom-apartment/page-{i}/k0c30349001l1700274?radius=64.0&ad=offering&address=3051+York+Regional+Rd+14%2C+Whitchurch-Stouffville%2C+ON+L4A+7X5%2C+Canada&ll=43.947208,-79.353803"
            x = requests.get(url,timeout=30)
            if x.status_code == 200:
                soup = BeautifulSoup(x.text, 'html.parser')
                info_container = soup.findAll("div", class_="info-container")
                description = soup.findAll("div",class_="description")
                price = soup.findAll("div",class_="price")
                location = soup.findAll("div",class_="location")
                for i in info_container:
                    urls = i.find('a')['href']
                    links.append("https://www.kijiji.ca"+urls)
                    titles = i.find('div',class_="title").text.strip()
                    title.append(titles)
                for p in price:
                    strs = ''.join(p.text.split())
                    prices.append(strs)
                
                for d in description:
                    strs = ' '.join(d.text.split())
                    desc.append(strs)
                for l in location:
                    locs = l.find('span')
                    strs = ' '.join(locs.text.split())
                    locations.append(strs)
            else:
                print(x.status_code)
        except Exception as e:
            print(e)
            break

    output["Title"] = title
    output["Location"] = locations
    output["Price"] = prices
    output["description"] = desc
    output["Links"] = links

    price_arr = post_processing(output)
    output = model_distribution(price_arr)
    write_to_db(output, "APARTMENT")


if __name__ == "__main__":
    bedroom_apartment()
    basement()