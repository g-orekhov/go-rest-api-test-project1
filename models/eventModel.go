package eventModel

import (
	"errors"
	"math/rand"
)

type Coords struct {
	Long float64 `json:"long"`
	Lat  float64 `json:"lat"`
}

type EventModel struct {
	Id               int      `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	Description      string   `json:"description"`
	Coords           Coords   `json:"coords"`
	Images           []string `json:"images"`
	Preview          string   `json:"prewiew"`
}

func Create(m *EventModel) (*EventModel, error) {
	for _, v := range data {
		if v.Id == m.Id {
			return nil, errors.New("object with given ID is allready exist")
		}
	}
	data = append(data, m)
	return m, nil
}

func Update(m *EventModel) (*EventModel, error) {
	for i, v := range data {
		if v.Id == m.Id {
			data[i] = m
			return data[i], nil
		}
	}
	return nil, errors.New("object with given ID is not exist")
}

func Delete(id int) (*EventModel, error) {
	for i, v := range data {
		if v.Id == id {
			deletedObject := data[i]
			copy(data[i:], data[i+1:])
			data = data[:len(data)-1]
			return deletedObject, nil
		}
	}
	return nil, errors.New("object with given ID is not exist")
}

func GetAll() ([]*EventModel, error) {
	return data, nil
}

func GetOne(id int) (*EventModel, error) {
	for i, v := range data {
		if v.Id == id {
			return data[i], nil
		}
	}
	return nil, errors.New("object with given ID is not exist")
}

func GetUniqueId() int {
OUTER:
	for {
		id := rand.Int()
		for _, v := range data {
			if v.Id == id {
				continue OUTER
			}
		}
		return id
	}
}

/* - - - - init section - - - - */
var data []*EventModel

func init() {
	data = []*EventModel{
		{
			Id:               1,
			Title:            "Some Title 1",
			ShortDescription: "Short 1",
			Description:      "Longer then short 1",
			Coords:           Coords{Long: 1.0, Lat: 2.2},
			Images:           []string{"img1", "img2"},
			Preview:          "Don't know what is a preview",
		},
		{
			Id:               2,
			Title:            "Some Title 2",
			ShortDescription: "Short 2",
			Description:      "Longer then short 2",
			Coords:           Coords{1, 1},
			Images:           []string{"img3", "img4"},
			Preview:          "Don't know what is a preview",
		},
	}
}
