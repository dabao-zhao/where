package where

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestEq(t *testing.T) {
	expr := Eq{"cate": "123"}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate = ?", query)
	assert.Equal(t, []interface{}{"123"}, args)

	expr = Eq{
		"cate": "123",
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate = ? AND id = ?", query)
	assert.Equal(t, []interface{}{"123", "456"}, args)
}

func TestNeq(t *testing.T) {
	expr := Neq{"cate": "123"}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate != ?", query)
	assert.Equal(t, []interface{}{"123"}, args)

	expr = Neq{
		"cate": "123",
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate != ? AND id != ?", query)
	assert.Equal(t, []interface{}{"123", "456"}, args)
}

func TestGt(t *testing.T) {
	expr := Gt{"cate": "123"}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate > ?", query)
	assert.Equal(t, []interface{}{"123"}, args)

	expr = Gt{
		"cate": "123",
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate > ? AND id > ?", query)
	assert.Equal(t, []interface{}{"123", "456"}, args)
}

func TestGtOrEq(t *testing.T) {
	expr := GtOrEq{"cate": "123"}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate >= ?", query)
	assert.Equal(t, []interface{}{"123"}, args)

	expr = GtOrEq{
		"cate": "123",
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate >= ? AND id >= ?", query)
	assert.Equal(t, []interface{}{"123", "456"}, args)
}

func TestLt(t *testing.T) {
	expr := Lt{"cate": "123"}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate < ?", query)
	assert.Equal(t, []interface{}{"123"}, args)

	expr = Lt{
		"cate": "123",
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate < ? AND id < ?", query)
	assert.Equal(t, []interface{}{"123", "456"}, args)
}

func TestLtOrEq(t *testing.T) {
	expr := LtOrEq{"cate": "123"}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate <= ?", query)
	assert.Equal(t, []interface{}{"123"}, args)

	expr = LtOrEq{
		"cate": "123",
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate <= ? AND id <= ?", query)
	assert.Equal(t, []interface{}{"123", "456"}, args)
}

func TestIn(t *testing.T) {
	expr := In{"cate": []string{"123", "456"}}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate IN (?)", query)
	assert.Equal(t, []interface{}{[]string{"123", "456"}}, args)

	expr = In{
		"cate": []string{"123", "456"},
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate IN (?) AND id IN (?)", query)
	assert.Equal(t, []interface{}{[]string{"123", "456"}, "456"}, args)
}

func TestLike(t *testing.T) {
	expr := Like{"cate": "tmp"}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "cate LIKE ?", query)
	assert.Equal(t, []interface{}{"tmp"}, args)

	expr = Like{
		"cate": "tmp",
		"id":   "456",
	}
	query, args = expr.toQueryAndArgs()
	assert.Equal(t, "cate LIKE ? AND id LIKE ?", query)
	assert.Equal(t, []interface{}{"tmp", "456"}, args)
}

func TestOr(t *testing.T) {
	expr := Or{Eq{"name": "1"}, Eq{"name": "2"}}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "(name = ? OR name = ?)", query)
	assert.Equal(t, []interface{}{"1", "2"}, args)
}

func TestAnd(t *testing.T) {
	expr := And{Eq{"name": "1"}, Eq{"name": "2"}}
	query, args := expr.toQueryAndArgs()
	assert.Equal(t, "(name = ? AND name = ?)", query)
	assert.Equal(t, []interface{}{"1", "2"}, args)
}

func TestBetween(t *testing.T) {
	between := Between{
		"date":  []interface{}{"2019-12", "2023-4"},
		"money": []interface{}{100, 1000},
	}
	query, args := between.toQueryAndArgs()
	assert.Equal(t, "date BETWEEN ? AND ? AND money BETWEEN ? AND ?", query)
	assert.Equal(t, []interface{}{"2019-12", "2023-4", 100, 1000}, args)
}

func TestToQueryAndArgs(t *testing.T) {
	exprs := []Expr{
		In{"cate": []string{"123", "456"}},
		Or{Eq{"name": "1"}, Eq{"name": "2"}},
		GtOrEq{"cate": "123"},
	}
	query, args := ToQueryAndArgs(exprs)
	assert.Equal(t, "cate IN (?) AND (name = ? OR name = ?) AND cate >= ?", query)
	assert.Equal(t, []interface{}{[]string{"123", "456"}, "1", "2", "123"}, args)

	exprs = []Expr{
		In{"cate": []string{"123", "456"}},
	}
	query, args = ToQueryAndArgs(exprs)
	assert.Equal(t, "cate IN (?)", query)
	assert.Equal(t, []interface{}{[]string{"123", "456"}}, args)

	query, args = ToQueryAndArgs(nil)
	assert.Equal(t, "", query)
	assert.Nil(t, args)
}

func TestGetSortedKeys(t *testing.T) {
	v := map[string]interface{}{
		"a":  1,
		"z":  1,
		"c":  1,
		"w":  1,
		"ay": 1,
	}

	s := getSortedKeys(v)
	assert.Equal(t, []string{"a", "ay", "c", "w", "z"}, s)
}

func TestGetSortedKeys2(t *testing.T) {
	v := map[string][]interface{}{
		"a":  nil,
		"z":  nil,
		"c":  nil,
		"w":  nil,
		"ay": nil,
	}

	s := getSortedKeys2(v)
	assert.Equal(t, []string{"a", "ay", "c", "w", "z"}, s)
}

func TestGorm(t *testing.T) {
	db, _, _ := getDBMock()

	var where []Expr
	where = append(where, Eq{"name": "da-bao"})
	where = append(where, Neq{"cate": 0})
	where = append(where, GtOrEq{"age": 10})
	where = append(where, Gt{"width": 100})
	where = append(where, Lt{"height": 200})
	where = append(where, LtOrEq{"max": 300})
	where = append(where, In{"num": []string{"1", "2", "3"}})
	where = append(where, Like{"hobby": "play"})
	where = append(where, Or{Eq{"sex": 1}, Eq{"sex": 2}})
	where = append(where, And{Eq{"period": 1}, Eq{"period_unit": 2}})
	where = append(where, Between{"date": []interface{}{"2019-12", "2023-4"}})

	query, args := ToQueryAndArgs(where)
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Table("test").Where(query, args...).Find(nil)
	})

	assert.Equal(t, "SELECT * FROM `test` WHERE "+
		"name = 'da-bao' AND "+
		"cate != 0 AND "+
		"age >= 10 AND "+
		"width > 100 AND "+
		"height < 200 AND "+
		"max <= 300 AND "+
		"num IN ('1','2','3') AND "+
		"hobby LIKE 'play' AND "+
		"(sex = 1 OR sex = 2) AND "+
		"(period = 1 AND period_unit = 2) AND "+
		"date BETWEEN '2019-12' AND '2023-4'", sql)
}

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	// mock一个*sql.DB对象，不需要连接真实的数据库
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	//测试时不需要真正连接数据库
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	return gdb, mock, nil
}
