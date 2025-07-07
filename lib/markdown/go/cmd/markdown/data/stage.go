package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
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

	const __write__local_time = "2025-07-07 05:03:52.547662 CEST"
	const __write__utc_time__ = "2025-07-07 03:03:52.547662 UTC"

	const __commitId__ = "0000000004"

	// Declaration of instances to stage

	__Content__000000_Dummy := (&models.Content{}).Stage(stage)

	// Setup of values

	__Content__000000_Dummy.Name = `Dummy`
	__Content__000000_Dummy.Content = `# Standalone Markdown Component

This is a **standalone** Angular component using "ngx-markdown"

## Features
- Renders markdown from string
- Supports remote markdown loading
- Fully configurable

- List item 1
- List item 2

![Local Angular Logo](assets/images/image.png)

<img src=/assets/images/star.svg width=50 />

# This is H1
This is plain text
## This is H2 with text format
Text formatting, such as **bold** and *italic*, code styles.
## Code Block

## List

- Bullet Item 1
- Bullet Item 2
- Bullet Item 3

1. Ordered Item 1
2. Ordered Item 2
3. Ordered Item 3

## CheckBox

- [ ] sample code
- [x] [Go](https://golang.org)
- [ ] ~~strikethrough~~

## Blockquote

> If you can dream it, you can do it.

### Horizontal Rule

---

## Table

| NAME  | AGE | COUNTRY |
|-------|-----|---------|
| David |  23 | USA     |
| John  |  30 | UK      |
| Bob   |  25 | Canada  |

## Image

![sample_image](/assets/images/star.svg)`

	// Setup of pointers
	// setup of Content instances pointers
}

