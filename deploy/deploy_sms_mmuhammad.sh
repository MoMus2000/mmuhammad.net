#! /bin/bash
echo "Enter SSH password:"

stty -echo
read pwd
stty echo

echo "Copying over the project files ..."
sshpass -p $pwd ssh root@mmuhammad.net  mkdir -p /root/sms.mmuhammad.net
sshpass -p $pwd rsync -a \
--exclude '.git' \
--exclude 'db/sms_mmuhammad.db' \
--exclude 'test_binary' \
--exclude 'test_queries' \
--exclude 'content' \
--exclude 'README.md' \
--exclude 'TODO.txt' \
--exclude 'visitors.txt' \
--exclude 'scripts/hg.py' \
--exclude 'scripts/deploy.sh' \
/Users/mmuhammad/Desktop/projects/mmuhammad.net/sms.mmuhammad.net/ root@mmuhammad.net:~/sms.mmuhammad.net
echo "Copied over the project files ..."
echo "Stopping services"
sshpass -p $pwd ssh root@mmuhammad.net systemctl stop sms_server
echo "Stopped services"
echo "Copying over service files"
sshpass -p $pwd ssh root@mmuhammad.net cp /root/sms.mmuhammad.net/services/sms_server.service /etc/systemd/system
sshpass -p $pwd ssh root@mmuhammad.net systemctl daemon-reload
echo "restarting sms_server server service"
sshpass -p $pwd ssh root@mmuhammad.net systemctl restart sms_server
echo "Deployment Complete"