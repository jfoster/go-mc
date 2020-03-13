package mcversions

import "time"

const ManifestURL = `https://launchermeta.mojang.com/mc/game/version_manifest.json`

type Manifest struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`
	Versions []struct {
		ID          ID        `json:"id"`
		Type        Type      `json:"type"`
		URL         string    `json:"url"`
		Time        time.Time `json:"time"`
		ReleaseTime time.Time `json:"releaseTime"`
	} `json:"versions"`
}

type Type string

const (
	ReleaseType  Type = "release"
	SnapshotType Type = "snapshot"
	BetaType     Type = "old_beta"
	AlphaType    Type = "old_alpha"
)

type ID string
