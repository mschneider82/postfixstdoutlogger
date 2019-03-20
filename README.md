[![Go Report Card](https://goreportcard.com/badge/github.com/mschneider82/postfixstdoutlogger)](https://goreportcard.com/report/github.com/mschneider82/postfixstdoutlogger) [![GoDoc](https://godoc.org/github.com/mschneider82/sharecmd?status.svg)](https://godoc.org/github.com/mschneider82/postfixstdoutlogger)

![gopher](gopher.png)

# Postfix Stdout Logger

This tool creates a unixsocket (e.g. `/dev/log`) and then runs `postfix start-fg`
all logs are logged to stdout.

Use this tool in a Dockerfile like this:

```

FROM ubuntu:19.04

MAINTAINER Matthias Schneider

RUN apt-get update && \
  echo "postfix postfix/mailname string example.com" | debconf-set-selections && \
  echo "postfix postfix/main_mailer_type string 'Internet Site'" | debconf-set-selections && \
  apt-get install curl postfix mailutils -y

RUN update-rc.d -f postfix remove

RUN postconf -e syslog_name=example-smtp
RUN postconf -e mynetworks=0.0.0.0/0

RUN cp /etc/host.conf /etc/hosts /etc/nsswitch.conf /etc/resolv.conf /etc/services /var/spool/postfix/etc

RUN curl -sfL https://raw.githubusercontent.com/mschneider82/postfixstdoutlogger/master/godownloader.sh | sh

# postfixstdoutlogger creates a unixsocket for logging and starts ["postfix", "start-fg"]
CMD ["./usr/bin/postfixstdoutlogger", "--overwrite", "--socketfile", "/dev/log"]

```