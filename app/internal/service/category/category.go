package category

import (
	"fmt"
	g "oj/app/global"
	"oj/app/internal/model"
	"sync"
)

type sCategory struct {
}

var (
	onceCategory sync.Once
	insCategory  *sCategory
)

//返回单例分类对象
func newSCategory() *sCategory {
	onceCategory.Do(func() {
		insCategory = &sCategory{}
	})
	return insCategory
}

func (s *sCategory) GetCategoryLists(page, size int, keyword string) ([]*model.Category, error) {
	sqlStr := "select id,name,parent_id,created_at from category where name like ? limit ?,?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[GetCategoryLists] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer stmt.Close()
	rows, err := stmt.Query(keyword, page, size)
	if err != nil {
		g.Logger.Errorf("[GetCategoryLists] querty failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	var categories []*model.Category
	for rows.Next() {
		var category model.Category
		err = rows.Scan(&category.Id, &category.Name, &category.ParentId, &category.CreatedAt)
		if err != nil {
			g.Logger.Errorf("[GetCategoryLists] scan failed,err:%v", err)
			return nil, fmt.Errorf("internal err")
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (s *sCategory) UpdateCategory(category model.Category) error {
	sqlStr := "update category set name =?,parent_id=? where id=?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[UpdateCategory] prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.Name, category.ParentId, category.Id)
	if err != nil {
		g.Logger.Errorf("[UpdateCategory] insert failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil

}

func (s *sCategory) CreateCategory(category model.Category) error {
	sqlStr := "insert into category(name,parent_id,created_at) values (?,?,?)"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[CreateCategory] prepare failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	defer stmt.Close()
	_, err = stmt.Exec(category.Name, category.ParentId, category.CreatedAt)
	if err != nil {
		g.Logger.Errorf("[CreateCategory] insert failed,err:%v", err)
		return fmt.Errorf("internal err")
	}
	return nil
}

func (s *sCategory) GetCategoryInfoById(id int) (*model.Category, error) {
	sqlStr := "select id,name,parent_id,created_at from category where id = ?"
	stmt, err := g.MysqlDB.Prepare(sqlStr)
	if err != nil {
		g.Logger.Errorf("[GetCategoryInfoById] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	var category model.Category
	err = stmt.QueryRow(id).Scan(&category.Id, &category.Name, &category.ParentId, &category.CreatedAt)
	if err != nil {
		if err.Error() == "record not find" {
			return nil, fmt.Errorf("category not exist")
		}
		g.Logger.Errorf("[GetCategoryInfoById] prepare failed,err:%v", err)
		return nil, fmt.Errorf("internal err")
	}
	defer stmt.Close()
	return &category, nil
}
