import json
import random
import time

import requests

# http://134.17.162.92:8000http://148.251.241.66:5050
var = "http://178.124.139.97:8000"
#var = "http://localhost:8000"

print(requests.get(var + "/currentVersion/").text)
print(requests.post(var + "/account/register/", json={
    'accountID': 1, 'login': 'client11@som.nu', 'auth_method': 'password', 'token': '', 'password': 'client111',
    'balance': 100}).text)

print(requests.post(var + "/account/register/", json={
    'accountID': 2, 'login': 'client22@som.nu', 'auth_method': 'password', 'token': '', 'password': 'client111',
    'balance': 100}).text)

s = json.loads(requests.post(var + "/account/login/", json={
    'accountID': 1, 'login': 'client11@som.nu', 'auth_method': 'password', 'token': '', 'password': 'client111'}).text)
print(s)
Token1 = s["token"]
print("Token1 " + Token1)
time.sleep(1)
s = json.loads(requests.post(var + "/account/login/", json={
    'accountID': 2, 'login': 'client22@som.nu', 'auth_method': 'password', 'token': '', 'password': 'client111'}).text)
time.sleep(1)
Token2 = s["token"]
print("Token2 " + Token2)
for _ in range(0, 3):
    ArenaID = random.randint(1, 99999999999)
    print("ARENA_ID", ArenaID)
    print("enter 1:", requests.post(var + "/arena/enter/", json={'token': Token1, 'arenaID': ArenaID}).text)
    time.sleep(1)
    print("status 1:", requests.post(var + "/arena/situation/", json={'token': Token1, 'arenaID': ArenaID}).text)
    print("enter 2:", requests.post(var + "/arena/enter/", json={'token': Token2, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("status 1:", requests.post(var + "/arena/situation/", json={'token': Token1, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("status 2:", requests.post(var + "/arena/situation/", json={'token': Token2, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("parry from 2 to 1", requests.post(var + "/parry/", json={
        'token': Token2, 'arenaID': ArenaID, "toAccountID": 1, "parryType": 'teamvictory', "betValue": 2}).text)
    time.sleep(2)
    temp = requests.post(var + "/arena/situation/", json={'token': Token1, 'arenaID': ArenaID}).text
    print("status 1", temp)
    s = json.loads(temp)
    betid = s["incoming"][0]["id"]
    print("2 declined", requests.post(var + "/parry/decline/",
                                      json={'token': Token2, "betID": betid, "arenaID": ArenaID}).text)
    time.sleep(2)
    print("status 1", requests.post(var + "/arena/situation/", json={'token': Token1, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("status 2", requests.post(var + "/arena/situation/", json={'token': Token2, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("parry from 2 to 1", requests.post(var + "/parry/", json={
        'token': Token2, 'arenaID': ArenaID, "toAccountID": 1, "parryType": 'teamvictory', "betValue": 2}).text)
    time.sleep(2)
    print("status 2", requests.post(var + "/arena/situation/", json={'token': Token2, 'arenaID': ArenaID}).text)
    time.sleep(2)
    temp = requests.post(var + "/arena/situation/", json={'token': Token1, 'arenaID': ArenaID}).text
    print("status 1", temp)
    s = json.loads(temp)
    betid = s["incoming"][0]["id"]
    time.sleep(2)
    print("1 activated", requests.post(var + "/parry/activate/",
                                       json={'token': Token1, "betID": betid, "arenaID": ArenaID}).text)
    time.sleep(2)
    print("status 1", requests.post(var + "/arena/situation/", json={'token': Token1, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("status 2", requests.post(var + "/arena/situation/", json={'token': Token2, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("quit 1", requests.post(var + "/arena/quit/", json={'token': Token1, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("status 2", requests.post(var + "/arena/situation/", json={'token': Token2, 'arenaID': ArenaID}).text)
    time.sleep(2)
    print("quit 2", requests.post(var + "/arena/quit/", json={'token': Token2, 'arenaID': ArenaID}).text)
    time.sleep(1)
    print("result 1", requests.post(var + "/arena/result/", json={
        'token': Token1, 'arenaID': ArenaID, "data": {"victory": True}}).text)
    time.sleep(2)
    print("result 2", requests.post(var + "/arena/result/", json={
        'token': Token2, 'arenaID': ArenaID, "data": {"victory": False}}).text)






#print(requests.post("http://localhost:8000/account/login/",
 #                   json={'accountID': 2000901, 'login': 'client', 'auth_method': 'password', 'token': '', 'password': '11'}).text)
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
 #                                                             "incoming":[],"rejected":[],"declined":[]}).text)
# print(requests.post("http://localhost:8000/parry", json={'token': Token2,
#                                                                'arenaID': ArenaID,
#                                                                "toAccountID": 2000902,
#                                                                "parryType": 'teamvictory',
#                                                                "betValue":2
#                                                                }).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token1,
#                                                                    'arenaID': ArenaID, "pending": [], "active": [],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)
# print(requests.post("http://localhost:8000/arena/situation", json={'token': Token2,
#                                                                    'arenaID': ArenaID, "pending": [], "active": [],
#                                                                    "incoming": [], "rejected": [], "declined": []}).text)
# for i in range(200,300):

#       print(requests.post(var+"/account/register/",
#                      json={'accountID': i, 'login': 'client'+str(i), 'auth_method': 'password', 'token': '', 'password': 'client111','balance':100}).text)


#       s = json.loads(requests.post(var+"/account/login/",
#                              json={'accountID': i,
#                                    'login': 'client'+str(i),
#                                    'auth_method': 'password',
#                                    'token': '', 'password': 'client111'}).text)
#       Token1 = s["token"]
#       tokens.append(Token1)
#       print(requests.post(var+"/arena/enter/", json={'token': Token1,
#                                                         'arenaID': i}).text)
#       print(i)
# for i in range(0,100):
#     print(requests.post(var+"/arena/enter/", json={'token': tokens[i],
#
# print(requests.get(var+"/currentVersion/").text)
# print(requests.post(var+"/account/register/",
#                      json={'accountID': 1, 'login': 'client', 'auth_method': 'password', 'token': '', 'password': 'client111', 'balance': 100}).text)


# {"accountIDs":["20388892","195331"],"pending":[],"incoming":[],"active":[{"id":3,"type":"teamvictory","fromAccountID":"195331","toAccountID":"20388892","betValue":28,"status":"active","arenaID":"12307775789379236","createdAt":"2018-11-10T10:12:04.131Z","updatedAt":"2018-11-10T10:12:23.171Z","deletedAt":null}],"rejected":[],"declined":[]}
# {"arenaID": 7623147585705008, "fromAccountID": 20388892, "toAccountID": 195331, "ID": 0, "parryTypeID": "teamvictory",
#                                                                                               "betValue": 32, "createdAt": "2018-11-13T12:07:36.658928856+03:00", "updatedAt": "2018-11-13T12:07:36.658928902+03:00", "deletedAt": "0001-01-01T00:00:00Z", "status": "incoming"}], "rejected": [], "declined": []}

# {"arena":{"victory":[{"id":3,"type":"teamvictory","fromAccountID":"195331","toAccountID":"20388892","betValue":28,"status":"paid","arenaID":"12307775789379236","createdAt":"2018-11-10T10:12:04.131Z","updatedAt":"2018-11-10T10:14:49.438Z","deletedAt":null}],"defeat":[],"balance":128},"status":"ok"}
# {"bet":{"id":2,"type":"teamvictory","fromAccountID":"20388892","toAccountID":"195331","betValue":16,"status":"rejected","arenaID":"12307775789379236","createdAt":"2018-11-10T10:11:50.547Z","updatedAt":"2018-11-10T10:12:03.955Z","deletedAt":null},"status":"ok"}
# {"accountIDs":["20388892","195331"],"pending":[{"id":3,"type":"teamvictory","fromAccountID":"195331","toAccountID":"20388892","betValue":28,"status":"pending","arenaID":"12307775789379236","createdAt":"2018-11-10T10:12:04.131Z","updatedAt":"2018-11-10T10:12:04.131Z","deletedAt":null}],"incoming":[],"active":[],"rejected":[],"declined":[]}
# print(requests.get(var+"/currentVersion").text)#http://148.251.241.66:5050
# tokens = []
#time.sleep(100)


##############
#print(requests.post(var+"/account/register/",
#                      json={'accountID': 2, 'login': 'client11@som.nu', 'auth_method': 'password', 'token': '', 'password': 'client111', 'balance': 100}).text)
