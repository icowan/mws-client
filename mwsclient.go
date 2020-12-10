/**
 * @Time: 2020/9/26 21:08
 * @Author: solacowa@gmail.com
 * @File: mwsclient
 * @Software: GoLand
 */

package mwsclient

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/kplcloud/request"
	"github.com/pkg/errors"
)

type RequestMethod struct {
	Method string
	Action string
	Path   string
	Date   string
}

var (
	endpoints = map[string]RequestMethod{
		"GetFeedSubmissionResult": {
			Method: http.MethodPost,
			Action: "GetFeedSubmissionResult",
			Path:   "/",
			Date:   "2009-01-01",
		},
		"ListOrders": {
			Method: http.MethodPost,
			Action: "ListOrders",
			Path:   "/Orders/2013-09-01",
			Date:   "2013-09-01",
		},
		"GetMatchingProductForId": {
			Method: http.MethodPost,
			Action: "GetMatchingProductForId",
			Path:   "/Products/2011-10-01",
			Date:   "2011-10-01",
		},
		"GetMatchingProduct": {
			Method: http.MethodPost,
			Action: "GetMatchingProduct",
			Path:   "/Products/2011-10-01",
			Date:   "2011-10-01",
		},
		"GetCompetitivePricingForASIN": {
			Method: http.MethodPost,
			Action: "GetCompetitivePricingForASIN",
			Path:   "/Products/2011-10-01",
			Date:   "2011-10-01",
		},
	}

	MarketplaceIds = map[string]string{
		"BR": "A2Q3Y263D00KWC",
		"CA": "A2EUQ1WTGCTBG2",
		"MX": "A1AM78C64UM0Y8",
		"US": "ATVPDKIKX0DER",
		"AE": "A2VIGQ35RCS4UG",
		"DE": "A1PA6795UKMFR9",
		"ES": "A1RKKUPIHCS9HS",
		"FR": "A13V1IB3VIYZZH",
		"GB": "A1F83G8C2ARO7P",
		"IN": "A21TJRUUN4KGV",
		"IT": "APJ6JRA9NG5V4",
		"TR": "A33AVAJ2PDY3EV",
		"AU": "A39IBJ37TRP1C6",
		"JP": "A1VC38T7YXB528",
		"CN": "AAHKV2X7AFYLW",
		"NL": "A1805IZSGTT6HS",
		"SA": "A17E79C6D8DWNP",
		"GE": "ARBP9OOSHTCHU",
	}

	MarketplaceEndpoints = map[string]string{
		"BR": "https://mws.amazonservices.com",
		"CA": "https://mws.amazonservices.ca",
		"MX": "https://mws.amazonservices.com.mx",
		"US": "https://mws.amazonservices.com",
		"AE": "https://mws.amazonservices.ae",
		"DE": "https://mws-eu.amazonservices.com",
		"ES": "https://mws-eu.amazonservices.com",
		"FR": "https://mws-eu.amazonservices.com",
		"GB": "https://mws-eu.amazonservices.com",
		"IN": "https://mws.amazonservices.in",
		"IT": "https://mws-eu.amazonservices.com",
		"TR": "https://mws-eu.amazonservices.com",
		"AU": "https://mws.amazonservices.com.au",
		"JP": "https://mws.amazonservices.jp",
		"CN": "https://mws.amazonservices.com.cn",
		"NL": "https://mws-eu.amazonservices.com",
		"SA": "https://mws-eu.amazonservices.com",
		"GE": "https://mws-eu.amazonservices.com",
	}
	MarketplaceIdUrls = map[string]string{
		"A2EUQ1WTGCTBG2": "mws.amazonservices.ca",
		"ATVPDKIKX0DER":  "mws.amazonservices.com",
		"A1AM78C64UM0Y8": "mws.amazonservices.com.mx",
		"A1PA6795UKMFR9": "mws-eu.amazonservices.com",
		"A1RKKUPIHCS9HS": "mws-eu.amazonservices.com",
		"A13V1IB3VIYZZH": "mws-eu.amazonservices.com",
		"A21TJRUUN4KGV":  "mws.amazonservices.in",
		"APJ6JRA9NG5V4":  "mws-eu.amazonservices.com",
		"A1F83G8C2ARO7P": "mws-eu.amazonservices.com",
		"A1VC38T7YXB528": "mws.amazonservices.jp",
		"AAHKV2X7AFYLW":  "mws.amazonservices.com.cn",
		"A39IBJ37TRP1C6": "mws.amazonservices.com.au",
		"A2Q3Y263D00KWC": "mws.amazonservices.com",
		"A1805IZSGTT6HS": "mws-eu.amazonservices.com",
		"ARBP9OOSHTCHU":  "mws-eu.amazonservices.com",
		"A17E79C6D8DWNP": "mws.amazonservices.com",
		"A33AVAJ2PDY3EV": "mws.amazonservices.com",
		"A19VAU5U5O7RUS": "mws-fe.amazonservices.com",
		"A2VIGQ35RCS4UG": "mws.amazonservices.ae",
	}
)

type OrderState string
type OrderChannel string

const (
	OrderStateUnshipped           OrderState = "Unshipped"
	OrderStatePartiallyShipped    OrderState = "PartiallyShipped"
	OrderStateShipped             OrderState = "Shipped"
	OrderStatePendingAvailability OrderState = "PendingAvailability"
	OrderStatePending             OrderState = "Pending"
	OrderStateInvoiceUnconfirmed  OrderState = "InvoiceUnconfirmed"
	OrderStateCanceled            OrderState = "Canceled"
	OrderStateUnfulfillable       OrderState = "Unfulfillable"

	OrderChannelMFN OrderChannel = "MFN"
	OrderChannelAFN OrderChannel = "AFN"

	dateFormat string = `Y-m-d\TH:i:s.\\0\\0\\0\\Z`
)

func (c OrderState) String() string {
	return string(c)
}

func (c OrderChannel) String() string {
	return string(c)
}

type ListOrderCallFun func(ctx context.Context, data []map[string]interface{}) error

type MwsClient interface {
	SetClient(ctx context.Context, authToken, accessKey, secretKey, sellerId, marketplaceId string) MwsClient

	// 返回您在指定时间段内所创建或更新的订单。
	// 该 ListOrders 操作可返回您在指定时间段内创建或更新的订单列表。您可以通过 CreatedAfter 参数或 LastUpdatedAfter 参数来指定时间段。
	// 您必须使用其中一个参数，但不可同时使用两个参数。您还可以通过应用筛选条件来缩小返回的订单列表范围。
	// 该 ListOrders 操作包括每个所返回订单的订单详情，其中包括 AmazonOrderId、 OrderStatus、 FulfillmentChannel 和 LastUpdateDate。
	ListOrders(ctx context.Context, fromDate time.Time, allMarketplaces bool, states []OrderState,
		fulfillmentChannel []OrderChannel, till *time.Time, callFun ListOrderCallFun) (res []Order, err error)

	// 根据 ASIN、GCID、SellerSKU、UPC、EAN、ISBN 和 JAN，返回商品及其属性列表。
	// 根据您指定的商品编码值列表，GetMatchingProductForId 操作会返回一个包含商品及其属性的列表。
	// 可能的商品编号包括：ASIN、GCID、SellerSKU、UPC、EAN、ISBN 和 JAN。
	GetMatchingProductForId(ctx context.Context, idType string, idList []string) (res []GetMatchingProductForIdResult, err error)

	// 根据 ASIN 值列表，返回商品及其属性列表。
	// 根据您指定的 ASIN 值列表，GetMatchingProduct 操作会返回一个包含商品及其属性的列表。此操作最多会返回十件商品。
	// 重要说明：新的 GetMatchingProductForId 操作具备GetMatchingProduct 操作的所有功能。
	// 为了实现向后兼容性，商品 API 部分包含了 GetMatchingProduct 操作，但是您应尽可能使用 GetMatchingProductForId 操作来代替 GetMatchingProduct 操作。
	GetMatchingProduct(ctx context.Context, ASINList []string) (res []GetMatchingProductResult, err error)

	// 根据 ASIN，返回商品的当前有竞争力的价格。
	// GetCompetitivePricingForASIN 操作与 GetCompetitivePricingForSKU 操作大体相同，但前者使用 MarketplaceId 和 ASIN 来唯一标识一件商品，且不会返回 SKUIdentifier 元素。
	// 如果您没有商品的 ASIN，必须先提交 ListMatchingProducts 操作，以避免歧义。
	GetCompetitivePricingForASIN(ctx context.Context, ASINList []string) (res []GetCompetitivePricingForASINResult, err error)

	// 返回上传数据处理报告及 Content-MD5 标头。
	// 该 GetFeedSubmissionResult 操作可返回上传数据处理报告及所返回 HTTP 正文的 Content-MD5 标头。
	GetFeedSubmissionResult(ctx context.Context, feedSubmissionId string) (err error)
}

type mwsClient struct {
	logger                                                         log.Logger
	traceId                                                        string
	host, authToken, accessKey, secretKey, sellerId, marketplaceId string
}

func (m *mwsClient) GetCompetitivePricingForASIN(ctx context.Context, ASINList []string) (res []GetCompetitivePricingForASINResult, err error) {
	defer func(begin time.Time) {
		_ = m.logger.Log(
			m.traceId, ctx.Value(m.traceId),
			"method", "GetMatchingProduct",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	params := url.Values{}
	params.Add("MarketplaceId", m.marketplaceId)
	for k, v := range ASINList {
		params.Add(fmt.Sprintf("ASINList.ASIN.%d", k+1), v)
	}

	ep := endpoints["GetCompetitivePricingForASIN"]

	var resp GetCompetitivePricingForASINResponse
	err = m.request(params, ep, &resp)
	if err != nil {
		return
	}

	return resp.GetCompetitivePricingForASINResult, nil
}

func (m *mwsClient) GetMatchingProduct(ctx context.Context, ASINList []string) (res []GetMatchingProductResult, err error) {
	defer func(begin time.Time) {
		_ = m.logger.Log(
			m.traceId, ctx.Value(m.traceId),
			"method", "GetMatchingProduct",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	params := url.Values{}
	params.Add("MarketplaceId", m.marketplaceId)
	for k, v := range ASINList {
		params.Add(fmt.Sprintf("ASINList.ASIN.%d", k+1), v)
	}
	ep := endpoints["GetMatchingProduct"]
	var resp GetMatchingProductResponse
	err = m.request(params, ep, &resp)
	if err != nil {
		return
	}
	return resp.GetMatchingProductResult, nil
}

func (m *mwsClient) GetMatchingProductForId(ctx context.Context, idType string, idList []string) (res []GetMatchingProductForIdResult, err error) {
	defer func(begin time.Time) {
		_ = m.logger.Log(
			m.traceId, ctx.Value(m.traceId),
			"method", "GetMatchingProductForId",
			"idType", idType,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	params := url.Values{}
	params.Add("MarketplaceId", m.marketplaceId)
	params.Add("IdType", idType)
	for k, v := range idList {
		params.Add(fmt.Sprintf("IdList.Id.%d", k+1), v)
	}
	ep := endpoints["GetMatchingProductForId"]
	var resp GetMatchingProductForIdResponse
	err = m.request(params, ep, resp)
	if err != nil {
		return
	}
	return resp.GetMatchingProductForIdResult, nil
}

func (m *mwsClient) ListOrders(ctx context.Context, fromDate time.Time, allMarketplaces bool, states []OrderState,
	fulfillmentChannel []OrderChannel, till *time.Time, callFun ListOrderCallFun) (res []Order, err error) {
	defer func(begin time.Time) {
		_ = m.logger.Log(
			m.traceId, ctx.Value(m.traceId),
			"method", "ListOrders",
			"till", till,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	params := url.Values{}
	params.Add("CreatedAfter", fromDate.UTC().Format(time.RFC3339))

	if till != nil {
		params.Set("CreatedBefore", till.UTC().Format(time.RFC3339))
	}

	var counter = 1
	for _, v := range states {
		params.Add(fmt.Sprintf("OrderStatus.Status.%d", counter), v.String())
		counter += 1
	}

	if allMarketplaces {
		counter = 1
		for k := range MarketplaceIdUrls {
			params.Add(fmt.Sprintf("MarketplaceId.Id.%d", counter), k)
			counter += 1
		}
	}

	if len(fulfillmentChannel) > 1 {
		counter = 1
		for _, v := range fulfillmentChannel {
			params.Add(fmt.Sprintf("FulfillmentChannel.Channel.%d", counter), v.String())
			counter += 1
		}
	} else {
		params.Add("FulfillmentChannel.Channel.1", "MFN")
	}

	ep := endpoints["ListOrders"]

	var resp ListOrdersResponse

	err = m.request(params, ep, &resp)
	if err != nil {
		return
	}

	return resp.ListOrdersResult.Orders.Order, nil
}

func (m *mwsClient) genAmazonUrl(ep RequestMethod, params url.Values) (finalUrl *url.URL, err error) {
	params.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	params.Add("AWSAccessKeyId", m.accessKey)
	params.Add("Action", ep.Action)
	params.Add("SellerId", m.sellerId)
	params.Add("SignatureMethod", "HmacSHA256")
	params.Add("SignatureVersion", "2")
	params.Add("Version", ep.Date)
	params.Add("MarketplaceId.Id.1", m.marketplaceId)
	if m.authToken != "" {
		params.Add("MWSAuthToken", m.authToken)
	}
	if params.Get("MarketplaceId") != "" {
		params.Del("MarketplaceId.Id.1")
	}
	if params.Get("MarketplaceIdList.Id.1") != "" {
		params.Del("MarketplaceId.Id.1")
	}

	result, err := url.Parse(MarketplaceIdUrls[m.marketplaceId])
	if err != nil {
		return
	}

	result.Host = MarketplaceIdUrls[m.marketplaceId]
	result.Scheme = "https"
	result.Path = ep.Path
	result.RawQuery = params.Encode()
	return result, nil
}

func (m *mwsClient) SetClient(ctx context.Context, authToken, accessKey, secretKey, sellerId, marketplaceId string) MwsClient {
	m.authToken = authToken
	m.accessKey = accessKey
	m.secretKey = secretKey
	m.sellerId = sellerId
	m.marketplaceId = marketplaceId
	return m
}

func (m *mwsClient) GetFeedSubmissionResult(ctx context.Context, feedSubmissionId string) (err error) {
	defer func(begin time.Time) {
		_ = m.logger.Log(
			m.traceId, ctx.Value(m.traceId),
			"method", "GetFeedSubmissionResult",
			"feedSubmissionId", feedSubmissionId,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	params := url.Values{}
	params.Add("FeedSubmissionId", feedSubmissionId)

	ep := endpoints["GetFeedSubmissionResult"]

	var resp ListOrdersResponse

	err = m.request(params, ep, &resp)
	if err != nil {
		return
	}

	return
}

func (m *mwsClient) signAmazonUrl(origUrl *url.URL, method string) (signedUrl string, err error) {
	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	//q, _ := url.QueryUnescape(origUrl.RawQuery)
	params := strings.Split(escapeUrl, "&")
	sort.StringsAreSorted(params)

	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("%s\n%s\n%s\n%s", method, origUrl.Host, origUrl.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(m.secretKey))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	origUrl.RawQuery = fmt.Sprintf("%s&Signature=%s", sortedParams, hash)

	return origUrl.String(), nil
}

func (m *mwsClient) request(params url.Values, ep RequestMethod, data interface{}) (err error) {
	genUrl, err := m.genAmazonUrl(ep, params)
	if err != nil {
		return
	}

	signed, err := m.signAmazonUrl(genUrl, ep.Method)
	if err != nil {
		return
	}

	u, err := url.Parse(signed)
	if err != nil {
		return
	}

	req := request.NewRequest(fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path), ep.Method)
	for k := range u.Query() {
		req = req.Param(k, u.Query().Get(k))
	}

	b, err := req.
		Header("Accept", "application/xml").
		Header("x-amazon-user-agent", "MCS/MwsClient/0.0.*").
		Do().Raw()
	if err != nil {
		var errResponse ErrorResponse
		if e := xml.NewDecoder(bytes.NewBuffer(b)).Decode(&errResponse); e != nil {
			err = errors.Wrap(err, err.Error())
		}
		err = errors.Wrap(err, fmt.Sprintf("code: %s, message: %s", errResponse.Error.Code, errResponse.Error.Message))
		return err
	}
	//if e := xml.NewDecoder(bytes.NewBuffer(b)).Decode(data); e != nil {
	if e := xml.Unmarshal(b, data); e != nil {
		err = errors.Wrap(e, "xml.NewDecode.Decode")
		return err
	}
	return nil
}

func NewMwsClient(logger log.Logger, traceId, host, authToken, accessKey, secretKey, sellerId, marketplaceId string) MwsClient {
	return &mwsClient{logger: logger, traceId: traceId, host: "mws.amazonservices.com"}
}
