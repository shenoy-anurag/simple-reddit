#!/bin/bash
export PATH=$PATH:/usr/local/go/bin

# Download all dependencies for simple-reddit's backend
cd /opt/simple-reddit/backend
go mod download

# Copy the systemd service file for simple-reddit's backend to systemd's system folder.
sudo mv /opt/simple-reddit/scripts/simple-reddit-backend.service /etc/systemd/system
sudo systemctl daemon-reload

# # Download all dependencies for simple-reddit's frontend
# cd /opt/simple-reddit/frontend/forum-app
# npm install
# npm install -g @angular/cli

# mkdir dist

set -xe

# Copy tar.gz file from S3 bucket to simple-reddit folder
aws s3 cp s3://anurags-personal/dist.tar.gz /opt/simple-reddit/frontend/forum-app

cd /opt/simple-reddit/frontend/forum-app

tar -xzvf dist.tar.gz

# Ensure the ownership permissions are correct.
chown -R ubuntu:ubuntu /opt/simple-reddit/frontend/forum-app/dist