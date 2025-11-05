package utils

type ManifestDir struct {
	Name       string
	IsDownload bool
	SubDir     []ManifestDir
	SubFiles   []ManifestFile
}

type ManifestFile struct {
	Name       string
	IsDownload bool
	NbChunks   int
}
