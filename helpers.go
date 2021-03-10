package querybuilder

import "fmt"

func parseColumns(columns []Column, options TableOptions) (cols string ){
	cols = ""
	for i, col := range columns {
		cols = cols + formatColumnByType(col, options)
		if i != len(columns) - 1 {
			cols = cols + ", "
		}
	}
	return
}

func formatColumnByType(column Column, options TableOptions) string {
	switch column.Type {
	case "date":
		return formatDateColumn(column.Name, options)
	default:
		return column.Name
	}
}

func formatDateColumn(colName string, options TableOptions) string {
	return fmt.Sprintf("TO_CHAR( TO_TIMESTAMP_TZ(%s, %s), %s) as %s", colName, options.SourceDateFormat, options.OutputDateFormat, colName)
}

func parseCondition(condition Condition) string {
	return fmt.Sprintf("%s %s %s", condition.Column, condition.Operator, condition.Value)
}