package gtools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

var Cmd cmdInterface

type cmdInterface interface {
	Cmd(name string, cmd ...string) (string, error)
	RunCMD(name string, cmd ...string) (string, error)
	FFmpeg(callback func(line string), args ...string) error
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
func (_cmd cmd) Cmd(name string, args ...string) (string, error) {
	// 返回一个 cmd 对象
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		return stderr.String(), err
	}
	return out.String(), nil

	// var out bytes.Buffer
	// var outErr bytes.Buffer

	// command := exec.Command(name, cmd...)

	// command.Stdout = &out
	// command.Stderr = &outErr

	// err := command.Start()
	// if err != nil {
	// 	return out.String(), err
	// }

	// err = command.Wait()
	// if err != nil {
	// 	return out.String(), err
	// }

	// return out.String(), err
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
			fmt.Println("tempoutput=", string(tempoutput))
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

		text := Turn.ConvertToString(outputBuff.String(), "gbk", "utf8")
		return strings.Trim(text, "\r\n"), err
	} else {
		outputbuff := bufio.NewReader(stdout)
		output, _, err := outputbuff.ReadLine()
		fmt.Println("tempoutput=", string(output))
		if err != nil {
			return "", err
		}

		text := Turn.ConvertToString(string(output), "gbk", "utf8")
		return strings.Trim(text, "\r\n"), nil
	}
}

/**
 * @description: ffmpeg专属的回调程序
 * @param {string} data
 * @return {*}
 */
func (_cmd cmd) FFmpeg(callback func(data string), args ...string) error {

	command := exec.Command("ffmpeg", args...)
	stderr, err := command.StderrPipe()
	if err != nil {
		return err
	}

	if err := command.Start(); err != nil {
		return err
	}

	//实时循环读取输出流中的一行内容
	reader := bufio.NewReader(stderr)
	for {
		temp := make([]byte, 512)
		n, err := reader.Read(temp)
		if err != nil || io.EOF == err {
			break
		}
		if n <= 1 {
			continue
		}
		callback(string(temp))
	}

	if err := command.Wait(); err != nil {
		return err
	}

	return nil
}
