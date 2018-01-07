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
        "db_name": "webscanner",
        "db_host": "mongodb://127.0.0.1:27017/"
    }
}

client = MongoClient(sys_config['database']['db_host'])
db_connect = client[sys_config['database']['db_name']]

def get_domain():
    pass
    f = open("domain.txt", "r")
    domain_range = []
    for each in f:
        domain_range.append(str(each).replace("\n",""))
    return domain_range

def get_domain_info(domain):
    for port in port_range:
        url = "http://" + domain + ":" + str(port)
        print url
        data = db_connect.domain_info.find({"url":url})
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
                        db_connect.domain_info.save({"url":url,"title":m,"create_time":current_time,"fingerprint":webEye_info,"source_domain":"58.com"})
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
    domain_range = get_domain()
    if len(domain_range) > 0:
        for domain in domain_range:
            print domain
            get_domain_info(domain)

if __name__ == '__main__':
    main()