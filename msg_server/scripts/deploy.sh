#! /bin/bash

echo "Enter SSH password:"

stty -echo
read pwd
stty echo

echo "Copying over the project files ..."
sshpass -p $pwd ssh root@mmuhammad.net  mkdir -p /root/msg_server
sshpass -p $pwd rsync -a \
--exclude '.git' \
--exclude 'scripts/deploy.sh' \
/Users/mmuhammad/Desktop/projects/mmuhammad.net/msg_server/ root@mmuhammad.net:~/msg_server
echo "Copied over the project files ..."
echo "Stopping services"
sshpass -p $pwd ssh root@mmuhammad.net systemctl stop msg_server
echo "Stopped services"
echo "Copying over service files"
sshpass -p $pwd ssh root@mmuhammad.net cp /root/msg_server/services/msg_server.service /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net systemctl daemon-reload
echo "restarting flask server service"
sshpass -p $pwd ssh root@mmuhammad.net systemctl restart msg_server
echo "Deployment Complete"