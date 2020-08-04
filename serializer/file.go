package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

//Receives a proto message and write a bin file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Could not serializer to binary: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("Could not write to file: %w", err)
	}
	return nil
}

//Read bin file of protomessage
func ReadProtobufFromBinaryFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("unable to marshal proto message to binary: %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("enable to write on file %w", err)

	}
	return nil
}
