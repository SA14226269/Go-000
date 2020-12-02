// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// dao层只考虑操作数据，处理数据及异常应该放在上层业务处理，应该抛给业务层，考虑业务逻辑与底层的解耦合，上层不关注具体何种数据库出错，应该将数据库错误wrap成一个自定义类似404notfound的标准数据库查询失败错误抛给上层处理。
// Dao query data info
package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

var SqlNotFound = errors.New("sqldata not found")

type Dao interface {
	Content(name string) (addr string, err error)
}

func (d *dao) Content(name string) (addr string, err error) {
	stmt, err := d.db.Prepare("select...")
	if err != nil {
		err = errors.Wrap(err, "prepare failed")
		return
	}

	err = stmt.QueryRow(name).Scan(&addr)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = errors.Wrap(SqlNotFound, "query failed")
		return
	}
	if err != nil {
		err = errors.Wrap(err, "query failed")
	}
	return
}
