#!/bin/bash
output_file_name="heartBeat_binary"
echo "Start Creating New Binary :: ${output_file_name} ..."
sleep 2
go build -o "${output_file_name}"
echo "Successfully Created New Binary :: ${output_file_name}"