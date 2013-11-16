package model

import (
	"github.com/lib/pq"

	"bytes"
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

type NullTime struct {
	pq.NullTime
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		ns.String = ""
		ns.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &ns.String)
	if err != nil {
		return err
	}
	ns.Valid = true
	return nil
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.String == "" && !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (nt *NullTime) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		nt.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &nt.Time)
	if err != nil {
		return err
	}
	nt.Valid = true
	return nil
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}
