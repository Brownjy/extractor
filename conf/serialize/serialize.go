package serialize

import (
	"errors"
	"extractor/conf"
	"fmt"
	"github.com/facebookgo/atomicfile"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
)

type ConfigPath string

// ErrNotInitialized is returned when we fail to read the config because the
// repo doesn't exist.
var ErrNotInitialized = errors.New("fce not initialized, please run 'fce init'")

// ReadConfigFile reads the config from `filename` into `cfg`.
func ReadConfigFile(filename ConfigPath, cfg any) error {
	f, err := os.Open(string(filename))
	if err != nil {
		if os.IsNotExist(err) {
			err = ErrNotInitialized
		}
		return err
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return fmt.Errorf("failure to decode config: %s", err)
	}
	return nil
}

// WriteConfigFile writes the config from `cfg` into `filename`.
func WriteConfigFile(filename ConfigPath, cfg any) error {
	err := os.MkdirAll(filepath.Dir(string(filename)), 0755)
	if err != nil {
		return err
	}

	f, err := atomicfile.New(string(filename), 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return encode(f, cfg)
}
func encode(w io.Writer, value any) error {
	buf, err := conf.Marshal(value)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	return err
}

func Load(filename ConfigPath) (*conf.Config, error) {
	var cfg conf.Config
	err := ReadConfigFile(filename, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
