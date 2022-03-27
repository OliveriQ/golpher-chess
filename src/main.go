package main

import (
	"fmt"
)

func main() {
	fmt.Print("Go chess engine\n\n")
	initialize_lookup_tables()

	parse_fen(start_position, 0)
	print_board()

}
