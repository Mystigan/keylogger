# keylogger
A simple key logger written in go.

## Disclaimer
The intention of this exercise is purely educational. The aim is to learn and understand Go and its libraries better while making some fun stuff.
Please use this program responsibly (once it becomes useful after a few iterations :) ) and at your own discretion. I shall not be held responsible for any use/misuse of this program and any resulting damages.

## Current status
Right now, all it does is use the [eiannone/keyboard](https://github.com/eiannone/keyboard) library to listen to key strokes from the terminal and save it to a text file.

## ToDO
Add more useful features like:
[] Send the recorded keystrokes/file over email.
[] Improve performance so that it consumes system resources minimally and stay under the radar.
[] At some point, write own code for mapping and recording the keystrokes.

## Expected goal(s) out of this activity/project
[] Get to understand Go better.
[] Learn how to use some of the libraries like os, os/exec, syscall, etc.
