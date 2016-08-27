package main

import (
	"github.com/syscrusher/fake"
)

const (
	COLUMN_ADDRESS_CONTINENT = "address.continent"
	COLUMN_ADDRESS_COUNTRY   = "address.country"
	COLUMN_ADDRESS_CITY      = "address.city"
	COLUMN_ADDRESS_PHONE     = "address.phone"
	COLUMN_ADDRESS_STATE     = "address.state"
	COLUMN_ADDRESS_STREET    = "address.street"
	COLUMN_ADDRESS_ZIP       = "address.zip"
)

type ColumnAddressContinentPayload struct {
	XmlType string `json:"xml"`
}

type ColumnAddressContinent struct {
	column  Column
	Payload ColumnAddressContinentPayload `json:"payload"`
}

func (column ColumnAddressContinent) Value(context *Context) (string, error) {
	return fake.Continent(), nil
}

func (column ColumnAddressContinent) Column() Column {
	return column.column
}

func (column ColumnAddressContinent) XmlType() string {
	return column.Payload.XmlType
}

type ColumnAddressCountryPayload struct {
	XmlType string `json:"xml"`
}

type ColumnAddressCountry struct {
	column  Column
	Payload ColumnAddressCountryPayload `json:"payload"`
}

func (column ColumnAddressCountry) Value(context *Context) (string, error) {
	return fake.Country(), nil
}

func (column ColumnAddressCountry) Column() Column {
	return column.column
}

func (column ColumnAddressCountry) XmlType() string {
	return column.Payload.XmlType
}

type ColumnAddressCityPayload struct {
	XmlType string `json:"xml"`
}

type ColumnAddressCity struct {
	column  Column
	Payload ColumnAddressCityPayload `json:"payload"`
}

func (column ColumnAddressCity) Value(context *Context) (string, error) {
	return fake.City(), nil
}

func (column ColumnAddressCity) Column() Column {
	return column.column
}

func (column ColumnAddressCity) XmlType() string {
	return column.Payload.XmlType
}

type ColumnAddressPhonePayload struct {
	XmlType string `json:"xml"`
}

type ColumnAddressPhone struct {
	column  Column
	Payload ColumnAddressPhonePayload `json:"payload"`
}

func (column ColumnAddressPhone) Value(context *Context) (string, error) {
	return fake.Phone(), nil
}

func (column ColumnAddressPhone) Column() Column {
	return column.column
}

func (column ColumnAddressPhone) XmlType() string {
	return column.Payload.XmlType
}

type ColumnAddressStatePayload struct {
	XmlType string `json:"xml"`
}
type ColumnAddressState struct {
	column  Column
	Payload ColumnAddressStatePayload `json:"payload"`
}

func (column ColumnAddressState) Value(context *Context) (string, error) {
	return fake.State(), nil
}

func (column ColumnAddressState) Column() Column {
	return column.column
}

func (column ColumnAddressState) XmlType() string {
	return column.Payload.XmlType
}

type ColumnAddressStreetPayload struct {
	XmlType string `json:"xml"`
}

type ColumnAddressStreet struct {
	column  Column
	Payload ColumnAddressStreetPayload `json:"payload"`
}

func (column ColumnAddressStreet) Value(context *Context) (string, error) {
	return fake.Street(), nil
}

func (column ColumnAddressStreet) Column() Column {
	return column.column
}

func (column ColumnAddressStreet) XmlType() string {
	return column.Payload.XmlType
}

type ColumnAddressZipPayload struct {
	XmlType string `json:"xml"`
}

type ColumnAddressZip struct {
	column  Column
	Payload ColumnAddressZipPayload `json:"payload"`
}

func (column ColumnAddressZip) Value(context *Context) (string, error) {
	return fake.Zip(), nil
}

func (column ColumnAddressZip) Column() Column {
	return column.column
}

func (column ColumnAddressZip) XmlType() string {
	return column.Payload.XmlType
}
