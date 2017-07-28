#!/usr/bin/env python

import socket   
import time   

address = ('127.0.0.1', 3333)    
a = open("h264")
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)  
s.connect(address)  

s.send("flf\n")
while True:
    d = a.read(128)
    if len(d) == 0:
        break
    s.send(d)
    time.sleep(0.01)




s.close()  
