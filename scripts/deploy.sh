#! /bin/bash

echo "Enter SSH password:"

stty -echo
read pwd
stty echo

echo "Copying over the project files ..."
sshpass -p $pwd scp -r /Users/a./Desktop/go/mustafa_m/ root@mmuhammad.net:~/
sshpass -p $pwd ssh root@mmuhammad.net rm -r /root/mustafa_m/.git
echo "Copied over the project files ..."

echo "Copying over service files"
sshpass -p $pwd ssh root@mmuhammad.net cp /root/mustafa_m/services/go_server.service /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net systemctl daemon-reload
echo "restarting go server service"
sshpass -p $pwd ssh root@mmuhammad.net systemctl restart go_server
echo "restarting caddy service"
sshpass -p $pwd ssh root@mmuhammad.net systemctl restart caddy
echo "Deployment Complete"