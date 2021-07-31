#!/bin/bash

echo "Rebooting device... " 
ssh pi@192.168.0.80 "sudo reboot"
echo "Done"
