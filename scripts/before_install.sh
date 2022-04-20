#!/bin/bash
# Delete the old directory as needed.
if [ -d /opt/simple-reddit ]; then
    rm -rf /opt/simple-reddit
fi

mkdir -vp /opt/simple-reddit

# Install AWS CLI
cd /opt/simple-reddit

WHICH_AWS=$(which aws)
echo $WHICH_AWS
if [ -z $WHICH_AWS ]; then
    echo "Installing aws cli..."
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    sudo ./aws/install
    source /home/ubuntu/aws/credentials.sh
    source /home/ubuntu/aws/config.sh
    aws configure set aws_access_key_id $AWS_ACCESS_KEY_ID
    aws configure set aws_secret_access_key $AWS_SECRET_ACCESS_KEY
    aws configure set default.region $AWS_DEFAULT_REGION
else
    echo "aws cli already installed"
fi
 