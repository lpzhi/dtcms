package models

type AuthMenu struct {
	Model
	Pid int			`json:"pid"`
	Name string 	`json:"name"`
	Sort int 		`json:"sort"`
	Route string	`json:"route"`
}

type TreeList struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Pid int			`json:"pid"`
	Sort int 		`json:"sort"`
	Route string	`json:"route"`
	Children []*TreeList	`json:"children"`
}

func (m *AuthMenu) MenuList() []*TreeList {
	 return  m.getMenu(0)
}



func (m *AuthMenu)getMenu(pid int)  []*TreeList{
	var menu []AuthMenu
	db.Debug().Where(&AuthMenu{Pid:pid}).Order("id desc").Find(&menu)

	treeList := []*TreeList{}

	for _, v := range menu{
		child := v.getMenu(v.ID)
		node := &TreeList{
			Id:v.ID,
			Name:v.Name,
			Sort:v.Sort,
			Route:v.Route,
			Pid:v.Pid,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList

	return nil
}