package main

type CreatePostRequest struct {
	RecordId     uint    `json:"record_id" binding:"required"`
	Title        string  `json:"title" binding:"required,max=10"`
	Price        float64 `json:"price" binding:"required"`
	Content      string  `json:"content" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	SupplierName string  `json:"supplier_name" binding:"required"`
	SupplierId   uint    `json:"supplier_id" binding:"required"`
	CompanyId    uint    `json:"company_id" binding:"required"`
}
