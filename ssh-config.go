package main
 
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getCommandArgs() (map[string]string) {
	var arguments = make(map[string]string)
	var configFile string
	flag.StringVar(&configFile, "f", "$HOME/.ssh/config", "Specify config file. Default is $HOME/.ssh/config")
	flag.Parse()
	arguments["configFile"] = configFile
	return arguments
}

func getFilePath() (filePath string) {
	// Return the SSH config file path
	// TODO use argument parser
	arguments := getCommandArgs()
	filePath = arguments["configFile"]
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal( err )
	}
	filePath = strings.Replace(filePath,"$HOME",homedir,1)
	return
}

func promptPath(reader *bufio.Reader) (*os.File, error) {
	fmt.Print("Config file not found. Please enter path or press enter to generate: ")
	text, err := reader.ReadString('\n')
	var file *os.File
	if text == "\n" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Print(err)
		}
		file, err = os.Create(filepath.Join(homedir, ".ssh", "config"))
		if err != nil {
			log.Print(err)
		}
	} else {
		file, err = os.Open(text)
		if err != nil {
			log.Print(err)
		}
	}
	return file, err
}

func chooseFromMenu (reader *bufio.Reader) int {
	fmt.Println("What would you like to do? (Default = 1)")
	fmt.Println("1. Generate default options.")
	fmt.Println("2. Set up a barebones config.")
	fmt.Println("3. Set up the deluxe config.")
	text, _ := reader.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(text))
	if (choice >= 1) && (choice <= 3) {
		return choice
	}
	return 1
}

func main() {
	fmt.Println("Welcome to ssh config!")
	file, err := os.Open(getFilePath())
	reader := bufio.NewReader(os.Stdin)
	for err != nil {
		file, err = promptPath(reader)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
 	option := chooseFromMenu(reader)
	fmt.Printf("You have chosen %v.\n", option)
	file.Close()
 
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}
