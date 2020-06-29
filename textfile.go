package textfiles

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Data struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
	//Msgslice []string ``
}
/*type Allmessages struct {
	Alldata []string
}
*/
var x = "test.txt"
var MSG []byte

func Createfile() {
	var _, err = os.Stat(x)
	if os.IsNotExist(err) {
		var file, err = os.Create(x)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}
func Writefile(y Data) {
	var file, err = os.OpenFile(x, os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	fmt.Fprintf(file, "%v\n", y)

	// Save file changes.
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("File Updated Successfully.")
}
func Readfile() []string {
	var file, err = os.OpenFile(x, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()
	var m []string
	//var msgs Allmessages
	reader := bufio.NewReader(file)
	for {
		line, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			break
		}
		m = append(m, line)
	}
	

	return m
	//	var z Allmessages
	//	fmt.Println(z)
}
