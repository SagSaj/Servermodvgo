import requests
#import log
#import json
#import time
Класс

{"accountIDs":["20388892","195331"],"pending":[],"incoming":[],"active":[{"id":3,"type":"teamvictory","fromAccountID":"195331","toAccountID":"20388892","betValue":28,"status":"active","arenaID":"12307775789379236","createdAt":"2018-11-10T10:12:04.131Z","updatedAt":"
2018-11-10T10:12:23.171Z","deletedAt":null}],"rejected":[],"declined":[]}
   {D&É²¬qER~@9{oûñBÀ¨;ºÉ°ÐaT_òPôò{"arena":{"victory":[{"id":3,"type":"teamvictory","fromAccountID":"195331","toAccountID":"20388892","betValue":28,"status":"paid","arenaID":"12307775789379236","createdAt":"2018-11-10T10:12:04.131Z","updatedAt":"2018-11-10T10:14:49.438Z","deletedAt":null}],"defeat":[],"balance":128},"status":"ok"}
{"bet":{"id":2,"type":"teamvictory","fromAccountID":"20388892","toAccountID":"195331","betValue":16,"status":"rejected","arenaID":"12307775789379236","createdAt":"2018-11-10T10:11:50.547Z","updatedAt":"2018-11-10T10:12:03.955Z","deletedAt":null},"status":"ok"}
{"accountIDs":["20388892","195331"],"pending":[{"id":3,"type":"teamvictory","fromAccountID":"195331","toAccountID":"20388892","betValue":28,"status":"pending","arenaID":"12307775789379236","createdAt":"2018-11-10T10:12:04.131Z","updatedAt":"2018-11-10T10:12:04.131Z","deletedAt":null}],"incoming":[],"active":[],"rejected":[],"declined":[]}
#print(requests.get("http://134.17.162.92:8000/currentVersion").text)#http://148.251.241.66:5050
#print(requests.get("http://134.17.162.92:8000/*").text)#http://148.251.241.66:5050
#print(requests.get("http://134.17.162.92:8000/.").text)#http://148.251.241.66:5050
#print(requests.get("http://134.17.162.92:8000/c").text)#http://148.251.241.66:5050
# print(requests.post("http://localhost:8000/registration", json = {'accountID':200090,'login':'client123','auth_method': 'password','token': '','password': '11' }).text)
#print(requests.post("http://localhost:8000/login", json = {'accountID':2000901,'login':'client','auth_method': 'password','token': '','password': '11' }).text)
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
# s = json.loads(requests.post("http://localhost:8000/login",
#                              json={'accountID': 2000902,
#                               'login': 'client111', 
#                               'auth_method': 'password', 
#                               'token': '', 'password': '11'}).text)
# Token1 = s["token"]
# s=json.loads(requests.post("http://localhost:8000/login", json = {'accountID':2000901,'login':'client','auth_method': 'password','token': '','password': '11' }).text)
# Token2 = s["token"]
# print(requests.post("http://localhost:8000/arena/enter", json={'token': Token1,
#                                                         'arenaID': 4372947891}).text)
# time.sleep(2)
# print(requests.post("http://localhost:8000/arena/enter", json={'token': Token2,
#                                                         'arenaID': 4372947891}).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token2,
#                                                                'arenaID': 4372947891,"pending":[],"active":[],
#                                                                "incoming":[],"rejected":[],"declined":[]}).text)
# print(requests.post("http://localhost:8000/parry", json={'token': Token2,
#                                                                'arenaID': 4372947891,
#                                                                "toAccountID": 2000902,
#                                                                "parryType": 'teamvictory',
#                                                                "betValue":2
#                                                                }).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token1,
#                                                                    'arenaID': 4372947891, "pending": [], "active": [],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token2,
#                                                                    'arenaID': 4372947891, "pending": [], "active": [],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)


