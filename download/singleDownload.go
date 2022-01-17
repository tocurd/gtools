/*
 * @Author: your name
 * @Date: 2021-07-30 21:32:46
 * @LastEditTime: 2022-01-17 22:56:23
 * @LastEditors: vscode
 * @Description: In User Settings Edit
 * @FilePath: \gtools\download\singleDownload.go
 */
package download

import (
	"io"
	"net/http"
	"os"
)

// var Download downloadInterface

type downloader struct {
	io.Reader
	Callback func(total int64, current int64)
	Total    int64
	Current  int64
}

/**
 * @description: 下载指定文件
 * @param {string} url 下载地址
 * @param {string} savePath 保存地址
 * @return {*}
 */
func (_download Download) DownloadFile(url string, savePath string, callback func(total int64, current int64)) error {
	_download.Callback = callback
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()
	fileDownloader := &downloader{
		Callback: callback,
		Reader:   resp.Body,
		Total:    resp.ContentLength,
	}
	if _, err := io.Copy(file, fileDownloader); err != nil {
		return err
	}

	return nil
}

/**
 * @description: 下载回调
 * @param {[]byte} p
 * @return {*}
 */
func (d *downloader) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	d.Current += int64(n)
	d.Callback(d.Total, d.Current*10000)
	return
}
