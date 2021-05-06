package sponsored

import (
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/lyticaa/lyticaa-data/internal/models"
)

var (
	reg, _ = regexp.Compile(`[^0-9\.]+`)
)

func (s *Sponsored) dateTime(row map[string]string) time.Time {
	var dateTime time.Time

	if days, err := strconv.Atoi(row["Date"]); err == nil {
		loc, _ := time.LoadLocation(os.Getenv("TZ"))
		dateTime = time.Date(1900, time.January, days, 0, 0, 0, 0, loc)
	} else {
		dateTime, _ = time.Parse("Jan 2, 2006", row["Date"])
	}

	return dateTime
}

func (s *Sponsored) portfolioName(row map[string]string) string {
	return row["Portfolio name"]
}

func (s *Sponsored) marketplaceID(row map[string]string) int64 {
	exchangeRateID := s.exchangeRateID(row)
	marketplaces := *models.LoadAmazonMarketplaces(s.db)
	for _, marketplace := range marketplaces {
		if marketplace.ExchangeRateID == exchangeRateID {
			return marketplace.ID
		}
	}

	return unknown
}

func (s *Sponsored) exchangeRateID(row map[string]string) int64 {
	for _, exchangeRate := range *s.exchangeRates() {
		if exchangeRate.Code == row["Currency"] {
			return exchangeRate.ID
		}
	}

	return unknown
}

func (s *Sponsored) campaignName(row map[string]string) string {
	return row["Campaign Name"]
}

func (s *Sponsored) adGroupName(row map[string]string) string {
	return row["Ad Group Name"]
}

func (s *Sponsored) sku(row map[string]string) string {
	return row["Advertised SKU"]
}

func (s *Sponsored) asin(row map[string]string) string {
	return row["Advertised ASIN"]
}

func (s *Sponsored) impressions(row map[string]string) int64 {
	impressions, _ := strconv.ParseInt(row["Impressions"], 10, 64)
	return impressions
}

func (s *Sponsored) clicks(row map[string]string) int64 {
	clicks, _ := strconv.ParseInt(row["Clicks"], 10, 64)
	return clicks
}

func (s *Sponsored) ctr(row map[string]string) float64 {
	ctr, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Click-Thru Rate (CTR)"], ""), 64)
	return ctr
}

func (s *Sponsored) cpc(row map[string]string) float64 {
	cpc, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Cost Per Click (CPC)"], ""), 64)
	return cpc
}

func (s *Sponsored) spend(row map[string]string) float64 {
	spend, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Spend"], ""), 64)
	return spend
}

func (s *Sponsored) totalSales(row map[string]string) float64 {
	totalSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Total Sales"], ""), 64)
	return totalSales
}

func (s *Sponsored) acos(row map[string]string) float64 {
	acos, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Total Advertising Cost of Sales (ACoS)"], ""), 64)
	return acos
}

func (s *Sponsored) roas(row map[string]string) float64 {
	roas, _ := strconv.ParseFloat(row["Total Return on Advertising Spend (RoAS)"], 64)
	return roas
}

func (s *Sponsored) totalOrders(row map[string]string) int64 {
	totalOrders, _ := strconv.ParseInt(row["7 Day Total Orders (#)"], 10, 64)
	return totalOrders
}

func (s *Sponsored) totalUnits(row map[string]string) int64 {
	totalUnits, _ := strconv.ParseInt(row["7 Day Total Units (#)"], 10, 64)
	return totalUnits
}

func (s *Sponsored) conversionRate(row map[string]string) float64 {
	conversionRate, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Conversion Rate"], ""), 64)
	return conversionRate
}

func (s *Sponsored) advertisedSKUUnits(row map[string]string) int64 {
	advertisedSKUUnits, _ := strconv.ParseInt(row["7 Day Advertised SKU Units (#)"], 10, 64)
	return advertisedSKUUnits
}

func (s *Sponsored) otherSKUUnits(row map[string]string) int64 {
	otherSKUUnits, _ := strconv.ParseInt(row["7 Day Other SKU Units (#)"], 10, 64)
	return otherSKUUnits
}

func (s *Sponsored) advertisedSKUSales(row map[string]string) float64 {
	advertisedSKUSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Advertised SKU Sales"], ""), 64)
	return advertisedSKUSales
}

func (s *Sponsored) otherSKUSales(row map[string]string) float64 {
	otherSKUSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Other SKU Sales"], ""), 64)
	return otherSKUSales
}
