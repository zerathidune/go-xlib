package xlib

import "fmt";
import "io";
import "encoding/binary";

type XAuthEntry struct {
    family uint16;
    addr_len uint16;
    address []byte;
    num_len uint16;
    num []byte;
    name_len uint16;
    name []byte;
    data_len uint16;
    data []byte;
}

func (xauth *XAuthEntry) String() string {
    str := fmt.Sprintf(
        "family : %v\n"
        "address (%d) : %s\n"
        "num (%d): %s\n"
        "name (%d): %s\n"
        "data (%d): %v\n",
         xauth.family, xauth.addr_len, xauth.address, xauth.num_len, xauth.num,
        xauth.name_len, xauth.name, xauth.data_len, xauth.data);
    return str;
}

func ReadLengthAndString(input io.Reader, end binary.ByteOrder) (uint16, []byte) {
    lenBuf := make([]byte, 2);
    input.Read(lenBuf);
    length := end.Uint16(lenBuf);
    if length == 0 {
        return 0, nil;
    }
    stringBuf := make([]byte, length);
    n, ok := input.Read(stringBuf);
    if n < int(length) {
        return 0, nil;
        fmt.Println(ok);
    }
    return length, stringBuf;
}

func ReadXAuthEntry (input io.Reader) *XAuthEntry {
    var xauth_entry XAuthEntry;
    end := binary.LittleEndian;
    lenBuf := make([]byte, 2);
    input.Read(lenBuf);
    xauth_entry.family = end.Uint16(lenBuf);
    xauth_entry.addr_len, xauth_entry.address = ReadLengthAndString(input,end);
    xauth_entry.num_len, xauth_entry.num = ReadLengthAndString(input, end);
    xauth_entry.name_len, xauth_entry.name = ReadLengthAndString(input, end);
    xauth_entry.data_len, xauth_entry.data = ReadLengthAndString(input, end);
    return &xauth_entry;
}
