package indexcreation

import (
	"encoding/json"
	"flag"
	//"fmt"
	//"strconv"
	"SearchBenchmarking/searchquery"
	"github.com/blevesearch/bleve"
	//"github.com/fatih/structs"
	//"github.com/blevesearch/bleve/index"
	//"github.com/blevesearch/bleve/search/query"
	"SearchBenchmarking/datamodel"
	//"github.com/blevesearch/bleve/search/searcher"
	//"reflect"
	//"github.com/blevesearch/bleve/analysis/analyzer/simple"
	"io/ioutil"
	"log"
	//	"os"
	"path/filepath"
	//	"runtime"
	//	"runtime/pprof"
	"time"
	//"regexp"
)

var jsonDir = flag.String("jsonDir", "data/", "json directory")
var indexPath = flag.String("index", "example.bleve", "index path")

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var batchSize = flag.Int("batchSize", 100, "batch size for indexing")

//var memprofile = flag.String("memprofile", "", "write mem profile to file")

//func main() {

//}
func QueryResponseTestFunction() {
	beerIndex, _ := Indexing()
	//	fmt.Println("err  during indexing is ", err)
	//fmt.Println("inside query response......")
	pmp := model.Pmp{
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
		EnabledOnDesktop:    "0",
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
	startTime := time.Now()
	searchquery.QueryResult(pmp, beerIndex)
	endTime := time.Now()
	log.Printf("time taken for searching query in milliseconds: %f", float64((endTime.Nanosecond()-startTime.Nanosecond()))/1000000)

}
func Indexing() (bleve.Index, error) {
	fmt.Println("inside indexing......")
	//fmt.Println("w")
	//flag.Parse()

	//log.Printf("GOMAXPROCS: %d", runtime.GOMAXPROCS(-1))

	beerIndex, err := bleve.Open(*indexPath)
	log.Printf("opening index file.... ")
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf("Creating new index...")
		// create a mapping
		indexMapping := bleve.NewIndexMapping()
		//fmt.Println(indexMapping)
		beerIndex, err = bleve.New(*indexPath, indexMapping)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("indexing documents...")
		err = IndexingFile(beerIndex)
		if err != nil {
			log.Fatal(err)
		}
		//		pprof.StopCPUProfile()
		//		if *memprofile != "" {
		//			f, err := os.Create(*memprofile)
		//			if err != nil {
		//				log.Fatal(err)
		//			}
		//			pprof.WriteHeapProfile(f)
		//			f.Close()
		//		}

	}

	//	chquery := bleve.NewTermQuery("USA")
	//	chquery.SetField("country")
	//	bhquery := bleve.NewTermQuery("-1")
	//	bhquery.SetField("wifitargeting")
	//	//query = bleve.NewDisjunctionQuery([]query.Query{chquery,bhquery})
	//	query := query.NewBooleanQuery([]query.Query{chquery}, []query.Query{bhquery}, []query.Query{bhquery})
	//	searchRequest := bleve.NewSearchRequest(query)
	//	searchResults, err := beerIndex.Search(searchRequest)
	//	fmt.Println("search results are : ", searchResults)
	fmt.Println("exiting indexing......")

	return beerIndex, err
}
func IndexingFile(i bleve.Index) error {
	log.Printf("opening the directory")
	// open the directory
	dirEntries, err := ioutil.ReadDir(*jsonDir)
	if err != nil {
		return err
	}

	// walk the directory entries for indexing
	log.Printf("Indexing documents...")
	count := 0
	startTime := time.Now()
	batch := i.NewBatch()
	batchCount := 0
	for _, dirEntry := range dirEntries {
		filename := dirEntry.Name()
		// read the bytes
		fmt.Println("dir is", *jsonDir+"/"+filename)
		jsonBytes, err := ioutil.ReadFile(*jsonDir + "/" + filename)
		if err != nil {
			//	fmt.Println("dir is", *jsonDir+"/"+filename)
			return err
		}
		// parse bytes as json
		var jsonDoc interface{}
		err = json.Unmarshal(jsonBytes, &jsonDoc)
		if err != nil {
			return err
		}
		ext := filepath.Ext(filename)
		docID := filename[:(len(filename) - len(ext))]
		batch.Index(docID, jsonDoc)
		batchCount++

		if batchCount >= *batchSize {
			err = i.Batch(batch)
			if err != nil {
				return err
			}
			batch = i.NewBatch()
			batchCount = 0
		}
		count++
		if count%1000 == 0 {
			indexDuration := time.Since(startTime)
			indexDurationSeconds := float64(indexDuration) / float64(time.Second)
			timePerDoc := float64(indexDuration) / float64(count)
			log.Printf("Indexed %d documents, in %.2fs (average %.2fms/doc)", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
		}
	}
	// flush the last batch
	if batchCount > 0 {
		err = i.Batch(batch)
		if err != nil {
			log.Fatal(err)
		}
	}
	indexDuration := time.Since(startTime)
	indexDurationSeconds := float64(indexDuration) / float64(time.Second)
	timePerDoc := float64(indexDuration) / float64(count)
	log.Printf("Indexed %d documents, in %.2fs (average %.2fms/doc)", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
	return nil
}
