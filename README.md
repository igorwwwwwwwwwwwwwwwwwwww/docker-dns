# docker-dns

`.docker` dns resolver.

resolves `$container.docker` to its respective ip. supports container ids as well as names. plays nicely with `docker-machine` on osx.

## install

    % ./install-resolver.sh
    % gvt restore

## run

    % eval $(docker-machine env)
    % go run docker_dns.go
    # in separate window
    % docker run -d --name nginx nginx
    % curl nginx.docker

## debug

    # you may need to restart your network
    # sudo ifconfig en0 down && sudo ifconfig en0 up
    % dns-sd -G v4 nginx.docker

## container

    % docker build -t docker-dns .

    % ./install-resolver.sh $(docker-machine ip)
    % docker run -v /var/run/docker.sock:/var/run/docker.sock -p 5300:5300/udp -it docker-dns
