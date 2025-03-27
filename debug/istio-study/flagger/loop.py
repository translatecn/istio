#!/usr/bin/env python3
import os
import time

while 1:
    time.sleep(1)
    os.system(f"bash info.sh > ./logs/{int(time.time())}.log  2>&1")
