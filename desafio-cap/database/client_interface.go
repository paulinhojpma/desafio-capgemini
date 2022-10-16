package database

// "time"

// IDataBase ..
type IDataBase interface {
	connectService(config *OptionsDBClient) error
	CreateSequence(sequence *Sequence) error
	GetInfoSequences() (*InfoSequence, error)
}

type Sequence struct {
	ID      int
	IsValid bool
	Letters string
}

type InfoSequence struct {
	Valids   int     `json:"count_valid"`
	Invalids int     `json:"count_invalid"`
	Ratio    float32 `json:"ratio"`
}

// OptionsCacheClient ..
type OptionsDBClient struct {
	URL    string `json:"url"`
	DBName string `json:"DBName"`
	Driver string
}

// ConfiguraCache
func (o *OptionsDBClient) ConfigDatabase() (*IDataBase, error) {

	var client IDataBase
	switch o.Driver {
	case "postgres":
		gormDB := &gormPostgres{}
		errGorm := gormDB.connectService(o)
		if errGorm != nil {
			return nil, errGorm
		}
		client = gormDB
	}
	return &client, nil
}
