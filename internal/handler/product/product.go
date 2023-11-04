package product

import (
	"net/http"
	validate "github.com/Marif226/melon/internal/lib/validator/product"
	"github.com/Marif226/melon/internal/model"
	"github.com/Marif226/melon/internal/service"
	"github.com/Marif226/melon/pkg/helper"
	"github.com/google/jsonapi"
)

type productHandler struct {
	service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{
		ProductService: productService,
	}
}

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request model.Product

	err := jsonapi.UnmarshalPayload(r.Body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Create(&request)
	if err != nil {
		helper.JsonapiError(w, []*jsonapi.ErrorObject{{
			Title: "Validation Error",
			Detail: err.Error(),
			Status: "400",
		}})
		return
	}

	response, err := h.ProductService.Create(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	err = jsonapi.MarshalPayload(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *productHandler) List(w http.ResponseWriter, r *http.Request) {
	var request model.ProductListRequest

	// err := jsonapi.UnmarshalPayload(r.Body, &request)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	response, err := h.ProductService.List(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	err = jsonapi.MarshalPayload(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}