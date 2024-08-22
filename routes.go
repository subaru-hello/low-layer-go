package main

type handleFunc func() string

var routes = map[string]handleFunc{
	"/say-hello": sayHello,
	"/goodbye": func() string { return "Goodbye World\n" },
}