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

func MappingSubCategory(code string) string {
	params := map[string]string{
		"":   "2",  //default
		"SI": "63", //slipper
		"PD": "78", //pad
		"M1": "65", //mat
		"17": "68", //box
		"BG": "68", //bag
		"8B": "71", //basket
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
	if lang != "TH" {
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
	if lang != "TH" {
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
        "value": "` + ReplaceAllSpecialChar(product.DIMENSION) + `"
      },
      {
        "attribute_code": "service_material2_th",
        "value": "` + product.MATERIAL2_TH + `"
      }
    `
}

func ReplaceAllSpecialChar(body string) string {
	input := strings.TrimSpace(body)
	// new
	// Replace \r\n with <br/>
	input = strings.ReplaceAll(input, "\r\n", "<br/>")
	// Replace \n with <br/>
	input = strings.ReplaceAll(input, "\n", "<br/>")
	// Replace \t with ' '
	input = strings.ReplaceAll(input, "\t", " ")

	// old
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
