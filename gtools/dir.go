package gtools

import (
	"os"
	"path"
)

var Dir dirInterface

type dirInterface interface {
	IsDir(path string) (bool, error)
	IsExist(path string) (bool, error)
	Mkdir(absDir string) error
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

// BuildDir 创建目录
func (_dir dir) Mkdir(absDir string) error {
	return os.MkdirAll(path.Dir(absDir), os.ModePerm) //生成多级目录
}
