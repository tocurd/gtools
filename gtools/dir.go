package gtools

import (
	"io/ioutil"
	"os"
	"path"
)

var Dir dirInterface

type dirInterface interface {
	IsDir(path string) (bool, error)
	IsExist(path string) (bool, error)
	Mkdir(absDir string) error
	Delete(absDir string) error
	GetPathDirs(absDir string) (re []string)
	GetPathFiles(absDir string) (re []string)
}

type dir struct{}

func init() {
	Dir = &dir{}
}

/**
 * @description: 文件夹是否存在
 * @param {string} path
 * @return {*}
 */
func (_dir dir) IsDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * @description: 文件夹是否存在
 * @param {string} path
 * @return {*}
 */
func (_dir dir) IsExist(path string) (bool, error) {
	return _dir.IsDir(path)
}

/**
 * @description: 创建目录
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) Mkdir(absDir string) error {
	return os.MkdirAll(path.Dir(absDir), os.ModePerm) //生成多级目录
}

/**
 * @description: 删除文件或文件夹
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) Delete(absDir string) error {
	return os.RemoveAll(absDir)
}

/**
 * @description: 获取目录所有文件夹
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) GetPathDirs(absDir string) (re []string) {
	if exist, _ := _dir.IsExist(absDir); exist {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

/**
 * @description: 获取目录所有文件
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) GetPathFiles(absDir string) (re []string) {
	if exist, _ := _dir.IsExist(absDir); exist {
		files, _ := ioutil.ReadDir(absDir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}
