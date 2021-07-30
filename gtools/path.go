package gtools

import (
	"os"
	"path/filepath"
	"strings"
)

var Path pathInterface

type pathInterface interface {
	GetModelPath() string
	GetCurrentDirectory() string
}

type path struct{}

func init() {
	Path = &path{}
}

/**
 * @description: 获取程序运行目录
 * @param {*}
 * @return {*}
 */
func (_path path) GetModelPath() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}

/**
 * @description: 获取exe所在目录
 * @param {*}
 * @return {*}
 */
func (_path path) GetCurrentDirectory() string {
	dir, _ := os.Executable()
	exPath := filepath.Dir(dir)
	return strings.Replace(exPath, "\\", "/", -1)
}
