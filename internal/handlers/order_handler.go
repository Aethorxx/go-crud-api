package handlers

import (
	"net/http"
	"strconv"

	"go-crud-api/internal/models"
	"go-crud-api/internal/services"

	"github.com/gin-gonic/gin"
)

// OrderHandler обрабатывает все запросы, связанные с заказами
// Использует middleware для проверки аутентификации
type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// CreateOrder создает новый заказ для текущего пользователя
// ID пользователя берется из JWT токена через middleware
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	// Получаем ID пользователя из URL
	userIDFromURL, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Получаем ID пользователя из JWT токена
	userIDFromToken := c.GetUint("user_id")

	// Проверяем, что ID из URL совпадает с ID из токена
	if uint(userIDFromURL) != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	var req models.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := &models.Order{
		UserID:   userIDFromToken,
		Product:  req.Product,
		Quantity: req.Quantity,
		Price:    req.Price,
	}

	if err := h.orderService.Create(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrder получает информацию о конкретном заказе
// Проверяет существование заказа и права доступа
func (h *OrderHandler) GetOrder(c *gin.Context) {
	// Получаем ID пользователя из URL
	userIDFromURL, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID пользователя"})
		return
	}

	// Получаем ID пользователя из JWT токена
	userIDFromToken := c.GetUint("user_id")

	// Проверяем, что ID из URL совпадает с ID из токена
	if uint(userIDFromURL) != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "нет доступа к заказам другого пользователя"})
		return
	}

	// Получаем ID заказа
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID заказа"})
		return
	}

	order, err := h.orderService.GetByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "заказ не найден"})
		return
	}

	// Проверяем принадлежность заказа пользователю
	if order.UserID != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "заказ принадлежит другому пользователю"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder обновляет существующий заказ
// Проверяет принадлежность заказа текущему пользователю
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	// Получаем ID пользователя из URL
	userIDFromURL, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID пользователя"})
		return
	}

	// Получаем ID пользователя из JWT токена
	userIDFromToken := c.GetUint("user_id")

	// Проверяем, что ID из URL совпадает с ID из токена
	if uint(userIDFromURL) != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "нет доступа к заказам другого пользователя"})
		return
	}

	// Получаем ID заказа
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID заказа"})
		return
	}

	// Проверяем существование заказа и принадлежность пользователю
	existingOrder, err := h.orderService.GetByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "заказ не найден"})
		return
	}

	// Проверяем принадлежность заказа пользователю
	if existingOrder.UserID != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "заказ принадлежит другому пользователю"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.ID = uint(orderID)
	order.UserID = userIDFromToken

	if err := h.orderService.Update(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteOrder удаляет заказ
// Возвращает 204 No Content при успешном удалении
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	// Получаем ID пользователя из URL
	userIDFromURL, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID пользователя"})
		return
	}

	// Получаем ID пользователя из JWT токена
	userIDFromToken := c.GetUint("user_id")

	// Проверяем, что ID из URL совпадает с ID из токена
	if uint(userIDFromURL) != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "нет доступа к заказам другого пользователя"})
		return
	}

	// Получаем ID заказа
	orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID заказа"})
		return
	}

	// Проверяем существование заказа и принадлежность пользователю
	existingOrder, err := h.orderService.GetByID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "заказ не найден"})
		return
	}

	// Проверяем принадлежность заказа пользователю
	if existingOrder.UserID != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "заказ принадлежит другому пользователю"})
		return
	}

	if err := h.orderService.Delete(uint(orderID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetUserOrders получает все заказы текущего пользователя
// Использует ID пользователя из JWT токена
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	// Получаем ID пользователя из URL
	userIDFromURL, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID пользователя"})
		return
	}

	// Получаем ID пользователя из JWT токена
	userIDFromToken := c.GetUint("user_id")

	// Проверяем, что ID из URL совпадает с ID из токена
	if uint(userIDFromURL) != userIDFromToken {
		c.JSON(http.StatusForbidden, gin.H{"error": "нет доступа к заказам другого пользователя"})
		return
	}

	orders, err := h.orderService.GetByUserID(userIDFromToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
