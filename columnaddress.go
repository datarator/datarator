package main

import (
	"github.com/syscrusher/fake"
)

const (
	columnAddressContinent = "address.continent"
	columnAddressCountry   = "address.country"
	columnAddressCity      = "address.city"
	columnAddressPhone     = "address.phone"
	columnAddressState     = "address.state"
	columnAddressStreet    = "address.street"
	columnAddressZip       = "address.zip"
)

type ColumnAddressContinent struct {
	TypedColumnBase
}

func (column ColumnAddressContinent) Value(chunk *Chunk) (string, error) {
	return fake.Continent(), nil
}

type ColumnAddressCountry struct {
	TypedColumnBase
}

func (column ColumnAddressCountry) Value(chunk *Chunk) (string, error) {
	return fake.Country(), nil
}

type ColumnAddressCity struct {
	TypedColumnBase
}

func (column ColumnAddressCity) Value(chunk *Chunk) (string, error) {
	return fake.City(), nil
}

type ColumnAddressPhone struct {
	TypedColumnBase
}

func (column ColumnAddressPhone) Value(chunk *Chunk) (string, error) {
	return fake.Phone(), nil
}

type ColumnAddressState struct {
	TypedColumnBase
}

func (column ColumnAddressState) Value(chunk *Chunk) (string, error) {
	return fake.State(), nil
}

type ColumnAddressStreet struct {
	TypedColumnBase
}

func (column ColumnAddressStreet) Value(chunk *Chunk) (string, error) {
	return fake.Street(), nil
}

type ColumnAddressZip struct {
	TypedColumnBase
}

func (column ColumnAddressZip) Value(chunk *Chunk) (string, error) {
	return fake.Zip(), nil
}
