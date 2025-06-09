package internal

import (
	"github.com/bulkashmak/gator-cli/internal/config"
	"github.com/bulkashmak/gator-cli/internal/database"
)

type State struct {
	Cfg *config.Config
	DB  *database.Queries
}
