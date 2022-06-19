FROM redis
COPY redis.conf /usr/local/etc/redis/redis.conf
# WORKDIR /
# RUN chmod +x docker-entrypoint.sh
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]

# FROM ruby:2.6.2-stretch

# RUN apt-get update \
#   && apt-get install -y nodejs \
#   && gem install redis-browser \
#   && rm -rf /var/lib/apt/lists/*