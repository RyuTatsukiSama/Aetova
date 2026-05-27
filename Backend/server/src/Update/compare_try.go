package update

import butcher "aetova/server/src/Butcher"

// This function will chop two zip, generate two manifestfile, compare manifest file, then compare chunk
func Compare(old string, new string) error {
	oldManifest, err := butcher.ChopeFileCaller("", old)
	if err != nil {
		return err
	}

	newManifest, err := butcher.ChopeFileCaller("", new)
	if err != nil {
		return err
	}

	CompareManifestFile(oldManifest, newManifest)

	return nil
}
