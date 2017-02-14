# Conway-s-Game-of-Life-in-Go
 Conway's Game of Life in Go for a 48 hours code challenge (starting at 11:30 PM Feb 13rd)


## Algorithm
When you check each cell's life status, the count of alive includes its self (total count)
> If total count on the board equals to 3, it is alive;
> If total count on the board equals to 4, it remains its original life status;
> otherwise it is dead.

I have board_a and board_b to flip between each other instead of creating a new board every time.