package querybuilder

import "fmt"

func SelectColumnsFromTableWhereCondition(columns []Column, tableName string, condition Condition, options TableOptions) string {
	cols := parseColumns(columns, options)
	cond := parseCondition(condition)
	whereCondition := ""
	if cond != "" {
		whereCondition = fmt.Sprintf("WHERE %s", cond)
	}

	return fmt.Sprintf("SELECT %s FROM %s %s", cols, tableName, whereCondition)
}

func UpdateTableSetColumnWherePrimaryKeyInIds(tableName string, updateStatement Condition, key string, ids string) string {
	statement := parseCondition(updateStatement)

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s IN %s", tableName, statement, key, ids)
}
