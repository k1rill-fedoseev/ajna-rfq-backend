package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	"ajna-rfq/internal/repo"
	"ajna-rfq/internal/service"
	"ajna-rfq/internal/types"
)

type OrdersListResponse struct {
	Items    []*repo.StoredOrder         `json:"items"`
	NextPage *OrdersListNextPageResponse `json:"nextPage,omitempty"`
}

type OrdersListNextPageResponse struct {
	CreatedBefore int64 `json:"createdBefore"`
}

func NewRouter(svc *service.Service) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/{chainId}/orders", func(w http.ResponseWriter, r *http.Request) {
		cs, err := svc.ChainService(r.PathValue("chainId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		q := r.URL.Query()
		filter := repo.Filter{}
		filter.LpOrder = q.Get("lpOrder") != "false"
		filter.Active = q.Get("active") != "false"
		if q.Has("maker") {
			addr := common.HexToAddress(q.Get("maker"))
			filter.Maker = &addr
		}
		if q.Has("taker") {
			addr := common.HexToAddress(q.Get("taker"))
			filter.Taker = &addr
		}
		for _, pool := range q["pool"] {
			if addr := common.HexToAddress(pool); addr != (common.Address{}) {
				filter.Pools = append(filter.Pools, addr)
			}
		}
		if num, err := strconv.ParseInt(q.Get("createdBefore"), 10, 64); err == nil {
			filter.CreatedBefore = &num
		}

		orders, nextCreatedBefore, err := cs.ListOrders(filter, 10)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := OrdersListResponse{Items: orders}
		if nextCreatedBefore != nil {
			res.NextPage = &OrdersListNextPageResponse{CreatedBefore: *nextCreatedBefore}
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("POST /api/v1/{chainId}/orders", func(w http.ResponseWriter, r *http.Request) {
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
