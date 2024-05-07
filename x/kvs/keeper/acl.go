package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"kvs/x/kvs/types"
)

// SetAcl set acl in the store
func (k Keeper) SetAcl(ctx sdk.Context, acl types.Acl) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AclKey))
	b := k.cdc.MustMarshal(&acl)
	store.Set([]byte{0}, b)
}

// GetAcl returns acl
func (k Keeper) GetAcl(ctx sdk.Context) (val types.Acl, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AclKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAcl removes acl from the store
func (k Keeper) RemoveAcl(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AclKey))
	store.Delete([]byte{0})
}


// CheckAddress check existence of address in store
func (k Keeper) CheckAddress(ctx sdk.Context, addr string) bool {
	acl, exist := k.GetAcl(ctx)
	if !exist {
		return false
	}
	isVerify := false 
	for _, aclAddr := range acl.Adresses {
		if aclAddr == addr { 
			isVerify = true 
			break
		}
	}
	return isVerify
}