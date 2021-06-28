package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	DBName    = "project"
	TableName = "account"
)

type TimeLine struct {
	ID       uint32
	Position string
	Value    string
}

const (
	mysqlCreateDatabase = iota
	mysqlTimeLineCreateTable
	mysqlTimeLineInsert
	mysqlTimelineDelete
	mysqlTimeLineUpdate
	mysqlTimeLineSelectAllUnDeleted
)

var (
	errInvalidUserCreateDefaultUser = errors.New("[user] invalid default user create ")

	TimeLineSQLString = []string{
		fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s`, DBName),
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s(
			id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			deleted BOOLEAN NOT NULL DEFAULT FALSE COMMENT '删除位',
			position VARCHAR(30) NOT NULL DEFAULT '' COMMENT'数据位置,左还是右',
			value TEXT NOT NULL DEFAULT '' COMMENT'timeline 每个节点具体内容'
		)ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`, DBName, TableName),
		fmt.Sprintf(`INSERT INTO %s.%s (deleted, position, value) VALUES (?, ?, ?)`, DBName, TableName),
		fmt.Sprintf(`UPDATE %s.%s SET deleted=? WHERE id=?`, DBName, TableName),
		fmt.Sprintf(`SELECT id, position, value FROM %s.%s WHERE deleted = ?`, DBName, TableName),
	}
)

func CreateDatabase(db *sql.DB) error {
	_, err := db.Exec(TimeLineSQLString[mysqlCreateDatabase])
	if err != nil {
		return err
	}

	return nil
}

func CreateTimeLine(db *sql.DB) error {
	_, err := db.Exec(TimeLineSQLString[mysqlTimeLineCreateTable])
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func InsertTimeLine(db *sql.DB, deleted bool, position, value string) error {
	result, err := db.Exec(TimeLineSQLString[mysqlTimeLineInsert], deleted, position, value)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidUserCreateDefaultUser
	}

	return nil
}

func DeleteTimeLine(db *sql.DB) error {
	result, err := db.Exec(TimeLineSQLString[mysqlTimelineDelete], true)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidUserCreateDefaultUser
	}

	return nil
}

func UpdateTimeLine(db *sql.DB, position, value string) error {
	result, err := db.Exec(TimeLineSQLString[mysqlTimeLineInsert], position, value)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalidUserCreateDefaultUser
	}

	return nil
}

func SelectAllUnDeletedTimeLine(db *sql.DB) ([]*TimeLine, error) {
	var (
		TimeLines []*TimeLine

		ID       uint32
		Position string
		Value    string
	)

	rows, err := db.Query(TimeLineSQLString[mysqlTimeLineSelectAllUnDeleted], false)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ID, &Position, &Value); err != nil {

			return nil, err
		}

		TimeLine := &TimeLine{
			ID:       ID,
			Position: Position,
			Value:    Value,
		}

		TimeLines = append(TimeLines, TimeLine)
	}

	return TimeLines, nil
}
