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

	const __write__local_time = "2025-08-29 13:52:07.315544 CEST"
	const __write__utc_time__ = "2025-08-29 11:52:07.315544 UTC"

	const __commitId__ = "0000000041"

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
	__Page__000007_Vera_Rubin := (&models.Page{}).Stage(stage)
	__Page__000008_From_Concept_to_Concrete_Plans := (&models.Page{}).Stage(stage)

	// Setup of values

	__Chapter__000000_The_Observatory.Name = `The Observatory`
	__Chapter__000000_The_Observatory.MardownContent = `This chapter provides an overview of the Vera C. Rubin Observatory, its history, and its key components.

<img src="../images/Large_Synoptic_Survey_Telescope_3_4_render_2013.png" alt="Large Synoptic Survey Telescope" style="width: 800px;">
`

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

	__Page__000007_Vera_Rubin.Name = `Vera Rubin`
	__Page__000007_Vera_Rubin.MardownContent = `The decision to name a major scientific facility is never taken lightly. It is an act of commemoration, a way to etch a name into the annals of science, ensuring a legacy that will inspire future generations. In December 2019, through an act of the United States Congress, the Large Synoptic Survey Telescope (LSST) was officially christened the Vera C. Rubin Observatory. This dedication was more than just an honor; it was a profound statement. For the first time, a major, publicly-funded national observatory in the United States was named after a woman, recognizing the monumental contributions of an astronomer who fundamentally altered our perception of the cosmos and tirelessly championed the cause of women in science.


Vera Cooper Rubin (1928-2016) was a force of nature in the field of astronomy, a brilliant and tenacious scientist whose work provided the most compelling evidence for the existence of one of the universe's most profound mysteries: dark matter. Her journey in science was one of perseverance against the systemic gender biases of her time. From her early fascination with the stars as a young girl in Washington, D.C., she pursued her passion with unwavering determination. Despite being the sole astronomy graduate from Vassar College in 1948 and facing obstacles that barred women from certain academic programs—the graduate astronomy program at Princeton being a notable example—she carved out an extraordinary career.




The work that cemented her legacy began in the 1960s and 70s. Collaborating with instrument-maker Kent Ford, Rubin undertook meticulous studies of the rotation of spiral galaxies, starting with the neighboring Andromeda Galaxy. At the time, the prevailing wisdom, based on Newtonian physics, was that the stars and gas in the outer regions of a galaxy should orbit its center more slowly than those closer in, much like the outer planets of our solar system move more slowly than the inner ones. This is because the visible matter—the luminous stars and glowing gas clouds—is concentrated towards the galactic center, and thus the gravitational pull should diminish with distance.


Rubin and Ford, using Ford's highly sensitive spectrometer, made a startling discovery. They found that the orbital velocities of stars in the outskirts of galaxies did not decrease as expected. Instead, they remained surprisingly constant, or "flat," as far out as they could be measured. This meant the stars at the edges were moving just as fast as those near the center. The implications were revolutionary. For these high-velocity stars to remain in their orbits without being flung into intergalactic space, there had to be a tremendous amount of unseen mass exerting a powerful gravitational grip. The visible matter simply wasn't enough to hold the galaxies together.


This "missing mass" problem had been postulated as early as the 1930s by Swiss astronomer Fritz Zwicky, but his observations of galaxy clusters were largely dismissed. It was Rubin's systematic and irrefutable observations of individual galaxies that brought the concept of dark matter from a fringe theory to a central pillar of modern cosmology. Her work provided the observational bedrock upon which our current understanding of galaxy formation and the large-scale structure of the universe is built. Today, the scientific consensus is that the universe is composed of only about 5% ordinary, visible matter—the stuff that makes up stars, planets, and people. Roughly 27% is dark matter, and the remaining 68% is the even more enigmatic dark energy.


Naming the observatory after Vera Rubin is therefore exceptionally fitting. One of the primary scientific missions of the Rubin Observatory is to map the distribution of dark matter with unprecedented detail. Using a technique called weak gravitational lensing, the observatory will measure the subtle distortions in the light from distant galaxies as it passes through the gravitational fields of intervening dark matter. This will create a vast, three-dimensional map of the dark matter scaffolding of the universe, allowing scientists to test cosmological models and perhaps even glean clues about the nature of the dark matter particle itself, a quest that eluded Rubin in her lifetime.


Furthermore, the naming honors Rubin's role as a passionate advocate for women in science. She was a mentor to countless aspiring astronomers and a vocal critic of the barriers that women faced. She fought for their inclusion in professional organizations and their access to observing facilities. Naming this next-generation observatory in her honor serves as a powerful symbol of progress and an inspiration for a new, more inclusive generation of scientists. The Vera C. Rubin Observatory is not just a testament to her groundbreaking discoveries; it is a continuation of her legacy of exploration, perseverance, and the relentless pursuit of cosmic truths`

	__Page__000008_From_Concept_to_Concrete_Plans.Name = `From Concept to Concrete Plans`
	__Page__000008_From_Concept_to_Concrete_Plans.MardownContent = `Every great scientific instrument begins as an idea, a response to the fundamental questions that drive human curiosity. The Vera C. Rubin Observatory, in its modern form, is a product of decades of technological advancement and evolving scientific priorities, but its conceptual roots lie in the long-held astronomical ambition to create a comprehensive, dynamic map of the night sky. For centuries, astronomers have compiled star charts and catalogs, but these were static snapshots. The idea of a telescope that could repeatedly survey the entire sky, not just to catalog objects but to monitor their changes over time, was a revolutionary leap.

The seeds for what would become the Rubin Observatory were sown in the 1990s, a period of fertile discussion within the astronomical community about the future of the field. Advances in computing, data storage, and detector technology—specifically the development of large-format charge-coupled devices (CCDs)—were making it possible to imagine a project of a scale that was previously pure science fiction. The core concept was to build a wide-field survey telescope that could achieve an unprecedented combination of depth, area, and cadence. It would need to see faint objects (depth), cover a huge swath of the sky (area), and do so repeatedly and rapidly (cadence).

One of the earliest and most influential proponents of this idea was the astrophysicist J. Anthony "Tony" Tyson. In the late 1990s and early 2000s, Tyson was deeply engaged in the study of dark matter through weak gravitational lensing. He recognized that to truly map the distribution of dark matter across the cosmos, a survey of immense proportions was necessary. The subtle distortions caused by weak lensing are statistically tiny, requiring the observation of hundreds of millions or even billions of galaxies to create a meaningful signal. Existing telescopes could either see a small patch of sky deeply or a large patch shallowly, but none could do both.

Tyson envisioned a new kind of facility, one designed specifically for this purpose. His initial concept, which he began sketching out, was for what he aptly named the "Dark Matter Telescope." This instrument would be an "etendue monster," a term of art in optics referring to the product of an optical system's collecting area and its field of view. A high etendue value is the key to a fast and efficient survey. The goal was to build a telescope that could survey the entire visible sky from its location every few nights, cataloging billions of objects and, crucially, noting any changes in their position or brightness.

The scientific case for such a survey was compelling and extended far beyond the study of dark matter. Astronomers recognized that a deep, time-domain survey would be a veritable discovery engine. It would create a massive, publicly accessible database that could be used for a vast range of scientific investigations:

A Solar System Census: By imaging the sky repeatedly, such a survey would be incredibly effective at finding moving objects. It could discover and track millions of asteroids and comets, providing an unprecedented inventory of our solar system and dramatically improving our ability to identify potentially hazardous near-Earth objects.

Exploring the Transient Sky: The universe is filled with objects that change, blink, and explode. Supernovae, cataclysmic variable stars, gamma-ray burst afterglows, and other transient events are crucial for understanding stellar evolution and cosmology. A high-cadence survey would catch these events in the act, providing alerts to the astronomical community for immediate follow-up. It also held the promise of discovering entirely new types of celestial phenomena.


Mapping the Milky Way: A deep survey of the entire sky would provide a stunningly detailed picture of our own galaxy. By cataloging the positions, brightnesses, and colors of billions of stars, it would reveal the Milky Way's structure, from its central bulge to its faint outer halo, providing invaluable data for understanding its formation and evolution.

This powerful, multi-faceted scientific potential began to attract a growing coalition of supporters. The concept evolved from the "Dark Matter Telescope" to the more encompassing "Large-Aperture Synoptic Survey Telescope," which eventually became the Large Synoptic Survey Telescope (LSST). The word "synoptic" was key, emphasizing the goal of providing a comprehensive view of the whole sky at once. The idea was no longer just a thought experiment; it was becoming a concrete project proposal, a vision for the future of astronomy that promised to transform not just one, but nearly every subfield of the science. The seeds had been planted, and the next challenge was to convince the wider scientific community and the funding agencies that this ambitious vision should become a reality.

`

	// Setup of pointers
	// setup of Chapter instances pointers
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000000_Overview_and_Mission)
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000001_The_Simonyi_Survey_Telescope)
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000002_The_Camera)
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000007_Vera_Rubin)
	__Chapter__000000_The_Observatory.Pages = append(__Chapter__000000_The_Observatory.Pages, __Page__000008_From_Concept_to_Concrete_Plans)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000003_Legacy_Survey_of_Space_and_Time_LSST_)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000004_Key_Science_Pillars)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000005_The_Data)
	__Chapter__000001_The_Science.Pages = append(__Chapter__000001_The_Science.Pages, __Page__000006_Accessing_the_Data)
	// setup of Content instances pointers
	__Content__000000_The_Vera_C_Rubin_Observatory.Chapters = append(__Content__000000_The_Vera_C_Rubin_Observatory.Chapters, __Chapter__000000_The_Observatory)
	__Content__000000_The_Vera_C_Rubin_Observatory.Chapters = append(__Content__000000_The_Vera_C_Rubin_Observatory.Chapters, __Chapter__000001_The_Science)
	// setup of Page instances pointers
}

