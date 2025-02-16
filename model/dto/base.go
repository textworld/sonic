package dto

import (
	"math"
	"reflect"

	"github.com/go-sonic/sonic/model/param"
	"github.com/go-sonic/sonic/util"
)

type BaseDTO struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BaseWpDTO struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Page struct {
	Content     interface{} `json:"content"`
	Pages       int         `json:"pages"`
	PrePage     int         `json:"pre_page"`
	NextPage    int         `json:"next_page"`
	Total       int64       `json:"total"`
	RPP         int         `json:"rpp"`
	PageNum     int         `json:"pageNum"`
	HasNext     bool        `json:"hasNext"`
	HasPrevious bool        `json:"hasPrevious"`
	IsFirst     bool        `json:"isFirst"`
	IsLast      bool        `json:"isLast"`
	IsEmpty     bool        `json:"isEmpty"`
	HasContent  bool        `json:"hasContent"`
}

func NewPage(content interface{}, totalCount int64, page param.Page) *Page {
	contentLen := 0
	r := reflect.ValueOf(content)

	if !r.IsNil() && r.Kind() != reflect.Slice {
		panic("not slice")
	} else {
		contentLen = r.Len()
	}
	totalPage := util.IfElse(page.PageSize == 0, 1, int(math.Ceil(float64(totalCount)/float64(page.PageSize)))).(int)
	dtoPage := &Page{
		Content:     content,
		Total:       totalCount,
		Pages:       totalPage,
		PrePage:     page.PageNum - 1,
		NextPage:    page.PageNum + 1,
		PageNum:     page.PageNum,
		RPP:         page.PageNum,
		HasNext:     page.PageNum < totalPage,
		HasPrevious: page.PageNum > 1,
		IsFirst:     page.PageNum == 1,
		IsLast:      page.PageNum == totalPage,
		IsEmpty:     contentLen == 0,
		HasContent:  contentLen > 0,
	}
	return dtoPage
}
