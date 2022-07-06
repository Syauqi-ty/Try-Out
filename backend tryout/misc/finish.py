import requests

# arr = [
#     'ovo-battle-royal-3-547-1606286030',
#     'ovo-battle-royal-3-560-1606286030',
#     'ovo-battle-royal-3-548-1606286030',
#     'ovo-battle-royal-3-549-1606286030',
#     'ovo-battle-royal-3-550-1606286030',
#     'ovo-battle-royal-3-561-1606286030',
#     'ovo-battle-royal-3-562-1606286030',
#     'ovo-battle-royal-3-551-1606286030',
#     'ovo-battle-royal-3-563-1606286030',
#     'ovo-battle-royal-3-552-1606286030',
#     'ovo-battle-royal-3-553-1606286030',
#     'ovo-battle-royal-3-564-1606286030',
#     'ovo-battle-royal-3-565-1606286030',
#     'ovo-battle-royal-3-566-1606286030',
#     'ovo-battle-royal-3-554-1606286030',
#     'ovo-battle-royal-3-567-1606286030',
#     'ovo-battle-royal-3-568-1606286030',
#     'ovo-battle-royal-3-569-1606286030',
#     'ovo-battle-royal-3-555-1606286030',
#     'ovo-battle-royal-3-556-1606286030',
#     'ovo-battle-royal-3-570-1606286030',
#     'ovo-battle-royal-3-546-1606286030',
#     'ovo-battle-royal-3-571-1606286030',
#     'ovo-battle-royal-3-572-1606286030',
#     'ovo-battle-royal-3-558-1606286030',
#     'ovo-battle-royal-3-559-1606286030',
#     'ovo-battle-royal-3-573-1606286030',
#     'ovo-battle-royal-3-557-1606286030'
# ]

arr = ['ovo-battle-royal-4-2-1608261095',
'ovo-battle-royal-4-26-1608261095',]

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