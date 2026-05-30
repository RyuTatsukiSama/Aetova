package update

import (
	"aetova/server/utils"
	"fmt"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func CompareManifestDir(old utils.ManifestDir, new utils.ManifestDir) utils.ManifestUDir {
	// Check sub dir
	manifestUdir := utils.ManifestUDir{
		Name:  new.Name,
		State: utils.NeedCheck,
	}

	for _, sNewDir := range new.SubDir {
		isBreaking := false
		for k, sOldDir := range old.SubDir {
			if sNewDir.Name == sOldDir.Name {
				manifestUdir.SubDir = append(manifestUdir.SubDir, CompareManifestDir(sNewDir, sOldDir))
				old.SubDir = append(old.SubDir[:k], old.SubDir[k+1:]...)
				isBreaking = true
				break
			}
		}
		if !isBreaking {
			// add new dir that need to be add
			manifestUdir.SubDir = append(manifestUdir.SubDir, MarkAddDir(sNewDir))
		}
	}

	// Old dir that need to be removed
	for _, sOldDir := range old.SubDir {
		manifestUdir.SubDir = append(manifestUdir.SubDir, MarkRemoveDir(sOldDir))
	}

	// check sub file
	for _, sNewFile := range new.SubFiles {
		isBreaking := false
		for k, sOldFile := range old.SubFiles {
			if sNewFile.Name == sOldFile.Name {
				manifestUdir.SubFiles = append(manifestUdir.SubFiles, utils.ManifestUFile{
					Name:  sNewFile.Name,
					State: utils.NeedCheck,
					Old:   sOldFile,
					New:   sNewFile,
				})
				old.SubFiles = append(old.SubFiles[:k], old.SubFiles[k+1:]...)
				isBreaking = true
				break
			}
		}
		if !isBreaking {
			// add new files that need to be add
			manifestUdir.SubFiles = append(manifestUdir.SubFiles, utils.ManifestUFile{
				Name:  sNewFile.Name,
				State: utils.Add,
				New:   sNewFile,
			})
		}
	}

	// Old files that need to be removed
	for _, sOldFile := range old.SubFiles {
		manifestUdir.SubFiles = append(manifestUdir.SubFiles, utils.ManifestUFile{
			Name:  sOldFile.Name,
			State: utils.Remove,
			Old:   sOldFile,
		})
	}

	return manifestUdir
}

func CheckManifestDir(uManifest *utils.ManifestUDir, path string) error {
	// dLog := docLogger.NewLoggerWithGOpts("server/cmd")

	path = fmt.Sprintf("%s/%s", path, uManifest.Name)

	// check all the sub dir, delete the one that don't have change
	var newSubDir []utils.ManifestUDir
	for _, subdDir := range uManifest.SubDir {
		if subdDir.State == utils.NeedCheck {
			err := CheckManifestDir(&subdDir, path)
			if err != nil {
				return err
			}
		}

		if subdDir.State != utils.None {
			newSubDir = append(newSubDir, subdDir)
		}
	}

	uManifest.SubDir = newSubDir

	// Check all the sub files delete the one that don't have change
	var newSubFiles []utils.ManifestUFile
	for _, subFile := range uManifest.SubFiles {
		if subFile.State == utils.NeedCheck {
			err := CheckFileChange(&subFile, path)
			if err != nil {
				return err
			}
		}

		if subFile.State != utils.None {
			newSubFiles = append(newSubFiles, subFile)
		}
	}

	uManifest.SubFiles = newSubFiles

	if len(uManifest.SubDir) == 0 && len(uManifest.SubFiles) == 0 {
		uManifest.State = utils.None
	} else {
		uManifest.State = utils.Change
	}

	return nil
}

func CheckFileChange(uManifest *utils.ManifestUFile, path string) error {
	dLog := docLogger.NewLoggerWithGOpts("server/cmf")

	chk_change, err := CompareChunks(uManifest.Old, uManifest.New, uManifest.New.Size > uManifest.Old.Size, path)
	if err != nil {
		return err
	}

	dLog.Info(fmt.Sprintf("Check file done %s nb of changes %d", uManifest.Name, len(chk_change)))

	if len(chk_change) > 0 {
		uManifest.Chk_changes = chk_change
		uManifest.State = utils.Change
		return nil
	}

	uManifest.State = utils.None
	return nil
}
