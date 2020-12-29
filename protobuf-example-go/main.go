package main

import (
	"fmt"
	"io/ioutil"
	"log"

	addressbookpb "github.com/protobuf-example-go/src/addressbook"

	complexpb "github.com/protobuf-example-go/src/complex"

	enumpb "github.com/protobuf-example-go/src/enum_example"

	simplepb "github.com/protobuf-example-go/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	// sm := doSimple()
	// readAndWriteDemo(sm)
	// jsonDemo(sm)

	// doEnum()
	// doComplex()

	readWriteAddressbookDemo()

}

func readWriteAddressbookDemo() {

	addressbook := addPersonToBook()

	writeToFile("addressbook.bin", addressbook)

	addressbook2 := &addressbookpb.AddressBook{}
	readFromFile("addressbook.bin", addressbook2)
	fmt.Println(addressbook2)

}

func addPersonToBook() *addressbookpb.AddressBook {
	person := addressbookpb.Person{
		Id:    1,
		Name:  "John",
		Email: "john@gmail.com",
		Phones: []*addressbookpb.Person_PhoneNumber{
			{Number: "123-456-789",
				Type: addressbookpb.Person_MOBILE},
		},
	}

	addressbook := addressbookpb.AddressBook{
		People: []*addressbookpb.Person{
			&person,
		},
	}

	return &addressbook
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_SUNDAY
	fmt.Println(em)
}

func jsonDemo(pb proto.Message) {
	smAsString := toJSON(pb)

	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println(sm2)
}

func toJSON(pb proto.Message) string {
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSOn", err)
		return ""
	}
	return string(out)
}

func fromJSON(in string, pb proto.Message) error {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("Can't unmarshal json", err)
		return err
	}
	return nil
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data is written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Can't read bytes", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My message",
		SampleList: []int32{1, 4, 7},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	return &sm

}
