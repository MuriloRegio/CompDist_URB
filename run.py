import os
import _thread as thread

MAX = 4
data = {}
status = {}

if os.path.exists("logs"):
	os.system("rm -rf logs")

os.mkdir("logs")

def run(_id, status):
    data[_id] = os.popen("go run main.go {} {}".format(_id, MAX)).read()
    status[_id] = 1

for i in range(MAX):
    thread.start_new_thread(run, (i,status))


while len(status) < MAX:
    pass

for k,v in data.items():
    with open("logs/log_proc#{}.log".format(k),'w') as outfile:
        outfile.write(v)
