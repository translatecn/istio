cat >/tmp/c.sh <<EOF
apt install net-tools -y
rm -rf /root/net
rm -rf /root/iptables
mkdir -p /root/net
mkdir -p /root/iptables
rm -rf /tmp/ns.sh
ps -ef |grep -v pause |awk -F ' ' '{print \$2}' |grep -v PID |xargs -I F echo "ls -al /proc/F/ns/pid && nsenter -t F -p -n netstat -lntpu > /root/net/F.txt " >> /tmp/ns.sh
ps -ef |grep -v pause |awk -F ' ' '{print \$2}' |grep -v PID |xargs -I F echo "ps -ef |grep F && echo F && nsenter -t F -p -n iptables-save > /root/iptables/F.txt" >> /tmp/ns.sh
bash /tmp/ns.sh
EOF

docker cp /tmp/c.sh koord-worker:/root/c.sh
docker exec -it koord-worker bash /root/c.sh

rm -rf ./net .iptables
docker cp koord-worker:/root/net ./net
docker cp koord-worker:/root/iptables ./iptables
