import sqlite3

conn = sqlite3.connect('../covid19.db')
cursor = conn.execute("select * from csse_confirmed")
names = [description[0] for description in cursor.description]
for i in range(len(names)):
	print(i)
	print(names[i])
