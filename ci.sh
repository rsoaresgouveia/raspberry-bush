#!/bin/bash

echo "Building go script..." 
env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/
echo "Building done. Now starting upload to raspberry..."
scp bin/raspberry-bush pi@192.168.0.80:program
echo "Upload Complete. Connecting to raspberry pi"
ssh pi@192.168.0.80 "echo \"Connection complete!\" && cd program && echo \"Inside folder /program\" && ./raspberry-bush "
echo "Done"
