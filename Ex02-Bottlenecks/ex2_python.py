# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
from threading import Lock

i=0
myLock=Lock()

def add():
    global i;
    myLock.acquire()
    for x in xrange(0,1000000):
	i=i+1
    myLock.release()    

def sub():
    global i;
    myLock.acquire()
    for x in xrange(0,1000000):
	i=i-1
    myLock.release()  
# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")


def main():
    add_t = Thread(target = add, args = (),)
    sub_t = Thread(target = sub, args = (),)

    add_t.start()
    sub_t.start()    

    add_t.join()
    sub_t.join()

    print(i)


main()
