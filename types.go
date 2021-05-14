package querybuilder

type Column struct {
	Name string
	Type string
}

type Condition struct {
	Column   string
	Operator string
	Value    string
}

type Table struct {
	Name            string
	Columns         []Column
	PrimaryKey      []string
	WhereCondition  Condition
	UpdateStatement Condition
	Options         TableOptions
	MasterTableName string
	MasterKey       string
}

type Config struct {
	Tables []Table
}

type TableOptions struct {
	SourceDateFormat      string
	OutputDateFormat      string
	MultipleRecordsPerKey bool
	GroupBy               string
}
