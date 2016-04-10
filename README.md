# docker-dns

`.docker` dns resolver.

resolves `$container.docker` to its respective ip. supports container ids as well as names. plays nicely with `docker-machine` on osx.

## install

    % sudo route -n add 172.17.0.0/16 $(docker-machine ip default)

    % ./install-resolver.sh
    % gvt restore

## run

    % eval $(docker-machine env)
    % go run docker_dns.go
    # in separate window
    % docker run -d --name nginx nginx
    % curl nginx.docker

## debug

    % dns-sd -G v4 nginx.docker

## container

    % docker build -t docker-dns .

    % ./install-resolver.sh $(docker-machine ip)
    % docker run -v /var/run/docker.sock:/var/run/docker.sock -p 5300:5300/udp -it docker-dns

---

After building this I found [dnsdock](https://github.com/tonistiigi/dnsdock).

    % sudo route -n add 172.17.0.0/16 $(docker-machine ip default)

    % ./install-resolver.sh $(docker-machine ip) 53
    % docker run -d -v /var/run/docker.sock:/var/run/docker.sock --name dnsdock -p 53:53/udp tonistiigi/dnsdock
    % ping dnsdock.docker

Much better.
