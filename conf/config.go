package conf

import (
	"extractor/conf/grafana"
	"extractor/conf/storage"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config all config info
type Config struct {
	// mongo
	Storage storage.Options
	// grafana
	Grafana grafana.Options
	// segment
	//Segment segment.Options
}

const (
	DefaultRepoName   = ".pando-dashboard"
	DefaultRepoRoot   = "./" + DefaultRepoName
	DefaultConfigFile = "config.yaml"
	EnvDir            = "EXTRACTOR_PATH"
)

// PathRoot get dir from env or default
func PathRoot() (string, error) {
	dir := os.Getenv(EnvDir)
	var err error
	if len(dir) == 0 {
		dir, err = homedir.Expand(DefaultRepoRoot)
	}
	return dir, err

}

// Path judge the path
func Path(configRoot, extension string) (string, error) {
	if len(configRoot) == 0 {
		dir, err := PathRoot()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, extension), nil
	}
	return filepath.Join(configRoot, extension), nil
}

// Filename get a filename string
func Filename(configRoot, userConfigFile string) (string, error) {
	if userConfigFile == "" {
		return Path(configRoot, DefaultConfigFile)
	}
	if filepath.Dir(userConfigFile) == "." {
		return Path(configRoot, userConfigFile)
	}
	return userConfigFile, nil
}
func Marshal(value any) ([]byte, error) {
	return yaml.Marshal(value)
}
