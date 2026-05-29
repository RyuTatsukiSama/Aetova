package utils

const (
	SizeChunk int64 = 4096 // bytes
)

type ManifestGame struct {
	Version uint
	Name    string
	Guid    uint
}

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
