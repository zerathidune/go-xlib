package xlib

import "fmt";
import "encoding/binary";
import "net";
import "io";
import "strings";

const (
    xBigEndian = 0x42;
    xLittleEndian = 0x6C;
	failed = 0;
	authenticate = 2;
	success = 1;
)

type initConnT struct {
    endianness byte;
    protocolMajorVersion uint16;
    protocolMinorVersion uint16;
	authProtoNameSize uint16;
	authProtoDataSize uint16;
    authProtoName string;
    authProtoData []byte;
}


func newInitConnT(end byte, major uint16, minor uint16) *initConnT {
	var init initConnT;
	init.endianness = end;
	init.protocolMajorVersion = major;
	init.protocolMinorVersion = minor;
    init.authProtoName = "MIT-MAGIC-COOKIE-1";
	init.authProtoNameSize = uint16(len(init.authProtoName));
	init.authProtoDataSize = 0;
	return &init;
}

func (init *initConnT) send (out io.Writer) {
	var data [12]byte;
	data[0] = init.endianness;
	var end binary.ByteOrder;
	if data[0] == xLittleEndian {
		end = binary.LittleEndian;
	}
	else {
		end = binary.BigEndian;
	}
	end.PutUint16(data[2:4], init.protocolMajorVersion);
	end.PutUint16(data[4:6], init.protocolMinorVersion);
	end.PutUint16(data[6:8], init.authProtoNameSize);
	end.PutUint16(data[8:10], init.authProtoDataSize);
	out.Write(&data);
    out.Write(strings.Bytes(init.authProtoName));
    padding := 4 - (init.authProtoNameSize % 4);
    namePadding := make([]byte, padding);
    out.Write(namePadding);
    padding = 4 - (init.authProtoDataSize % 4);
    dataPadding := make([]byte, padding);
    out.Write(dataPadding);
}


type initConnResponseT struct {
	status uint8;
	protocolMajorVersion uint16;
	protocolMinorVersion uint16;
	reason []byte;
	reasonLength uint8;
}

func initConnResponseTRead(in io.Reader, end binary.ByteOrder) *initConnResponseT {
	var buf [256]byte;
	var response initConnResponseT;
	in.Read(&buf);
	if buf[0] == failed {
		response.status = failed;
		response.reasonLength = buf[1];
		response.reason = make([]byte,  response.reasonLength);
		var i uint8;
		for i = 0; i < response.reasonLength; i++ {
			response.reason[i] = buf[i+2];
		}
		end.Uint16(buf[i+1:i+3]);
		end.Uint16(buf[i+3:i+5]);
	}
	return &response;
}

func (response *initConnResponseT) String() string {
	if response.status == failed {
		return fmt.Sprintf("Failed!\nReason : %s\n", response.reason);
	}
	return "not a failure!";
}

type Display struct {

}

func XDisplayOpen(s string)  {
	init := newInitConnT(xLittleEndian, 11, 0);
	socket, addr_err := net.ResolveUnixAddr("unix", "/tmp/.X11-unix/X0");
	fmt.Println("ResolveUnixAddr err :", addr_err);
	conn, dial_err := net.DialUnix("unix", nil, socket);
	fmt.Println("DialUnix err :", dial_err);
	init.send(conn);
	response := initConnResponseTRead(conn, binary.LittleEndian);
	fmt.Println(response);
	conn.Close();
}
