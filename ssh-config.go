package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getFilePath(args []string) (filePath string) {
	// Return the SSH config file path
	// TODO make this a proper parsed argument thingy
	argLength := len(args[1:])
	if argLength > 0 {
		filePath = args[1]
	} else {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal( err )
		}
		filePath = filepath.Join(homedir,".ssh","config")
	}
	return
}

func main() {
	file, err := os.Open(getFilePath(os.Args))
 
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
