package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "2025-06-27 06:46:20.210247 CEST"
	const __write__utc_time__ = "2025-06-27 04:46:20.210247 UTC"

	const __commitId__ = "0000000003"

	// Declaration of instances to stage

	__Chapter__000000_Getting_Started := (&models.Chapter{}).Stage(stage)
	__Chapter__000001_Advanced_Topics := (&models.Chapter{}).Stage(stage)

	__Content__000000_My_Awesome_Go_Book := (&models.Content{}).Stage(stage)

	__Page__000000_First_Steps := (&models.Page{}).Stage(stage)

	// Setup of values

	__Chapter__000000_Getting_Started.Name = `Getting Started`
	__Chapter__000000_Getting_Started.MardownContent = `This is the introduction to the first chapter.
We'll cover the basics here.

This demonstrates how new lines are generated`

	__Chapter__000001_Advanced_Topics.Name = `Advanced Topics`
	__Chapter__000001_Advanced_Topics.MardownContent = `Moving onto more advanced concepts in the second chapter.`

	__Content__000000_My_Awesome_Go_Book.Name = `My Awesome Go Book`
	__Content__000000_My_Awesome_Go_Book.MardownContent = `The Vera C. Rubin Observatory, formerly known as the Large Synoptic Survey Telescope (LSST), is an astronomical observatory in Chile. Its main task is an astronomical survey of the entire available southern sky every few nights, creating a time-lapse record over ten years, the Legacy Survey of Space and Time (also abbreviated LSST).[2][3][4] The observatory is located on the El Peñón peak of Cerro Pachón, a 2,682-meter-high (8,799 ft) mountain in Coquimbo Region, in northern Chile, alongside the existing Gemini South and Southern Astrophysical Research Telescopes.[5] The Rubin Observatory base facility is located about 100 kilometres (62 miles) away from the observatory by road, in the city of La Serena. The observatory is named for Vera Rubin, an American astronomer who pioneered discoveries about galactic rotation rates.

Vera C. Rubin Observatory is a joint initiative of the U.S. National Science Foundation (NSF) and the U.S. Department of Energy's Office of Science and is operated jointly by NSF NOIRLab and SLAC National Accelerator Laboratory.[6]

The Rubin Observatory houses the Simonyi Survey Telescope, a wide-field reflecting telescope with an 8.4-meter primary mirror that can photograph the entire available sky every few nights.[7] The telescope uses a variant of three-mirror anastigmat, which allows the compact telescope to deliver sharp images over a very wide 3.5-degree-diameter field of view. Images are recorded by a 3.2-gigapixel charge-coupled device imaging (CCD) camera, the largest digital camera ever constructed.[8]

The Rubin Observatory was proposed in 2001 as the LSST, and construction of the mirror began (with private funds) in 2007. The LSST then became the top-ranked large ground-based project in the 2010 Astrophysics Decadal Survey, and the project officially began construction on 1 August 2014, when the United States National Science Foundation (NSF) authorized the FY2014 portion ($27.5 million) of its construction budget.[9] Funding comes from the NSF, the United States Department of Energy, and private funding raised by the dedicated international non-profit organization, the LSST Discovery Alliance.[10] Operations are under the management of the Association of Universities for Research in Astronomy (AURA).[11] The total construction cost was expected to be about $680 million.[12]

Site construction began on 14 April 2015 with the ceremonial laying of the first stone.[13][14] The first on-sky observations with the engineering camera occurred on 24 October 2024,[15] while system first light images were released 23 June 2025. Full survey operations are planned to begin later in 2025, due to COVID-related schedule delays.[16] Data is scheduled to become fully public after two years.[17]`
	__Content__000000_My_Awesome_Go_Book.ContentPath = `content`
	__Content__000000_My_Awesome_Go_Book.OutputPath = `public`
	__Content__000000_My_Awesome_Go_Book.LayoutPath = `../../defaults/layouts`
	__Content__000000_My_Awesome_Go_Book.StaticPath = `../../defaults/static`
	__Content__000000_My_Awesome_Go_Book.Target = models.FILE

	__Page__000000_First_Steps.Name = `First Steps`
	__Page__000000_First_Steps.MardownContent = `Here is the detailed content for the first page within Chapter 1.

You can use standard Markdown formatting:

* Lists
* Are
* Easy

And include "code blocks".

Here is the logo embedded using HTML:

<img src="../../images/gong logo.svg" alt="My Logo" style="height: 200px; width: auto;">

You can place text around it.

| Header 1 | Header 2 | Header 3 |
| -------- | -------- | -------- |
| Cell 1-1 | Cell 1-2 | Cell 1-3 |
| Cell 2-1 | Cell 2-2 | Cell 2-3 |
| Cell 2-1 | Cell 2-2 | Cell 2-3 |

Legend of the table`

	// Setup of pointers
	// setup of Chapter instances pointers
	__Chapter__000000_Getting_Started.Pages = append(__Chapter__000000_Getting_Started.Pages, __Page__000000_First_Steps)
	// setup of Content instances pointers
	__Content__000000_My_Awesome_Go_Book.Chapters = append(__Content__000000_My_Awesome_Go_Book.Chapters, __Chapter__000000_Getting_Started)
	__Content__000000_My_Awesome_Go_Book.Chapters = append(__Content__000000_My_Awesome_Go_Book.Chapters, __Chapter__000001_Advanced_Topics)
	// setup of Page instances pointers
}

