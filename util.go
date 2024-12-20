package dynamodb

import (
	"strconv"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
)

// toProto converts a byte array given its manifest into a valid proto message
func toProto(manifest string, bytea []byte) (*anypb.Any, error) {
	mt, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(manifest))
	if err != nil {
		return nil, err
	}

	pm := mt.New().Interface()
	err = proto.Unmarshal(bytea, pm)
	if err != nil {
		return nil, err
	}

	if cast, ok := pm.(*anypb.Any); ok {
		return cast, nil
	}
	return nil, fmt.Errorf("failed to unpack message=%s", manifest)
}

func parseDynamoUint64(element types.AttributeValue) uint64 {
	n, _ := strconv.ParseUint(element.(*types.AttributeValueMemberN).Value, 10, 64)
	return n
}

func parseDynamoInt64(element types.AttributeValue) int64 {
	n, _ := strconv.ParseInt(element.(*types.AttributeValueMemberN).Value, 10, 64)
	return n
}
