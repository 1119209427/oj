package testCase

import (
	"fmt"
	g "oj/app/global"
	"oj/app/internal/model"
	"sync"
)

type sTestCase struct{}

var (
	onceTestCase sync.Once
	insTestCase  *sTestCase
)

func newTestCaseService() *sTestCase {
	onceTestCase.Do(
		func() {
			insTestCase = &sTestCase{}
		})
	return insTestCase
}

func (s *sTestCase) GetTestCase(problemId int) ([]*model.TestCase, error) {
	//goland:noinspection SqlResolve
	sqlStr := "select id,problem_id,input,output from test_case where id = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[GetTestCase] prepare failed:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer stmt.Close()
	var testCases []*model.TestCase
	rows, err := stmt.Query(problemId)
	if err != nil {
		g.Logger.Errorf("[GetTestCase] query failed:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	for rows.Next() {
		var testCase model.TestCase
		err = rows.Scan(&testCase.Id, &testCase.ProblemId, &testCase.Input, &testCase.Output)
		if err != nil {
			g.Logger.Errorf("[GetTestCase] scan failed:%v", err)
			return nil, fmt.Errorf("internal err")
		}
		testCases = append(testCases, &testCase)
	}
	return testCases, nil

}

func (s *sTestCase) CreateTestCase(testCase *model.TestCase) error {
	//goland:noinspection SqlResolve
	sqlStr := "insert into test_case(problem_id,input,output) values(?,?,?)"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CreateTestCase] prepare failed:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(testCase.ProblemId, testCase.Input, testCase.Output)
	if err != nil {
		g.Logger.Errorf("[CreateTestCase] insert failed:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sTestCase) UpdateTestCase(testCase *model.TestCase) error {
	//goland:noinspection SqlResolve
	sqlStr := "update test_case set problem_id=?,input=?,output=? where id=?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[UpdateTestCase prepare failed:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(testCase.ProblemId, testCase.Input, testCase.Output, testCase.Id)
	if err != nil {
		g.Logger.Errorf("[UpdateTestCase] insert failed:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}
