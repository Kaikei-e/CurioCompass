FROM elixir:1.16.1

RUN apt-get update && apt-get install -y inotify-tools

ADD /src /usr/src/app

RUN cd /usr/src/app && \
mix deps.get && \
mix ecto.setup && \
mix compile

EXPOSE 4000