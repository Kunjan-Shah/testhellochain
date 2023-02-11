package types

const (
	// ModuleName defines the module name
	ModuleName = "estimator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_estimator"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ApiHitsKey      = "ApiHits/value/"
	ApiHitsCountKey = "ApiHits/count/"
)

const (
	ApiDataKey      = "ApiData/value/"
	ApiDataCountKey = "ApiData/count/"
)
