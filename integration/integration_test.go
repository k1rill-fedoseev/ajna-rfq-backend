package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"ajna-rfq/internal/config"
	"ajna-rfq/internal/contract"
	"ajna-rfq/internal/repo"
	"ajna-rfq/internal/repo/sqlite3"
	"ajna-rfq/internal/server"
	"ajna-rfq/internal/service"
	"ajna-rfq/internal/types"
)

const user1PK = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
const user2PK = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"

func TestIntegration(t *testing.T) {
	cmd := exec.Command("anvil", "--load-state", "./state.json")
	assert.NoError(t, cmd.Start())
	defer cmd.Process.Kill()
	time.Sleep(time.Second)

	rpcURL := "http://localhost:8545"
	rpc, err := ethclient.Dial(rpcURL)
	assert.NoError(t, err, "failed to dial rpc")

	cfg := &config.Config{Chains: []config.ChainConfig{{
		RPC:            rpcURL,
		Factory:        common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"),
		RFQ:            common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		UpdateInterval: time.Second * 2,
	}}}
	repository, err := sqlite3.NewOrdersRepoSQLite3(":memory:")
	if err != nil {
		t.Fatalf("failed to start repo: %v", err)
	}
	svc, err := service.NewService(cfg, repository)
	if err != nil {
		t.Fatalf("failed to init service: %v", err)
	}
	t.Logf("starting server")
	testServer := httptest.NewServer(server.NewRouter(svc))
	defer testServer.Close()

	t.Run("test submit lp order", func(t *testing.T) { InternalTestSubmitLPOrder(t, testServer.URL) })
	time.Sleep(time.Second)
	t.Run("test submit lp order with taker", func(t *testing.T) { InternalTestSubmitLPOrderWithTaker(t, testServer.URL) })
	time.Sleep(time.Second)
	t.Run("test submit quote order", func(t *testing.T) { InternalTestSubmitQuoteOrder(t, testServer.URL) })
	time.Sleep(time.Second)
	t.Run("test submit quote order with taker", func(t *testing.T) { InternalTestSubmitQuoteOrderWithTaker(t, testServer.URL) })
	time.Sleep(time.Second)

	t.Run("test fetch lp orders", func(t *testing.T) {
		res := fetchOrders(t, testServer.URL, "maker=0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
		assert.Nil(t, res.NextPage)
		assert.Len(t, res.Items, 2)
		assert.Equal(t, decimal.RequireFromString("4000000000000000000"), res.Items[0].RemainingMakeAmount)
		assert.Equal(t, decimal.RequireFromString("4000000000000000000"), res.Items[0].ApprovedAmount)
		assert.Equal(t, decimal.RequireFromString("5000000000000000000"), res.Items[1].RemainingMakeAmount)
		assert.Equal(t, decimal.RequireFromString("5000000000000000000"), res.Items[1].ApprovedAmount)
		assert.Equal(t, res, fetchOrders(t, testServer.URL, "maker=0x70997970C51812dc3A010C7d01b50e0d17dc79C8&pool=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"))
		assert.Equal(t, res.Items[1:], fetchOrders(t, testServer.URL, "").Items)
		assert.Equal(t, res, fetchOrders(t, testServer.URL, "taker=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"))
		assert.Equal(t, res, fetchOrders(t, testServer.URL, "taker=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC&pool=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"))
	})

	t.Run("test fetch quote orders", func(t *testing.T) {
		res := fetchOrders(t, testServer.URL, "maker=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC&lpOrder=false")
		assert.Nil(t, res.NextPage)
		assert.Len(t, res.Items, 2)
		assert.Equal(t, decimal.RequireFromString("2000000000000000000"), res.Items[0].RemainingMakeAmount)
		assert.Equal(t, decimal.RequireFromString("2000000000000000000"), res.Items[0].ApprovedAmount)
		assert.Equal(t, decimal.RequireFromString("3000000000000000000"), res.Items[1].RemainingMakeAmount)
		assert.Equal(t, decimal.RequireFromString("3000000000000000000"), res.Items[1].ApprovedAmount)
		assert.Equal(t, res, fetchOrders(t, testServer.URL, "maker=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC&pool=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0&lpOrder=false"))
		assert.Equal(t, res.Items[1:], fetchOrders(t, testServer.URL, "lpOrder=false").Items)
		assert.Equal(t, res, fetchOrders(t, testServer.URL, "taker=0x70997970C51812dc3A010C7d01b50e0d17dc79C8&lpOrder=false"))
		assert.Equal(t, res, fetchOrders(t, testServer.URL, "taker=0x70997970C51812dc3A010C7d01b50e0d17dc79C8&pool=0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0&lpOrder=false"))
	})

	id := new(hexutil.Big)
	assert.NoError(t, rpc.Client().Call(id, "evm_snapshot"))

	t.Run("test fill lp order", func(t *testing.T) { InternalTestFillLPOrder(t, testServer.URL, rpc) })
	time.Sleep(time.Second * 3)
	t.Run("test check filled lp order", func(t *testing.T) {
		res := fetchOrders(t, testServer.URL, "maker=0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
		assert.Len(t, res.Items, 2)
		assert.Equal(t, decimal.RequireFromString("0"), res.Items[0].RemainingMakeAmount)
		assert.Empty(t, fetchOrders(t, testServer.URL, "maker=0x70997970C51812dc3A010C7d01b50e0d17dc79C8&active=true").Items)
	})

	assert.NoError(t, rpc.Client().Call(nil, "evm_revert", id))

	t.Run("test fill quote order", func(t *testing.T) { InternalTestFillQuoteOrder(t, testServer.URL, rpc) })
	time.Sleep(time.Second * 3)
	t.Run("test check filled quote order", func(t *testing.T) {
		res := fetchOrders(t, testServer.URL, "maker=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC&lpOrder=false")
		assert.Len(t, res.Items, 2)
		assert.Equal(t, decimal.RequireFromString("0"), res.Items[0].RemainingMakeAmount)
		assert.Equal(t, res.Items[1:], fetchOrders(t, testServer.URL, "maker=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC&lpOrder=false&active=true").Items)
	})
}

func InternalTestSubmitLPOrder(t *testing.T, url string) {
	order := types.Order{
		LpOrder:       true,
		Maker:         common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
		Taker:         nil,
		Pool:          common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"),
		Index:         1234,
		MakeAmount:    decimal.RequireFromString("5000000000000000000"),
		MinMakeAmount: decimal.RequireFromString("5000000000000000000"),
		Expiry:        decimal.NewFromInt(time.Now().Add(time.Hour).Unix()),
		Price:         decimal.RequireFromString("990000000000000000"),
	}
	resp := signAndSubmit(t, order, user1PK, url)
	assert.Equal(t, http.StatusOK, resp.StatusCode, readBody(resp))
}

func InternalTestSubmitLPOrderWithTaker(t *testing.T, url string) {
	taker := common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
	order := types.Order{
		LpOrder:       true,
		Maker:         common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
		Taker:         &taker,
		Pool:          common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"),
		Index:         1234,
		MakeAmount:    decimal.RequireFromString("4000000000000000000"),
		MinMakeAmount: decimal.RequireFromString("4000000000000000000"),
		Expiry:        decimal.NewFromInt(time.Now().Add(time.Hour).Unix()),
		Price:         decimal.RequireFromString("980000000000000000"),
	}
	resp := signAndSubmit(t, order, user1PK, url)
	assert.Equal(t, http.StatusOK, resp.StatusCode, readBody(resp))
}

func InternalTestSubmitQuoteOrder(t *testing.T, url string) {
	order := types.Order{
		LpOrder:       false,
		Maker:         common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
		Taker:         nil,
		Pool:          common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"),
		Index:         1000,
		MakeAmount:    decimal.RequireFromString("3000000000000000000"),
		MinMakeAmount: decimal.RequireFromString("3000000000000000000"),
		Expiry:        decimal.NewFromInt(time.Now().Add(time.Hour).Unix()),
		Price:         decimal.RequireFromString("970000000000000000"),
	}
	resp := signAndSubmit(t, order, user2PK, url)
	assert.Equal(t, http.StatusOK, resp.StatusCode, readBody(resp))
}

func InternalTestSubmitQuoteOrderWithTaker(t *testing.T, url string) {
	taker := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	order := types.Order{
		LpOrder:       false,
		Maker:         common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
		Taker:         &taker,
		Pool:          common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"),
		Index:         1000,
		MakeAmount:    decimal.RequireFromString("2000000000000000000"),
		MinMakeAmount: decimal.RequireFromString("2000000000000000000"),
		Expiry:        decimal.NewFromInt(time.Now().Add(time.Hour).Unix()),
		Price:         decimal.RequireFromString("960000000000000000"),
	}
	resp := signAndSubmit(t, order, user2PK, url)
	assert.Equal(t, http.StatusOK, resp.StatusCode, readBody(resp))
}

func InternalTestFillLPOrder(t *testing.T, url string, rpc *ethclient.Client) {
	order := fetchOrders(t, url, "maker=0x70997970C51812dc3A010C7d01b50e0d17dc79C8").Items[0]
	fillOrder(t, rpc, order, user2PK)
}

func InternalTestFillQuoteOrder(t *testing.T, url string, rpc *ethclient.Client) {
	order := fetchOrders(t, url, "maker=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC&lpOrder=false").Items[0]
	fillOrder(t, rpc, order, user1PK)
}

func fillOrder(t *testing.T, rpc *ethclient.Client, order *repo.StoredOrder, privateKey string) {
	tx, err := contract.NewRFQTransactor(order.RFQ, rpc)
	assert.NoError(t, err, "failed to create transactor")
	fillAmount, _ := new(big.Int).SetString("10000000000000000000", 10)
	minFillAmount, _ := new(big.Int).SetString("1000000000000000000", 10)
	pk, err := crypto.HexToECDSA(privateKey)
	assert.NoError(t, err, "failed to get private key")
	opts := &bind.TransactOpts{
		From: crypto.PubkeyToAddress(pk.PublicKey),
		Signer: func(_ common.Address, tx *types2.Transaction) (*types2.Transaction, error) {
			return types2.SignTx(tx, types2.NewLondonSigner(big.NewInt(31337)), pk)
		},
	}
	taker := order.TakerOrZero()
	txOrder := contract.IAjnaRFQOrder{
		LpOrder:       order.LpOrder,
		Maker:         order.Maker,
		Taker:         taker,
		Pool:          order.Pool,
		Index:         big.NewInt(order.Index),
		MakeAmount:    order.MakeAmount.BigInt(),
		MinMakeAmount: order.MinMakeAmount.BigInt(),
		Expiry:        order.Expiry.BigInt(),
		Price:         order.Price.BigInt(),
	}
	sentTx, err := tx.FillOrder(opts, txOrder, order.Signature, taker, big.NewInt(1234), fillAmount, minFillAmount, txOrder.Expiry)
	assert.NoError(t, err, "failed to fill order")
	time.Sleep(time.Second)
	receipt, err := rpc.TransactionReceipt(context.Background(), sentTx.Hash())
	assert.NoError(t, err, "failed to fetch tx receipt")
	assert.Equal(t, types2.ReceiptStatusSuccessful, receipt.Status)
}

func fetchOrders(t *testing.T, baseUrl string, params string) *server.OrdersListResponse {
	resp, err := http.Get(baseUrl + "/api/v1/31337/orders?" + params)
	assert.NoError(t, err, "failed to get")
	defer resp.Body.Close()
	res := new(server.OrdersListResponse)
	err = json.NewDecoder(resp.Body).Decode(res)
	assert.NoError(t, err, "failed to decode")
	return res
}

func signAndSubmit(t *testing.T, order types.Order, privateKey, baseUrl string) *http.Response {
	hash, _, err := order.Hash(big.NewInt(31337), common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"))
	assert.NoError(t, err, "failed to hash order")
	pk, err := crypto.HexToECDSA(privateKey)
	assert.NoError(t, err, "failed to get private key")
	sig, err := crypto.Sign(hash.Bytes(), pk)
	sig[64] += 27
	assert.NoError(t, err, "failed to sign order")
	body, err := json.Marshal(types.SignedOrder{
		Signature: sig,
		Order:     order,
	})
	assert.NoError(t, err, "failed to marshal")
	resp, err := http.Post(baseUrl+"/api/v1/31337/orders", "application/json", bytes.NewBuffer(body))
	assert.NoError(t, err, "failed to post")
	return resp
}

func readBody(res *http.Response) string {
	defer res.Body.Close()
	msg, _ := io.ReadAll(res.Body)
	return string(msg)
}
