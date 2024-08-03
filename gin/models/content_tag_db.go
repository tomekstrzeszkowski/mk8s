package models

import (
	"database/sql"
	"fmt"
)

const dbuser = ""
const dbpass = ""
const dbname = ""
const dbhost = "mysql"

func ListContentTags() []ContentTag {
	db, err := sql.Open("mysql", "dbname="+dbname+" user="+dbuser+" password="+dbpass+" host="+dbhost+" sslmode=disable")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()
	content_tags := []ContentTag{}
	results, err := db.Query("SELECT * FROM \"ContentTag\"")
	if err != nil {
		fmt.Println("Err", err.Error())
	}
	for results.Next() {
		var content_tag ContentTag
		results.Scan(&content_tag.ID)
		content_tags = append(content_tags, content_tag)
	}
	return content_tags
}
