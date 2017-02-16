# Conway's Game of Life in Go Language
The challenge was to write Conway's Game of Life in one of several languages, include Elm, Rust, Elixir, Clojure, Swift, 
or Go within a 48 hours. 
The final solution will be executable in some format and, ideally, will include automated unit tests.
I started at 11:30 PM Feb 13rd and finished at 9:30 pm Feb 15th.

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

# Installation

## Running the Game on Mac OS X
Clone the repository to an empty folder on your local machine with this command:  `git clone git@github.com:Graciexia/Conway-s-Game-of-Life-in-Go.git`.

Then just type `./board` in the `Conway-s-Game-of-Life-in-Go`
directory and the program should run.

## Running the Game on other Operating Systems
First, install `go` on your local machine. To do so,
download the install files (https://golang.org/dl/) appropriate for your operating system and then follow the instructions.

Then, clone the repository to an empty folder on your local machine with this command:  `git clone git@github.com:Graciexia/Conway-s-Game-of-Life-in-Go.git`.

You may also need to set up an environment variable to point to the project path with:
`export GOPATH=dir_path`, substituting `dir_path` with the full folder location
you cloned to above.

Lastly, in the `Conway-s-Game-of-Life-in-Go` directory, 
just type `go run board.go`, and it will run the game with all parameters set to their defaults (Pentadecathlon pattern).
