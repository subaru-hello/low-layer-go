package main

type handleFunc func() string

var routes = map[string]handleFunc{
	"/": func() string { return "nice to meet you, this is a root.\n" },
	"/say-hello": sayHello,
	"/goodbye": func() string { return "Goodbye World\n" },
	"/cgm": check_goroutine_mechanism,
}