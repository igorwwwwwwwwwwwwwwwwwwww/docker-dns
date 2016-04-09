#!/bin/bash
# https://blog.skipfault.com/2013/04/22/domain-specific-dns-under-osx/
# https://github.com/tonistiigi/dnsdock

HOST=${1:-127.0.0.1}

sudo mkdir -p /etc/resolver
echo -e "nameserver $HOST\nport 5300" | sudo tee /etc/resolver/docker > /dev/null
sudo touch /etc/resolver/docker
dscacheutil -flushcache
