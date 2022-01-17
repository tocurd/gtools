package download

type DownloadProgress struct {
	Total   int64
	Current int64
}
type Download struct {
	Callback func(total int64, current int64)
}

type DownloadInterface interface {
	DownloadFile(url string, filePath string, callback func(total int64, current int64)) error
}
