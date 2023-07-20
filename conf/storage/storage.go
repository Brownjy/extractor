package storage

type StorageRW uint8

const (
	R StorageRW = iota
	W
	RW
)

type Options struct {
	DBType      string    `yaml:"DbType"`
	DSN         string    `yaml:"DSN"`
	ReadOrWrite StorageRW `yaml:"RW"`
}

func DefaultOptions() Options {
	return Options{
		DBType:      "mongo",
		DSN:         "mongodb://127.0.0.1:27017/?directConnection=true",
		ReadOrWrite: RW,
	}
}
