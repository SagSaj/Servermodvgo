import requests
#import log
import random
import json
import time
var = "http://134.17.161.217:8000"#http://134.17.162.92:8000http://148.251.241.66:5050
ArenaID = 100000001

tokens = []
i=100
for i in range(200,200+i):

      print(requests.post(var+"/account/register/",
                     json={'accountID': i, 'login': 'client'+str(i), 'auth_method': 'password', 'token': '', 'password': 'client111','balance':100}).text)


      s = json.loads(requests.post(var+"/account/login/",
                             json={'accountID': i,
                                   'login': 'client'+str(i),
                                   'auth_method': 'password',
                                   'token': '', 'password': 'client111'}).text)
      Token1 = s["token"]
      tokens.append(Token1)
      print(requests.post(var+"/arena/enter/", json={'token': Token1,
                                                        'arenaID': i}).text)
      print(i)
    print(requests.post(var+"/arena/enter/", json={'token': tokens[i-200],
                                                            'arenaID':  i+random.randint(0, 10)+1000000}).text)