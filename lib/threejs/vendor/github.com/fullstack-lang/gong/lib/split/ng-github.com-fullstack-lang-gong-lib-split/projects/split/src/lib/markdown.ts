// generated code - do not edit

import { MarkdownAPI } from './markdown-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Markdown {

	static GONGSTRUCT_NAME = "Markdown"

	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	StackName: string = ""

	// insertion point for pointers and slices of pointers declarations

	CreatedAt?: string
	DeletedAt?: string
}

export function CopyMarkdownToMarkdownAPI(markdown: Markdown, markdownAPI: MarkdownAPI) {

	markdownAPI.CreatedAt = markdown.CreatedAt
	markdownAPI.DeletedAt = markdown.DeletedAt
	markdownAPI.ID = markdown.ID

	// insertion point for basic fields copy operations
	markdownAPI.Name = markdown.Name
	markdownAPI.StackName = markdown.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyMarkdownAPIToMarkdown update basic, pointers and slice of pointers fields of markdown
// from respectively the basic fields and encoded fields of pointers and slices of pointers of markdownAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMarkdownAPIToMarkdown(markdownAPI: MarkdownAPI, markdown: Markdown, frontRepo: FrontRepo) {

	markdown.CreatedAt = markdownAPI.CreatedAt
	markdown.DeletedAt = markdownAPI.DeletedAt
	markdown.ID = markdownAPI.ID

	// insertion point for basic fields copy operations
	markdown.Name = markdownAPI.Name
	markdown.StackName = markdownAPI.StackName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
