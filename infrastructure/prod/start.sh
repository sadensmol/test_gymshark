#!/bin/bash

pkill test_gymshark 
mv ./test_gymshark_new ./test_gymshark
chmod +x ./test_gymshark
./godotenv ./test_gymshark