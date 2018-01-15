package searchquery

import (
	//"fmt"
	"github.com/blevesearch/bleve"

	//"github.com/blevesearch/bleve/search/query"
	//"strconv"
	"SearchBenchmarking/datamodel"
	//"log"
	//"github.com/fatih/structs"
	//"github.com/blevesearch/bleve/index"
	//	"flag"
	"reflect"
)

func QueryResult(pmp datamodel.Pmp, index bleve.Index) (*bleve.SearchResult, error) {

	//m := structs.Map(pmp)
	//fmt.Println("inside query result......")
	totalQuery := bleve.NewBooleanQuery()
	mapping := datamodel.GetMetaDataMapping()
	v := reflect.ValueOf(pmp)
	for i := 0; i < v.NumField(); i++ {

		jsonValue := mapping[v.Type().Field(i).Name]

		val := v.Field(i).Interface().(string)
		if len(val) > 0 && val != "nil" {
			//			fmt.Print("keys is : ", jsonValue)
			//			fmt.Print(" Value is : ", val)
			//			fmt.Println()
			matchQuery := bleve.NewMatchQuery(val)
			matchQuery.SetField(jsonValue)
			nilQuery := bleve.NewMatchQuery("nil")
			nilQuery.SetField(jsonValue)

			paramQuery := bleve.NewDisjunctionQuery(matchQuery, nilQuery)

			totalQuery.AddMust(paramQuery)

		}

	}

	searchRequest := bleve.NewSearchRequest(totalQuery)
	searchResults, err := index.Search(searchRequest)
	return searchResults, err
	//log.Printf("result is %s ", searchResults)
	//fmt.Println("search results are : ", searchResults)
	//fmt.Println("exiting query result......")
	//
	//	integrationMatchQuery := bleve.NewMatchQuery(pmp.Integration)
	//	integrationMatchQuery.SetField("integration")
	//	integrationNilQuery := bleve.NewMatchQuery("nil")
	//	integrationNilQuery.SetField("integration")
	//	integrationCompleteQuery := bleve.NewDisjunctionQuery(integrationMatchQuery, integrationNilQuery)
	//
	//	completeConjuctiveQuery := bleve.NewConjunctionQuery(countryCompleteQuery, adsizeCompleteQuery,
	//		domainBundleCompleteQuery, adFormatCompleteQuery, gpsCompleteQuery, integrationCompleteQuery)
	//
	//	searchRequest := bleve.NewSearchRequest(completeConjuctiveQuery)
	//
	//	searchResults, _ := index.Search(searchRequest)
	//	fmt.Println("search results for new match  query : ", searchResults)

}
