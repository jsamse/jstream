package jstream

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Decoder struct {
	r       io.Reader
	r2      io.Reader
	buf     *bytes.Buffer
	b       []byte
	started bool
	err     error
	done    bool
}

func (d *Decoder) Decode(v interface{}) bool {
	if d.done {
		return false
	}
	if !d.started {
		d.started = true
		if _, err := d.r.Read(d.b); err != nil {
			d.err = err
			return false
		}
		if d.b[0] != '[' {
			d.err = fmt.Errorf("Expected first byte to be '[', was '%v'.", d.b[0])
			return false
		}
		d.r2 = d.r
	}
	dec := json.NewDecoder(d.r2)
	if err := dec.Decode(v); err != nil {
		d.err = err
		return false
	}
	if _, err := io.Copy(d.buf, dec.Buffered()); err != nil {
		panic(err)
	}
	d.r2 = io.MultiReader(d.buf, d.r)
	for {
		if _, err := d.r2.Read(d.b); err != nil {
			d.err = err
			return false
		}
		if d.b[0] == ',' {
			break
		}
		if d.b[0] == ']' {
			d.done = true
			break
		}
	}
	return true
}

func (d *Decoder) Err() error {
	return d.err
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:   r,
		buf: bytes.NewBuffer(nil),
		b:   make([]byte, 1),
	}
}
