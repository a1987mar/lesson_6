package main

import (
	"fmt"
	"lesson4/pkg/documentstore"
	"lesson4/pkg/users"
	"log/slog"
)

func main() {
	slog.Info("App started")
	//marshalExample()
	//unmarshalExample()
	lesson6()
}

func marshalExample() {
	s := &documentstore.MyStruct{X: 15}
	doc, err := documentstore.MarshalDocument(s)
	if err != nil {
		fmt.Printf("failed to marshal document: %+v\n", err)
		return
	}
	fmt.Printf("marshaled document: %+v\n", doc)
}

func unmarshalExample() {
	doc := &documentstore.Document{Fields: map[string]documentstore.DocumentField{}}
	doc.Fields["X"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeNumber,
		Value: int(32),
	}

	s := &documentstore.MyStruct{}
	err := documentstore.UnmarshalDocument(doc, s)
	if err != nil {
		fmt.Printf("failed to unmarshal document: %+v\n", err)
		return
	}
	fmt.Printf("unmarshaled document: %+v\n", s)
}

func lesson6() {

	slog.Info("Creating user")
	s := users.NewService()
	d1 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d1.Fields["id-1"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "setup.exe",
	}

	cfg1 := documentstore.CollectionConfig{PrimaryKey: "id-1"}
	_, err := s.CreateUser("id-1", "UserTest-1", cfg1, &d1)
	if err != nil {
		slog.Error("Failed to create", err)
	}
	slog.Info("user created successfully")

	slog.Info("Creating user")
	s2 := users.NewService()
	d2 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d2.Fields["id-2"] = documentstore.DocumentField{
		Type:  documentstore.DocumentFieldTypeString,
		Value: "main.go",
	}
	cfg2 := documentstore.CollectionConfig{PrimaryKey: "id-2"}
	_, err = s2.CreateUser("id-2", "UserTest-2", cfg2, &d2)
	if err != nil {
		slog.Error("Failed to create", err)
	}
	slog.Info("user created successfully")

	//d := []byte{123, 10, 32, 34, 99, 111, 108, 108, 101, 99, 116, 105, 111, 110, 115, 34, 58, 32, 123, 10, 32, 34, 105, 100, 45, 49, 34, 58, 32, 123, 10, 32, 34, 100, 111, 99, 117, 109, 101, 110, 116, 115, 34, 58, 32, 123, 10, 32, 34, 105, 100, 45, 49, 34, 58, 32, 123, 10, 32, 34, 102, 105, 101, 108, 100, 115, 34, 58, 32, 123, 10, 32, 34, 105, 100, 45, 49, 34, 58, 32, 123, 10, 32, 34, 116, 121, 112, 101, 34, 58, 32, 34, 115, 116, 114, 105, 110, 103, 34, 44, 10, 32, 34, 118, 97, 108, 117, 101, 34, 58, 32, 34, 115, 101, 116, 117, 112, 46, 101, 120, 101, 34, 10, 32, 125, 10, 32, 125, 10, 32, 125, 10, 32, 125, 44, 10, 32, 34, 99, 111, 110, 102, 105, 103, 34, 58, 32, 123, 10, 32, 34, 99, 103, 103, 34, 58, 32, 34, 105, 100, 45, 49, 34, 10, 32, 125, 10, 32, 125, 10, 32, 125, 10, 32, 125}
	//
	//doc, err := documentstore.NewStoreFromDump(d)
	//if err != nil {
	//	zap.Error(err)
	//}
	//
	//fmt.Println(doc.Dump())
	sF, err := documentstore.NewStoreFromFile("id-2")
	if err != nil {
		slog.Info(err.Error())
	}
	fmt.Printf("%+v \n", sF)
	slog.Info("App done")
}
