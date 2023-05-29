package database

import (
	sq "github.com/Masterminds/squirrel"
)

// PSQL is a prepared placeholder for postgresql
var PSQL = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
