FROM alpine:3.13.7

ENV OPERATOR=/usr/local/bin/gerrit-operator \
    USER_UID=1001 \
    USER_NAME=gerrit-operator \
    HOME=/home/gerrit-operator

RUN apk add --no-cache ca-certificates==20191127-r5 \
                       openssh-client==8.4_p1-r4 \
                       git==2.30.2-r0

# install operator binary
COPY ./dist/go-binary ${OPERATOR}

COPY build/bin /usr/local/bin
COPY build/configs /usr/local/configs

RUN chmod u+x /usr/local/bin/user_setup && \
    chmod ugo+x /usr/local/bin/entrypoint && \
    /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
