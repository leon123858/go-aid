package scanner

import (
	our_chain_rpc "github.com/leon123858/go-aid/service/rpc"
)

// ListUnspent list unspent utxo
func ListUnspent(chain *our_chain_rpc.Bitcoind, addressList []string, confirm int) (result []our_chain_rpc.Unspent, err error) {
	chainInfo, err := chain.GetBlockChainInfo()
	if err != nil {
		return
	}
	// can not index block 0, so start from 1
	utxoMap := make(map[string]our_chain_rpc.Unspent)
	for i := 1; i < chainInfo.Blocks-confirm; i++ {
		var blockHash string
		blockHash, err = chain.GetBlockHash(uint64(i))
		if err != nil {
			return
		}
		var blockInfo our_chain_rpc.BlockInfo
		blockInfo, err = chain.GetBlock(blockHash)
		if err != nil {
			return
		}
		for _, tx := range blockInfo.Tx {
			var txInfo our_chain_rpc.Transaction
			txInfo, err = chain.GetRawTransaction(tx)
			if err != nil {
				return
			}
			for index, vout := range txInfo.Vout {
				if vout.ScriptPubKey.Type == "pubkey" && vout.ScriptPubKey.Addresses != nil {
					utxoMap[tx] = our_chain_rpc.Unspent{
						Txid:          tx,
						Vout:          index,
						Address:       vout.ScriptPubKey.Addresses[0],
						Amount:        vout.Value,
						Confirmations: txInfo.Confirmations,
					}
				}
			}
			for _, vin := range txInfo.Vin {
				if vin.Coinbase != "" {
					continue
				}
				if _, ok := utxoMap[vin.Txid]; ok {
					delete(utxoMap, vin.Txid)
				}
			}
		}
	}
	if len(addressList) == 0 {
		for _, v := range utxoMap {
			result = append(result, v)
		}
		return
	}
	for _, v := range utxoMap {
		// if v.Address in addressList
		for _, address := range addressList {
			if v.Address == address {
				result = append(result, v)
				break
			}
		}
	}
	return
}
