package transaction

import (
	"github.com/lyticaa/lyticaa-data/internal/models"
)

const (
	unknown = int64(1)
)

func (t *Transaction) Process(rows []map[string]string, userID string) []error {
	var errors []error

	formatted := t.Format(rows, userID)
	for _, item := range formatted {
		err := t.Save(item)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (t *Transaction) Format(rows []map[string]string, userID string) []models.AmazonCustomTransaction {
	var transactions []models.AmazonCustomTransaction
	for _, row := range rows {
		transactions = append(transactions, models.AmazonCustomTransaction{
			DateTime:                   t.dateTime(row),
			UserID:                     userID,
			SettlementID:               t.settlementID(row),
			AmazonTransactionTypeID:    t.amazonTransactionType(row),
			OrderID:                    t.orderID(row),
			SKU:                        t.sku(row),
			Quantity:                   t.quantity(row),
			AmazonMarketplaceID:        t.amazonMarketplace(row),
			AmazonFulfillmentID:        t.amazonFulfillment(row),
			AmazonTaxCollectionModelID: t.amazonTaxCollectionModel(row),
			ProductSales:               t.productSales(row),
			ProductSalesTax:            t.productSalesTax(row),
			ShippingCredits:            t.shippingCredits(row),
			ShippingCreditsTax:         t.shippingCreditsTax(row),
			GiftwrapCredits:            t.giftwrapCredits(row),
			GiftwrapCreditsTax:         t.giftwrapCreditsTax(row),
			PromotionalRebates:         t.promotionalRebates(row),
			PromotionalRebatesTax:      t.promotionalRebatesTax(row),
			MarketplaceWithheldTax:     t.marketplaceWithheldTax(row),
			SellingFees:                t.sellingFees(row),
			FBAFees:                    t.fbaFees(row),
			OtherTransactionFees:       t.otherTransactionFees(row),
			Other:                      t.other(row),
			Total:                      t.total(row),
		})
	}

	return transactions
}

func (t *Transaction) Save(transaction models.AmazonCustomTransaction) error {
	err := transaction.Save(t.db)
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) AmazonTransactionTypes() *[]models.AmazonTransactionType {
	return models.LoadAmazonTransactionTypes(t.db)
}

func (t *Transaction) amazonTransactionTypeIDByName(name string) (int64, bool) {
	txnTypes := *t.AmazonTransactionTypes()
	for _, txnType := range txnTypes {
		if txnType.Name == name {
			return txnType.ID, true
		}
	}

	return unknown, false
}

func (t *Transaction) AmazonMarketplaces() *[]models.AmazonMarketplace {
	return models.LoadAmazonMarketplaces(t.db)
}

func (t *Transaction) amazonMarketplaceIDByName(name string) (int64, bool) {
	marketplaces := *t.AmazonMarketplaces()
	for _, marketplace := range marketplaces {
		if marketplace.Name == name {
			return marketplace.ID, true
		}
	}

	return unknown, false
}

func (t *Transaction) AmazonFulfillments() *[]models.AmazonFulfillment {
	return models.LoadAmazonFulfillments(t.db)
}

func (t *Transaction) amazonFulfillmentIDByName(name string) (int64, bool) {
	fulfillments := *t.AmazonFulfillments()
	for _, fulfillment := range fulfillments {
		if fulfillment.Name == name {
			return fulfillment.ID, true
		}
	}

	return unknown, false
}

func (t *Transaction) AmazonTaxCollectionModels() *[]models.AmazonTaxCollectionModel {
	return models.LoadAmazonTaxCollectionModels(t.db)
}

func (t *Transaction) amazonTaxCollectionModelIDByName(name string) (int64, bool) {
	taxCollectionModels := *t.AmazonTaxCollectionModels()
	for _, taxCollectionModel := range taxCollectionModels {
		if taxCollectionModel.Name == name {
			return taxCollectionModel.ID, true
		}
	}

	return unknown, false
}
