// Copyright 2020 FoxyUtils ehf. All rights reserved.
package main

import (
	"fmt"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/presentation"
)

const licenseKey = `
-----BEGIN UNIDOC LICENSE KEY-----
Free trial license keys are available at: https://unidoc.io/
-----END UNIDOC LICENSE KEY-----
`

func init() {
	err := license.SetLicenseKey(licenseKey, `Company Name`)
	if err != nil {
		panic(err)
	}
}

func main() {
	ppt, err := presentation.Open("source.pptx")
	if err != nil {
		panic(err)
	}
	defer ppt.Close()
	slide := ppt.Slides()[0] // taking the first slide

	// Getting the list of text boxes
	tbs := slide.GetTextBoxes() // getting all textboxes

	for _, tb := range tbs {
		for _, p := range tb.X().TxBody.P {
			for _, tr := range p.EG_TextRun {
				fmt.Println(tr.R.T)
			}
		}
	}

	// Editing the existing text box
	tb := tbs[0] // taking first of them
	run := tb.X().TxBody.P[0].EG_TextRun[0].R // taking the first run of the first paragraph
	run.T = "Edited TextBox text" // changing the text of the run

	// creating a new text box
	newTb := slide.AddTextBox()
	newTb.SetOffsetX(measurement.Inch * 5)
	newTb.SetOffsetY(measurement.Inch * 4)

	newPara := newTb.AddParagraph()
	newRun := newPara.AddRun()
	newRun.SetText("New TextBox text")

	ppt.SaveToFile("mod.pptx")
}
