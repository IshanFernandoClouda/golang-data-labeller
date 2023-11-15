I thought using go routines would speed up the process. I was wrong. With go routines it took around 1min20seconds to process a file with 5000 lines. 

When I didn't use any go routines. The same file was processed in 20seconds

Go routine code https://github.com/IshanFernandoClouda/golang-data-labeller/releases/tag/multi_threaded
