// Package main
// Author: Samoei Oloo
// Created: 2025-03-28
// License: None
//
// This script is responsible for the main execution of this project

package main

import (
	"fmt"
	"time"

	"github.com/samoei-ftw/tasko/config"
)

func main() {
	fmt.Println("Running container on port 8080...")
	// Load environment variables
	config.Load()
	for {
		time.Sleep(10 * time.Second) // keep container alive
	}
}
