/*
 * File: /gorm/joins/main.go                                                   *
 * Project: go-demo                                                            *
 * Created At: Friday, 2022/06/17 , 13:20:02                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Sunday, 2022/06/19 , 14:13:55                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package model

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/plugin/soft_delete"
)

// Model common model
type Model struct {
	ID         uint32                `gorm:"primary_key" json:"id"`
	CreatedBy  string                `json:"created_by"`
	ModifiedBy string                `json:"modified_by"`
	CreatedAt  uint32                `json:"created_at"`
	ModifiedAt uint32                `json:"modified_at"`
	IsDel      soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
}

type Article struct {
	Model `gorm:"embedded"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

// Tag Model
type Tag struct {
	Model `gorm:"embedded"`
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type ArticleWithTags struct {
	Article `gorm:"embedded"`
	Tags    []Tag `json:"tags" gorm:"many2many:blog_article_tag;joinForeignKey:article_id;omitempty;Constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

type ArticleList struct {
	List  []ArticleWithTags
	Pager *Pager
}

type ArticleTag struct {
	Model

	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (at ArticleTag) TableName() string {
	return "blog_article_tag"
}

// func (atw *ArticleWithTags) BeforeCreate(tx *gorm.DB) (err error) {
// 	cols := []clause.Column{}
// 	colsNames := []string{}
// 	for _, field := range tx.Statement.Schema.Fields {
// 		cols = append(cols, clause.Column{Name: field.DBName})
// 		colsNames = append(colsNames, field.DBName)
// 	}
// 	tx.Statement.AddClause(clause.Returning{})

// 	tx.Statement.AddClause(clause.OnConflict{
// 		Columns: cols,
// 		// DoUpdates: clause.AssignmentColumns(colsNames),
// 		DoNothing: true,
// 	})
// 	return nil
// }

// func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
// 	cols := []clause.Column{}
// 	colsNames := []string{}
// 	for _, field := range tx.Statement.Schema.Fields {
// 		cols = append(cols, clause.Column{Name: field.DBName})
// 		colsNames = append(colsNames, field.DBName)
// 	}
// 	tx.Statement.AddClause(clause.Returning{})
// 	tx.Statement.AddClause(clause.OnConflict{
// 		Columns: cols,
// 		// DoUpdates: clause.AssignmentColumns(colsNames),
// 		DoNothing: true,
// 	})
// 	return nil
// }
