FROM alpine:3.6

MAINTAINER Tien Thanh <tthanh.smt@gmail.com>

RUN apk add --no-cache ca-certificates

COPY view /app/view
COPY yoblog /app

ENV OAUTH2_CLIENT_ID=1943526969221067
ENV OAUTH2_SCOPE=public_profile,email

EXPOSE 8080

WORKDIR /app

ENTRYPOINT ["./yoblog"]
