package transaction

import (
	"strconv"
	"strings"
	"time"
)

func (t *Transaction) stripComma(data string) string {
	return strings.Replace(data, ",", "", -1)
}

func (t *Transaction) amazonTransactionType(row map[string]string) int64 {
	transactionType, _ := t.amazonTransactionTypeIDByName(row["type"])
	return transactionType
}

func (t *Transaction) orderID(row map[string]string) string {
	return row["order id"]
}

func (t *Transaction) sku(row map[string]string) string {
	return row["sku"]
}

func (t *Transaction) description(row map[string]string) string {
	return row["description"]
}

func (t *Transaction) amazonMarketplace(row map[string]string) int64 {
	marketplace, _ := t.amazonMarketplaceIDByName(strings.ToLower(row["marketplace"]))
	return marketplace
}

func (t *Transaction) amazonFulfillment(row map[string]string) int64 {
	fulfillment, _ := t.amazonFulfillmentIDByName(row["fulfillment"])
	return fulfillment
}

func (t *Transaction) amazonTaxCollectionModel(row map[string]string) int64 {
	taxCollectionModel, _ := t.amazonTaxCollectionModelIDByName(row["tax collection model"])
	return taxCollectionModel
}

func (t *Transaction) dateTime(row map[string]string) time.Time {
	dt, _ := time.Parse("Jan 2, 2006 3:04:05 PM MST", row["date/time"])
	return dt
}

func (t *Transaction) settlementID(row map[string]string) int64 {
	settlementID, _ := strconv.ParseInt(row["settlement id"], 10, 64)
	return settlementID
}

func (t *Transaction) quantity(row map[string]string) int64 {
	quantity, _ := strconv.ParseInt(row["quantity"], 10, 64)
	return quantity
}

func (t *Transaction) productSales(row map[string]string) float64 {
	productSales, _ := strconv.ParseFloat(t.stripComma(row["product sales"]), 64)
	return productSales
}

func (t *Transaction) productSalesTax(row map[string]string) float64 {
	productSalesTax, _ := strconv.ParseFloat(t.stripComma(row["product sales tax"]), 64)
	return productSalesTax
}

func (t *Transaction) shippingCredits(row map[string]string) float64 {
	shippingCredits, _ := strconv.ParseFloat(t.stripComma(row["shipping credits"]), 64)
	return shippingCredits
}

func (t *Transaction) shippingCreditsTax(row map[string]string) float64 {
	shippingCreditsTax, _ := strconv.ParseFloat(t.stripComma(row["shipping credits tax"]), 64)
	return shippingCreditsTax
}

func (t *Transaction) giftwrapCredits(row map[string]string) float64 {
	giftwrapCredits, _ := strconv.ParseFloat(t.stripComma(row["gift wrap credits"]), 64)
	return giftwrapCredits
}

func (t *Transaction) giftwrapCreditsTax(row map[string]string) float64 {
	giftwrapCreditsTax, _ := strconv.ParseFloat(t.stripComma(row["giftwrap credits tax"]), 64)
	return giftwrapCreditsTax
}

func (t *Transaction) promotionalRebates(row map[string]string) float64 {
	promotionalRebates, _ := strconv.ParseFloat(t.stripComma(row["promotional rebates"]), 64)
	return promotionalRebates
}

func (t *Transaction) promotionalRebatesTax(row map[string]string) float64 {
	promotionalRebatesTax, _ := strconv.ParseFloat(t.stripComma(row["promotional rebates tax"]), 64)
	return promotionalRebatesTax
}

func (t *Transaction) marketplaceWithheldTax(row map[string]string) float64 {
	marketplaceWithheldTax, _ := strconv.ParseFloat(t.stripComma(row["marketplace withheld tax"]), 64)
	return marketplaceWithheldTax
}

func (t *Transaction) sellingFees(row map[string]string) float64 {
	sellingFees, _ := strconv.ParseFloat(t.stripComma(row["selling fees"]), 64)
	return sellingFees
}

func (t *Transaction) fbaFees(row map[string]string) float64 {
	fbaFees, _ := strconv.ParseFloat(t.stripComma(row["fba fees"]), 64)
	return fbaFees
}

func (t *Transaction) otherTransactionFees(row map[string]string) float64 {
	otherTransactionFees, _ := strconv.ParseFloat(t.stripComma(row["other transaction fees"]), 64)
	return otherTransactionFees
}

func (t *Transaction) other(row map[string]string) float64 {
	other, _ := strconv.ParseFloat(t.stripComma(row["other"]), 64)
	return other
}

func (t *Transaction) total(row map[string]string) float64 {
	total, _ := strconv.ParseFloat(t.stripComma(row["total"]), 64)
	return total
}
