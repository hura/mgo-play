// vim: ts=2 sw=2 sts=2 sta ai si noci
package main

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

type Author struct {
	Id   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"t"             json:"name"`
}

type Books []Book
type Book struct {
	// omitempty so mongodb itself will generate the id if omitted
	Id      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title   string        `bson:"t"             json:"title"`
	Authors []Author      `bson:"i"             json:"chapters"`
}

func main() {
	logout := log.New(os.Stdout, "MGO: ", log.Lshortfile)
	mgo.SetLogger(logout)
	mgo.SetDebug(false)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db := session.DB("bookdb")
	coll := db.C("bookcoll")

	// Insert some books
	mine := Book{
		bson.NewObjectId(),
		"Gang of four thingy",
		[]Author{{
			bson.NewObjectId(),
			"Author 1",
		}, {
			bson.NewObjectId(),
			"Author 2",
		}},
	}
	coll.Insert(&mine)

	var alldocs Books
	coll.Find(bson.M{}).All(&alldocs)
	str, _ := json.MarshalIndent(alldocs, "", " ")
	fmt.Printf("%s\n", str)

	c, err := coll.Count()
	fmt.Printf("Total documents: %d\n", c)
	fmt.Println("Dropping entire collection...")
	coll.DropCollection()
	fmt.Println("Done")
}
