#! /usr/bin/python
# -*- coding: utf-8 -*-
# Author: bipabo1l
import optparse
import os
import webEye

from pymongo import MongoClient
import re
import requests
from IPy import IP
import time

ISOTIMEFORMAT = "%Y-%m-%d %X"
port_range = [80,8080,8088,8888,8000 ]
int_ip = lambda x: '.'.join([str(x/(256**i)%256) for i in range(3,-1,-1)])
ip_int = lambda x:sum([256**j*int(i) for j,i in enumerate(x.split('.')[::-1])])

sys_config = {
    "database": {
        "db_host": "mongodb://127.0.0.1:27017/"
        ,"db_name": "webscanner",
  }
}

client = MongoClient(sys_config['database']['db_host'])
db_connect = client[sys_config['database']['db_name']]

def get_IP():
    pass
    f = open("ip.txt", "r")
    ip_range = []
    for each in f:
        ips = IP(each)
        for x in ips:
            if str(x).endswith(".0"):
                pass
            else:
                ip_range.append(str(x))
    return ip_range

def get_ips(ip1,ip2):
    ip_range = []
    ip1_num = ip_int(ip1)
    ip2_num = ip_int(ip2)
    for i in range(ip1_num,ip2_num+1):
        ip_range.append(int_ip(i))
    print ip_range
    return ip_range

def get_ip_range(s):
    return s.split(".")[-4]+"."+s.split(".")[-3]+"."+s.split(".")[-2]+".0/24"

def get_ip_info(ip):
    for port in port_range:
        url = "http://" + ip + ":" + str(port)
        print url
        data = db_connect.ip_info.find({"url":url})
        total = data.count()
        if total == 0:
            try:
                req = requests.get(url,timeout=3,allow_redirects=True)
                s = req.text
                code = req.status_code
                if code == 200 or code == 302:
                    regex =  '<title>(.*)</title>'
                    webEye_res = webEye.WebEye(url)
                    webEye_res.run()
                    webEye_info = list(webEye_res.cms_list)
                    print webEye_info
                    for m in re.findall(regex, s):
                        print m + str(port)
                        current_time = time.strftime(ISOTIMEFORMAT, time.localtime(time.time()))
                        db_connect.ip_info.save({"url":url,"title":m,"create_time":current_time,"ip_range":get_ip_range(ip),"fingerprint":webEye_info})
                        continue
            except requests.exceptions.ConnectionError:
                pass
            except requests.exceptions.ConnectTimeout:
                NETWORK_STATUS = False
                pass
            except requests.exceptions.Timeout:
                REQUEST_TIMEOUT = True
                pass


def main():
    ip_range = get_IP()
    print ip_range
    if len(ip_range) > 0:
        for ip in ip_range:
            print ip
            get_ip_info(ip)

if __name__ == '__main__':
    main()