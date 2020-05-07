package model

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Record interface {
	IsNewRecord() bool
	IsValid() bool
	C() Collection
	id() bson.ObjectId

	setIsNewRocord(bool)
	setCreatedAt(time.Time)
	setUpdatedAt(time.Time)
}

type Collection interface {
	Query(query func(c *mgo.Collection))
}

func Save(r Record, m ...bson.M) (err error) {
	if !r.id().Valid() {
		return IDInvalidErr
	}

	if !r.IsValid() {
		return RecordInvalidErr
	}

	r.C().Query(func(c *mgo.Collection) {
		t := time.Now()
		if r.IsNewRecord() {
			r.setCreatedAt(t)
			r.setUpdatedAt(t)

			err = c.Insert(r)
			if err == nil {
				r.setIsNewRocord(false)
			}
		} else {
			var update bson.M

			if len(m) > 0 {
				update = m[0]
			} else {
				update = bson.M{}
			}

			update["updated_at"] = t

			err = c.UpdateId(r.id(), bson.M{
				"$set": update,
			})
		}
	})

	return
}

func Destroy(r Record) (err error) {
	r.C().Query(func(mc *mgo.Collection) {
		err = mc.RemoveId(r.id())
	})

	return
}

func FindById(c Collection, id string, r interface{}) (err error) {
	if !bson.IsObjectIdHex(id) {
		return IDInvalidErr
	}

	c.Query(func(mc *mgo.Collection) {
		err = mc.FindId(bson.ObjectIdHex(id)).One(r)
	})

	return
}

func FindOne(c Collection, query bson.M, r interface{}) (err error) {
	c.Query(func(mc *mgo.Collection) {
		err = mc.Find(query).One(r)
	})
	return
}

func FindAll(c Collection, query bson.M, r interface{}) (err error) {
	c.Query(func(mc *mgo.Collection) {
		err = mc.Find(query).All(r)
	})
	return
}

func Update(r Record, query bson.M, update bson.M) (err error) {
	if !r.IsNewRecord() {
		return RecordInvalidErr
	}

	r.C().Query(func(mc *mgo.Collection) {
		err = mc.Update(query, bson.M{
			"$set": update,
		})
	})
	return
}

func FindByLimit(c Collection, query bson.M, total int, r interface{}, sort ...string) (err error) {
	sortStr := "-_id" //按数据库id降序
	if len(sort) > 0 {
		sortStr = sort[0]
	}

	c.Query(func(mc *mgo.Collection) {
		if total == -1 {
			err = mc.Find(query).Sort(sortStr).All(r)
		} else {
			err = mc.Find(query).Sort(sortStr).Limit(total).All(r)
		}
	})
	return
}

func FindByPage(c Collection, perPage, pageNum int, query bson.M, r interface{}, sortStr ...string) (err error) {
	skipNum := 0
	if pageNum > 0 {
		skipNum = (pageNum - 1) * perPage
	}

	var sort string
	if len(sortStr) == 0 {
		sort = "-_id"
	} else {
		sort = sortStr[0]
	}

	c.Query(func(mc *mgo.Collection) {
		err = mc.Find(query).Skip(skipNum).Limit(perPage).Sort(sort).All(r)
	})
	return
}

func InsertBatch(c Collection, docs ...interface{}) (err error) {
	c.Query(func(mc *mgo.Collection) {
		err = mc.Insert(docs)
	})
	return
}

func DestroyBatch(c Collection, query bson.M) (err error) {
	c.Query(func(mc *mgo.Collection) {
		_, err = mc.RemoveAll(query)
	})

	return
}

func CountBy(c Collection, query bson.M) (err error, num int) {
	c.Query(func(mc *mgo.Collection) {
		num, err = mc.Find(query).Count()
	})

	return
}
