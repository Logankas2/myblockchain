package types

const (
	// ModuleName defines the module name
	ModuleName = "pow"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pow"

    
)

var (
	ParamsKey = []byte("p_pow")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
