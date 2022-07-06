import csv
import requests

def load():
    global arr
    file = "students.csv"
    arr = [line for line in(csv.reader(open(file)))]
    return arr

file2 = open("coba.txt","w")
load()

# print(arr)
file2.writelines(arr[0])
students={
        "id" :0,
        "username" : "wibupo",
        "email" : "vaghan123@gmail.com",
        "name" : "wibu po",
        "phone" : "082120812801",
        "student_class" : "12",
        "password" : "caplinos",
        "bright" : 0,
        "target" :0,
        "school": "",
        "circle_id":0,
        "createdAt":"",
        "updatedAt":""
    }
scores={
"id" :0,
"student_id":0,
"test_id":0,
"type":"",
"x1":0,
"x2":0,
"x3":0,
"x4":0,
"x5":0,
"x6":0,
"x7":0,
"x8":0,
"x9":0,
"x10":0,
"x11":0,
"x12":0,
"x13":0,
"x14":0,
"x15":0,
"x16":0,
"x17":0,
"x18":0,
"x19":0,
"x20":0,
"score":0,}



r = requests.post(url="http://localhost:5000/api/v2/student/",
    headers={
        "Content-Type": "application/json",
    },
    json=students)

print(f"STATUS: {r.status_code}, RESP: {r.text}")
