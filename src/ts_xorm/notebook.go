package ts_xorm

type Notebook struct {
	Name         string    `xorm:"varchar(255) 'name'"`
	Namespace    string    `xorm:"varchar(255) 'namespace'"`
	Image        string    `xorm:"-"`
	MountStorage string    `xorm:"-"`
	Resource     *Resource `xorm:"-"`
	Id           int64     `xorm:"pk autoincr comment('ID') bigint 'id'"`
	DisplayName  string    `xorm:"varchar(255) 'display_name'"`
	NbcId        int64     `xorm:"not null bigint 'nbc_id'"`
}

func (n *Notebook) TableName() string {
	return "notebook"
}

func (n *Notebook) IsInNBCollection() bool {
	return n.NbcId != 0
}

type Resource struct {
	Cpu    string
	Memory string
	Gpus   *Gpus
}

type Gpus struct {
	Num    string
	Vendor string
}
