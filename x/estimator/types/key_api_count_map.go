package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ApiCountMapKeyPrefix is the prefix to retrieve all ApiCountMap
	ApiCountMapKeyPrefix = "ApiCountMap/value/"
)

// ApiCountMapKey returns the store key to retrieve a ApiCountMap from the index fields
func ApiCountMapKey(
	creator string,
) []byte {
	var key []byte

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
