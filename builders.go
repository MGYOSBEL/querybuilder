package querybuilder

import "fmt"

func SelectColumnsFromTableWhereCondition(columns []Column, tableName string, condition Condition, options TableOptions) string {
	cols := parseColumns(columns, options)
	cond := parseCondition(condition)
	whereCondition := ""
	if !isEmpty(cond) {
		whereCondition = fmt.Sprintf("WHERE %s", cond)
	}
	return fmt.Sprintf("SELECT %s FROM %s %s", cols, tableName, whereCondition)
}

func SelectColumnsFromTableWhereAllConditions(columns []Column, tableName string, conditions []Condition, options TableOptions) string {
	cols := parseColumns(columns, options)
	cond := parseConditions(conditions, "AND")
	whereCondition := ""
	if !isEmpty(cond) {
		whereCondition = fmt.Sprintf("WHERE %s", cond)
	}
	return fmt.Sprintf("SELECT %s FROM %s %s", cols, tableName, whereCondition)
}

func SelectColumnsFromTableWhereAtLeastOneCondition(columns []Column, tableName string, conditions []Condition, options TableOptions) string {
	cols := parseColumns(columns, options)
	cond := parseConditions(conditions, "OR")
	whereCondition := ""
	if !isEmpty(cond) {
		whereCondition = fmt.Sprintf("WHERE %s", cond)
	}
	return fmt.Sprintf("SELECT %s FROM %s %s", cols, tableName, whereCondition)
}

func UpdateTableSetColumnWherePrimaryKeyInIds(tableName string, updateStatement Condition, key string, ids string) (string, error) {
	statement := parseCondition(updateStatement)

	if isEmpty(statement) {
		return "", fmt.Errorf("Empty statement")
	}
	return fmt.Sprintf("UPDATE %s SET %s WHERE %s IN %s", tableName, statement, key, ids), nil
}
