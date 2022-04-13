#!/bin/bash
# Delete the old directory as needed.
if [ -d /opt/simple-reddit ]; then
    rm -rf /opt/simple-reddit
fi

mkdir -vp /opt/simple-reddit
