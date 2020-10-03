#!/bin/bash

ln -s ~/brightness/brightness.service /etc/systemd/system/brightness.service

systemctl enable brightness.service
systemctl start brightness.service 
