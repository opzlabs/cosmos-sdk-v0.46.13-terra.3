package store

import (
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/store/cache"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/store/rootmulti"
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db, log.NewNopLogger())
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
