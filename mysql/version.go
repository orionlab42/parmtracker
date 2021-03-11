package mysql

// Update creates a table for version control - for now not necessary to use
func Update() {
	query := `create table if not exists version (
	k varchar(255) NOT NULL default '',
	v varchar(255) NOT NULL default '',
	PRIMARY KEY (k)
	);`

	c := GetInstance()
	_, e := c.Exec(query)
	if e != nil {
		panic(e)
	}
}
