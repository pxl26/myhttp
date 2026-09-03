package main

import (
	"encoding/json"
	"fmt"
	"myhttp/internal/http"
	"strings"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(rw http.ResponseWriter, req http.Request) {
	path := req.Path()
	fmt.Println(path)
	switch {
	case strings.HasPrefix(path, "/greet"):
		handleGreet(rw, req)
	case strings.HasPrefix(path, "/web"):
		handleWeb(rw, req)
	default:
		handleNotFound(rw, req)
	}
}

func handleGreet(rw http.ResponseWriter, req http.Request) {
	params := req.RequestParams()
	type HelloMsg struct {
		Name string
		Age  int
	}
	msg := &HelloMsg{
		Name: "hello " + params["name"],
		Age:  24,
	}
	resBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	_, err = rw.Write(resBytes)
	if err != nil {
		panic(err)
	}
}

func handleWeb(rw http.ResponseWriter, req http.Request) {
	username := "you"
	if name, exist := req.RequestParams()["name"]; exist {
		username = name
	}
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
    		<meta charset="UTF-8">
    		<title>Ahead to game</title>
		</head>
		<body>
    		<h1>Welcome %s, my fucking buddy :D</h1>
		</body>
		</html>
	`, username)
	_, err := rw.Write([]byte(html))
	if err != nil {
		panic(err)
	}
}

func handleNotFound(rw http.ResponseWriter, _ http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
    		<meta charset="UTF-8">
    		<title>NOT FOUND</title>
		</head>
		<body>
    		<h1>Kidding huh?</h1>
			<div>There are no fucking route as your stupid request !!!!!!</div>
		</body>
		</html>
	`
	_, err := rw.Write([]byte(html))
	if err != nil {
		panic(err)
	}
}
