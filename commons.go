package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"sync"

	"github.com/jiegemena/gotools/stringtools"
)

type Cmd1 struct {
	Cmd   string `json:"cmd"`
	Path  string `json:"path"`
	Name  string `json:"name"`
	State int    `json:"state"`
	// 1 普通任务，2 定时任务，3 循环任务
	RunType int `json:"runtype"`
}

type CmdList struct {
	Cmds []Cmd1 `json:"cmds"`
}

var (
	allCmd  map[string]Cmd1   = make(map[string]Cmd1)
	runStr  map[string]string = make(map[string]string)
	runNums int               = 0
	mutex   sync.Mutex
)

func getRunStrLen() int {
	mutex.Lock()
	defer mutex.Unlock()
	return len(runStr)
}

func setRunStr(key, val1 string) {
	mutex.Lock()
	defer mutex.Unlock()
	runStr[key] = val1
}

func getRunStr(key string) (string, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	b, err := runStr[key]
	return b, err
}

// 检查配置新增
func getNewConfig() bool {
	c, err := ReadAll("config.json")
	if err != nil {
		fmt.Println("%w", err)
		return false
	}

	var tmpjson CmdList
	json.Unmarshal(c, &tmpjson)
	if tmpjson.Cmds == nil {
		return false
	}

	vt := false

	mutex.Lock()
	defer mutex.Unlock()

	if len(allCmd) == 0 {
		for _, v := range tmpjson.Cmds {
			if v.State == 1 {
				allCmd[v.Name] = v
				vt = true
			} else {
				fmt.Println("服务", v.Name, "未开启")
			}
		}
	} else {
		for _, v := range tmpjson.Cmds {
			_, ok := allCmd[v.Name]
			if ok {
				fmt.Println("服务", v.Name, "存在")
			} else {
				if allCmd[v.Name].State == 1 {
					allCmd[v.Name] = v
					vt = true
				} else {
					fmt.Println("服务", v.Name, "未开启")
				}
			}
		}
	}
	return vt
}

func removeRun(vname string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(runStr, vname)
	fmt.Println("退出服务", vname)
}

func execCommand(cmdStr string, runpath string, runname string) {
	defer removeRun(runname)
	RunCommand(cmdStr, runpath)
}

func RunCommand(cmdStr string, runpath string) bool {
	list := strings.Split(cmdStr, " ")
	cmd := exec.Command(list[0], list[1:]...)
	// cmd := exec.Command(commandName, params...)
	cmd.Dir = runpath
	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		var data1 []byte = []byte(line)
		fmt.Println(stringtools.ConvertByte2String(data1, stringtools.GB18030))
	}

	cmd.Wait()
	return true
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
