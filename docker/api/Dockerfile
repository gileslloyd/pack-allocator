FROM php:7.4-fpm-alpine

RUN apk add --no-cache --update --virtual buildDeps autoconf curl

RUN set -ex \
  && apk --no-cache add \
    libmcrypt-dev perl bc \
    alpine-sdk libtool make g++ icu-dev

RUN docker-php-ext-install opcache bcmath sockets intl \
        && docker-php-ext-enable sockets intl

WORKDIR /var/www/vhosts/api