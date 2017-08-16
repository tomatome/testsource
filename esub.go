// test
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	file := fmt.Sprintf("%s%s", "c:/unischeduler/jhw", os.Getenv("LSB_SUB_PARM_FILE"))
	dat, err := ioutil.ReadFile(file)
	fmt.Print(string(dat))
	b := new(bytes.Buffer)
	env := os.Environ()
	fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	er := fmt.Sprintf("%v", err)
	fd.WriteString(er)
	fd.WriteString(string(dat))
	fd.WriteString("\r\n==========================================\r\n")
	fd_time := time.Now().Format("2006-01-02 15:04:05")
	b.WriteString(fd_time)
	b.WriteString("\r\n")
	for _, e := range env {
		b.WriteString(e)
		b.WriteString("\r\n")
	}
	fd.Write(b.Bytes())
	fd.Close()
}
