package client_api

type DownloadStatus byte

const (
	StatusNeedDownload DownloadStatus = 0
	StatusDownloading
	StatusDownloaded
)
