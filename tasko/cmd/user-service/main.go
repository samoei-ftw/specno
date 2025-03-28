/* 28 March 2025
* Samoei Oloo
* Specno Technical Assessment
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Running container on port 8080...")
	for {
		time.Sleep(10 * time.Second) // keep container alive
	}
}
