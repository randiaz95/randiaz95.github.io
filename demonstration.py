# -*- coding: utf-8 -*-
"""
Created on Tue Jun  6 21:33:28 2017

@author: Randizo

Meta Data:
    
Music:Chopin


Position for IDE: (720,145)
"""

import webbrowser as wb
import pyautogui as pg
import time

# This loop stores the positions of the mouse
while(True):
    x,y = pg.position()
    print("X: ",x," Y: ", y)

# This loop counts down, so that the user gets the screen ready.
for i in list(range(4))[::-1]:
    print(i+1)
    time.sleep(1)
    
def ctrl_sting(a):
    pg.keyDown("ctrl")
    pg.keyDown(a)
    pg.keyUp(a)
    pg.keyUp("ctrl")

"""
wb.open("youtube.com")
ctrl_sting("c")
"""
