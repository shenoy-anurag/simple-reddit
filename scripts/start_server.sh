#!/bin/bash
cp /home/ubuntu/.env /opt/simple-reddit/backend
cd /opt/simple-reddit/backend
sudo /usr/local/go/bin/go build -o simple-reddit-build main.go  # due to permission denied error
echo "Running server..."
# nohup ./simple-reddit-build & disown  # doesn't work as it is blocking in AWS CodeDeploy
sudo systemctl start simple-reddit-backend.service
echo "Server is now up!"

# # Build Angular app
# ng build
# sudo cp -r dist/ /var/www/simple-reddit/
# sudo service nginx restart

# Copy Angular build folder dist to nginx sites-enabled's simple-reddit site's location
sudo cp -r /opt/simple-reddit/frontend/forum-app/dist/ /var/www/simple-reddit/

sudo service nginx restart
