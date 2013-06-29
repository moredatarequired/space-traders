package lib

import (
	"encoding/csv"
	"os"
	"strconv"
	"io"
)

type Star struct {
	Id int
	Name, Class string
	X, Y, Z, Magnitude float64
}

func (s *Star) Position() (x, y, z float64) { return s.X, s.Y, s.Z }

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

func readFromCSV(reader io.Reader) ([]*Star, error) {
	r := csv.NewReader(reader)
	var stars []*Star
	var err error
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		if record[2] != "1" { continue }
		if star, err := newStar(record); err == nil {
			stars = append(stars, star)
		} else { return nil, err }
	}
	if err == nil { return stars, nil }
	return nil, err
}

func readFromFile(f string) ([]*Star, error) {
	reader, err := os.Open(f)
	if err != nil { return nil, err }
	defer reader.Close()
	return readFromCSV(reader)
}
