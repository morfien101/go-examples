# Go Web Execution engine

## Description

This demo application can be used to wrap a binary on a server and give it a web interface. This is useful if you have something that you need to execute remotely but don't want or have access to login to the server.

## Warning

`This is a demo application and therefore there could be bugs and naughty ways of doing things. I have deliberately tried to make the code easy to read over efficient code. Please make sure you understand the code before copy and pasting it into your environments.`

## Demo Highlights

This demo shows how to wrap a http server around a binary. We wrap a local ping in the demo code because it is easy on CPU and takes time to execute. This allows the student to see the pages change as the ping executes and completes.

To use the demo you need to run the go_execute.go file.

```bash
go run ./go_execute
```

Then in a browser or curl call the web page.

```bash
curl http://localhost
curl http://localhost/jobs
curl http://localhost/jobs/1
```

The code is heavily commented. Please read it for further details.