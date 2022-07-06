mb = [
    547,
    560,
    548,
    549,
    550,
    561,
    562,
    551,
    563,
    552,
    553,
    564,
    565,
    566,
    554,
    567,
    568,
    569,
    555,
    556,
    570,
    546,
    571,
    572,
    558,
    559,
    573,
    557
]

testing=[3030]
cmsaintek = [281,275,272,274,273,261,265,276,270,263]
cmsoshum = [268,283,277,375,282,279,271,264,267,269,266,278,280]
cm = [265,
261,
273,
266,
267,
262,
268,
269,
270,
271,
263,
375,
281,
272,
376,
284,
275,
285,
276,
277,
274,
278,
279,
280,
264,
282,
283]

ga19=[
2384,
2406,
2409,
2446,
2451,
2555,
2571
]

ga24br=[
779,
3110
]

ga24b=[
3138,
3131,
3133,
2604
]

testid = 15
testslug = "battle-royal-5-saintek"
subtest=["pu","ppu","pk","pmm","eng"]
subtestsaintek=["pu","ppu","pk","pmm","eng","ma","fi","ki","bi"]
file = open("sql.txt","w")

import uuid
import time
import datetime
import requests

def my_random_string(string_length=10):
    random = str(uuid.uuid4()) 
    random = random.lower() 
    random = random.replace("-","") 
    return random[0:string_length] 

for one in cmsaintek:
    t = int(time.time());
    file.writelines(f"INSERT INTO payments ( external_id, payer_id, test_id, amount, method, status, createdAt, updatedAt, finishedAt ) VALUES ( 'ovo-{testslug}-{one}-{t}', {one}, 15, 0, 'partnership', 'COMPLETED', '{datetime.datetime.now()}', '{datetime.datetime.now()}', '{datetime.datetime.now()}' );\n")
    
    for subtes in subtestsaintek:
        file.writelines(f"INSERT INTO scores ( student_id , test_id , type , x1 , x2 , x3 , x4 , x5 , x6 , x7 , x8 , x9 , x10 , x11 , x12 , x13 , x14 , x15 , x16 , x17 , x18 , x19 , x20 ) VALUES ( {one} , {testid} , '{subtes}' , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 , 0 );\n")
    
    username = my_random_string(8)
    password = my_random_string(12)
    file.writelines(f"INSERT INTO test_auths ( test_id, student_id, username_to, password_to ) VALUES ( {testid}, {one}, '{username}', '{password}' );\n")

arr = []

for one in cmsaintek:
    t = int(time.time());
    arr.append(f"'ovo-{testslug}-{one}-{t}'")

for one in arr:
    r = requests.post(url="https://studybuddy.id/api/v2/payment/ovo/callback", 
    headers={
        "Content-Type": "application/json",
        "X-CALLBACK-TOKEN": "d916916038c07a017259a26015b32c515fd150b8277c09c0f750d9f5afa3031f"
    },
    json={
        "id": "a962e7ab-9bb6-4e39-aa67-e43c565d7a7d",
        "event": "ewallet.payment",
        "phone": "081234567890",
        "amount": 1,
        "status": "COMPLETED",
        "created": "2020-02-01T01:02:03.456Z",
        "business_id": "58cd618ba0464eb64acdb246",
        "external_id": one,
        "ewallet_type": "OVO"
    })

    print(f"STATUS: {r.status_code}, RESP: {r.text}")