package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonCustomTransaction struct {
	ID                         int64     `db:"id"`
	UserID                     string    `db:"user_id"`
	DateTime                   time.Time `db:"date_time"`
	DateRange                  string    `db:"date_range"`
	SettlementID               int64     `db:"settlement_id"`
	AmazonTransactionType      string    `db:"amazon_transaction_type"`
	AmazonTransactionTypeID    int64     `db:"amazon_transaction_type_id"`
	OrderID                    string    `db:"order_id"`
	SKU                        string    `db:"sku"`
	Quantity                   int64     `db:"quantity"`
	AmazonMarketplace          string    `db:"amazon_marketplace"`
	AmazonMarketplaceID        int64     `db:"amazon_marketplace_id"`
	AmazonFulfillmentID        int64     `db:"amazon_fulfillment_id"`
	AmazonTaxCollectionModelID int64     `dn:"amazon_tax_collection_model_id"`
	ProductSales               float64   `db:"product_sales"`
	ProductSalesTax            float64   `db:"product_sales_tax"`
	ShippingCredits            float64   `db:"shipping_credits"`
	ShippingCreditsTax         float64   `db:"shipping_credits_tax"`
	GiftwrapCredits            float64   `db:"giftwrap_credits"`
	GiftwrapCreditsTax         float64   `db:"giftwrap_credits_tax"`
	PromotionalRebates         float64   `db:"promotional_rebates"`
	PromotionalRebatesTax      float64   `db:"promotional_rebates_tax"`
	MarketplaceWithheldTax     float64   `db:"marketplace_withheld_tax"`
	SellingFees                float64   `db:"selling_fees"`
	FBAFees                    float64   `db:"fba_fees"`
	OtherTransactionFees       float64   `db:"other_transaction_fees"`
	Other                      float64   `db:"other"`
	Total                      float64   `db:"total"`
	CreatedAt                  time.Time `db:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
}

func LoadAmazonMarketplacesByDateRange(userID string, filter *Filter, db *sqlx.DB) *[]AmazonCustomTransaction {
	var transactions []AmazonCustomTransaction

	query := `SELECT
       DISTINCT(m.name) AS amazon_marketplace FROM amazon_transactions AS t
           LEFT JOIN amazon_marketplaces AS m ON t.amazon_marketplace_id = m.id
           LEFT JOIN amazon_transaction_types tt ON t.amazon_transaction_type_id = tt.id WHERE tt.name IN ('Order', 'Refund')
                                                                                           AND t.user_id = $1
                                                                                           AND t.date_time BETWEEN $2 AND $3`
	_ = db.Select(
		&transactions,
		query,
		userID,
		filter.StartDate,
		filter.EndDate,
	)

	return &transactions
}

func LoadAmazonOrdersByDateRange(userID, view string, db *sqlx.DB) *[]AmazonCustomTransaction {
	var transactions []AmazonCustomTransaction

	query := `SELECT
       '%v' AS date_range,
       t.*,
       tt.name AS amazon_transaction_type,
       m.name AS amazon_marketplace FROM amazon_transactions_%v AS t
           LEFT JOIN amazon_marketplaces AS m ON t.amazon_marketplace_id = m.id
           LEFT JOIN amazon_transaction_types tt ON t.amazon_transaction_type_id = tt.id WHERE tt.name IN ('Order', 'Refund')
                                                                                           AND t.user_id = $1`
	_ = db.Select(
		&transactions,
		fmt.Sprintf(query, view, view),
		userID,
	)

	return &transactions
}

func LoadAmazonUsers(db *sqlx.DB) *[]User {
	var users []User

	query := `SELECT user_id FROM amazon_users`
	_ = db.Select(&users, query)

	return &users
}

func LoadAmazonProduct(userID, sku, marketplace string, db *sqlx.DB) *Product {
	var product Product

	query := `SELECT
       sku,
       user_id,
       marketplace FROM amazon_products WHERE user_id = $1
                                          AND sku = $2
                                          AND marketplace = $3`
	_ = db.QueryRow(query, userID, sku, marketplace).Scan(
		&product.SKU,
		&product.UserID,
		&product.Description,
		&product.Marketplace,
	)

	return &product
}

func (t *AmazonCustomTransaction) Save(db *sqlx.DB) error {
	query := `INSERT INTO amazon_custom_transactions (
                                        user_id,
                                        date_time,
                                        settlement_id,
                                        amazon_transaction_type_id,
                                        order_id,
                                        sku,
                                        quantity,
                                        amazon_marketplace_id,
                                        amazon_fulfillment_id,
                                        amazon_tax_collection_model_id,
                                        product_sales,
                                        product_sales_tax,
                                        shipping_credits,
                                        shipping_credits_tax,
                                        giftwrap_credits,
                                        giftwrap_credits_tax,
                                        promotional_rebates,
                                        promotional_rebates_tax,
                                        marketplace_withheld_tax,
                                        selling_fees,
                                        fba_fees,
                                        other_transaction_fees,
                                        other,
                                        total)
                                        VALUES (
                                                :user_id,
                                                :date_time,
                                                :settlement_id,
                                                :amazon_transaction_type_id,
                                                :order_id,
                                                :sku,
                                                :quantity,
                                                :amazon_marketplace_id,
                                                :amazon_fulfillment_id,
                                                :amazon_tax_collection_model_id,
                                                :product_sales,
                                                :product_sales_tax,
                                                :shipping_credits,
                                                :shipping_credits_tax,
                                                :giftwrap_credits,
                                                :giftwrap_credits_tax,
                                                :promotional_rebates,
                                                :promotional_rebates_tax,
                                                :marketplace_withheld_tax,
                                                :selling_fees,
                                                :fba_fees,
                                                :other_transaction_fees,
                                                :other,
                                                :total)
                                                ON CONFLICT (user_id, date_time, settlement_id, amazon_transaction_type_id, order_id, sku)
                                                    DO UPDATE SET quantity = :quantity,
                                                                  amazon_marketplace_id = :amazon_marketplace_id,
                                                                  amazon_fulfillment_id = :amazon_fulfillment_id,
                                                                  amazon_tax_collection_model_id = :amazon_tax_collection_model_id,
                                                                  product_sales = :product_sales,
                                                                  product_sales_tax = :product_sales_tax,
                                                                  shipping_credits = :shipping_credits,
                                                                  shipping_credits_tax = :shipping_credits_tax,
                                                                  giftwrap_credits = :giftwrap_credits,
                                                                  giftwrap_credits_tax = :giftwrap_credits_tax,
                                                                  promotional_rebates = :promotional_rebates,
                                                                  promotional_rebates_tax = :promotional_rebates_tax,
                                                                  marketplace_withheld_tax = :marketplace_withheld_tax,
                                                                  selling_fees = :selling_fees,
                                                                  fba_fees = :fba_fees,
                                                                  other_transaction_fees = :other_transaction_fees,
                                                                  other = :other,
                                                                  total = :total,
                                                                  updated_at = NOW()`
	_, err := db.NamedExec(query, map[string]interface{}{
		"user_id":                        t.UserID,
		"date_time":                      t.DateTime,
		"settlement_id":                  t.SettlementID,
		"amazon_transaction_type_id":     t.AmazonTransactionTypeID,
		"order_id":                       t.OrderID,
		"sku":                            t.SKU,
		"quantity":                       t.Quantity,
		"amazon_marketplace_id":          t.AmazonMarketplaceID,
		"amazon_fulfillment_id":          t.AmazonFulfillmentID,
		"amazon_tax_collection_model_id": t.AmazonTaxCollectionModelID,
		"product_sales":                  t.ProductSales,
		"product_sales_tax":              t.ProductSalesTax,
		"shipping_credits":               t.ShippingCredits,
		"shipping_credits_tax":           t.ShippingCreditsTax,
		"giftwrap_credits":               t.GiftwrapCredits,
		"giftwrap_credits_tax":           t.GiftwrapCreditsTax,
		"promotional_rebates":            t.PromotionalRebates,
		"promotional_rebates_tax":        t.PromotionalRebatesTax,
		"marketplace_withheld_tax":       t.MarketplaceWithheldTax,
		"selling_fees":                   t.SellingFees,
		"fba_fees":                       t.FBAFees,
		"other_transaction_fees":         t.OtherTransactionFees,
		"other":                          t.Other,
		"total":                          t.Total,
	})

	if err != nil {
		return err
	}

	return nil
}
