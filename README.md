# Go Learning Project

I've been wanting to learn some Go
I initially had these as separate repositories, but decided to make them into a single monorepo after project 1 therefore I'm missing the commit history for poject 0 and 1.

## Resources

- [Beginner tutorial](https://www.youtube.com/watch?v=yyUHQIec83I)
- [11 Projects](https://www.youtube.com/watch?v=jFfo23yIWac)

## Project 0: Hello world/booking app

Very first dive into Go.

1. Input and output
2. Project structure (files and packages)
3. Intro to concurrency

## Project 1: simple static web server

Simple web server
3 routes:

- "/": static file serve
- "/hello": api-like handler function
- "/form": handler function, improved with GET redirect

## project 2: simple CRUD api

Simple movie database crud API
No actual database -> in memory structs + slices
External library: Gorilla Mux
5 routes:

- get all (GET)
- get by id (GET)
- create (POST)
- update (PUT)
- delete (DELETE)

2 endpoints:

- "/movies"
- "/movies/[id]

Testable using Postman
