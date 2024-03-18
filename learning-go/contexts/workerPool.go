package contexts

import (
	"context"
	"fmt"
)

func workerPool() {
	// Смотреть про каналы!
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println(ctx)
}
