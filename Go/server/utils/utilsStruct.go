package utils

const (
	SizeChunk int64 = 4096 // bytes
)

type ManifestDir struct {
	Name     string
	SubDir   []ManifestDir
	SubFiles []ManifestFile
}

type ManifestFile struct {
	Name     string
	NbChunks int64
	Size     int64
}
