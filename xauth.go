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
    str := fmt.Sprintf("family : %v\naddress : %s\nnum : %s"
        "\nname : %s\ndata %v\n", xauth.family, xauth.address, xauth.num,
        xauth.name, xauth.data);
    return str;
}

func ReadXAuthEntry (input io.Reader) *XAuthEntry {
    var xauth_entry XAuthEntry;
    end := binary.LittleEndian;
    lenBuf := make([]byte, 2);
    input.Read(lenBuf);
    xauth_entry.family = end.Uint16(lenBuf);
    input.Read(lenBuf);
    xauth_entry.addr_len = end.Uint16(lenBuf);
    xauth_entry.address = make([]byte, xauth_entry.addr_len);
    input.Read(xauth_entry.address);
    input.Read(lenBuf);
    xauth_entry.num_len = end.Uint16(lenBuf);
    xauth_entry.num = make([]byte, xauth_entry.num_len);
    input.Read(xauth_entry.num);
    input.Read(lenBuf);
    xauth_entry.name_len = end.Uint16(lenBuf);
    xauth_entry.name = make([]byte, xauth_entry.name_len);
    input.Read(xauth_entry.name);
    input.Read(lenBuf);
    xauth_entry.data_len = end.Uint16(lenBuf);
    xauth_entry.data = make([]byte, xauth_entry.data_len);
    input.Read(xauth_entry.data);
    return &xauth_entry;
}
