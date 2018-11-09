package main

import "bytes"
import "errors"
import "fmt"
import "encoding/binary"
import "unsafe"

func Serialize(i interface{}) {
	buf := new(bytes.Buffer)
	var err error
	switch i.(type) {
	case uint8, uint16, uint32, uint64, uintptr:
		size := unsafe.Sizeof(i)
		if size%8 == 0 && size <= 256 {
			err = binary.Write(buf, binary.BigEndian, i)
		} else {
			err = errors.New("Size was not a multiple of 8/ less than 256", size)
		}
	case string:
		for _, v := range i.(string) {
			err = binary.Write(buf, binary.BigEndian, v)
			if err != nil {
				break
			}
		}
	case bool:
		err = binary.Write(buf, binary.BigEndian, i)
	default:
		fmt.Println("???")
	}
	if err != nil {
		fmt.Println("binary.Write failed", err)
	} else {
		fmt.Printf("%x\n", buf.Bytes())
	}
}

func main() {
	Serialize(uint64(800))

	Serialize("A string called stupidity")

	Serialize(false)
	Serialize(true)

}
