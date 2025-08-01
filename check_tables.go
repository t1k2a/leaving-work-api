package main

import (
	"fmt"
	"leaving-work-api/db"
)

func main() {
	db.Init()
	
	// テーブルの存在確認
	var tables []string
	db.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)
	
	fmt.Println("Existing tables:")
	for _, table := range tables {
		fmt.Println("-", table)
	}
	
	// work_recordsテーブルの構造確認
	var columns []struct {
		ColumnName string `gorm:"column:column_name"`
		DataType   string `gorm:"column:data_type"`
	}
	
	db.DB.Raw(`
		SELECT column_name, data_type 
		FROM information_schema.columns 
		WHERE table_name = 'work_records' 
		ORDER BY ordinal_position
	`).Scan(&columns)
	
	fmt.Println("\nwork_records table structure:")
	for _, col := range columns {
		fmt.Printf("- %s: %s\n", col.ColumnName, col.DataType)
	}
}