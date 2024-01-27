// generated code - do not edit

import { LayerDB } from './layer-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Rect } from './rect'
import { Text } from './text'
import { Circle } from './circle'
import { Line } from './line'
import { Ellipse } from './ellipse'
import { Polyline } from './polyline'
import { Polygone } from './polygone'
import { Path } from './path'
import { Link } from './link'
import { RectLinkLink } from './rectlinklink'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Layer {

	static GONGSTRUCT_NAME = "Layer"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Display: boolean = false
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Rects: Array<Rect> = []
	Texts: Array<Text> = []
	Circles: Array<Circle> = []
	Lines: Array<Line> = []
	Ellipses: Array<Ellipse> = []
	Polylines: Array<Polyline> = []
	Polygones: Array<Polygone> = []
	Paths: Array<Path> = []
	Links: Array<Link> = []
	RectLinkLinks: Array<RectLinkLink> = []
}

export function CopyLayerToLayerDB(layer: Layer, layerDB: LayerDB) {

	layerDB.CreatedAt = layer.CreatedAt
	layerDB.DeletedAt = layer.DeletedAt
	layerDB.ID = layer.ID

	// insertion point for basic fields copy operations
	layerDB.Display = layer.Display
	layerDB.Name = layer.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	layerDB.LayerPointersEncoding.Rects = []
	for (let _rect of layer.Rects) {
		layerDB.LayerPointersEncoding.Rects.push(_rect.ID)
	}

	layerDB.LayerPointersEncoding.Texts = []
	for (let _text of layer.Texts) {
		layerDB.LayerPointersEncoding.Texts.push(_text.ID)
	}

	layerDB.LayerPointersEncoding.Circles = []
	for (let _circle of layer.Circles) {
		layerDB.LayerPointersEncoding.Circles.push(_circle.ID)
	}

	layerDB.LayerPointersEncoding.Lines = []
	for (let _line of layer.Lines) {
		layerDB.LayerPointersEncoding.Lines.push(_line.ID)
	}

	layerDB.LayerPointersEncoding.Ellipses = []
	for (let _ellipse of layer.Ellipses) {
		layerDB.LayerPointersEncoding.Ellipses.push(_ellipse.ID)
	}

	layerDB.LayerPointersEncoding.Polylines = []
	for (let _polyline of layer.Polylines) {
		layerDB.LayerPointersEncoding.Polylines.push(_polyline.ID)
	}

	layerDB.LayerPointersEncoding.Polygones = []
	for (let _polygone of layer.Polygones) {
		layerDB.LayerPointersEncoding.Polygones.push(_polygone.ID)
	}

	layerDB.LayerPointersEncoding.Paths = []
	for (let _path of layer.Paths) {
		layerDB.LayerPointersEncoding.Paths.push(_path.ID)
	}

	layerDB.LayerPointersEncoding.Links = []
	for (let _link of layer.Links) {
		layerDB.LayerPointersEncoding.Links.push(_link.ID)
	}

	layerDB.LayerPointersEncoding.RectLinkLinks = []
	for (let _rectlinklink of layer.RectLinkLinks) {
		layerDB.LayerPointersEncoding.RectLinkLinks.push(_rectlinklink.ID)
	}

}

// CopyLayerDBToLayer update basic, pointers and slice of pointers fields of layer
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layerDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayerDBToLayer(layerDB: LayerDB, layer: Layer, frontRepo: FrontRepo) {

	layer.CreatedAt = layerDB.CreatedAt
	layer.DeletedAt = layerDB.DeletedAt
	layer.ID = layerDB.ID

	// insertion point for basic fields copy operations
	layer.Display = layerDB.Display
	layer.Name = layerDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	layer.Rects = new Array<Rect>()
	for (let _id of layerDB.LayerPointersEncoding.Rects) {
		let _rect = frontRepo.map_ID_Rect.get(_id)
		if (_rect != undefined) {
			layer.Rects.push(_rect!)
		}
	}
	layer.Texts = new Array<Text>()
	for (let _id of layerDB.LayerPointersEncoding.Texts) {
		let _text = frontRepo.map_ID_Text.get(_id)
		if (_text != undefined) {
			layer.Texts.push(_text!)
		}
	}
	layer.Circles = new Array<Circle>()
	for (let _id of layerDB.LayerPointersEncoding.Circles) {
		let _circle = frontRepo.map_ID_Circle.get(_id)
		if (_circle != undefined) {
			layer.Circles.push(_circle!)
		}
	}
	layer.Lines = new Array<Line>()
	for (let _id of layerDB.LayerPointersEncoding.Lines) {
		let _line = frontRepo.map_ID_Line.get(_id)
		if (_line != undefined) {
			layer.Lines.push(_line!)
		}
	}
	layer.Ellipses = new Array<Ellipse>()
	for (let _id of layerDB.LayerPointersEncoding.Ellipses) {
		let _ellipse = frontRepo.map_ID_Ellipse.get(_id)
		if (_ellipse != undefined) {
			layer.Ellipses.push(_ellipse!)
		}
	}
	layer.Polylines = new Array<Polyline>()
	for (let _id of layerDB.LayerPointersEncoding.Polylines) {
		let _polyline = frontRepo.map_ID_Polyline.get(_id)
		if (_polyline != undefined) {
			layer.Polylines.push(_polyline!)
		}
	}
	layer.Polygones = new Array<Polygone>()
	for (let _id of layerDB.LayerPointersEncoding.Polygones) {
		let _polygone = frontRepo.map_ID_Polygone.get(_id)
		if (_polygone != undefined) {
			layer.Polygones.push(_polygone!)
		}
	}
	layer.Paths = new Array<Path>()
	for (let _id of layerDB.LayerPointersEncoding.Paths) {
		let _path = frontRepo.map_ID_Path.get(_id)
		if (_path != undefined) {
			layer.Paths.push(_path!)
		}
	}
	layer.Links = new Array<Link>()
	for (let _id of layerDB.LayerPointersEncoding.Links) {
		let _link = frontRepo.map_ID_Link.get(_id)
		if (_link != undefined) {
			layer.Links.push(_link!)
		}
	}
	layer.RectLinkLinks = new Array<RectLinkLink>()
	for (let _id of layerDB.LayerPointersEncoding.RectLinkLinks) {
		let _rectlinklink = frontRepo.map_ID_RectLinkLink.get(_id)
		if (_rectlinklink != undefined) {
			layer.RectLinkLinks.push(_rectlinklink!)
		}
	}
}
