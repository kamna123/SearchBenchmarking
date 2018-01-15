package searchquery

import (
	//	"strings"
	"SearchBenchmarking/datamodel"
	"testing"
	//	"fmt"
	"flag"
	"github.com/blevesearch/bleve"
)

var indexPath = flag.String("index", "example.bleve", "index path")
var index, _ = bleve.Open(*indexPath)

func benchmarkQueryResult(b *testing.B, pmp datamodel.Pmp) {
	b.ReportAllocs()
	b.SetBytes(2)
	//fmt.Println("beer index path is ", *indexPath)

	//fmt.Println("inside query response......")
	//b.N=10
	for n := 0; n < b.N; n++ {
		//fmt.Println("n value is ",n)
		// for i :=0 ;i<50;i++{
		//	startTime := time.Now()
		QueryResult(pmp, index)
		//	endTime := time.Now()
		//	log.Printf("time taken for searching query in milliseconds: %f", float64((endTime.Nanosecond()-startTime.Nanosecond()))/1000000)
		//	fmt.Println("exiting query response ......")
	}
}

func BenchmarkQueryResult1(b *testing.B) {
	pmp := datamodel.Pmp{
		Country:             "USA",
		Adsize:              "250x300",
		DomainBundle:        "castcom.protectedio.app",
		AdFormats:           "maxvdo",
		Gps:                 "false",
		Ifa:                 "nil",
		Integration:         "Chocolate SSP",
		Clickable:           "false",
		Instream:            "true",
		Incentivized:        "nil",
		Autoplay:            "nil",
		Sound:               "true",
		Skippable:           "true",
		Viewability:         "nil",
		Moat:                "true",
		AdsizeDesk:          "small",
		AdunitDesk:          "Interstital",
		Apporwebtargetting:  "-1",
		Categories:          "Arts & Entertainment",
		CategoriesInclude:   "1",
		DomainBundleInclude: "1",
		AdultContent:        "1",
		DeviceOSinclude:     "1",
		DeviceOS:            "Android",
		DeviceID:            "IFA",
		DeviceCategory:      "-2",
		DeviceMake:          "1&1",
		DeviceModel:         "Android",
		DeviceOSVersion:     "8.3",
		CarrierWhitelist:    "& Aguiar Ltda Me",
		Ispwhitelist:        "nil",
		WifiTargeting:       "-1",
		Postalcode:          "151001",
		Dma:                 "nil",
		DpExpression:        "nil",
		EnabledOnDesktop:    "true",
		EnabledOnMobile:     "nil",
		EnabledOnTablet:     "nil",
		DesktopSelected:     "nil",
		MobSelected:         "nil",
		TabSelected:         "nil",
		DeviceTypeError:     "false",
		States:              "nil",
		Cities:              "nil",
		DesignId:            "nil",
		Site:                "nil",
	}
	benchmarkQueryResult(b, pmp)
}

func BenchmarkQueryResult2(b *testing.B) {
	pmp := datamodel.Pmp{
		Country:             "nil",
		Adsize:              "nil",
		DomainBundle:        "nil",
		AdFormats:           "nil",
		Gps:                 "nil",
		Ifa:                 "nil",
		Integration:         "nil",
		Clickable:           "nil",
		Instream:            "nil",
		Incentivized:        "nil",
		Autoplay:            "nil",
		Sound:               "nil",
		Skippable:           "nil",
		Viewability:         "nil",
		Moat:                "nil",
		AdsizeDesk:          "nil",
		AdunitDesk:          "nil",
		Apporwebtargetting:  "nil",
		Categories:          "nil",
		CategoriesInclude:   "nil",
		DomainBundleInclude: "nil",
		AdultContent:        "nil",
		DeviceOSinclude:     "nil",
		DeviceOS:            "nil",
		DeviceID:            "nil",
		DeviceCategory:      "nil",
		DeviceMake:          "nil",
		DeviceModel:         "nil",
		DeviceOSVersion:     "nil",
		CarrierWhitelist:    "nil",
		Ispwhitelist:        "nil",
		WifiTargeting:       "nil",
		Postalcode:          "nil",
		Dma:                 "nil",
		DpExpression:        "nil",
		EnabledOnDesktop:    "nil",
		EnabledOnMobile:     "nil",
		EnabledOnTablet:     "nil",
		DesktopSelected:     "nil",
		MobSelected:         "nil",
		TabSelected:         "nil",
		DeviceTypeError:     "false",
		States:              "nil",
		Cities:              "nil",
		DesignId:            "nil",
		Site:                "nil",
	}
	benchmarkQueryResult(b, pmp)
}
func BenchmarkQueryResult3(b *testing.B) {
	pmp := datamodel.Pmp{
		Country:             "nil",
		Adsize:              "nil",
		DomainBundle:        "nil",
		AdFormats:           "nil",
		Gps:                 "nil",
		Ifa:                 "nil",
		Integration:         "nil",
		Clickable:           "nil",
		Instream:            "nil",
		Incentivized:        "nil",
		Autoplay:            "nil",
		Sound:               "false",
		Skippable:           "nil",
		Viewability:         "nil",
		Moat:                "nil",
		AdsizeDesk:          "nil",
		AdunitDesk:          "nil",
		Apporwebtargetting:  "nil",
		Categories:          "nil",
		CategoriesInclude:   "nil",
		DomainBundleInclude: "nil",
		AdultContent:        "nil",
		DeviceOSinclude:     "nil",
		DeviceOS:            "nil",
		DeviceID:            "nil",
		DeviceCategory:      "nil",
		DeviceMake:          "nil",
		DeviceModel:         "nil",
		DeviceOSVersion:     "nil",
		CarrierWhitelist:    "nil",
		Ispwhitelist:        "nil",
		WifiTargeting:       "nil",
		Postalcode:          "nil",
		Dma:                 "nil",
		DpExpression:        "nil",
		EnabledOnDesktop:    "nil",
		EnabledOnMobile:     "1",
		EnabledOnTablet:     "1",
		DesktopSelected:     "nil",
		MobSelected:         "nil",
		TabSelected:         "nil",
		DeviceTypeError:     "false",
		States:              "nil",
		Cities:              "nil",
		DesignId:            "nil",
		Site:                "nil",
	}
	benchmarkQueryResult(b, pmp)
}
