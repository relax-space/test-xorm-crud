package main

type Fruit struct {
	Id     int             `json:"id" xorm:"pk autoincr"`
	Code string `json:"code" xorm:"unique"`
}

func (d *Fruit) Update(id int) (int64, error) {
	return Db.Where("id=?", id).Update(d)
}

func (d *Fruit) Create() (int64, error) {
	return Db.Insert(d)
}