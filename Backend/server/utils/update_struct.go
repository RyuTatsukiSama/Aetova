package utils

type State int

const (
	Add State = iota
	Remove
	Change
)

type ManifestUDir struct {
	Name     string
	State    State
	SubFiles []ManifestUFile
}

type ManifestUFile struct {
	Name        string
	State       State
	chk_changes []int
}
