package main

import (
	"os"
)

func Write(context []byte, path string) error {
	err := os.WriteFile(path, context, 0644)
	if err != nil {
		return err
	}
	return nil
}
