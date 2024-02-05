package handlers

// func AddProduct(c *gin.Context) {
// 	var input domain.Product
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	err := repository.GetProductByName(input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = repository.CreateProduct(input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "creating product failed"})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"success": "product created succesfuly"})
// 	}

// }

// func ListProducts(c *gin.Context) {
// 	products, err := repository.GetAllProduct()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"Products": products})
// 	}

// }
