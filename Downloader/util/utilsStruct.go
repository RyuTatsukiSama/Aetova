package util

type ManifestDir struct {
	Name       string
	IsDownload bool
	SubDir     []ManifestDir
	SubFiles   []ManifestFile
}

type ManifestFile struct {
	Name       string
	IsDownload bool
	Chunks     []Chunk
}

type Chunk struct {
	Name       string
	IsDownload bool
}
