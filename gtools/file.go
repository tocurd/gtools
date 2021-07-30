/*
 * @Author: your name
 * @Date: 2021-07-30 15:04:34
 * @LastEditTime: 2021-07-30 15:11:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gtools\gtools\file.go
 */
package gtools

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"errors"
)

var File fileInterface

type fileInterface interface {
	GetPathDirs(string) []string
	GetPathFiles(string) []string
	ReadLine(string, int) (string, error)
	GetFileLineCount(string) (int, error)
	IsExist(string) bool
	Delete(string) error
	Write(string, []string, bool) bool
	Read(string) []string
	Move(string, string) error
	Copy(string, string) error
}

type file struct{}

func init() {
	File = &file{}
}

// GetPathDirs 获取目录所有文件夹
func (_file file) GetPathDirs(absDir string) (re []string) {
	if _file.IsExist(absDir) {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// GetPathFiles 获取目录所有文件
func (_file file) GetPathFiles(absDir string) (re []string) {
	if _file.IsExist(absDir) {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

/**
 * @description: 读取文件指定行数
 * @param {int} lineNumber
 * @return {*}
 */
func (_file file) ReadLine(path string, lineNumber int) (string, error) {
	file, _ := os.Open(path)
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text(), nil
		}
		lineCount++
	}
	defer file.Close()
	return "", errors.New("超过总行数")
}

/**
 * @description: 获取文件总行数
 * @param {string} path
 * @return {*}
 */
func (_file file) GetFileLineCount(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	fd := bufio.NewReader(file)
	count := 0
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		count++

	}
	return count, nil
}

/**
 * @description: 检查目录或者文件是否存在
 * @param {string} filename
 * @return {*}
 */
func (_file file) IsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
 * @description: 删除文件或文件夹
 * @param {*}
 * @return {*}
 */
func (_file file) Delete(absDir string) error {
	return os.RemoveAll(absDir)
}

/**
 * @description: 写入文件
 * @param {*}
 * @return {*}
 */
func (_file file) Write(fname string, src []string, isClear bool) bool {
	// dirUtil.Mkdir(fname)
	// flag := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	// if !isClear {
	// 	flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
	// }
	// f, err := os.OpenFile(fname, flag, 0666)
	// if err != nil {
	// 	return false
	// }
	// defer f.Close()

	// for _, v := range src {
	// 	f.WriteString(v)
	// 	f.WriteString("\r\n")
	// }

	return true
}

/**
 * @description: 读取文件内容
 * @param {string} fname
 * @return {*}
 */
func (_file file) Read(fname string) (src []string) {
	f, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		return []string{}
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		src = append(src, string(line))
	}

	return src
}

/**
 * @description: 移动文件或目录
 * @param {*}
 * @return {*}
 */
func (_file file) Move(from, to string) error {
	return os.Rename(from, to)
}

/**
 * @description: 拷贝文件或目录
 * @param {*} src
 * @param {string} des
 * @return {*}
 */
func (_file file) Copy(src, des string) error {
	// if !_file.IsExist(des) {
	// return
	// dirUtil.Mkdir(des)
	// }
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		return err
	}
	defer desFile.Close()

	_, err = io.Copy(desFile, srcFile)
	return err
}
