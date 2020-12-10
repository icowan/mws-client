/**
 * @Time: 2020/9/26 22:10
 * @Author: solacowa@gmail.com
 * @File: mwsclient_test
 * @Software: GoLand
 */

package mwsclient

import (
	"context"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/term"
	"os"
	"testing"
	"time"
)

var (
	logger = term.NewLogger(os.Stdout, kitlog.NewLogfmtLogger, colorFunc())
)

func colorFunc() func(keyvals ...interface{}) term.FgBgColor {
	return func(keyvals ...interface{}) term.FgBgColor {
		for i := 0; i < len(keyvals)-1; i += 2 {
			if keyvals[i] != "level" {
				continue
			}
			val := fmt.Sprintf("%v", keyvals[i+1])
			switch val {
			case "debug":
				return term.FgBgColor{Fg: term.DarkGray}
			case "info":
				return term.FgBgColor{Fg: term.Blue}
			case "warn":
				return term.FgBgColor{Fg: term.Yellow}
			case "error":
				return term.FgBgColor{Fg: term.Red}
			case "crit":
				return term.FgBgColor{Fg: term.Gray, Bg: term.DarkRed}
			default:
				return term.FgBgColor{}
			}
		}
		return term.FgBgColor{}
	}
}

func getClient() MwsClient {
	logger = kitlog.WithPrefix(logger, "ts", kitlog.TimestampFormat(func() time.Time {
		return time.Now()
	}, "2006-01-02 15:04:05"))
	return NewMwsClient(logger, "tract-id", "", "", "", "", "", "")
}

func TestMwsClient_GetMatchingProductForId(t *testing.T) {
	client := getClient()
	res, err := client.SetClient(context.Background(),
		"amzn.mws.7.......",
		"AKIAJKDL.........",
		"jZ53cYxd0/............",
		"A2DEW.........",
		"A1PA6795UKMFR9",
	).GetMatchingProductForId(context.Background(), "ASIN", []string{"B082WFZR5M", "B07SCJ6KZK"})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestMwsClient_ListOrders(t *testing.T) {
	client := getClient()
	fromDate := time.Now().AddDate(0, 0, -7) // 7天前
	res, err := client.SetClient(context.Background(),
		"amzn.mws.7.......",
		"AKIAJKDL.........",
		"jZ53cYxd0/............",
		"A2DEW.........",
		"A1PA6795UKMFR9",
	).ListOrders(context.Background(), fromDate, false, []OrderState{
		OrderStateUnshipped,
		OrderStatePartiallyShipped,
		OrderStateShipped,
		OrderStatePendingAvailability,
		OrderStatePending,
		OrderStateInvoiceUnconfirmed,
		OrderStateCanceled,
		OrderStateUnfulfillable,
	}, []OrderChannel{
		OrderChannelAFN,
		OrderChannelMFN,
	}, nil, nil)
	if err != nil {
		t.Error(err)
	}

	for _, v := range res {
		t.Log(v.AmazonOrderId)
	}
}

func TestMwsClient_GetFeedSubmissionResult(t *testing.T) {
	client := getClient()
	err := client.SetClient(context.Background(),
		"amzn.mws.7.......",
		"AKIAJKDL.........",
		"jZ53cYxd0/............",
		"A2DEW.........",
		"A1PA6795UKMFR9",
	).GetFeedSubmissionResult(context.Background(), "50006018604")
	if err != nil {
		t.Error(err)
	}
}
