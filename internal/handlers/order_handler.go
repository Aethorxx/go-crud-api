package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-crud-api/internal/models"
	"go-crud-api/internal/services"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// Create создает новый заказ
func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "неверный формат данных", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(r.Context(), &order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetByID получает заказ по ID
func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if order == nil {
		http.Error(w, "заказ не найден", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// Update обновляет данные заказа
func (h *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "неверный формат данных", http.StatusBadRequest)
		return
	}
	order.ID = uint(id)

	if err := h.service.Update(r.Context(), &order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// Delete удаляет заказ
func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "неверный ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// List получает список всех заказов
func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

// GetByUserID получает все заказы пользователя
func (h *OrderHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		http.Error(w, "неверный ID пользователя", http.StatusBadRequest)
		return
	}

	orders, err := h.service.GetByUserID(r.Context(), uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}
