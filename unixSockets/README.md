# Unix Sockets as a health check end point

## Warning

`This is a demo application and therefore there could be bugs and naughty ways of doing things. I have deliberately tried to make the code easy to read over efficient code. Please make sure you understand the code before copy and pasting it into your environments.`

## Why

Health checks are important. They tell us about the application and in production platforms they are a god send. However the standard way to do this is via a HTTP(S) end point. Which is cool and everything until you make an application that doesn't need HTTP to do its main job. Now you are bringing in an attack vector that you just don't need plus a bunch of code that you don't want to manage.

Consider a different approch. You could use a unix socket to expose the state to a local caller. This means that you don't need to worry about attacks over HTTP or libraries that give you HTTP. You could use something to proxy the calls to a HTTP endpoint. However as a developer thats not really your problem. And thats a good thing.

Follow the unix mantra. Make your app do its thing the best it can.

## Demo Highlights

This demo is all about unix sockets and how we can use them as an alternative to HTTP endpoints. We will create a server that listens to requests. Responds as needed.

We will also create a client that poll the health check endpoint.

To make it fun we will combine them into a single application for the demo.

To run the code you will need go installed.

```bash
go main.go
```

The output will look something like this (You would have different dates):

```text
2018/03/27 15:51:07 Starting a handler for the connection.
2018/03/27 15:51:07 Server got: GET /healthz
2018/03/27 15:51:07 Server responded with 42 bytes.
Client got: {"ticks":49,"epoch_last_tick":1522162267}
```

If you have linux you can also use `socat` to talk to the socket.

```bash
socat - UNIX-CONNECT:/tmp/app.sock <<< "GET /healthz"
# output > {"ticks":31,"epoch_last_tick":1522162013}
```