package sponsored

import (
	"github.com/lyticaa/lyticaa-data/internal/models"
)

const (
	unknown = int64(1)
)

func (s *Sponsored) Process(rows []map[string]string, userID string) []error {
	var errors []error

	formatted := s.Format(rows, userID)
	for _, item := range formatted {
		err := s.Save(item)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (s *Sponsored) Format(rows []map[string]string, userID string) []models.AmazonSponsoredProduct {
	var sponsoredProducts []models.AmazonSponsoredProduct
	for _, row := range rows {
		sponsoredProducts = append(sponsoredProducts, models.AmazonSponsoredProduct{
			UserID:              userID,
			DateTime:            s.dateTime(row),
			PortfolioName:       s.portfolioName(row),
			AmazonMarketplaceID: s.marketplaceID(row),
			CampaignName:        s.campaignName(row),
			AdGroupName:         s.adGroupName(row),
			SKU:                 s.sku(row),
			ASIN:                s.asin(row),
			Impressions:         s.impressions(row),
			Clicks:              s.clicks(row),
			CTR:                 s.ctr(row),
			CPC:                 s.cpc(row),
			Spend:               s.spend(row),
			TotalSales:          s.totalSales(row),
			ACoS:                s.acos(row),
			RoAS:                s.roas(row),
			TotalOrders:         s.totalOrders(row),
			TotalUnits:          s.totalUnits(row),
			ConversionRate:      s.conversionRate(row),
			AdvertisedSKUUnits:  s.advertisedSKUUnits(row),
			OtherSKUUnits:       s.otherSKUUnits(row),
			AdvertisedSKUSales:  s.advertisedSKUSales(row),
			OtherSKUSales:       s.otherSKUSales(row),
		})
	}

	return sponsoredProducts
}

func (s *Sponsored) Save(sponsoredProduct models.AmazonSponsoredProduct) error {
	err := sponsoredProduct.Save(s.db)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sponsored) exchangeRates() *[]models.ExchangeRate {
	return models.LoadExchangeRates(s.db)
}
