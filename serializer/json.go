package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

//converts protomessage to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		//if true change ieps to int for enums
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       " ",
		//if false swith to camel case
		OrigName: true,
	}

	return marshaler.MarshalToString(message)
}
