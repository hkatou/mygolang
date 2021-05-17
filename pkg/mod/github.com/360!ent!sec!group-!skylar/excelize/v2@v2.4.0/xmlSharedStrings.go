// Copyright 2016 - 2021 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to
// and read from XLSX / XLSM / XLTM files. Supports reading and writing
// spreadsheet documents generated by Microsoft Excel™ 2007 and later. Supports
// complex components by high compatibility, and provided streaming API for
// generating or reading data from a worksheet with huge amounts of data. This
// library needs Go version 1.15 or later.

package excelize

import (
	"encoding/xml"
	"strings"
)

// xlsxSST directly maps the sst element from the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main. String values may
// be stored directly inside spreadsheet cell elements; however, storing the
// same value inside multiple cell elements can result in very large worksheet
// Parts, possibly resulting in performance degradation. The Shared String Table
// is an indexed list of string values, shared across the workbook, which allows
// implementations to store values only once.
type xlsxSST struct {
	XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main sst"`
	Count       int      `xml:"count,attr"`
	UniqueCount int      `xml:"uniqueCount,attr"`
	SI          []xlsxSI `xml:"si"`
}

// xlsxSI (String Item) is the representation of an individual string in the
// Shared String table. If the string is just a simple string with formatting
// applied at the cell level, then the String Item (si) should contain a
// single text element used to express the string. However, if the string in
// the cell is more complex - i.e., has formatting applied at the character
// level - then the string item shall consist of multiple rich text runs which
// collectively are used to express the string.
type xlsxSI struct {
	T          *xlsxT             `xml:"t,omitempty"`
	R          []xlsxR            `xml:"r"`
	RPh        []*xlsxPhoneticRun `xml:"rPh"`
	PhoneticPr *xlsxPhoneticPr    `xml:"phoneticPr"`
}

// String extracts characters from a string item.
func (x xlsxSI) String() string {
	if len(x.R) > 0 {
		var rows strings.Builder
		for _, s := range x.R {
			if s.T != nil {
				rows.WriteString(s.T.Val)
			}
		}
		return rows.String()
	}
	if x.T != nil {
		return x.T.Val
	}
	return ""
}

// xlsxR represents a run of rich text. A rich text run is a region of text
// that share a common set of properties, such as formatting properties. The
// properties are defined in the rPr element, and the text displayed to the
// user is defined in the Text (t) element.
type xlsxR struct {
	RPr *xlsxRPr `xml:"rPr"`
	T   *xlsxT   `xml:"t"`
}

// xlsxT directly maps the t element in the run properties.
type xlsxT struct {
	XMLName xml.Name `xml:"t"`
	Space   xml.Attr `xml:"space,attr,omitempty"`
	Val     string   `xml:",chardata"`
}

// xlsxRPr (Run Properties) specifies a set of run properties which shall be
// applied to the contents of the parent run after all style formatting has been
// applied to the text. These properties are defined as direct formatting, since
// they are directly applied to the run and supersede any formatting from
// styles.
type xlsxRPr struct {
	RFont     *attrValString `xml:"rFont"`
	Charset   *attrValInt    `xml:"charset"`
	Family    *attrValInt    `xml:"family"`
	B         *string        `xml:"b"`
	I         *string        `xml:"i"`
	Strike    *string        `xml:"strike"`
	Outline   *string        `xml:"outline"`
	Shadow    *string        `xml:"shadow"`
	Condense  *string        `xml:"condense"`
	Extend    *string        `xml:"extend"`
	Color     *xlsxColor     `xml:"color"`
	Sz        *attrValFloat  `xml:"sz"`
	U         *attrValString `xml:"u"`
	VertAlign *attrValString `xml:"vertAlign"`
	Scheme    *attrValString `xml:"scheme"`
}

// RichTextRun directly maps the settings of the rich text run.
type RichTextRun struct {
	Font *Font
	Text string
}
