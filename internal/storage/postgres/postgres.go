package postgres

type postgres struct {
	cfg Config
}

func new(cfg Config) storage.Storage {
	return &postgers {
		cfg: cfg,
	}
}