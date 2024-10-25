package main

func foo() {
	/* LINE TO BE REMOVED IF compiler option is -target=lite
	log.Println("jojo")
	TO BE REMOVED IF compiler option is -target=lite */

	/* LINE TO BE REMOVED IF compiler option is -target=orm
	log.Println("jaja")
	LINE TO BE REMOVED IF compiler option is -target=orm */
}
