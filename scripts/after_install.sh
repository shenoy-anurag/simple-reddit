#!/bin/bash
export PATH=$PATH:/usr/local/go/bin

# Download all dependencies for simple-reddit's backend
cd /opt/simple-reddit/backend
go mod download

# Copy the systemd service file for simple-reddit's backend to systemd's system folder.
sudo mv /opt/simple-reddit/scripts/simple-reddit-backend.service /etc/systemd/system
sudo systemctl daemon-reload

# Download all dependencies for simple-reddit's frontend
cd /opt/simple-reddit/frontend/forum-app
npm install
npm install -g @angular/cli

mkdir dist
