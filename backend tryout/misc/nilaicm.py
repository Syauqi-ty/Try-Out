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

file = open("sqlcm.txt","w")
subtest=["pu","ppu","pk","pmm","eng"]

for subtes in cm:   
    for one in subtest:
        file.writelines(f"select s.name,u.type,u.score from students s, scores u where s.id = {subtes} and u.student_id = s.id and u.test_id = 13 and u.type = '{one}';\n")
