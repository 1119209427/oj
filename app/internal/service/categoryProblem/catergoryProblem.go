package categoryProblem

import (
	"fmt"
	g "oj/app/global"
	"oj/app/internal/model"
	"sync"
)

type sCategoryProblem struct{}

var (
	insCategoryProblem  *sCategoryProblem
	onceCategoryProblem sync.Once
)

func newCategoryProblemService() *sCategoryProblem {
	onceCategoryProblem.Do(
		func() {
			insCategoryProblem = &sCategoryProblem{}
		})
	return insCategoryProblem
}

func (s *sCategoryProblem) GetCategoryIdByProblemId(Pid int) ([]int, error) {
	sqlStr := "select id,category_id,problem_id from category_problem where problem_id = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[GetCategoryIdByProblemId] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer stmt.Close()
	rows, err := stmt.Query(Pid)
	if err != nil {
		if err.Error() == "recorder not found" {
			return nil, fmt.Errorf("category not exist")
		}
		g.Logger.Errorf("[GetCategoryIdByProblemId] query failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var cProblem model.CategoryProblem
		err = rows.Scan(&cProblem.Id, &cProblem.CategoryId, &cProblem.ProblemId)
		if err != nil {
			g.Logger.Errorf("[GetCategoryIdByProblemId] scan failed,err:%v", err)
			return nil, fmt.Errorf("internal err")
		}
		ids = append(ids, cProblem.CategoryId)
	}
	return ids, nil
}
