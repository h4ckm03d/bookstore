package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Millisecond)
	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	ctx = context.WithValue(ctx, "key", "value")
	req = req.WithContext(ctx)
	ctx.Value("key")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	fmt.Println("Response received, status code:", res.StatusCode)
}
