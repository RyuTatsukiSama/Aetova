package utils

type State int

const (
	None State = iota
	Add
	Remove
	Change
	NeedCheck
)

type ManifestUDir struct {
	Name     string
	State    State
	SubDir   []ManifestUDir
	SubFiles []ManifestUFile
}

type ManifestUFile struct {
	Name        string
	State       State
	Chk_changes []int
	Old         ManifestFile
	New         ManifestFile
}
