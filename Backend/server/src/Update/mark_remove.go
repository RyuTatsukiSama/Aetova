package update

import "aetova/server/utils"

func MarkRemoveDir(ManifestDir utils.ManifestDir) utils.ManifestUDir {

	manifestUdir := utils.ManifestUDir{
		Name:  ManifestDir.Name,
		State: utils.Remove,
	}
	// Mark all sub dir to remove
	for _, v := range ManifestDir.SubDir {
		manifestUdir.SubDir = append(manifestUdir.SubDir, MarkRemoveDir(v))
	}

	// Mark all sub files to remove
	for _, v := range manifestUdir.SubFiles {
		manifestUdir.SubFiles = append(manifestUdir.SubFiles, utils.ManifestUFile{
			Name:  v.Name,
			State: utils.Remove,
		})
	}

	return manifestUdir
}
