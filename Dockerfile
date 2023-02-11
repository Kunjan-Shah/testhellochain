# Dockerfile to build a Ubuntu 20.04 compatible node binary using Docker

FROM ubuntu:focal

RUN apt-get update -y
RUN apt-get upgrade -y

RUN apt-get install curl -y
RUN apt-get install wget -y
# Install tools
RUN wget https://dl.google.com/go/go1.19.1.linux-amd64.tar.gz
RUN tar -xvf go1.19.1.linux-amd64.tar.gz

ENV PATH "$PATH:/go/bin"

RUN curl https://get.ignite.com/cli! | bash

# Copy chain 
WORKDIR /testhellochain
COPY . .

CMD ignite chain build

