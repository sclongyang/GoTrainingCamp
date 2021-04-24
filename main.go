package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	//模拟DAO层查询
	result, err := dao()
	if errors.Cause(err) == sql.ErrNoRows {
		//TODO
	}else if err != nil{
		//TODO
	}
	fmt.Println("%+v", result) //模拟处理查询结果
}

func dao() (result []interface{}, err error) {
	result = []interface{}{}
	err = query()

	//1.若需求是返回空slice即可, 则降级返回nil
	//if err == sql.ErrNoRows{
	//	err = nil
	//	return
	//}

	//2.若需求是:上层要处理该err, 则需wrap后抛给上层, 为了既携带上下文和堆栈方便人阅读, 又携带原始err才能==判断
	if err != nil {
		err = errors.Wrap(err, "from dao")
		return
	}

	//doSomething
	return
}

func query() error {
	return sql.ErrNoRows
}

