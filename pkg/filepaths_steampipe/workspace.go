package filepaths_steampipe

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/turbot/pipe-fittings/constants/runtime"
)

// mod related constants
const (
	WorkspaceDataDir            = ".steampipe"
	WorkspaceModDir             = "mods"
	WorkspaceModShadowDirPrefix = ".mods."
	WorkspaceConfigFileName     = "workspace.spc"
	WorkspaceIgnoreFile         = ".steampipeignore"
	ModFileName                 = "mod.sp"
	DefaultVarsFileName         = "steampipe.spvars"
	WorkspaceLockFileName       = ".mod.cache.json"
)

func WorkspaceModPath(workspacePath string) string {
	return path.Join(workspacePath, WorkspaceDataDir, WorkspaceModDir)
}

func WorkspaceModShadowPath(workspacePath string) string {
	return path.Join(workspacePath, WorkspaceDataDir, fmt.Sprintf("%s%s", WorkspaceModShadowDirPrefix, runtime.ExecutionID))
}

func IsModInstallShadowPath(dirName string) bool {
	return strings.HasPrefix(dirName, WorkspaceModShadowDirPrefix)
}

func WorkspaceLockPath(workspacePath string) string {
	return path.Join(workspacePath, WorkspaceLockFileName)
}

func DefaultVarsFilePath(workspacePath string) string {
	return path.Join(workspacePath, DefaultVarsFileName)
}

func ModFilePath(modFolder string) string {
	modFilePath := filepath.Join(modFolder, ModFileName)
	return modFilePath
}
