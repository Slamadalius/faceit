package listener

import (
	"fmt"
)

func work(workerID int) {
	defer wg.Done()

	for request := range que {
		fmt.Println(request.UserID, request.Action)
	}

	fmt.Printf("worker id: %d finished \n", workerID)
}
