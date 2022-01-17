package gtools

import (
	"github.com/tocurd/gtools/crc"
	"github.com/tocurd/gtools/download"
)

var Download download.DownloadInterface = download.Download{}
var CRC crc.CRCInterface = crc.CRC{}
