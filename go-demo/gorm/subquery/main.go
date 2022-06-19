/*
 * File: /gorm/subquery/main.go                                                *
 * Project: go-demo                                                            *
 * Created At: Friday, 2022/06/17 , 14:04:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/06/18 , 06:43:12                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "go-demo/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:toor@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// var articles []model.Article
	sbq := db.Debug().Model(&model.Article{}).Select("id")
	// SELECT `id` FROM `blog_article` WHERE `blog_article`.`is_del` = 0
	// fmt.Println(len(articles))
	// fmt.Printf("%+v\n", articles[0])
	// fmt.Printf("%#v\n", articles[0])
	// fmt.Println(json.MarshalIndent(articles[0], "", "    "))
	// db.Model(&model.Article{}).
	// fmt.Println(articles[0])
	var tags []model.Tag
	var at []model.ArticleTag
	// db.Debug().Table("blog_article_tag").Select("tag_id").Where("article_id IN (?)", sbq).Find(&tags)

	db.Debug().Table("blog_article_tag").Where("article_id IN (?)", sbq).Find(&at)
	// SELECT * FROM `blog_article_tag` WHERE article_id IN (SELECT `id` FROM `blog_article` WHERE `blog_article`.`is_del` = 0) AND `blog_article_tag`.`is_del` = 0	// select `tag_id` from `blog_article` where `article_id` =

	fmt.Println(len(tags))
	fmt.Println(at[0])
	for _, t := range at {
		fmt.Println(t.ArticleID, t.TagID)
	}

}
