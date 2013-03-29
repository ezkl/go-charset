package charset

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
)

func init() {
	registerClass("cp949", fromCp949, toCp949)
}

var (
	ErrorCp949DatInvalid        = errors.New("broken cp949.dat")
	ErrorInvalidCp949Code       = errors.New("invalid byte seq. for cp949")
	ErrorCantRepresentWithCp949 = errors.New("cant represent with cp949")
)

type cp949Code struct {
	cp949 [2]byte
	utf8  [3]byte
}

type translateCp949 struct {
	table   []*cp949Code
	scratch []byte
}

type translateFromCp949 struct {
	translateCp949
}

// XXX: what meanig of eof
func (p *translateFromCp949) Translate(data []byte, eof bool) (int, []byte, error) {
	if p.table == nil {
		panic("table to translate from cp949 not exists")
	}

	p.scratch = p.scratch[:0]
	n := 0
	for len(data) > 0 {
		if data[0]&0x80 == 0 {
			p.scratch = append(p.scratch, data[0])
			data = data[1:]
			n += 1
			continue
		}

		fi := sort.Search(len(p.table), func(i int) bool {
			c := p.table[i]
			if c.cp949[0] < data[0] {
				return false
			}
			if c.cp949[0] == data[0] && c.cp949[1] >= data[1] {
				return true
			}
			return false
		})

		f := p.table[fi]
		if f.cp949[0] != data[0] || f.cp949[1] != data[1] {
			return n, p.scratch, ErrorInvalidCp949Code
		}
		data = data[2:]
		n += 2

		p.scratch = append(p.scratch, f.utf8[0], f.utf8[1], f.utf8[2])
	}
	return n, p.scratch, nil
}

type translateToCp949 struct {
	translateCp949
}

func (p *translateToCp949) Translate(data []byte, eof bool) (int, []byte, error) {
	p.scratch = p.scratch[:0]
	n := 0
	for len(data) > 0 {
		if data[0]&0x80 == 0 {
			p.scratch = append(p.scratch, data[0])
			data = data[1:]
			n += 1
			continue
		}

		fi := sort.Search(len(p.table), func(i int) bool {
			c := p.table[i]
			if c.utf8[0] < data[0] {
				return false
			}
			if c.utf8[1] < data[1] {
				return false
			}
			if c.utf8[0] == data[0] && c.utf8[1] == data[1] &&
				c.utf8[2] >= data[2] {
				return true
			}
			return false
		})

		f := p.table[fi]
		if f.utf8[0] != data[0] || f.utf8[1] != data[1] || f.utf8[2] != data[2] {
			return n, p.scratch, ErrorCantRepresentWithCp949
		}
		data = data[3:]
		n += 3

		p.scratch = append(p.scratch, f.cp949[0], f.cp949[1])
	}
	return n, p.scratch, nil
}

func fromCp949(arg string) (Translator, error) {
	// XXX: need cache? what's purpose of key of the cache?

	// load dat file and create translate table
	// TODO:
	dat, err := readFile("cp949.dat")
	buf := bytes.NewBuffer(dat)
	var start, count int

	var tbl []*cp949Code

	copyTo3BytesArray := func(dst *[3]byte, src []byte) {
		dst[0], dst[1], dst[2] = src[0], src[1], src[2]
	}

	_, err = fmt.Fscanf(buf, "0x%04x %d", &start, &count)
	for err == nil {
		var line []byte
		fmt.Fscanln(buf, &line)
		if len(line) != 3*count {
			return nil, ErrorCp949DatInvalid
		}
		for i := 0; i < count; i++ {
			var c cp949Code
			c.cp949[0] = byte((start + i) >> 8)
			c.cp949[1] = byte((start + i) & 0xff)
			copyTo3BytesArray(&c.utf8, line[i*3:i*3+3])
			tbl = append(tbl, &c)
		}
		_, err = fmt.Fscanf(buf, "0x%04x %d", &start, &count)
	}

	if err != nil {
		return nil, err
	}
	// TODO: do I need check it sorted?

	return &translateFromCp949{translateCp949{table: tbl}}, nil
}

func toCp949(arg string) (Translator, error) {
	// XXX: need cache? what's purpose of key of the cache?

	// TODO: sort table from translateFromCp949
	return &translateToCp949{translateCp949{table: nil}}, nil
}
