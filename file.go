/*
 * @Author: your name
 * @Date: 2021-07-30 15:04:34
 * @LastEditTime: 2022-01-14 22:05:50
 * @LastEditors: vscode
 * @Description: In User Settings Edit
 * @FilePath: \serveri:\project\gtools\file.go
 */
package gtools

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"errors"
)

var File fileInterface

type fileInterface interface {
	GetPathDirs(string) []string
	GetPathFiles(string) []string
	ReadLine(string, int) (string, error)
	Read(string) ([]string, error)
	ReadAll(string) (string, error)

	GetFileLineCount(string) (int, error)
	IsExist(string) bool
	Delete(string) error
	Remove(string) error
	Write(string, string, bool) bool
	Create(string, []byte, bool) error
	Move(string, string) error
	Copy(string, string) error
}

type file struct{}

func init() {
	File = &file{}
}

// GetPathDirs 获取目录所有文件夹
func (_file file) GetPathDirs(absDir string) (re []string) {
	absDir = Path.ParsePath(absDir)
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
	absDir = Path.ParsePath(absDir)
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
	path = Path.ParsePath(path)
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
	path = Path.ParsePath(path)
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
func (_file file) IsExist(path string) bool {
	path = Path.ParsePath(path)
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
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
	absDir = Path.ParsePath(absDir)
	return os.RemoveAll(absDir)
}

/**
 * @description: 删除文件或文件夹
 * @param {*}
 * @return {*}
 */
func (_file file) Remove(absDir string) error {
	return _file.Delete(absDir)
}

/**
 * @description: 写入文件
 * @param {*}
 * @return {*}
 */
func (_file file) Write(path string, src string, isClear bool) bool {
	path = Path.ParsePath(path)
	Dir.Mkdir(filepath.Dir(path))
	flag := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !isClear {
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
	}

	f, err := os.OpenFile(path, flag, 0666)
	if err != nil {
		return false
	}
	defer f.Close()

	f.WriteString(src)
	return true
}

/**
 * @description: 写入文件
 * @param {*}
 * @return {*}
 */
func (_file file) Create(path string, src []byte, isClear bool) error {
	path = Path.ParsePath(path)
	Dir.Mkdir(filepath.Dir(path))
	flag := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !isClear {
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
	}

	f, err := os.OpenFile(path, flag, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(src)
	return nil
}

/**
 * @description: 读取文件内容
 * @param {string} fname
 * @return {*}
 */
func (_file file) ReadAll(path string) (src string, err error) {
	path = Path.ParsePath(path)
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	src = ""
	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		src += string(line) + "\r\n"
	}

	src = src[0 : len(src)-2]

	return src, nil
}

/**
 * @description: 读取文件内容
 * @param {string} fname
 * @return {*}
 */
func (_file file) Read(path string) (src []string, err error) {
	path = Path.ParsePath(path)
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return []string{}, err
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

	return src, nil
}

/**
 * @description: 移动文件或目录
 * @param {*}
 * @return {*}
 */
func (_file file) Move(from, to string) error {
	from = Path.ParsePath(from)
	to = Path.ParsePath(to)
	return os.Rename(from, to)
}

/**
 * @description: 拷贝文件或目录
 * @param {*} src
 * @param {string} des
 * @return {*}
 */
func (_file file) Copy(src, des string) error {
	src = Path.ParsePath(src)
	des = Path.ParsePath(des)
	if !_file.IsExist(des) {
		Dir.Mkdir(des)
	}
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
