FROM ubuntu:19.04

MAINTAINER Matthias Schneider <matthias.schneider@retarus.de>

RUN apt-get update && \
  echo "postfix postfix/mailname string example.com" | debconf-set-selections && \
  echo "postfix postfix/main_mailer_type string 'Internet Site'" | debconf-set-selections && \
  apt-get install curl postfix mailutils -y

RUN update-rc.d -f postfix remove

RUN postconf -e syslog_name=example-smtp
RUN postconf -e mynetworks=0.0.0.0/0

# ubuntu sets chroot=y on default, to disable chroot use sed:
#RUN sed -i 's/smtp      inet  n       -       y/smtp      inet  n       -       n/g' /etc/postfix/master.cf
#if chroot then this copy is needed:
RUN cp /etc/host.conf /etc/hosts /etc/nsswitch.conf /etc/resolv.conf /etc/services /var/spool/postfix/etc

RUN curl -sfL https://raw.githubusercontent.com/mschneider82/postfixstdoutlogger/master/godownloader.sh | sh 

# postfixstdoutlogger creates a unixsocket for logging and starts ["postfix", "start-fg"]
CMD ["./usr/bin/postfixstdoutlogger", "--overwrite", "--socketfile", "/dev/log"]

