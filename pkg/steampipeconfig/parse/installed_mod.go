package parse

import (
	"github.com/Masterminds/semver/v3"
	"github.com/turbot/steampipe/pkg/steampipeconfig/modconfig"
)

type InstalledMod struct {
	Mod     *modconfig.Mod
	Version *semver.Version
}
