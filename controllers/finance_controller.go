package controllers

import (
	"net/http"
	"warteg-system-backend/database"
	"warteg-system-backend/models"

	"github.com/gin-gonic/gin"
)

// --- Wallet Handlers ---

func GetWallets(c *gin.Context) {
	userID, _ := c.Get("userID")
	var wallets []models.Wallet
	database.DB.Where("user_id = ?", userID).Find(&wallets)
	c.JSON(http.StatusOK, wallets)
}

func CreateWallet(c *gin.Context) {
	var wallet models.Wallet
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := c.Get("userID")
	wallet.UserID = userID.(uint)

	if err := database.DB.Create(&wallet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wallet"})
		return
	}
	c.JSON(http.StatusCreated, wallet)
}

// --- Category Handlers ---

func GetCategories(c *gin.Context) {
	userID, _ := c.Get("userID")
	var categories []models.Category
	database.DB.Where("user_id = ? OR user_id = 0", userID).Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := c.Get("userID")
	category.UserID = userID.(uint)

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	c.JSON(http.StatusCreated, category)
}

// --- Transaction Handlers ---

func GetTransactions(c *gin.Context) {
	userID, _ := c.Get("userID")
	var transactions []models.Transaction
	
	query := database.DB.Preload("Category").Preload("Wallet").Where("user_id = ?", userID)
	
	// Optional filtering by type
	tType := c.Query("type")
	if tType != "" {
		query = query.Where("type = ?", tType)
	}

	query.Order("date desc").Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var req models.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	
	transaction := models.Transaction{
		UserID:      userID.(uint),
		WalletID:    req.WalletID,
		CategoryID:  req.CategoryID,
		Amount:      req.Amount,
		Type:        req.Type,
		Description: req.Description,
		Date:        req.Date,
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record transaction"})
		return
	}

	// Preload for response
	database.DB.Preload("Category").Preload("Wallet").First(&transaction, transaction.ID)
	c.JSON(http.StatusCreated, transaction)
}

func GetSummary(c *gin.Context) {
	userID, _ := c.Get("userID")
	
	var summary struct {
		TotalIncome  float64 `json:"total_income"`
		TotalExpense float64 `json:"total_expense"`
		Balance      float64 `json:"balance"`
	}

	database.DB.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, "income").
		Select("COALESCE(SUM(amount), 0)").Scan(&summary.TotalIncome)

	database.DB.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ?", userID, "expense").
		Select("COALESCE(SUM(amount), 0)").Scan(&summary.TotalExpense)

	summary.Balance = summary.TotalIncome - summary.TotalExpense
	
	c.JSON(http.StatusOK, summary)
}

func SeedUserDefaults(userID uint) {
	// 1. Create Default Wallets
	wallets := []models.Wallet{
		{Name: "Cash", UserID: userID},
		{Name: "QRIS", UserID: userID},
	}
	for _, w := range wallets {
		database.DB.Create(&w)
	}

	// 2. Create Default Categories
	categories := []models.Category{
		{Name: "Penjualan", Type: "income", Icon: "cart", UserID: userID},
		{Name: "Modal", Type: "income", Icon: "wallet", UserID: userID},
		{Name: "Belanja Pasar", Type: "expense", Icon: "basket", UserID: userID},
		{Name: "Gaji Karyawan", Type: "expense", Icon: "people", UserID: userID},
		{Name: "Operasional (Gas/Listrik)", Type: "expense", Icon: "flame", UserID: userID},
	}
	for _, cat := range categories {
		database.DB.Create(&cat)
	}
}
