FROM alpine:3.16.2

ENV OPERATOR=/usr/local/bin/jenkins-operator \
    USER_UID=1001 \
    USER_NAME=jenkins-operator \
    HOME=/home/jenkins-operator

RUN apk add --no-cache ca-certificates=20220614-r0 \
                       openssh-client==9.0_p1-r2

# install operator binary
COPY ./dist/go-binary ${OPERATOR}

COPY build/bin /usr/local/bin
COPY build/configs /usr/local/configs

RUN chmod u+x /usr/local/bin/user_setup && \
    chmod ugo+x /usr/local/bin/entrypoint && \
    /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
