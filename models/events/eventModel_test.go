package eventModel

import "testing"

func TestSave(t *testing.T) {
	l := len(data)
	m := &EventModel{
		Id:               2,
		Title:            "Some Title 3",
		ShortDescription: "Short 3",
		Description:      "Longer then short 3",
		Coords:           Coords{1, 1},
		Images:           []string{"img5", "img6"},
		Preview:          "Don't know what is a preview",
	}
	_, err1 := Save(m)
	if err1 == nil {
		t.Errorf("Already exist error was expected")
	}
	m.Id = 3
	_, err2 := Save(m)
	if err2 != nil {
		t.Errorf(err2.Error())
	}
	if l+1 != len(data) {
		t.Errorf("len(data) = %v; want %v", len(data), l+1)
	}
}

func TestDelete(t *testing.T) {
	l := len(data)
	_, err1 := Delete(5)
	if err1 == nil {
		t.Errorf("Not found error was expected")
	}
	_, err2 := Delete(1)
	if err2 != nil {
		t.Errorf(err2.Error())
	} else if l-1 != len(data) {
		t.Errorf("len(data) = %v; want %v", len(data), l-1)
	}
	if id := data[0].Id; id != 2 {
		t.Errorf("Expected data id == 2, got %v", id)
	}
}

func TestGetAll(t *testing.T) {
	v, _ := GetAll()
	if len(v) != 2 {
		t.Errorf("Expected len(data) == 2, got %v", len(v))
	}
}

func TestGetOne(t *testing.T) {
	_, err1 := GetOne(5)
	if err1 == nil {
		t.Errorf("Not found error was expected")
	}
	v, err2 := GetOne(2)
	if err2 != nil {
		t.Errorf(err2.Error())
	} else if v.Id != 2 {
		t.Errorf("Expexted id to be 2, got %v", v.Id)
	}
}
