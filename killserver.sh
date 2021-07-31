#!/bin/bash

echo "Finding service and killing the service" 
ssh pi@192.168.0.80 "pkill -f ./raspberry-bush"
echo "Done"
