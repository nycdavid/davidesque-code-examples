package how_to_debug_a_goroutine

import (
	"fmt"
	"strconv"
	"sync"
)

func HeavyProcess(wg *sync.WaitGroup, ticket int) {
	// Insert breakpoint here
	defer wg.Done()
	fmt.Println("Executing Heavy Process...")
	fmt.Println("ticket " + strconv.Itoa(ticket))
	fmt.Print("")
}
