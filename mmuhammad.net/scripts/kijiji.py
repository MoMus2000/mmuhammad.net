import requests 
from bs4 import BeautifulSoup
import pandas as pd
from tqdm import tqdm
import time
import re
import numpy as np
import statistics
import sqlite3
from distfit import distfit
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
    if len(dataset) < 50:
        price_arr = np.array(price_arr)
    else:
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

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/markham-york-region/2+bedrooms-basement+apartment/page-{i}/c37l1700274a27949001a29276001?ll=43.872291%2C-79.482061&address=266+Lady+Valentina+Ave%2C+Maple%2C+ON+L6A+0E1%2C+Canada&ad=offering&radius=68.0"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/markham-york-region/2+bedrooms-apartment/page-{i}/c37l1700274a27949001a29276001?ll=43.872291%2C-79.482061&address=266+Lady+Valentina+Ave%2C+Maple%2C+ON+L6A+0E1%2C+Canada&ad=offering&radius=68.0"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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

def two_bedroom_st_catharines():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/st-catharines/2+bedrooms-apartment/page-{i}/c37l80016a27949001a29276001?radius=10.0&address=St.+Catharines%2C+ON&ll=43.159375,-79.246863"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_ST_CATHARINES_APARTMENT")


def two_bedroom_basement_st_catharines():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/st-catharines/2+bedrooms-basement+apartment/page-{i}/c37l80016a27949001a29276001?radius=10.0&address=St.+Catharines%2C+ON&ll=43.159375,-79.246863"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_ST_CATHARINES_BASEMENT")

def two_bedroom_apartment_durham():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/oshawa-durham-region/2+bedrooms-apartment/page-{i}/c37l1700275a27949001a29276001"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_DURHAM_APARTMENT")

def two_bedroom_basement_durham():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/oshawa-durham-region/2+bedrooms-basement+apartment/page-{i}/c37l1700275a27949001a29276001"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_DURHAM_BASEMENT")

def two_bedroom_apartment_hamilton():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/hamilton/2+bedrooms-apartment/page-{i}/c37l80014a27949001a29276001?ll=43.255721%2C-79.871102&address=Hamilton%2C+ON&radius=20.0"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_HAMILTON_APARTMENT")

def two_bedroom_basement_hamilton():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/hamilton/2+bedrooms-basement+apartment/page-{i}/c37l80014a27949001a29276001?ll=43.255721%2C-79.871102&address=Hamilton%2C+ON&radius=20.0"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_HAMILTON_BASEMENT")

def two_bedroom_apartment_windsor():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/windsor-area-on/2+bedrooms-apartment/page-{i}/c37l1700220a27949001a29276001?ll=42.314937%2C-83.036363&address=Windsor%2C+ON&radius=20.0"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_WINDSOR_APARTMENT")

def two_bedroom_basement_windsor():
    title= []
    desc= []
    add = []
    links =[]
    prices = []
    locations = []

    output = pd.DataFrame()

    url_set = set()

    for i in tqdm(range(1,int(200))):
        try:
            time.sleep(3)
            url = f"https://www.kijiji.ca/b-apartments-condos/windsor-area-on/2+bedrooms-basement+apartment/page-{i}/c37l1700220a27949001a29276001?ll=42.314937%2C-83.036363&address=Windsor%2C+ON&radius=20.0"
            x = requests.get(url,timeout=30)
            url_set.add(x.url)
            if url not in url_set and i>1:
                print("Skipping ", i)
                print(x.url, url)
                break
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
    write_to_db(output, "TWO_BD_WINDSOR_BASEMENT")

if __name__ == "__main__":
    bedroom_apartment()
    basement()
    two_bedroom_st_catharines()
    two_bedroom_basement_st_catharines()
    two_bedroom_apartment_durham()
    two_bedroom_basement_durham()
    two_bedroom_apartment_hamilton()
    two_bedroom_basement_hamilton()
    two_bedroom_apartment_windsor()
    two_bedroom_basement_windsor()