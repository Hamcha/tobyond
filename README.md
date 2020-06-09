# BYOND redirector

Got a chat where you can't link to BYOND? Just use this!

## Getting started

There's one already hosted if you don't want to bother yourself:

Replace byond://**yourserver.yourdomain:12345** with https://byond.ovo.ovh/**yourserver.yourdomain:12345**

## Install and deploy

`go get github.com/hamcha/tobyond` or use the docker image `hamcha/tobyond`.

Needs a env var `BIND` to something like `:1234`, otherwise defaults to port 8080!

HTTP-only, just put it behind a reverse proxy.
