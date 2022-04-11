#!/bin/bash
cd /opt/simple-reddit/backend
go build -o simple-reddit-build main.go
nohup ./simple-reddit-go & disown
