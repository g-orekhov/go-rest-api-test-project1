package db

// Таблицу создал ещё до того как заюзал GORM
// так что надо сразу учится адапитровать GORM под кастомные таблицы.
//
// CREATE TABLE IF NOT EXISTS public.events
// (
//     id bigint NOT NULL DEFAULT nextval('events_id_seq'::regclass),
//     title character varying(50) COLLATE pg_catalog."default",
//     short_desc text COLLATE pg_catalog."default",
//     description text COLLATE pg_catalog."default",
//     images character varying(40)[] COLLATE pg_catalog."default",
//     preview text COLLATE pg_catalog."default",
//     coords geography(Point,4326),
//     CONSTRAINT id PRIMARY KEY (id)
// )

import (
	"database/sql/driver"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"gorm.io/gorm/clause"
)

/*
Geo coords record in DB - POINT(LONG, LAT)
Methods:
String() string - retuns point string to store in format "SRID=4326;POINT(LONG, LAT)"
Scan(any) error - scans values from DB respons
Value() string -  returns point to store in DB
*/
type Coords struct {
	Long float64 `json:"long"`
	Lat  float64 `json:"lat"`
}

func (c Coords) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", c.Long, c.Lat)
}

func (c *Coords) Scan(src any) error {
	// В БД формат point(Lon,Lat)::geometry
	// по умолчанию БД возврашает точку в формате - EWKB или что-то вроде (гугл не помогает)
	// Структура кторую удалось распарсить:
	// первый байт - порядок байт Big/Little-endian (у нас он 1 - little-endian)
	// 4 байта - тип обьекта (но у нас почему то 3 байта тип, а последний байт Х/З что)
	// 4 байта - SRID
	// 8 + 8 байт - значения точки
	// Значения читаю как UINT и затем конвертирую в FLOAT64. ParseFloat выдаёт бред какой-то

	//Интерфейс src - string, но нужно поискать более правильный способ работы с байтами
	str := src.(string)
	if len(str) != 50 {
		//TODO: может ли точка иметь другой размер, что в src - если в бд null? Узнать - обработать.
		c.Lat = 0
		c.Long = 0
	}
	reverse := func(str string, start int, bytes int) string {
		var ret string
		for i := start + (bytes * 2) - 2; i >= start; i -= 2 {
			ret += str[i : i+2]
		}
		return ret
	}
	//TODO: Разобраться возможны ли тут ошибки? Если да - обработать.
	xUint, _ := strconv.ParseUint(reverse(str, 18, 8), 16, 64)
	yUint, _ := strconv.ParseUint(reverse(str, 34, 8), 16, 64)
	c.Long = math.Float64frombits(xUint)
	c.Lat = math.Float64frombits(yUint)
	return nil
}

func (c Coords) Value() (driver.Value, error) {
	return c.String(), nil
}

// У массива по умолчанию вид [1,2] а в PostgreSQL надо {1,2}
// Загуглил решение с кучей гемороя в запросах
// не нашел решения и решил что лучше реализовать интерфейсы Scaner/Valuer
// скажите мне если есть более простое решение
type Images []string

func (im Images) String() string {
	ret := strings.Join(im, ",")
	return "{" + ret + "}"
}

func (im *Images) Scan(src any) error {
	str := src.(string)
	str = str[1 : len(str)-1]
	if len(str) > 0 {
		*im = strings.Split(str, ",")
	} else {
		*im = make([]string, 0)
	}
	return nil
}

func (im Images) Value() (driver.Value, error) {
	return im.String(), nil
}

type Event struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"shortDescription" gorm:"column:short_desc"`
	Description      string `json:"description"`
	Coords           Coords `json:"coords"`
	Images           Images `json:"images"`
	Preview          string `json:"preview"`
}

func (db *DB) EventCreate(event *Event) (*Event, error) {
	result := db.pool.Create(event)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errDataBase
	}
	return event, nil
}

func (db *DB) EventGetAll() ([]Event, error) {
	//TODO: retrieve pagination data
	count := 10
	events := make([]Event, 0, count)
	result := db.pool.Find(&events)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errDataBase
	}
	return events, nil
}

func (db *DB) EventGetOne(id int) (*Event, error) {
	event := new(Event)
	result := db.pool.Where("id = ?", id).First(event)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errNotFound
	}
	if result.RowsAffected > 0 {
		log.Println(event)
		return event, nil
	}
	return nil, nil
}

func (db *DB) EventDelete(id int) (*Event, error) {
	event := new(Event)
	result := db.pool.Clauses(clause.Returning{}).Delete(event, id)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Println(result.Error)
		return nil, errNotFound
	}
	return event, nil
}
