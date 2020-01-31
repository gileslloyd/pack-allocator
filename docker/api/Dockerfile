FROM php:7.4-fpm-alpine
RUN set -ex \
  && apk --no-cache add \
    postgresql-dev
RUN docker-php-ext-install opcache pdo pdo_pgsql \
        && docker-php-ext-enable pdo pdo_pgsql

WORKDIR /var/www/vhosts/app