#!/bin/bash

echo "Building go script..." 
env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/
echo "Finding service and killing the service" 
ssh pi@192.168.0.82 "sudo pkill -f ./raspberry-bush"
echo "Building done. Now starting upload to raspberry..."
scp bin/raspberry-bush pi@192.168.0.82:program
echo "Upload Complete. Connecting to raspberry pi"
ssh pi@192.168.0.82 "echo \"Connection complete!\" && cd program && echo \"Inside folder /program\" && sudo ./raspberry-bush "
echo "Done"
