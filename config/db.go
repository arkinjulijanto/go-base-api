package config

func InitDBDsn(c Config) string {
	dsn := c.DB_USER + `:` +
		c.DB_PASS + `@tcp(` +
		c.DB_HOST + `:` +
		c.DB_PORT + `)/` +
		c.DB_NAME +
		`?charset=` + c.DB_CHARSET +
		`&parseTime=` + c.DB_PARSE_TIME +
		`&loc=` + c.DB_LOCAL

	return dsn
}
