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

//GetSingleNewsCrud ----
func GetSingleNewsCrud(portalID string) (NewsPortal, error) {
	var portal NewsPortal
	session, err := newDBSession()
	if err != nil {
		return NewsPortal{}, nil
	}
	coll := session.DB("NewsPortal").C("NP-table")
	err = coll.Find(bson.M{"_id": portalID}).One(&portal)

	if err != nil {
		return NewsPortal{}, nil
	}

	return portal, nil
}

//DeleteNewsCrud ----
func DeleteNewsCrud(portalID string) error {

	session, err := newDBSession()
	if err != nil {
		return nil
	}
	coll := session.DB("NewsPortal").C("NP-table")
	err = coll.Remove(bson.M{"_id": portalID})

	if err != nil {
		return err
	}
	return nil
}

//UpdateNewsCrud ----
func UpdateNewsCrud(portalID string, portal NewsPortal) (NewsPortal, error) {

	session, err := newDBSession()
	if err != nil {
		return NewsPortal{}, nil
	}
	portal.ID = portalID
	coll := session.DB("NewsPortal").C("NP-table")
	selector := bson.M{"_id": portalID}
	err = coll.Update(selector, bson.M{"$set": portal})
	if err != nil {
		return NewsPortal{}, err
	}
	Updated, err := GetSingleNewsCrud(portalID)
	if err != nil {
		return NewsPortal{}, err
	}
	// fmt.Println("Updated===>", Updated)
	return Updated, nil

}
func generateUUID() string {
	return uuid.New().String()
}
