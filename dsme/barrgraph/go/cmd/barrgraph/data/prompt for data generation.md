Appended files:
- stage.go
- poster with barrgraph.png
- modernism.go

I want to do chronological flowcharts to map the genealogy and evolution of art movements, inspired by the methodology of **Alfred Barr**. His diagram for the 1936 exhibition *Cubism and Abstract Art* remains a canonical example of data visualization in art history.

I want to do the same with the modernism movement in design, starting the graph in 1900 with the english art & craft (affordable design) and finish it around 1965 with the scandinavians masters. 

I need you to help with a list of "movement", "artefact types" and "artist". "Movements and Artists can be related to places. Next I need a list of "influences" between those elements. The point is that you have to generate the data in a structured format that can be understood by a tool named barrgraph.

In practice, I want you to: 
- list in a markdown file "movement", "artefact types", "artist" and "place" of the modernism movement with the Dates for movement and artists
- understand the barrgraph format. Parse the enclosed stage.go file and understand carefully how it was translated by the tool barrgraph into the enclosed picture "poster with barrgraph.png"
- generate the abstract and the concrete syntax (the shapes) data in a format that is understandable by the tool barrgraph.
- update the enclosed modernism.go file with generated elements
- respect strict __<StructName>__<8-digit-index>_ nomenclature for variable declarations
- respect single-item appending for slices.