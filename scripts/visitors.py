import ipinfo
import socket
import os

token = os.environ.get("IP_ADDRESS_TOKEN")

handler = ipinfo.getHandler(token)

ip_addresses = open("~/mustafa_m/visitors.txt", "r")

ip_addresses = ip_addresses.readlines()

for i in range(0, len(ip_addresses)):
    ip = ip_addresses[i].split(" ")[-1]
    try:
        socket.inet_aton(ip)
        details = handler.getDetails(ip)
        print(details.city)
    except Exception as e:
        pass