package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

/*************************************************************************

Parser tool to extract data from the GHC Rate Engine XLSX

For help run: <program> -h

`go run cmd/parse_ratengine_data_ghc/*.go -h`

Rate Engine XLSX sections this tool will be parsing:

1) 1b) Service Areas

2) Domestic Price Tabs
        2a) Domestic Linehaul Prices
	    2b) Domestic Service Area Prices
	    2c) Other Domestic Prices

3) International Price Tabs
        3a) OCONUS to OCONUS Prices
	    3b) CONUS to OCONUS Prices
	    3c) OCONUS to CONUS Prices
	    3d) Other International Prices
	    3e) Non-Standard Loc'n Prices

4) Mgmt., Coun., Trans. Prices Tab
        4a) Mgmt., Coun., Trans. Prices

5) Other Prices Tabs
        5a) Access. and Add. Prices
	    5b) Price Escalation Discount


--------------------------------------------------------------------------

Rate Engine XLSX sheet tabs:

0: 	Guide to Pricing Rate Table
1: 	Total Evaluated Price
2: 	Submission Checklist
3: 	1a) Directions
4: 	1b) Service Areas
5: 	Domestic Price Tabs >>>
6: 	2a) Domestic Linehaul Prices
7: 	2b) Dom. Service Area Prices
8: 	2c) Other Domestic Prices
9: 	International Prices Tables >>>
10: 3a) OCONUS to OCONUS Prices
11: 3b) CONUS to OCONUS Prices
12: 3c) OCONUS to CONUS Prices
13: 3d) Other International Prices
14: 3e) Non-Standard Loc'n Prices
15:	Other Price Tables
16: 4a) Mgmt., Coun., Trans. Prices
17: 5a) Access. and Add. Prices
18: 5b) Price Escalation Discount
19: Domestic Linehaul Data
20: Domestic Move Count
21: Domestic Avg Weight
22: Domestic Avg Milage
23: Domestic Price Calculation >>>
24: Domestic Linehaul Calculation
25: Domestic SA Price Calculation
26: NTS Packing Calculation
27: Int'l Price Calculation >>>
28: OCONUS to OCONUS Calculation
29: CONUS to OCONUS Calculation
30: OCONUS to CONUS Calculation
31: Other Int'l Prices Calculation
32: Non-Standard Loc'n Calculation
33: Other Calculations >>>
34: Mgmt., Coun., Trans., Calc
35: Access. and Add. Calculation


 *************************************************************************/

/*************************************************************************

To add new parser functions to this file:

	a.) (optional) Add new verify function for your processing must match signature verifyXlsxSheet
	b.) Add new process function to process XLSX data sheet must match signature processXlsxSheet
	c.) Update initDataSheetInfo() with a.) and b.)
		The index must match the sheet index in the XLSX that you aim to process

You should not have to update the main() or  process() functions. Unless you
intentionally are modifying the pattern of how the processing functions are called.

 *************************************************************************/
const sharedNumEscalationYearsToProcess int = 1
const xlsxSheetsCountMax int = 35

type processXlsxSheet func(paramConfig, int) error
type verifyXlsxSheet func(paramConfig, int) error

type xlsxDataSheetInfo struct {
	description    *string
	process        *processXlsxSheet
	verify         *verifyXlsxSheet
	outputFilename *string //do not include suffix see func generateOutputFilename for details
}

var xlsxDataSheets []xlsxDataSheetInfo

// initDataSheetInfo: When adding new functions for parsing sheets, must add new xlsxDataSheetInfo
// defining the parse function
//
// The index MUST match the sheet that is being processed. Refer to file comments or XLSX to
// determine the correct index to add.
func initDataSheetInfo() {
	xlsxDataSheets = make([]xlsxDataSheetInfo, xlsxSheetsCountMax, xlsxSheetsCountMax)

	// 6: 	2a) Domestic Linehaul Prices
	xlsxDataSheets[6] = xlsxDataSheetInfo{
		description:    stringPointer("2a) Domestic Linehaul Prices"),
		outputFilename: stringPointer("2a_domestic_linehaul_prices"),
		process:        &parseDomesticLinehaulPrices,
		verify:         &verifyDomesticLinehaulPrices,
	}

	// 7: 	2b) Dom. Service Area Prices
	xlsxDataSheets[7] = xlsxDataSheetInfo{
		description:    stringPointer("2b) Dom. Service Area Prices"),
		outputFilename: stringPointer("2b_domestic_service_area_prices"),
		process:        &parseDomesticServiceAreaPrices,
		verify:         &verifyDomesticServiceAreaPrices,
	}

}

type paramConfig struct {
	processAll   bool
	showOutput   bool
	xlsxFilename *string
	xlsxSheets   []string
	saveToFile   bool
	runTime      time.Time
	xlsxFile     *xlsx.File
	runVerify    bool
}

func xlsxSheetsUsage() string {
	message := "Provide comma separated string of sequential sheet index numbers starting with 0:\n"
	message += "\t e.g. '-xlsxSheets=\"6,7,11\"'\n"
	message += "\t      '-xlsxSheets=\"6\"'\n"
	message += "\n"
	message += "Available sheets for parsing are: \n"

	for i, v := range xlsxDataSheets {
		if v.process != nil {
			description := ""
			if v.description != nil {
				description = *v.description
			}
			message += fmt.Sprintf("%d:  %s\n", i, description)
		}
	}

	return message
}

func main() {
	initDataSheetInfo()
	params := paramConfig{}
	params.runTime = time.Now()

	filename := flag.String("filename", "", "Filename including path of the XLSX to parse for Rate Engine GHC import")
	all := flag.Bool("all", true, "Parse entire Rate Engine GHC XLSX")
	sheets := flag.String("xlsxSheets", "", xlsxSheetsUsage())
	display := flag.Bool("display", false, "Display output of parsed info")
	saveToFile := flag.Bool("save", false, "Save output to CSV file")
	runVerify := flag.Bool("verify", true, "Default is true, if false skip sheet format verification")

	flag.Parse()

	// Process command line params

	params.processAll = false
	if all != nil && *all == true {
		params.processAll = true
	}

	// option `xlsxSheets` will override `all` flag
	if sheets != nil && len(*sheets) > 0 {
		// If processes based on `xlsxSheets` indices provided as arguments
		// process those and do not run all
		params.processAll = false
		params.xlsxSheets = strings.Split(*sheets, ",")
	}

	params.xlsxFilename = filename
	if filename != nil {
		log.Printf("Importing file %s\n", *filename)
	} else {
		log.Fatalf("Did not receive an XLSX filename to parse, missing -filename\n")
	}

	xlsxFile, err := xlsx.OpenFile(*params.xlsxFilename)
	params.xlsxFile = xlsxFile
	if err != nil {
		log.Fatalf("Failed to open file %s with error %v\n", *params.xlsxFilename, err)
	}

	params.showOutput = false
	if display != nil && *display == true {
		params.showOutput = true
	}

	params.saveToFile = false
	if saveToFile != nil && *saveToFile == true {
		params.saveToFile = true
	}

	params.runVerify = false
	if runVerify != nil {
		params.runVerify = *runVerify
	}

	// Must be after processing config param
	// Run the process function

	if params.processAll == true {
		for i, x := range xlsxDataSheets {
			if x.process != nil {
				err := process(params, i)
				if err != nil {
					log.Fatalf("Error processing xlsxDataSheets %v\n", err.Error())
				}
			}
		}
	} else {
		for _, v := range params.xlsxSheets {
			index, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("Bad xlsxSheets index provided %v\n", err)
			}
			if index < len(xlsxDataSheets) {
				err = process(params, index)
				if err != nil {
					log.Fatalf("Error processing %v\n", err)
				}
			} else {
				log.Fatalf("Error processing index %d, not in range of slice xlsxDataSheets\n", index)
			}
		}
	}
}

// process: is the main process function. It will call the
// appropriate verify and process functions based on what is defined
// in the xlsxDataSheets array
//
// Should not need to edit this function when adding new processing functions
//     to add new processing functions update:
//         a.) add new verify function for your processing
//         b.) add new process function for your processing
//         c.) update initDataSheetInfo() with a.) and b.)
func process(params paramConfig, sheetIndex int) error {
	xlsxInfo := xlsxDataSheets[sheetIndex]
	var description string
	if xlsxInfo.description != nil {
		description = *xlsxInfo.description
		log.Printf("Processing sheet index %d with description %s\n", sheetIndex, description)
	} else {
		log.Printf("Processing sheet index %d with missing description\n", sheetIndex)
	}

	// Call verify function
	if params.runVerify == true {
		if xlsxInfo.verify != nil {
			var callFunc verifyXlsxSheet
			callFunc = *xlsxInfo.verify
			err := callFunc(params, sheetIndex)
			if err != nil {
				log.Printf("%s verify error: %v\n", description, err)
				return errors.Wrapf(err, " verify error for sheet index: %d with description: %s", sheetIndex, description)
			}
		} else {
			log.Printf("No verify function for sheet index %d with description %s\n", sheetIndex, description)
		}
	} else {
		log.Print("Skip running the verify functions")
	}

	// Call process function
	if xlsxInfo.process != nil {
		var callFunc processXlsxSheet
		callFunc = *xlsxInfo.process
		err := callFunc(params, sheetIndex)
		if err != nil {
			log.Printf("%s process error: %v\n", description, err)
			return errors.Wrapf(err, " process error for sheet index: %d with description: %s", sheetIndex, description)
		}
	} else {
		log.Fatalf("Missing process function for sheet index %d with description %s\n", sheetIndex, description)
	}

	// Verification and Process completed
	log.Printf("Completed processing sheet index %d with description %s\n", sheetIndex, description)
	return nil
}