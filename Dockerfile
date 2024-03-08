FROM alpine:3.15

WORKDIR /opt/app
COPY ./bin/userapi ./bin/userapi

CMD ./bin/userapi run