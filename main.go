package main

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/gocolly/colly"
)

func fileReader(filename string) string {
	data, err := os.ReadFile(filename)
	errorHandler(err)

	return string(data)
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func getElvuiVersion() string {
	c := colly.NewCollector()
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println("Find links:", e)
		fmt.Println(reflect.TypeOf(e))
	})

	c.Visit("https://www.tukui.org/welcome.php")

	return "Catching webSite"
}

func checkFile(filename string) bool {
	file, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !file.IsDir()

}

func checkPackageVersion(lastVersion string, addonsPath string) bool {
	fmt.Println("Receive last version", lastVersion)

	_, err := os.Stat(addonsPath)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("File exist, procced to install")
	time.Sleep(5 * time.Second)
	return true
}

func main() {

	var fileContainPath string = "file_path.txt"

	response := getElvuiVersion()

	isFile := checkFile(fileContainPath)
	if isFile {
		pathForInstall := fileReader(fileContainPath)
		checkPackageVersion(response, pathForInstall)
	} else {
		fmt.Println("No file exist ...")
		time.Sleep(5 * time.Second)
		return
	}
}
