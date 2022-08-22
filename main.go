package main

import (
	"fmt"
	"go-setup-elvui/utils"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func fileReader(filename string) string {
	data, err := os.ReadFile(filename)
	utils.ErrorHandler(err)

	return string(data)
}

func getElvuiVersion() string {
	c := colly.NewCollector()
	//var tableRefs []string
	//var elvuiVersion string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("href"), "elvui") {

			elvuiContains := strings.TrimSpace(e.Text)

			fmt.Println(elvuiContains)

			re := regexp.MustCompile("[0-9]+")

			fmt.Println(re.FindAllString(elvuiContains, -1))
		}
		// fmt.Println("Find links:", strings.Split(e.Attr("href"), "elvui"))
	})

	err := c.Visit("https://www.tukui.org/welcome.php")
	if err != nil {
		return ""
	}

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
		fmt.Println("End program. Please create file_path.txt ....")
		time.Sleep(5 * time.Second)
		return
	}
}
