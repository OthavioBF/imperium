package api

import (
	"net/http"

	"github.com/othavioBF/imperium/internal/jsonutils"
)

func (api *Api) GetproductById(w http.ResponseWriter, r *http.Request) {
	_, problems, err := jsonutils.DecodeValidJson[product.GetProductByIdReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, problems)
		return
	}

	product, err := api.ProductService.HandleGetProductById(r.Context())
	if err != nil {
		api.Logger.Error("product handler", "error")
	}

	jsonutils.EncodeJson(w, r, http.StatusOK, product)
}

func (api *Api) GetAllproductsByUnity(w http.ResponseWriter, r *http.Request) {
	_, problems, err := jsonutils.DecodeValidJson[product.GetAllproductsByUnityReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, problems)
		return
	}

	api.ProductService.GetAllProductsByUnity(r.Context())
}
