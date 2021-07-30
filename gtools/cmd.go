package gtools

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"strings"
)

var Cmd cmdInterface

type cmdInterface interface {
	Cmd(name string, cmd ...string) (string, error)
	RunCMD(name string, cmd ...string) (string, error)
}

type cmd struct{}

func init() {
	Cmd = &cmd{}
}

/**
 * @description: 运行CMD
 * @param {string} name
 * @param {...string} cmd
 * @return {*}
 */
func (_cmd cmd) Cmd(name string, cmd ...string) (string, error) {

	var out bytes.Buffer
	var outErr bytes.Buffer

	command := exec.Command(name, cmd...)

	command.Stdout = &out
	command.Stderr = &outErr

	err := command.Start()
	if err != nil {
		return outErr.String(), err
	}

	err = command.Wait()
	if err != nil {
		return outErr.String(), err
	}

	return out.String(), err
}

/**
 * @description: 运行CMD
 * @param {string} name
 * @param {...string} cmd
 * @return {*}
 */
func (_cmd cmd) RunCMD(name string, cmd ...string) (string, error) {
	command := exec.Command(name, cmd...)
	stdout, err := command.StdoutPipe() // 获取命令输出内容
	if err != nil {
		return "", err
	}
	if err := command.Start(); err != nil { //开始执行命令
		return "", err
	}

	useBufferIO := false
	if !useBufferIO {
		var outputBuff bytes.Buffer
		for {
			tempoutput := make([]byte, 256)
			n, err := stdout.Read(tempoutput)
			if err != nil {
				if err == io.EOF { //读取到内容的最后位置
					break
				}
				return "", err
			}
			if n > 0 {
				outputBuff.Write(tempoutput[:n])
			}
		}

		text := turnUtil.ConvertToString(outputBuff.String(), "gbk", "utf8")
		return strings.Trim(text, "\r\n"), err
	} else {
		outputbuff := bufio.NewReader(stdout)
		output, _, err := outputbuff.ReadLine()
		if err != nil {
			return "", err
		}

		text := turnUtil.ConvertToString(string(output), "gbk", "utf8")
		return strings.Trim(text, "\r\n"), nil
	}
}
