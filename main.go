package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

func run() error {
	ctx := context.Background()

	yc, err := ycNew(ctx, os.Getenv("YANDEX_TOKEN"))
	if err != nil {
		return fmt.Errorf("unable to initialize sdk: %w", err)
	}

	gid, err := yc.GetGroupID(ctx, *args.Folder, *args.Group)
	if err != nil {
		return fmt.Errorf("unable to get instance group id: %w", err)
	}

	for {
		ins, err := yc.GetGroupMembers(ctx, gid)
		if err != nil {
			return fmt.Errorf("unable to get instance group members: %w", err)
		}

		initializing := false
		for _, i := range ins {
			if i.Fqdn == "" {
				initializing = true
				break
			}
		}

		if initializing {
			continue
		}

		for _, i := range ins {
			fmt.Println(i.Fqdn)
		}

		break
	}

	return nil
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
