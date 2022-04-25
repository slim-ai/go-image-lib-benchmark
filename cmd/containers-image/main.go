package main

import "context"

func main() {
	_ = CopyContainersImage(context.Background(), "", "", 1)
}
