/**
 * @Time: 2020/11/29 10:38
 * @Author: solacowa@gmail.com
 * @File: mwclient_response
 * @Software: GoLand
 */

package mwsclient

import "encoding/xml"

type ErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Error   struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"Type"`
		Code    string `xml:"Code"`
		Message string `xml:"Message"`
	} `xml:"Error"`
	RequestId string `xml:"RequestId"`
}

type GetCompetitivePricingForASINResult struct {
	Text    string `xml:",chardata"`
	ASIN    string `xml:"ASIN,attr"`
	Status  string `xml:"status,attr"`
	Product struct {
		Text        string `xml:",chardata"`
		Identifiers struct {
			Text            string `xml:",chardata"`
			MarketplaceASIN struct {
				Text          string `xml:",chardata"`
				MarketplaceId string `xml:"MarketplaceId"`
				ASIN          string `xml:"ASIN"`
			} `xml:"MarketplaceASIN"`
		} `xml:"Identifiers"`
		CompetitivePricing struct {
			Text              string `xml:",chardata"`
			CompetitivePrices struct {
				Text             string `xml:",chardata"`
				CompetitivePrice struct {
					Text               string `xml:",chardata"`
					BelongsToRequester string `xml:"belongsToRequester,attr"`
					Condition          string `xml:"condition,attr"`
					Subcondition       string `xml:"subcondition,attr"`
					CompetitivePriceId string `xml:"CompetitivePriceId"`
					Price              struct {
						Text        string `xml:",chardata"`
						LandedPrice struct {
							Text         string `xml:",chardata"`
							CurrencyCode string `xml:"CurrencyCode"`
							Amount       string `xml:"Amount"`
						} `xml:"LandedPrice"` // 降价
						ListingPrice struct {
							Text         string `xml:",chardata"`
							CurrencyCode string `xml:"CurrencyCode"`
							Amount       string `xml:"Amount"`
						} `xml:"ListingPrice"` // 挂牌价
						Shipping struct {
							Text         string `xml:",chardata"`
							CurrencyCode string `xml:"CurrencyCode"`
							Amount       string `xml:"Amount"`
						} `xml:"Shipping"`
					} `xml:"Price"`
				} `xml:"CompetitivePrice"`
			} `xml:"CompetitivePrices"`
			NumberOfOfferListings struct {
				Text              string `xml:",chardata"`
				OfferListingCount []struct {
					Text      string `xml:",chardata"`
					Condition string `xml:"condition,attr"`
				} `xml:"OfferListingCount"`
			} `xml:"NumberOfOfferListings"`
		} `xml:"CompetitivePricing"`
		SalesRankings struct {
			Text      string `xml:",chardata"`
			SalesRank []struct {
				Text              string `xml:",chardata"`
				ProductCategoryId string `xml:"ProductCategoryId"`
				Rank              string `xml:"Rank"`
			} `xml:"SalesRank"`
		} `xml:"SalesRankings"`
	} `xml:"Product"`
}

type GetCompetitivePricingForASINResponse struct {
	XMLName                            xml.Name                             `xml:"GetCompetitivePricingForASINResponse"`
	Text                               string                               `xml:",chardata"`
	Xmlns                              string                               `xml:"xmlns,attr"`
	GetCompetitivePricingForASINResult []GetCompetitivePricingForASINResult `xml:"GetCompetitivePricingForASINResult"`
	ResponseMetadata                   struct {
		Text      string `xml:",chardata"`
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}

type GetMatchingProductForIdResult struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"Id,attr"`
	IdType   string `xml:"IdType,attr"`
	Status   string `xml:"status,attr"`
	Products struct {
		Text    string `xml:",chardata"`
		Ns2     string `xml:"ns2,attr"`
		Product struct {
			Text        string `xml:",chardata"`
			Identifiers struct {
				Text            string `xml:",chardata"`
				MarketplaceASIN struct {
					Text          string `xml:",chardata"`
					MarketplaceId string `xml:"MarketplaceId"`
					ASIN          string `xml:"ASIN"`
				} `xml:"MarketplaceASIN"`
			} `xml:"Identifiers"`
			AttributeSets struct {
				Text           string `xml:",chardata"`
				ItemAttributes struct {
					Text            string `xml:",chardata"`
					Lang            string `xml:"lang,attr"`
					Binding         string `xml:"Binding"`
					Brand           string `xml:"Brand"`
					Color           string `xml:"Color"`
					Label           string `xml:"Label"`
					Manufacturer    string `xml:"Manufacturer"`
					PartNumber      string `xml:"PartNumber"`
					ProductGroup    string `xml:"ProductGroup"`
					ProductTypeName string `xml:"ProductTypeName"`
					Publisher       string `xml:"Publisher"`
					SmallImage      struct {
						Text   string `xml:",chardata"`
						URL    string `xml:"URL"`
						Height struct {
							Text  string `xml:",chardata"`
							Units string `xml:"Units,attr"`
						} `xml:"Height"`
						Width struct {
							Text  string `xml:",chardata"`
							Units string `xml:"Units,attr"`
						} `xml:"Width"`
					} `xml:"SmallImage"`
					Studio string `xml:"Studio"`
					Title  string `xml:"Title"`
				} `xml:"ItemAttributes"`
			} `xml:"AttributeSets"`
			Relationships struct {
				Text            string `xml:",chardata"`
				VariationParent struct {
					Text        string `xml:",chardata"`
					Identifiers struct {
						Text            string `xml:",chardata"`
						MarketplaceASIN struct {
							Text          string `xml:",chardata"`
							MarketplaceId string `xml:"MarketplaceId"`
							ASIN          string `xml:"ASIN"`
						} `xml:"MarketplaceASIN"`
					} `xml:"Identifiers"`
				} `xml:"VariationParent"`
			} `xml:"Relationships"`
			SalesRankings string `xml:"SalesRankings"`
		} `xml:"Product"`
	} `xml:"Products"`
}

type GetMatchingProductForIdResponse struct {
	XMLName                       xml.Name                        `xml:"GetMatchingProductForIdResponse"`
	Text                          string                          `xml:",chardata"`
	Xmlns                         string                          `xml:"xmlns,attr"`
	GetMatchingProductForIdResult []GetMatchingProductForIdResult `xml:"GetMatchingProductForIdResult"`
	ResponseMetadata              struct {
		Text      string `xml:",chardata"`
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}

type GetMatchingProductResult struct {
	Text    string `xml:",chardata"`
	ASIN    string `xml:"ASIN,attr"`
	Status  string `xml:"status,attr"`
	Product struct {
		Text        string `xml:",chardata"`
		Ns2         string `xml:"ns2,attr"`
		Identifiers struct {
			Text            string `xml:",chardata"`
			MarketplaceASIN struct {
				Text          string `xml:",chardata"`
				MarketplaceId string `xml:"MarketplaceId"`
				ASIN          string `xml:"ASIN"`
			} `xml:"MarketplaceASIN"`
		} `xml:"Identifiers"`
		AttributeSets struct {
			Text           string `xml:",chardata"`
			ItemAttributes struct {
				Text            string `xml:",chardata"`
				Lang            string `xml:"lang,attr"`
				Brand           string `xml:"Brand"`
				Color           string `xml:"Color"`
				Label           string `xml:"Label"`
				Manufacturer    string `xml:"Manufacturer"`
				PackageQuantity string `xml:"PackageQuantity"`
				PartNumber      string `xml:"PartNumber"`
				ProductGroup    string `xml:"ProductGroup"`
				ProductTypeName string `xml:"ProductTypeName"`
				Publisher       string `xml:"Publisher"`
				Size            string `xml:"Size"`
				SmallImage      struct {
					Text   string `xml:",chardata"`
					URL    string `xml:"URL"`
					Height struct {
						Text  string `xml:",chardata"`
						Units string `xml:"Units,attr"`
					} `xml:"Height"`
					Width struct {
						Text  string `xml:",chardata"`
						Units string `xml:"Units,attr"`
					} `xml:"Width"`
				} `xml:"SmallImage"`
				Studio string `xml:"Studio"`
				Title  string `xml:"Title"`
			} `xml:"ItemAttributes"`
		} `xml:"AttributeSets"`
		Relationships struct {
			Text            string `xml:",chardata"`
			VariationParent struct {
				Text        string `xml:",chardata"`
				Identifiers struct {
					Text            string `xml:",chardata"`
					MarketplaceASIN struct {
						Text          string `xml:",chardata"`
						MarketplaceId string `xml:"MarketplaceId"`
						ASIN          string `xml:"ASIN"`
					} `xml:"MarketplaceASIN"`
				} `xml:"Identifiers"`
			} `xml:"VariationParent"`
		} `xml:"Relationships"`
		SalesRankings struct {
			Text      string `xml:",chardata"`
			SalesRank []struct {
				Text              string `xml:",chardata"`
				ProductCategoryId string `xml:"ProductCategoryId"`
				Rank              string `xml:"Rank"`
			} `xml:"SalesRank"`
		} `xml:"SalesRankings"`
	} `xml:"Product"`
}

type GetMatchingProductResponse struct {
	XMLName                  xml.Name                   `xml:"GetMatchingProductResponse"`
	Text                     string                     `xml:",chardata"`
	Xmlns                    string                     `xml:"xmlns,attr"`
	GetMatchingProductResult []GetMatchingProductResult `xml:"GetMatchingProductResult"`
	ResponseMetadata         struct {
		Text      string `xml:",chardata"`
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}

type Order struct {
	Text                   string `xml:",chardata"`
	LatestShipDate         string `xml:"LatestShipDate"`
	OrderType              string `xml:"OrderType"`
	PurchaseDate           string `xml:"PurchaseDate"`
	AmazonOrderId          string `xml:"AmazonOrderId"`
	BuyerEmail             string `xml:"BuyerEmail"`
	LastUpdateDate         string `xml:"LastUpdateDate"`
	IsReplacementOrder     string `xml:"IsReplacementOrder"`
	NumberOfItemsShipped   string `xml:"NumberOfItemsShipped"`
	ShipServiceLevel       string `xml:"ShipServiceLevel"`
	OrderStatus            string `xml:"OrderStatus"`
	SalesChannel           string `xml:"SalesChannel"`
	ShippedByAmazonTFM     string `xml:"ShippedByAmazonTFM"`
	IsBusinessOrder        string `xml:"IsBusinessOrder"`
	NumberOfItemsUnshipped string `xml:"NumberOfItemsUnshipped"`
	LatestDeliveryDate     string `xml:"LatestDeliveryDate"`
	PaymentMethodDetails   struct {
		Text                string   `xml:",chardata"`
		PaymentMethodDetail []string `xml:"PaymentMethodDetail"`
	} `xml:"PaymentMethodDetails"`
	IsGlobalExpressEnabled string `xml:"IsGlobalExpressEnabled"`
	IsSoldByAB             string `xml:"IsSoldByAB"`
	BuyerName              string `xml:"BuyerName"`
	EarliestDeliveryDate   string `xml:"EarliestDeliveryDate"`
	IsPremiumOrder         string `xml:"IsPremiumOrder"`
	OrderTotal             struct {
		Text         string `xml:",chardata"`
		Amount       string `xml:"Amount"`
		CurrencyCode string `xml:"CurrencyCode"`
	} `xml:"OrderTotal"`
	EarliestShipDate               string `xml:"EarliestShipDate"`
	MarketplaceId                  string `xml:"MarketplaceId"`
	DefaultShipFromLocationAddress struct {
		Text                         string `xml:",chardata"`
		City                         string `xml:"City"`
		PostalCode                   string `xml:"PostalCode"`
		IsAddressSharingConfidential string `xml:"isAddressSharingConfidential"`
		StateOrRegion                string `xml:"StateOrRegion"`
		Phone                        string `xml:"Phone"`
		CountryCode                  string `xml:"CountryCode"`
		Name                         string `xml:"Name"`
		AddressLine1                 string `xml:"AddressLine1"`
	} `xml:"DefaultShipFromLocationAddress"`
	FulfillmentChannel string `xml:"FulfillmentChannel"`
	PaymentMethod      string `xml:"PaymentMethod"`
	ShippingAddress    struct {
		Text                         string `xml:",chardata"`
		City                         string `xml:"City"`
		AddressType                  string `xml:"AddressType"`
		PostalCode                   string `xml:"PostalCode"`
		IsAddressSharingConfidential string `xml:"isAddressSharingConfidential"`
		Phone                        string `xml:"Phone"`
		CountryCode                  string `xml:"CountryCode"`
		Name                         string `xml:"Name"`
		AddressLine2                 string `xml:"AddressLine2"`
		StateOrRegion                string `xml:"StateOrRegion"`
		AddressLine1                 string `xml:"AddressLine1"`
	} `xml:"ShippingAddress"`
	IsISPU                       string `xml:"IsISPU"`
	IsPrime                      string `xml:"IsPrime"`
	ShipmentServiceLevelCategory string `xml:"ShipmentServiceLevelCategory"`
}

type ListOrdersResponse struct {
	XMLName          xml.Name `xml:"ListOrdersResponse"`
	Text             string   `xml:",chardata"`
	Xmlns            string   `xml:"xmlns,attr"`
	ListOrdersResult struct {
		Text   string `xml:",chardata"`
		Orders struct {
			Text  string  `xml:",chardata"`
			Order []Order `xml:"Order"`
		} `xml:"Orders"`
		CreatedBefore string `xml:"CreatedBefore"`
	} `xml:"ListOrdersResult"`
	ResponseMetadata struct {
		Text      string `xml:",chardata"`
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}
