# Functions returning Function and Fuctions consuming functions

## Description

In this demo we will look at functions that return functions. We will also then look at functions that take functions and why we might use this. This is an advanced topic about memory scoping and behaviour changes but I will try keep it as simple as possible.

## Warning

`This is a demo application and therefore there could be bugs and naughty ways of doing things. I have deliberately tried to make the code easy to read over efficient code. Please make sure you understand the code before copy and pasting it into your environments.`

## Demo Highlights

This demo starts with showing a simple counter using returned functions. I show how memory is scoped to the function even though it is outside the function.

We then move onto a more realistic example of wher this may be used.
We have a logger package that provides the ability to override the formatting of your logs. We make use of this and create a structed logger but also make use of the same logger package.

To use the demo you need to run all the `.go` files.

```bash
go run ./*.go
```

Output from the above:

```txt
Static: 0
Dynamic: 1
Static: 0
Dynamic: 2
Static: 0
Dynamic: 3
Static: 0
Dynamic: 4
Static: 0
Dynamic: 5
Static: 0
Dynamic: 6
Static: 0
Dynamic: 7
Static: 0
Dynamic: 8
Static: 0
Dynamic: 9
Static: 0
Dynamic: 10


Loggers - Debug1
{"log_type":"Debug","log_id":0,"log_epoch":1522063863,"log_message":"Oh No! It'sall going wrong!","go_version":"go1.9.3","go_number_of_goroutines":1,"os":"linux"}
Debug: Oh No! It's all going wrong!


Loggers - Debug2
Debug: Hide the evidence!!!


Loggers - Warning
{"log_type":"WARN","log_id":1,"log_epoch":1522063863,"log_message":"Things couldbe going better"}
Warning: Things could be going better
```

This demo will just exit at the end.

The code is heavily commented. Please read it for further details.