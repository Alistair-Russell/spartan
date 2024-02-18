package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Project to test
	src := client.Host().Directory(".")

	_, err = client.Container().
		From("golangci/golangci-lint:v1.56").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"golangci-lint", "run", "-v"}).
		Stdout(ctx)
	if err != nil {
		panic(err)
	}

	_, err = client.Container().
		From("golang:1.22").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"go", "test"}).
		WithExec([]string{"go", "build"}).
		Stdout(ctx)
	if err != nil {
		panic(err)
	}
}
