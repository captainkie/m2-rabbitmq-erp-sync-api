package helpers

import (
	"strings"

	"github.com/captainkie/websync-api/types/request"
)

// Categories ID
// Default: 2 => Default Category
// Find in Shop : 82 => required all product
// Bed Linen: 41 => BD
// White Filling: 42 => SL
// Bath Linen: 43 => JB
// Art & Decor: 44 => BR,DC,DN,TA
// Floor Covering: 45 => RU
// Ethan Allen: 47 => JH
// Mattress & Bed Ensemles: 48 => MU,S7,BX,ST,HB

var AttributeSetID int = 16
var DefaultWeight int = 0
var WebsiteID int = 1

func MappingCategory(brand, code string) string {
	// Ethan Allen
	if brand == "JH" {
		return "47"
	} else {
		// Santas
		params := map[string]string{
			"":   "2",
			"BR": "44",
			"DC": "44",
			"DN": "44",
			"TA": "44",
			"BD": "41",
			"JB": "43",
			"MU": "48",
			"S7": "48",
			"BX": "48",
			"ST": "48",
			"HB": "48",
			"RU": "45",
			"SL": "42",
		}

		if params[code] == "" {
			return "2"
		} else {
			return params[code]
		}
	}
}

// func MappingSubCategory(code string) string {
// 	params := map[string]string{
// 		"": "2",
// 	}

// 	if params[code] == "" {
// 		return "2"
// 	} else {
// 		return params[code]
// 	}
// }

func MappingSubCategory(code string) string {
	params := map[string]string{
		"":   "2",  //default
		"FI": "49", //fitted sheet
		"FA": "49", //flat sheet
		"QC": "51", //duvet cover
		"PC": "52", //pillow case
		"BC": "53", //bolster case
		"ES": "54", //euro sham
		"CC": "54", //cushion case
		"BP": "54", //breakfast pillowcase
		"PT": "54", //neckroll cover
		"B6": "54", //boudoir pillowcase
		"TR": "55", //throw
		"TE": "55", //throw blanket
		"11": "55", //blanket
		"BE": "55", //bed spread
		"PA": "56", //pajama
		"NA": "57", //bundle set
		"UN": "58", //surprising unbox
		"TW": "59", //towel
		"BA": "61", //bath robe
		"BZ": "62", //bathmat
		"SI": "63", //slipper
		"Z1": "64", //shoulder bag
		"PW": "72", //pillow
		"CS": "73", //cushion
		"PB": "73", //baby pillow
		"PO": "73", //neckroll
		"B5": "73", //boudoir pillow
		"BL": "74", //bolster
		"DU": "75", //duvet
		"OM": "76", //over mattress
		"FR": "77", //feather bed
		"PD": "78", //pad
		"F5": "65", //floor mat
		"M1": "65", //mat
		"JR": "66", //jute rug
		"RU": "66", //rug
		"RG": "66", //bathroom rug set
		"LM": "67", //lamp
		"AS": "68", //ashtray
		"RN": "68", //runner
		"X1": "68", //cool pack case
		"17": "68", //box
		"8T": "68", //topper
		"AD": "68", //candle stand
		"BO": "68", //bowl
		"O2": "68", //oil burner
		"WT": "68", //wood tray
		"C8": "68", //cake stand
		"D6": "68", //diffuser
		"HH": "68", //box heart hammerd
		"M2": "68", //magazine rack
		"VA": "68", //vase
		"X2": "68", //hot pack case
		"0T": "68", //trivet
		"0Z": "68", //paper bin
		"1A": "68", //trencher
		"20": "68", //candle
		"2I": "68", //tea set
		"2T": "68", //tea bag squeezer
		"77": "68", //mirror
		"D3": "68", //decorative cutlery
		"DI": "68", //dish
		"JB": "68", //jewels box
		"UM": "68", //umbrella
		"W1": "68", //wine accessaries
		"2R": "68", //tray
		"CG": "68", //cup
		"CL": "68", //clock
		"F8": "68", //figuring
		"GL": "68", //glass
		"TU": "68", //tumbler
		"7C": "68", //chess
		"91": "68", //oval platter
		"AH": "68", //candle holder
		"BG": "68", //bag
		"BJ": "68", //bookmark
		"GB": "68", //ballerina
		"JU": "68", //jug
		"MK": "68", //mask
		"PZ": "68", //placemat
		"S9": "68", //sauce stand
		"IB": "68", //ice bucket
		"IN": "68", //salt/pepper
		"PL": "68", //plate
		"RH": "68", //rack
		"30": "68", //coasters
		"J1": "68", //breaded basket
		"NK": "68", //napkin
		"TS": "68", //tablecloth
		"P9": "69", //picture
		"PH": "70", //photoframe
		"8B": "71", //basket
		"SE": "71", //soap dispenser
		"TT": "71", //toothbrush holder
		"0J": "71", //soap disp with scrubber holder
		"0I": "71", //scrubber holder
		"TA": "71", //towel rack
		"SD": "71", //soap dish
	}

	if params[code] == "" {
		return "2"
	} else {
		return params[code]
	}
}

func MappingFindInShopCategory() string {
	return "82"
}

func MappingCustomAttr(product request.AddUpdateProductRequest, attributes, lang, token string) string {
	if lang == "EN" {
		return `[
      {
        "attribute_code": "short_description",
        "value": "` + ReplaceAllSpecialChar(product.SHORT_DESC_EN) + `"
      },
      {
        "attribute_code": "description",
        "value": "` + ReplaceAllSpecialChar(product.DESC_EN) + `"
      },` +
			attributes +
			MappingServiceAttr(product) +
			`]`
	} else {
		return `[
      {
        "attribute_code": "short_description",
        "value": "` + ReplaceAllSpecialChar(product.SHORT_DESC_TH) + `"
      },
      {
        "attribute_code": "description",
        "value": "` + ReplaceAllSpecialChar(product.DESC_TH) + `"
      },` +
			attributes +
			MappingServiceAttr(product) +
			`]`
	}
}

func MappingConfigurableCustomAttr(product request.AddUpdateProductRequest, attributes, lang, token string) string {
	if lang == "EN" {
		return `[
      {
        "attribute_code": "short_description",
        "value": "` + ReplaceAllSpecialChar(product.SHORT_DESC_EN) + `"
      },
      {
        "attribute_code": "description",
        "value": "` + ReplaceAllSpecialChar(product.DESC_EN) + `"
      },` +
			attributes +
			`]`
	} else {
		return `[
      {
        "attribute_code": "short_description",
        "value": "` + ReplaceAllSpecialChar(product.SHORT_DESC_TH) + `"
      },
      {
        "attribute_code": "description",
        "value": "` + ReplaceAllSpecialChar(product.DESC_TH) + `"
      },` +
			attributes +
			`]`
	}
}

func MappingServiceAttr(product request.AddUpdateProductRequest) string {
	return `
      {
        "attribute_code": "service_prod_tname",
        "value": "` + ReplaceAllSpecialChar(product.PROD_TNAME) + `"
      },
      {
        "attribute_code": "service_prod_ename",
        "value": "` + ReplaceAllSpecialChar(product.PROD_ENAME) + `"
      },
      {
        "attribute_code": "service_uom_code",
        "value": "` + product.UOM_CODE + `"
      },
      {
        "attribute_code": "service_bar_code",
        "value": "` + product.BAR_CODE + `"
      },
      {
        "attribute_code": "service_pdgrp_code",
        "value": "` + product.PDGRP_CODE + `"
      },
      {
        "attribute_code": "service_pdbrnd_code",
        "value": "` + product.PDBRND_CODE + `"
      },
      {
        "attribute_code": "service_pdtype_code",
        "value": "` + product.PDTYPE_CODE + `"
      },
      {
        "attribute_code": "service_pddsgn_code",
        "value": "` + product.PDDSGN_CODE + `"
      },
      {
        "attribute_code": "service_pdsize_code",
        "value": "` + product.PDSIZE_CODE + `"
      },
      {
        "attribute_code": "service_pdcolor_code",
        "value": "` + product.PDCOLOR_CODE + `"
      },
      {
        "attribute_code": "service_pdmisc_code",
        "value": "` + product.PDMISC_CODE + `"
      },
      {
        "attribute_code": "service_pdbrnd_desc",
        "value": "` + product.PDBRND_DESC + `"
      },
      {
        "attribute_code": "service_pddsgn_desc",
        "value": "` + product.PDDSGN_DESC + `"
      },
      {
        "attribute_code": "service_pdcolor_desc",
        "value": "` + product.PDCOLOR_DESC + `"
      },
      {
        "attribute_code": "service_pdmisc_desc",
        "value": "` + product.PDMISC_DESC + `"
      },
      {
        "attribute_code": "service_ref_no",
        "value": "` + product.REF_NO + `"
      },
      {
        "attribute_code": "service_prod_st",
        "value": "` + product.PROD_ST + `"
      },
      {
        "attribute_code": "service_prod_class",
        "value": "` + product.PROD_CLASS + `"
      },
      {
        "attribute_code": "service_prod_line",
        "value": "` + product.PROD_LINE + `"
      },
      {
        "attribute_code": "service_prod_type",
        "value": "` + product.PROD_TYPE + `"
      },
      {
        "attribute_code": "service_pdmodel_code",
        "value": "` + product.PDMODEL_CODE + `"
      },
      {
        "attribute_code": "service_color2",
        "value": "` + product.COLOR2 + `"
      },
      {
        "attribute_code": "service_material2",
        "value": "` + product.MATERIAL2 + `"
      },
      {
        "attribute_code": "service_dimension",
        "value": "` + product.DIMENSION + `"
      },
      {
        "attribute_code": "service_material2_th",
        "value": "` + product.MATERIAL2_TH + `"
      }
    `
}

func ReplaceAllSpecialChar(body string) string {
	input := strings.TrimSpace(body)
	// Replace \r with \\r
	input = strings.ReplaceAll(input, "\r", "\\r")
	// Replace \n with \\n
	input = strings.ReplaceAll(input, "\n", "\\n")
	// Replace \t with \\t
	input = strings.ReplaceAll(input, "\t", "\\t")
	// Replace \t with \\t
	input = strings.ReplaceAll(input, "\"", "&quot;")

	return input
}

func ReplaceAllQuot(body string) string {
	input := body
	// Replace \r with \\r
	input = strings.ReplaceAll(strings.TrimSpace(body), `"`, "")

	return input
}

func PadString(input string, length int, padChar rune) string {
	if len(input) >= length {
		return input[:length]
	}
	padding := make([]rune, length-len(input))
	for i := range padding {
		padding[i] = padChar
	}
	return input + string(padding)
}
