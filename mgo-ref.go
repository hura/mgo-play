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

type Book struct {
	Id        bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	Title     string          `bson:"t"             json:"title"`
	AuthorIds []bson.ObjectId `bson:"AuthorIds"      json:"authorids"`
}

type AssembledBooks struct {
	Book
	Authors []Author
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
	cbooks := db.C("bookcoll")
	cauthors := db.C("authorcoll")

	aids := []bson.ObjectId{bson.NewObjectId(), bson.NewObjectId()}
	authors := []Author{{
		aids[0],
		"Author 1",
	}, {
		aids[1],
		"Author 2",
	}}
	cauthors.Insert(authors[0])
	cauthors.Insert(authors[1])
	// Insert some books
	mine := Book{
		bson.NewObjectId(),
		"Gang of four thingy",
		aids,
	}
	cbooks.Insert(&mine)

	var assembl []AssembledBooks
	cauthors.Find(bson.M{}).All(&assembl)
	str1, _ := json.MarshalIndent(assembl, "", " ")
	fmt.Printf("%s\n", str1)

	var allauthors []Author
	cauthors.Find(bson.M{}).All(&allauthors)
	str, _ := json.MarshalIndent(allauthors, "", " ")
	fmt.Printf("%s\n", str)
	var allbooks []Book
	cbooks.Find(bson.M{}).All(&allbooks)
	str, _ = json.MarshalIndent(allbooks, "", " ")
	fmt.Printf("%s\n", str)

	fmt.Println("Dropping all collections...")
	cauthors.DropCollection()
	cbooks.DropCollection()
	fmt.Println("Done")
}
