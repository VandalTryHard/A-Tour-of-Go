// сканирует ТОЛЬКО папку my_directory
package scanFiles

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Scan(dir1 string) string {
	files, err := ioutil.ReadDir(dir1)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("File:", file.Name())
		}
	}
	return "Scan work"
}
