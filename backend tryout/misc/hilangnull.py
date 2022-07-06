soal=[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20]

file = open("null.txt","w")

for one in soal:
    file.writelines(f"update scores set x{one} = 0 where x{one} is null;\n")