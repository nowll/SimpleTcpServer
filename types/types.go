package types

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"io"
)

const BinaryTypes uint8 = iota + 1

// kontrak interface
type Payload interface {
	io.WriterTo
	io.ReaderFrom
	Bytes() Binary
}

type Binary []byte

func (b Binary) Bytes() Binary {
	return b
}

func (b Binary) WriteTo(r io.Writer) (int64, error) {

	err := binary.Write(r, binary.BigEndian, BinaryTypes)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = binary.Write(r, binary.BigEndian, int32(len(b)))

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	n, err := r.Write(b)

	return int64(n + 5), err

}

func (b *Binary) ReadFrom(w io.Reader) (int64, error) {
	var typ uint8

	err := binary.Read(w, binary.BigEndian, &typ)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var size int32

	err = binary.Read(w, binary.BigEndian, &size)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	*b = make(Binary, size)
	n, errRead := w.Read(*b)

	return int64(n + 5), errRead
}

func Decode(r io.Reader) (Payload, error) {
	var typ uint8

	err := binary.Read(r, binary.BigEndian, &typ)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	payload := new(Binary)

	_, err = payload.ReadFrom(io.MultiReader(bytes.NewReader(Binary{typ}), r))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return payload, err

}
