import requests
import log
import json
#print(requests.get("http://localhost:8000/currentVersion").text)
print(requests.post("http://localhost:8000/registration", json = {'accountID':200090,'login':'client123','auth_method': 'password','token': '','password': '11' }).text)
print(requests.post("http://localhost:8000/login", json = {'accountID':2000901,'login':'client','auth_method': 'password','token': '','password': '11' }).text)
print(requests.post("http://localhost:8000/login",
                    json={'accountID': 2000902, 'login': 'client111', 'auth_method': 'password', 'token': '', 'password': '11'}).text)
print(requests.post("http://localhost:8000/login",
                    json={'accountID': 200090, 'login': 'client123', 'auth_method': 'password', 'token': '', 'password': '11'}).text)
print(requests.post("http://localhost:8000/login",
                    json={'accountID': 2000902, 'login': '', 'auth_method': 'token', 'token': 'oHMgFekGar0=4058585359', 'password': ''}).text)
#print(requests.post("http://localhost:8000/balance", json = {'Token':''}).text)
#print(requests.get("http://localhost:8000/StatsAllPersons").text)
#print(requests.get("http://localhost:8000/StatsActivePersons").text)
#print(requests.get("http://localhost:8000/StatAllBets").text)
# s = json.loads(requests.post("http://localhost:8000/login",
#                              json={'accountID': '2000901', 'login': 'client111', 'auth_method': 'password', 'token': '', 'password': '11'}).json())
# Token1 = s["token"]
# s=json.loads(requests.post("http://localhost:8000/login", json = {'accountID':'2000901','login':'client','auth_method': 'password','token': '','password': '11' }).json())
# Token2 = s["token"]
# print(requests.post("http://localhost:8000/arena/enter", json={'token': Token1,
#                                                         'arenaID': 4372947891}))
# print(requests.post("http://localhost:8000/arena/enter", json={'token': Token2,
#                                                         'arenaID': 4372947891}))

