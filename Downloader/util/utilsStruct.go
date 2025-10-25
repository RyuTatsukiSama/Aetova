package util

type ManifestFile struct {
	Name       string
	IsDownload bool
}

type ManifestDir struct {
	Name       string
	IsDownload bool
	Sub        []interface{}
}
