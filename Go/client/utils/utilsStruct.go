package utils

type ManifestDir struct {
	Name     string
	SubDir   []ManifestDir
	SubFiles []ManifestFile
}

type ManifestFile struct {
	Name     string
	NbChunks int
}
