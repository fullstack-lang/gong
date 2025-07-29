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

	const __write__local_time = "2025-07-29 08:07:07.966147 EDT"
	const __write__utc_time__ = "2025-07-29 12:07:07.966147 UTC"

	const __commitId__ = "0000000034"

	// Declaration of instances to stage

	__Chapter__000000_The_Observatory := (&models.Chapter{}).Stage(stage)
	__Chapter__000001_The_Science := (&models.Chapter{}).Stage(stage)

	__Content__000000_The_Vera_C_Rubin_Observatory := (&models.Content{}).Stage(stage)

	__Page__000000_Overview_and_Mission := (&models.Page{}).Stage(stage)
	__Page__000001_The_Simonyi_Survey_Telescope := (&models.Page{}).Stage(stage)
	__Page__000002_The_Camera := (&models.Page{}).Stage(stage)
	__Page__000003_Legacy_Survey_of_Space_and_Time_LSST_ := (&models.Page{}).Stage(stage)
	__Page__000004_Key_Science_Pillars := (&models.Page{}).Stage(stage)
	__Page__000005_The_Data := (&models.Page{}).Stage(stage)
	__Page__000006_Accessing_the_Data := (&models.Page{}).Stage(stage)

	// Setup of values

	__Chapter__000000_The_Observatory.Name = `The Observatory`
	__Chapter__000000_The_Observatory.MardownContent = `This chapter provides an overview of the Vera C. Rubin Observatory, its history, and its key components.`

	__Chapter__000001_The_Science.Name = `The Science`
	__Chapter__000001_The_Science.MardownContent = `The Legacy Survey of Space and Time (LSST) will generate an unprecedented dataset, enabling a wide array of scientific discoveries.
`

	__Content__000000_The_Vera_C_Rubin_Observatory.Name = `The Vera C. Rubin Observatory`
	__Content__000000_The_Vera_C_Rubin_Observatory.MardownContent = `The Vera C. Rubin Observatory, formerly known as the Large Synoptic Survey Telescope (LSST), is an astronomical observatory in Chile. Its main task is an astronomical survey of the entire available southern sky every few nights, creating a time-lapse record over ten years, the Legacy Survey of Space and Time (also abbreviated LSST).[2][3][4] The observatory is located on the El Peñón peak of Cerro Pachón, a 2,682-meter-high (8,799 ft) mountain in Coquimbo Region, in northern Chile, alongside the existing Gemini South and Southern Astrophysical Research Telescopes.[5] The Rubin Observatory base facility is located about 100 kilometres (62 miles) away from the observatory by road, in the city of La Serena. The observatory is named for Vera Rubin, an American astronomer who pioneered discoveries about galactic rotation rates.

Vera C. Rubin Observatory is a joint initiative of the U.S. National Science Foundation (NSF) and the U.S. Department of Energy's Office of Science and is operated jointly by NSF NOIRLab and SLAC National Accelerator Laboratory.[6]

The Rubin Observatory houses the Simonyi Survey Telescope, a wide-field reflecting telescope with an 8.4-meter primary mirror that can photograph the entire available sky every few nights.[7] The telescope uses a variant of three-mirror anastigmat, which allows the compact telescope to deliver sharp images over a very wide 3.5-degree-diameter field of view. Images are recorded by a 3.2-gigapixel charge-coupled device imaging (CCD) camera, the largest digital camera ever constructed.[8]

The Rubin Observatory was proposed in 2001 as the LSST, and construction of the mirror began (with private funds) in 2007. The LSST then became the top-ranked large ground-based project in the 2010 Astrophysics Decadal Survey, and the project officially began construction on 1 August 2014, when the United States National Science Foundation (NSF) authorized the FY2014 portion ($27.5 million) of its construction budget.[9] Funding comes from the NSF, the United States Department of Energy, and private funding raised by the dedicated international non-profit organization, the LSST Discovery Alliance.[10] Operations are under the management of the Association of Universities for Research in Astronomy (AURA).[11] The total construction cost was expected to be about $680 million.[12]

Site construction began on 14 April 2015 with the ceremonial laying of the first stone.[13][14] The first on-sky observations with the engineering camera occurred on 24 October 2024,[15] while system first light images were released 23 June 2025. Full survey operations are planned to begin later in 2025, due to COVID-related schedule delays.[16] Data is scheduled to become fully public after two years.[17]`
	__Content__000000_The_Vera_C_Rubin_Observatory.ContentPath = `content`
	__Content__000000_The_Vera_C_Rubin_Observatory.OutputPath = `public`
	__Content__000000_The_Vera_C_Rubin_Observatory.LayoutPath = `../../defaults/layouts`
	__Content__000000_The_Vera_C_Rubin_Observatory.StaticPath = `../../defaults/static`
	__Content__000000_The_Vera_C_Rubin_Observatory.Target = models.FILE
	__Content__000000_The_Vera_C_Rubin_Observatory.VersionInfo = `1.9.8`

	__Page__000000_Overview_and_Mission.Name = `Overview and Mission`
	__Page__000000_Overview_and_Mission.MardownContent = `The Vera C. Rubin Observatory, formerly known as the Large Synoptic Survey Telescope (LSST), is an astronomical observatory currently under construction in Chile. Its main task is to carry out a ten-year survey of the southern sky, the Legacy Survey of Space and Time (LSST).

The mission is to build a well-understood and calibrated system that will produce a vast astronomical dataset for a broad range of scientific investigations. The observatory is designed to address four main science areas:

1.  Probing Dark Energy and Dark Matter.
2.  Taking an inventory of the Solar System.
3.  Exploring the transient optical sky.
4.  Mapping the Milky Way.`

	__Page__000001_The_Simonyi_Survey_Telescope.Name = `The Simonyi Survey Telescope`
	__Page__000001_The_Simonyi_Survey_Telescope.MardownContent = `The heart of the Rubin Observatory is the Simonyi Survey Telescope, an 8.4-meter telescope with a unique three-mirror design. This design provides a very wide 3.5-degree field of view, allowing it to survey the entire available sky in just a few nights.

Key features include:
- A large aperture to collect significant light.
- A compact design to enable rapid slewing and repositioning.
- A state-of-the-art active optics system to maintain image quality.`

	__Page__000002_The_Camera.Name = `The Camera`
	__Page__000002_The_Camera.MardownContent = `The telescope is equipped with the largest digital camera ever constructed for astronomy. It has 3.2 gigapixels and will produce images of exceptional quality.

| Feature          | Specification                |
|------------------|------------------------------|
| Resolution       | 3,200 megapixels (3.2 GP)    |
| Sensor Type      | High-sensitivity CCDs        |
| Filter System    | 6 filters (u, g, r, i, z, y) |
| Cooldown        | Cryogenic cooling            |

The camera's data will be read out in just two seconds, a critical feature for the rapid cadence of the survey.`

	__Page__000003_Legacy_Survey_of_Space_and_Time_LSST_.Name = `Legacy Survey of Space and Time (LSST)`
	__Page__000003_Legacy_Survey_of_Space_and_Time_LSST_.MardownContent = `The LSST is a planned 10-year survey of the southern sky that will take place at the Vera C. Rubin Observatory. It will image the sky over and over again, producing a time-lapse movie of the universe.

The survey is expected to generate approximately 20 terabytes of data per night, leading to a total database of over 15 petabytes by the end of the survey. This "Big Data" challenge requires sophisticated data management and processing techniques.
`

	__Page__000004_Key_Science_Pillars.Name = `Key Science Pillars`
	__Page__000004_Key_Science_Pillars.MardownContent = `The LSST is designed to make transformative discoveries in four key areas:

1.  **Dark Energy and Dark Matter:** By observing billions of galaxies, the LSST will map the distribution of dark matter and measure the effects of dark energy on the expansion of the universe.
2.  **Solar System Inventory:** The survey will catalog millions of small bodies in our solar system, from near-Earth asteroids to distant Kuiper Belt objects.
3.  **The Transient Sky:** Rubin Observatory will detect billions of transient events, such as supernovae, variable stars, and potentially new and unexpected phenomena.
4.  **Milky Way Mapping:** The LSST will provide a detailed map of the structure and evolution of our own galaxy, the Milky Way.
`

	__Page__000005_The_Data.Name = `The Data`
	__Page__000005_The_Data.MardownContent = `he data management system for the Rubin Observatory is a massive undertaking. It is responsible for:
- Transporting data from the telescope in Chile to the archive center in the United States.
- Processing the raw images to produce scientifically useful data products.
- Archiving the data and making it accessible to users.
- Providing tools and services for data analysis.`

	__Page__000006_Accessing_the_Data.Name = `Accessing the Data`
	__Page__000006_Accessing_the_Data.MardownContent = `The primary data products will be accessible through the Rubin Science Platform (RSP), a web-based portal. The RSP will provide a suite of tools for querying the database, visualizing images, and performing scientific analysis.

Data will be released in annual data releases. After a proprietary period for scientists in the US and Chile, all data products will become fully public to the world.
`

	// Setup of pointers
	// setup of Chapter instances pointers
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000000_Overview_and_Mission)
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000001_The_Simonyi_Survey_Telescope)
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000002_The_Camera)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000003_Legacy_Survey_of_Space_and_Time_LSST_)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000004_Key_Science_Pillars)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000005_The_Data)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000006_Accessing_the_Data)
	// setup of Content instances pointers
	__Content__000000_The_Vera_C_Rubin_Observatory.Chapters = append(__Content__000000_The_Vera_C_Rubin_Observatory.Chapters, __Chapter__000000_The_Observatory)
	__Content__000000_The_Vera_C_Rubin_Observatory.Chapters = append(__Content__000000_The_Vera_C_Rubin_Observatory.Chapters, __Chapter__000001_The_Science)
	// setup of Page instances pointers
}

