#!/usr/bin/env python

import socket   
import time   

address = ('127.0.0.1', 3333)    
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)  
s.connect(address)  
s.send("abc")
a = open("h264")
data = a.read()
s.send(data)

s.close()  

