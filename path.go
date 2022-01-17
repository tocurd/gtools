package gtools

import (
	"os"
	"path/filepath"
	"strings"
)

var Path pathInterface

type pathInterface interface {
	GetModelPath() string
	ParsePath(path string) string
	GetBasePath() string
	GetCurrentDirectory() string
}

type selfPath struct{}

func init() {
	Path = &selfPath{}
}

/**
 * @description: 解析用户输入的路径
 * @param {string} path
 * @return {*}
 */
func (_path selfPath) ParsePath(path string) string {
	base := _path.GetBasePath()
	return strings.Replace(path, "@/", base+"/", 1)
}

/**
 * @description: 获取程序当前可执行路径
 * @param {*}
 * @return {*}
 */
func (_path selfPath) GetBasePath() string {
	basePath := _path.GetCurrentDirectory()
	if strings.Contains(basePath, "/Temp") {
		basePath = _path.GetModelPath()
	}
	return basePath
}

/**
 * @description: 获取程序运行目录
 * @param {*}
 * @return {*}
 */
func (_path selfPath) GetModelPath() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}

/**
 * @description: 获取exe所在目录
 * @param {*}
 * @return {*}
 */
func (_path selfPath) GetCurrentDirectory() string {
	dir, _ := os.Executable()
	exPath := filepath.Dir(dir)
	return strings.Replace(exPath, "\\", "/", -1)
}
