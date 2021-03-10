package querybuilder

import "fmt"

func SelectColumnsFromTableWhereCondition(columns []Column, tableName string ,condition Condition, options TableOptions) string {
	cols := parseColumns(columns, options)
	cond := parseCondition(condition)

	return fmt.Sprintf("SELECT %s FROM %s WHERE %s", cols, tableName, cond)
}

func UpdateTableSetColumnWherePrimaryKeyInIds(tableName string, updateStatement Condition, key string, ids string) string {
	statement := parseCondition(updateStatement)

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s IN %s", tableName, statement, key, ids)
}