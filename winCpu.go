// test
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type CpuInfo struct {
	logicalProcessorCount int
	numaNodeCount         int
	processorCoreCount    int
	processorL1CacheCount int
	processorL2CacheCount int
	processorL3CacheCount int
	processorPackageCount int
}

func main() {
	cpu := new(CpuInfo)
	i := 0
	ExecCmd := "wmic cpu get NumberOfCores,NumberOfLogicalProcessors"
	cmd := exec.Command("cmd.exe", "/C", ExecCmd)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return
	}

	cmd.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		line = strings.TrimSpace(line)
		if i == 0 || len(line) == 0 {
			i++
			continue
		}

		for j, v := range strings.SplitN(line, " ", 2) {
			val, _ := strconv.Atoi(strings.TrimSpace(v))
			if j == 0 {
				cpu.processorCoreCount += val
			} else if j == 1 {
				cpu.logicalProcessorCount += val
			}
		}
		i++
		cpu.processorPackageCount++
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("exec<%s> err: %v\n", ExecCmd, err)
		fmt.Println("can't collect cpu info...")
		os.Exit(-1)
	}

	fmt.Println("socketnum:", cpu.processorPackageCount)
	fmt.Println("corenum:", cpu.processorCoreCount)
	fmt.Println("threadnum:", cpu.logicalProcessorCount)
}
