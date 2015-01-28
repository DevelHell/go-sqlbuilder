package sqlbuilder

var dialect Dialect

type Statement interface {
	ToSql() (query string, attrs []interface{}, err error)
}

type serializable interface {
	serialize(b *builder)
}

type Clause interface {
	serializable
}

type Expression interface {
	serializable
}

func SetDialect(opt Dialect) {
	dialect = opt
}

func init() {
	// initial setup
	SetDialect(&sqliteDialect{})
}
