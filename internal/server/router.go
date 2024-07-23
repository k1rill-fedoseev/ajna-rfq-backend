package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"ajna-rfq/internal/service"
	"ajna-rfq/internal/types"
)

func NewRouter(svc *service.Service) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/{chainId}/make-orders", func(w http.ResponseWriter, r *http.Request) {
		cs, err := svc.ChainService(r.PathValue("chainId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		q := r.URL.Query()
		maker := common.HexToAddress(q.Get("maker"))
		pools := make([]common.Address, len(q["pool"]))
		for _, pool := range q["pool"] {
			if addr := common.HexToAddress(pool); addr != (common.Address{}) {
				pools = append(pools, addr)
			}
		}
		var createdBefore *uint64
		if num, err := strconv.ParseUint(q.Get("createdBefore"), 10, 64); err == nil {
			createdBefore = &num
		}

		orders, err := cs.ListMakeOrders(maker, pools, createdBefore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(orders)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	mux.HandleFunc("GET /api/v1/{chainId}/take-orders", func(w http.ResponseWriter, r *http.Request) {
		cs, err := svc.ChainService(r.PathValue("chainId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		q := r.URL.Query()
		var taker *common.Address
		if addr := common.HexToAddress(q.Get("taker")); addr != (common.Address{}) {
			taker = &addr
		}
		pools := make([]common.Address, len(q["pool"]))
		for _, pool := range q["pool"] {
			if addr := common.HexToAddress(pool); addr != (common.Address{}) {
				pools = append(pools, addr)
			}
		}
		var createdBefore *uint64
		if num, err := strconv.ParseUint(q.Get("createdBefore"), 10, 64); err == nil {
			createdBefore = &num
		}

		orders, err := cs.ListTakeOrders(taker, pools, createdBefore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(orders)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	mux.HandleFunc("PUT /api/v1/{chainId}/orders", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		cs, err := svc.ChainService(r.PathValue("chainId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		order := types.SignedOrder{}
		err = json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = cs.ValidateOrder(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = cs.UpdateOrder(order, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return mux
}
