package mcversions

import "time"

type Versions []Version

type Version struct {
	Arguments struct {
		Game []interface{} `json:"game"`
		Jvm  []interface{} `json:"jvm"`
	} `json:"arguments"`
	AssetIndex FileInfo `json:"assetIndex"`
	Assets     string   `json:"assets"`
	Downloads  struct {
		Client         FileInfo `json:"client"`
		ClientMappings FileInfo `json:"client_mappings"`
		Server         FileInfo `json:"server"`
		ServerMappings FileInfo `json:"server_mappings"`
	} `json:"downloads"`
	ID        ID `json:"id"`
	Libraries []struct {
		Name  string `json:"name"`
		Rules []struct {
			Action string `json:"action"`
			Os     struct {
				Name string `json:"name"`
			} `json:"os"`
		} `json:"rules,omitempty"`
		Downloads struct {
			Artifact    FileInfo `json:"artifact"`
			Classifiers struct {
				Javadoc        FileInfo `json:"javadoc"`
				NativesLinux   FileInfo `json:"natives-linux"`
				NativesMacos   FileInfo `json:"natives-macos"`
				NativesOsx     FileInfo `json:"natives-osx"`
				NativesWindows FileInfo `json:"natives-windows"`
				Sources        FileInfo `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Extract struct {
			Exclude []string `json:"exclude"`
		} `json:"extract,omitempty"`
		Natives struct {
			Linux   string `json:"linux"`
			Osx     string `json:"osx"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
	} `json:"libraries"`
	Logging struct {
		Client struct {
			Argument string `json:"argument"`
			File     struct {
				ID   string `json:"id"`
				Sha1 string `json:"sha1"`
				Size int32  `json:"size"`
				URL  string `json:"url"`
			} `json:"file"`
			Type string `json:"type"`
		} `json:"client"`
	} `json:"logging"`
	MainClass              string    `json:"mainClass"`
	MinimumLauncherVersion int32     `json:"minimumLauncherVersion"`
	ReleaseTime            time.Time `json:"releaseTime"`
	Time                   time.Time `json:"time"`
	Type                   Type      `json:"type"`
}

type FileInfo struct {
	Sha1 string `json:"sha1"`
	Size int32  `json:"size"`
	URL  string `json:"url"`

	ID        *ID     `json:"id,omitempty"`
	Path      *string `json:"path,omitempty"`
	TotalSize *int32  `json:"totalSize,omitempty"`
}
