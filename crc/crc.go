package crc

type CRC struct {
}

type CRCInterface interface {
	CRC16(pucFrame []byte) int
}
