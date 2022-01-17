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
	Clear(absDir string) error
	GetPathDirs(absDir string) (re []string)
	GetPathFiles(absDir string) (re []string)
}

// configPath := path.Join(gtools.Path.GetCurrentDirectory(), "config.yaml") // 先找本程序文件夹
// if !gtools.File.IsExist(configPath) {
// 	configPath = path.Join(gtools.Path.GetModelPath(), "config.yaml")
// 	if !gtools.File.IsExist(configPath) {
// 		logger.SystemLog.Fatal("config.yaml not exit. using default config")
// 	}
// }

type dir struct{}

func init() {
	Dir = &dir{}
}

/**
 * @description: 内部处理
 * @param {string} path
 * @return {*}
 */
func (_dir dir) dirPathParse(path string) string {
	if path == "" {
		return ""
	}
	if string(path[len(path)-1]) != "/" {
		return Path.ParsePath(path) + "/"
	}
	return Path.ParsePath(path)

}

/**
 * @description: 文件夹是否存在
 * @param {string} path
 * @return {*}
 */
func (_dir dir) IsDir(path string) (bool, error) {
	path = _dir.dirPathParse(path)
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
	path = _dir.dirPathParse(path)
	return _dir.IsDir(path)
}

/**
 * @description: 创建目录
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) Mkdir(absDir string) error {
	absDir = _dir.dirPathParse(absDir)

	if flag, _ := _dir.IsExist(absDir); !flag {
		return os.MkdirAll(path.Dir(absDir), os.ModePerm) //生成多级目录
	}

	return nil
}

/**
 * @description: 清空指定文件夹
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) Clear(absDir string) error {
	absDir = _dir.dirPathParse(absDir)
	dir, err := ioutil.ReadDir(absDir)
	if err != nil {
		return err
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{absDir, d.Name()}...))
	}

	return nil
}

/**
 * @description: 删除文件或文件夹
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) Delete(absDir string) error {
	absDir = _dir.dirPathParse(absDir)
	return os.RemoveAll(absDir)
}

/**
 * @description: 获取目录所有文件夹
 * @param {string} absDir
 * @return {*}
 */
func (_dir dir) GetPathDirs(absDir string) (re []string) {
	absDir = _dir.dirPathParse(absDir)
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
	absDir = _dir.dirPathParse(absDir)
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
