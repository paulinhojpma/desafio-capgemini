package database

import (
	"fmt"
	"testing"
)

func InitDataBase() (IDataBase, error) {
	DBOpt := &OptionsDBClient{
		URL:    "postgresql://root:luke@localhost:5434/app",
		Driver: "postgres",
	}
	db, err := DBOpt.ConfigDatabase()
	return *db, err

}

func TestCreateSequenceValid(t *testing.T) {
	db, _ := InitDataBase()
	sequence := &Sequence{
		IsValid: true,
		Letters: "DUHHDB,DUBUHD,UBUUHU,BHBDHH,UDBDUH,UHDDDD",
	}

	err := db.CreateSequence(sequence)
	if err != nil {
		t.Error("Expect nil, got ", err)
	}

}
func TestCreateSequenceInValid(t *testing.T) {
	db, _ := InitDataBase()
	sequence := &Sequence{
		IsValid: false,
		Letters: "DUHHDB,DUBUDD,UBUUHU,BHBDHH,UDBDUH,UHDHHD",
	}

	err := db.CreateSequence(sequence)
	if err != nil {
		t.Error("Expect nil, got ", err)
	}

}

func TestGetInfoSequence(t *testing.T) {
	db, _ := InitDataBase()

	info, err := db.GetInfoSequences()
	fmt.Println(info.Valids)
	fmt.Println(info.Invalids)
	fmt.Println(info.Ratio)
	if err == nil {
		t.Error("Expect nil, got ", err)
	}
}
