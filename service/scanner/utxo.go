package scanner

import (
	ourchainrpc "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
)

// ListUnspent list unspent utxo
func ListUnspent(chain *ourchainrpc.Bitcoind, db *sqlite.Client, addressList []string, confirm int) (result *[]ourchainrpc.Unspent, err error) {
	curLocalChain := newChain(clientWrapper{
		ChainType: LOCAL,
		DB:        db,
	})
	curRemoteChain := newChain(clientWrapper{
		ChainType: REMOTE,
		RPC:       chain,
	})
	err = curLocalChain.InitChainStep()
	if err != nil {
		return
	}
	err = curRemoteChain.InitChainStep()
	if err != nil {
		return
	}
	err = curLocalChain.SyncLength(curRemoteChain)
	if err != nil {
		return
	}
	result, err = curLocalChain.GetUnspent(addressList, confirm)
	return
}
