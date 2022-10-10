#! /bin/bash

echo "Enter SSH password:"

stty -echo
read pwd
stty echo

echo "Copying over the project files ..."
sshpass -p $pwd ssh root@mmuhammad.net  mkdir -p /root/mustafa_m
sshpass -p $pwd rsync -a \
--exclude '.git' \
--exclude 'db/*' \
--exclude 'test_binary' \
--exclude 'test_queries' \
--exclude 'content' \
--exclude 'README.md' \
--exclude 'TODO.txt' \
--exclude 'visitors.txt' \
--exclude 'scripts/hg.py' \
--exclude 'scripts/deploy.sh' \
/Users/mmuhammad/Desktop/projects/mmuhammad.net/mmuhammad.net/ root@mmuhammad.net:~/mustafa_m
echo "Copied over the project files ..."
echo "Stopping services"
sshpass -p $pwd ssh root@mmuhammad.net systemctl stop go_server
sshpass -p $pwd ssh root@mmuhammad.net systemctl stop caddy
echo "Stopped services"
echo "Copying over service files"
sshpass -p $pwd ssh root@mmuhammad.net cp /root/mustafa_m/services/go_server.service /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net cp /root/mustafa_m/services/finance.service /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net cp /root/mustafa_m/services/finance.timer /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net cp /root/mustafa_m/services/kijiji.service /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net cp /root/mustafa_m/services/kijiji.timer /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net cp /root/mustafa_m/services/flask_server.service /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net systemctl daemon-reload
echo "restarting go server service"
sshpass -p $pwd ssh root@mmuhammad.net systemctl restart go_server
echo "restarting flask server service"
sshpass -p $pwd ssh root@mmuhammad.net systemctl restart flask_server
echo "restarting caddy service"
sshpass -p $pwd ssh root@mmuhammad.net systemctl restart caddy
echo "Deployment Complete"