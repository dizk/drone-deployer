# Docker image for deployer Drone plugin

FROM alpine:3.3
RUN apk add --update php-cli php-phar wget \
    && wget http://deployer.org/deployer.phar -q \
    && mv deployer.phar /bin/dep \
    && chmod +x /bin/dep \
    && apk del wget \
    && rm -rf /var/cache/apk/*

ADD drone-deployer /bin/
ENTRYPOINT ["/bin/drone-deployer"]
