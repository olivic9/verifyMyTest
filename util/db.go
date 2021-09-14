package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

func NewDb() (*gorm.DB, error) {

	dbm, err := gorm.Open("mysql", "root:root@(mysql:3306)/verify-my-test?utf8mb4&parseTime=True&loc=UTC")
	if err != nil {
		Fatal(FatalError, err)
		return nil, err
	}
	dbm.LogMode(false)

	return dbm, nil
}

//ps: below could be a huge improvement for gorm

type Paginator struct {
	DB      *gorm.DB
	OrderBy []string
	Page    string
	PerPage string
}

type Data struct {
	TotalRecords int         `json:"total_records"`
	CurrentPage  string      `json:"current_page"`
	TotalPages   int64       `json:"total_pages"`
	Records      interface{} `json:"records"`
}

func (p *Paginator) paginate(dataSource interface{}, field string, arg string) *Data {
	db := p.DB

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var output Data
	var count int
	var offset int64

	go countRecords(db, dataSource, done, &count, field, arg)

	if p.Page == "1" {
		offset = 0
	} else {
		tmpPage, _ := strconv.ParseInt(p.Page, 10, 32)
		tmpPerPage, _ := strconv.ParseInt(p.PerPage, 10, 32)
		offset = (tmpPage - 1) * tmpPerPage
	}

	if field != "" && arg != "" {
		query := fmt.Sprintf("%s = ?", field)
		db.Where(query, arg).Find(dataSource)
	} else {
		db.Limit(p.PerPage).Offset(offset).Find(dataSource)
	}

	<-done

	output.TotalRecords = count
	output.Records = dataSource
	output.CurrentPage = p.Page
	output.TotalPages = getTotalPages(p.PerPage, count)

	return &output
}

func countRecords(db *gorm.DB, countDataSource interface{}, done chan bool, count *int, field string, arg string) {
	if field != "" && arg != "" {
		query := fmt.Sprintf("%s = ?", field)
		db.Where(query, arg).Model(countDataSource).Count(count)
	} else {
		db.Model(countDataSource).Count(count)
	}

	done <- true
}

type Parameters struct {
	// Pagination page
	Page string
	//Results per page
	PerPage string
	//Sort field
	SortField string
	//Sort direction
	SortDirection string
}

type FindParameters struct {
	// Search criteria field
	Field string
	// Search criteria argument
	Arg string
}

func getTotalPages(perPage string, totalRecords int) int64 {
	perPageInt, _ := strconv.ParseInt(perPage, 10, 32)
	totalPages := float64(totalRecords) / float64(perPageInt)
	return int64(float64(totalPages) + float64(1.0))
}

func GetPage(db *gorm.DB, model interface{}, findParameters FindParameters, parameters Parameters) *Data {

	sort := fmt.Sprintf("%s %s ", parameters.SortField, parameters.SortDirection)

	orderBy := []string{sort}
	paginator := Paginator{DB: db, OrderBy: orderBy, Page: parameters.Page, PerPage: parameters.PerPage}
	result := paginator.paginate(model, findParameters.Field, findParameters.Arg)

	return result

}
