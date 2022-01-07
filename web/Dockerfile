FROM node:slim as build

ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update
RUN apt-get -qq install python3 make gcc g++

WORKDIR /root/
COPY . .
RUN yarn
RUN yarn build

CMD ["yarn", "serve"]
