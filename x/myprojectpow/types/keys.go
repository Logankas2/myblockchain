package types

const (
	// ModuleName defines the module name
	ModuleName = "myprojectpow"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_myprojectpow"

    
)

var (
	ParamsKey = []byte("p_myprojectpow")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
