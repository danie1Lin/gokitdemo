package prototool

import (
	types "github.com/gogo/protobuf/types"
)

// Maybe generic or gomarco can save us time
func PInt32(v *types.Int32Value) *int32 {
	if v == nil {
		return nil
	}
	return &v.Value
}

func OptionInt32(v *int32) *types.Int32Value {
	if v == nil {
		return nil
	}
	return &types.Int32Value{Value: *v}
}
