package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getFilePath() string {
	argLength := len(os.Args[1:])
	returnValue := "null"
	if argLength > 0 {
		returnValue = os.Args[1]
	} else {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal( err )
		}
		returnValue = filepath.Join(homedir,".ssh","config")
	}
	return returnValue
}

func main() {
	file, err := os.Open(getFilePath())
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}
