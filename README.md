# Conway's Game of Life in Go Language
The challenge was to write Conway's Game of Life in one of several languages, include Elm, Rust, Elixir, Clojure, Swift, 
or Go within a 48 hours. 
The final solution will be executable in some format and, ideally, will include automated unit tests.
I started at 11:30 PM Feb 13rd

## Problem Approach
##### analyze the problem to be resolved, as well as constraints (time, tools, etc.)
##### look at different languages and determine which one to use
##### learn about language
##### download development tools (including IDE)
##### set-up development environment and validate correctness
##### research and decide on algorithmic approach
##### write code
* determine objects for object oriented design
* start creating basic object and methods
* test methods as they are developed
* build full algorithm, and try to get to minimal viable product
* build unit tests
* refactor code

## Algorithm
Upon research, one of the easier algorithms involve counting all the neighbors around a given cell (including itself):
* If total count equals 3, then the considered cell is alive
* If total count equals 4, then the considered cell remains unchanged
* Otherwise, the cell is dead

Besides that, the board needs two copies at any given time - one for the current state of the board, and one for
creating the next generation. The avoid the need to create a new board every time, saving memory and avoiding
unnecessary garbage collection (at the expense of slightly more complicated code).

## Installation
Conle this repo to your local machine.

## Run
Just typing ` go run board.go  ` under `github.com/Graciexia/Conway-s-Game-of-Life-in-Go` derectory, it will run a game with all parameters set to their defaults(Pentadecathlon pattern).