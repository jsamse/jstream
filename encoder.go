package jstream

import (
	"encoding/json"
	"errors"
	"io"
)

type Encoder struct {
	w       io.Writer
	started bool
	closed  bool
}

func (e *Encoder) Encode(v interface{}) error {
	if e.started {
		if _, err := e.w.Write([]byte{','}); err != nil {
			return err
		}
	} else {
		_, err := e.w.Write([]byte{'['})
		if err != nil {
			return err
		}
		e.started = true
	}
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = e.w.Write(b)
	return err
}

func (e *Encoder) Close() error {
	if e.closed {
		return errors.New("Already closed")
	}
	e.closed = true
	if !e.started {
		_, err := e.w.Write([]byte{'['})
		if err != nil {
			return err
		}
	}
	_, err := e.w.Write([]byte{']'})
	return err
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}
