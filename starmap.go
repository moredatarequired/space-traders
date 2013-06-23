package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

type Star struct {
	Id int
	Name, Class string
	X, Y, Z, Magnitude float64
}

func atof(s string) (float64, error) { return strconv.ParseFloat(s, 64) }

func newStar(s []string) (*Star, error) {
	star := new(Star)
	var err error
	if star.Id, err = strconv.Atoi(s[0]); err != nil { return nil, err }
	star.Name = s[3]
	star.Class = s[11]
	if star.X, err = atof(s[13]); err != nil { return nil, err }
	if star.Y, err = atof(s[14]); err != nil { return nil, err }
	if star.Z, err = atof(s[15]); err != nil { return nil, err }
	if star.Magnitude, err = atof(s[16]); err != nil { return nil, err }
	return star, nil
}

func readFromFile(f string) (err error) {
	reader, err := os.Open(f)
	if err != nil { return err }
	defer reader.Close()
	r := csv.NewReader(reader)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		if record[2] != "" {
			star, _ := newStar(record)
			fmt.Println(star)
		}
	}
	return err
}

func main() {
	readFromFile("HabHYG.csv")

	return
}
