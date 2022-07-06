import uuid
import time
import datetime

cm=[2,26]

def my_random_string(string_length=10):
    random = str(uuid.uuid4()) 
    random = random.lower() 
    random = random.replace("-","") 
    return random[0:string_length] 

for one in cm:
    t = int(time.time())
    username = my_random_string(8)
    password = my_random_string(12)
    print(f"INSERT INTO test_auths ( test_id, student_id, username_to, password_to ) VALUES ( {testid}, {one}, '{username}', '{password}' );")