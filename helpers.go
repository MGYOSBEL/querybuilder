package querybuilder

import (
	"fmt"
	"strings"
)

func parseColumns(columns []Column, options TableOptions) (cols string) {
	cols = ""
	for i, col := range columns {
		cols = cols + formatColumnByType(col, options)
		if i != len(columns)-1 {
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

func parseConditions(conditions []Condition, operator string) string {
	cond_str := ""
	for idx, cond := range conditions {
		cond_str = cond_str + parseCondition(cond)
		if len(conditions) > idx {
			op_str := fmt.Sprintf(" %s ", operator)
			cond_str = cond_str + op_str
		}
	}
	return cond_str
}

func isEmpty(str string) bool {
	trimmedStr := strings.ReplaceAll(str, " ", "")
	return len(trimmedStr) == 0
}
