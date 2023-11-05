package controller

import (
	"github.com/labstack/echo"
	model "github.com/leon123858/go-aid/utils/modal"
	our_chain_rpc "github.com/leon123858/go-aid/utils/rpc"
	"net/http"
)

func GenerateChainGetController(chain *our_chain_rpc.Bitcoind, which string) echo.HandlerFunc {
	switch which {
	case "getUnspent":
		return func(ctx echo.Context) error {
			list, err := chain.ListUnspent()
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   list,
			})
		}
	case "getBalance":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			balance, err := chain.GetBalance(address, 1)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   balance,
			})
		}
	case "getPrivateKey":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			privateKey, err := chain.DumpPrivKey(address)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   privateKey,
			})
		}
	case "getTransaction":
		return func(ctx echo.Context) error {
			txid := ctx.QueryParam("txid")
			tx, err := chain.GetTransaction(txid)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   tx,
			})
		}
	default:
		return nil
	}
}

func GenerateChainPostController(chain *our_chain_rpc.Bitcoind, which string) echo.HandlerFunc {
	switch which {
	case "generateBlock":
		return func(ctx echo.Context) error {
			blockIds, err := chain.GenerateBlock(1)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   blockIds,
			})
		}
	case "DumpContractMessage":
		return func(ctx echo.Context) error {
			req := new(model.ContractRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			message, err := chain.DumpContractMessage(req.Address, req.Arguments)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   message,
			})
		}
	default:
		return nil
	}
}