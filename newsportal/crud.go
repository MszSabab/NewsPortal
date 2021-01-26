package newsportal

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

//db connection
func newDBSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	return session, nil
}

//CreateNewsCrud ----
func CreateNewsCrud(portal NewsPortal) (NewsPortal, error) {
	portal.ID = generateUUID()
	session, err := newDBSession()
	if err != nil {
		return NewsPortal{}, nil
	}

	coll := session.DB("NewsPortal").C("NP-table")
	err = coll.Insert(&portal)

	if err != nil {
		return NewsPortal{}, nil
	}
	return portal, nil

}

//GetAllNewsCrud ----
func GetAllNewsCrud() ([]NewsPortal, error) {
	var portal []NewsPortal
	session, err := newDBSession()
	if err != nil {
		return []NewsPortal{}, nil
	}
	coll := session.DB("NewsPortal").C("NP-table")
	err = coll.Find(bson.M{}).All(&portal)

	if err != nil {
		return []NewsPortal{}, nil
	}

	return portal, nil
}

func generateUUID() string {
	return uuid.New().String()
}
