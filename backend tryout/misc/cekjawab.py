b = 'select * from scores where test_id ="'
t_id = 14
b = b + str(t_id) + '" and ( '

for i in range(20):
    a = "x"+str(i+1)
    if(i==19):
        a = a + ' = "1"'     
    else:   
        a = a + ' = "1" or '
    b = b + a
b = b + " );"
print(b)