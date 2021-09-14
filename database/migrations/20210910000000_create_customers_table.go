package migrations

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Customer struct {
	ID        uint64 `gorm:"primary_key;auto_increment"`
	Name      string `gorm:"type:varchar(100);index:name"`
	Age       uint32 `gorm:"type:int(5);index:age"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"type:varchar(255);index:password"`
	Address   string `gorm:"->:false;<-:create;type:varchar(255);index:address"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func Up_20210910000000(txn *gorm.DB) {
	txn.CreateTable(&Customer{})
}

func Down_20210910000000(txn *gorm.DB) {
	txn.DropTable(&Customer{})
}
