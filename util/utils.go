package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// LBtU8 converts 1 byte to uint8. Errors out if byte slice length is greater
// than 1. Uses little endian
// LBtU8은 길이 1의 바이트 슬라이스를 uint8으로 변환합니다. 길이가 1 보다 긴
// 바이트 슬라이스를 받으면 에러가 납니다. Little Endian을 사용합니다.
func LBtU8(b []byte) (uint8, error) {
	if len(b) != 1 {
		return 0, fmt.Errorf("BtU8 byte slice greater than 1" +
			"BtU8에 준 byte slice가 1바이트가 아닙니다.")
	}
	var i uint8
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &i)
	return i, nil
}

// LBtU16 converts 2 bytes to uint16. Errors out if byte slice length is greater
// than 2. Uses little endian.
// LBtU16은 길이 2의 바이트 슬라이스를 uint16으로 변환합니다. 길이가 2 보다 긴
// 바이트 슬라이스를 받으면 에러가 납니다. Little Endian을 사용합니다.
func LBtU16(b []byte) (uint16, error) {
	if len(b) != 2 {
		return 0, fmt.Errorf("BtU16 byte slice greater than 2" +
			"BtU16에 준 byte slice가 2바이트가 아닙니다.")
	}
	var i uint16
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &i)
	return i, nil
}

// LBtU32 converts 4 bytes to uint32. Errors out if byte slice length is greater
// than 4. Uses little endian.
// LBtU32은 길이 4의 바이트 슬라이스를 uint32으로 변환합니다. 길이가 4 보다 긴
// 바이트 슬라이스를 받으면 에러가 납니다. Little endian을 사용합니다.
func LBtU32(b []byte) (uint32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("BtU32 byte slice greater than 4" +
			"BtU32에 준 byte slice가 2바이트가 아닙니다.")
	}
	var i uint32
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &i)
	return i, nil
}

// BtU32 converts 4 bytes to uint32. Errors out if byte slice length is greater
// than 4. Uses little endian.
// BtU32은 길이 4의 바이트 슬라이스를 uint32으로 변환합니다. 길이가 4 보다 긴
// 바이트 슬라이스를 받으면 에러가 납니다. Little endian을 사용합니다.
func BtU32(b []byte) (uint32, error) {
	if len(b) != 4 {
		return 0, fmt.Errorf("BtU32 byte slice greater than 4" +
			"BtU32에 준 byte slice가 2바이트가 아닙니다.")
	}
	var i uint32
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &i)
	return i, nil
}
