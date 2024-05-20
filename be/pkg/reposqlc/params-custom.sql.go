package reposqlc

type ColumnCustomParams struct {
	ColumnName string
	Ascending  bool
	Limit      int
	Offset     int
}

const (
	ASC  = "ASC"
	DESC = "DESC"
)
