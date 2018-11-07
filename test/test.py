import requests
import log
import json
import time
#print(requests.get("http://localhost:8000/currentVersion").text)
# print(requests.post("http://localhost:8000/registration", json = {'accountID':200090,'login':'client123','auth_method': 'password','token': '','password': '11' }).text)
# print(requests.post("http://localhost:8000/login", json = {'accountID':2000901,'login':'client','auth_method': 'password','token': '','password': '11' }).text)
# print(requests.post("http://localhost:8000/login",
#                     json={'accountID': 2000902, 'login': 'client111', 'auth_method': 'password', 'token': '', 'password': '11'}).text)
# print(requests.post("http://localhost:8000/login",
#                     json={'accountID': 200090, 'login': 'client123', 'auth_method': 'password', 'token': '', 'password': '11'}).text)
# print(requests.post("http://localhost:8000/login",
#                     json={'accountID': 2000902, 'login': '', 'auth_method': 'token', 'token': 'oHMgFekGar0=4058585359', 'password': ''}).text)
#print(requests.post("http://localhost:8000/balance", json = {'Token':''}).text)
#print(requests.get("http://localhost:8000/StatsAllPersons").text)
#print(requests.get("http://localhost:8000/StatsActivePersons").text)
#print(requests.get("http://localhost:8000/StatAllBets").text)
s = json.loads(requests.post("http://localhost:8000/login",
                             json={'accountID': 2000902,
                              'login': 'client111', 
                              'auth_method': 'password', 
                              'token': '', 'password': '11'}).text)
Token1 = s["token"]
s=json.loads(requests.post("http://localhost:8000/login", json = {'accountID':2000901,'login':'client','auth_method': 'password','token': '','password': '11' }).text)
Token2 = s["token"]
print(requests.post("http://localhost:8000/arena/enter", json={'token': Token1,
                                                        'arenaID': 4372947891}).text)

print(requests.post("http://localhost:8000/arena/enter", json={'token': Token2,
                                                        'arenaID': 4372947891}).text)
print(requests.post("http://localhost:8000/arena/situation", json={'token': Token2,
                                                               'arenaID': 4372947891,"pending":[],"active":[],
                                                               "incoming":[],"rejected":[],"declined":[]}).text)
time.sleep(2)
print(requests.post("http://localhost:8000/parry", json={'token': Token2,
                                                               'arenaID': 4372947891,
                                                               "toAccountID": 2000902,
                                                               "parryType": 'teamvictory',
                                                               "betValue":2
                                                               }).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token1,
#                                                                    'arenaID': 4372947891, "pending": [], "active": [],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token2,
#                                                                    'arenaID': 4372947891, "pending": [], "active": [],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)
# time.sleep(2)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token1,
#                                                                    'arenaID': 4372947891, "pending": [], "active": [{"arenaID": 4372947891, "accountID": 2000901, "parryTypeID": "teamvictory", "betValue": 2}],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token2,
#                                                                    'arenaID': 4372947891, "pending": [], "active": [],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)
time.sleep(2)
print(requests.post("http://localhost:8000/arena/situation", json={'token': Token1,
                                                                   'arenaID': 4372947891, "pending": [], "active": [],
                                                                   "incoming": [], "rejected": [], "declined": [{"arenaID": 4372947891, "accountID": 2000901, "parryTypeID": "teamvictory", "betValue": 2}]}).text)
print(requests.post("http://localhost:8000/arena/situation", json={'token': Token2,
                                                                   'arenaID': 4372947891, "pending": [], "active": [],
                                                                   "incoming": [], "rejected": [], "declined": []}).text)



