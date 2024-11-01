// generated code - do not edit

import { LayerAPI } from './layer-api'
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

export function CopyLayerToLayerAPI(layer: Layer, layerAPI: LayerAPI) {

	layerAPI.CreatedAt = layer.CreatedAt
	layerAPI.DeletedAt = layer.DeletedAt
	layerAPI.ID = layer.ID

	// insertion point for basic fields copy operations
	layerAPI.Display = layer.Display
	layerAPI.Name = layer.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	layerAPI.LayerPointersEncoding.Rects = []
	for (let _rect of layer.Rects) {
		layerAPI.LayerPointersEncoding.Rects.push(_rect.ID)
	}

	layerAPI.LayerPointersEncoding.Texts = []
	for (let _text of layer.Texts) {
		layerAPI.LayerPointersEncoding.Texts.push(_text.ID)
	}

	layerAPI.LayerPointersEncoding.Circles = []
	for (let _circle of layer.Circles) {
		layerAPI.LayerPointersEncoding.Circles.push(_circle.ID)
	}

	layerAPI.LayerPointersEncoding.Lines = []
	for (let _line of layer.Lines) {
		layerAPI.LayerPointersEncoding.Lines.push(_line.ID)
	}

	layerAPI.LayerPointersEncoding.Ellipses = []
	for (let _ellipse of layer.Ellipses) {
		layerAPI.LayerPointersEncoding.Ellipses.push(_ellipse.ID)
	}

	layerAPI.LayerPointersEncoding.Polylines = []
	for (let _polyline of layer.Polylines) {
		layerAPI.LayerPointersEncoding.Polylines.push(_polyline.ID)
	}

	layerAPI.LayerPointersEncoding.Polygones = []
	for (let _polygone of layer.Polygones) {
		layerAPI.LayerPointersEncoding.Polygones.push(_polygone.ID)
	}

	layerAPI.LayerPointersEncoding.Paths = []
	for (let _path of layer.Paths) {
		layerAPI.LayerPointersEncoding.Paths.push(_path.ID)
	}

	layerAPI.LayerPointersEncoding.Links = []
	for (let _link of layer.Links) {
		layerAPI.LayerPointersEncoding.Links.push(_link.ID)
	}

	layerAPI.LayerPointersEncoding.RectLinkLinks = []
	for (let _rectlinklink of layer.RectLinkLinks) {
		layerAPI.LayerPointersEncoding.RectLinkLinks.push(_rectlinklink.ID)
	}

}

// CopyLayerAPIToLayer update basic, pointers and slice of pointers fields of layer
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layerAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayerAPIToLayer(layerAPI: LayerAPI, layer: Layer, frontRepo: FrontRepo) {

	layer.CreatedAt = layerAPI.CreatedAt
	layer.DeletedAt = layerAPI.DeletedAt
	layer.ID = layerAPI.ID

	// insertion point for basic fields copy operations
	layer.Display = layerAPI.Display
	layer.Name = layerAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Rects)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Rects);
		return;
	}

	layer.Rects = new Array<Rect>()
	for (let _id of layerAPI.LayerPointersEncoding.Rects) {
		let _rect = frontRepo.map_ID_Rect.get(_id)
		if (_rect != undefined) {
			layer.Rects.push(_rect!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Texts)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Texts);
		return;
	}

	layer.Texts = new Array<Text>()
	for (let _id of layerAPI.LayerPointersEncoding.Texts) {
		let _text = frontRepo.map_ID_Text.get(_id)
		if (_text != undefined) {
			layer.Texts.push(_text!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Circles)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Circles);
		return;
	}

	layer.Circles = new Array<Circle>()
	for (let _id of layerAPI.LayerPointersEncoding.Circles) {
		let _circle = frontRepo.map_ID_Circle.get(_id)
		if (_circle != undefined) {
			layer.Circles.push(_circle!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Lines)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Lines);
		return;
	}

	layer.Lines = new Array<Line>()
	for (let _id of layerAPI.LayerPointersEncoding.Lines) {
		let _line = frontRepo.map_ID_Line.get(_id)
		if (_line != undefined) {
			layer.Lines.push(_line!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Ellipses)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Ellipses);
		return;
	}

	layer.Ellipses = new Array<Ellipse>()
	for (let _id of layerAPI.LayerPointersEncoding.Ellipses) {
		let _ellipse = frontRepo.map_ID_Ellipse.get(_id)
		if (_ellipse != undefined) {
			layer.Ellipses.push(_ellipse!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Polylines)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Polylines);
		return;
	}

	layer.Polylines = new Array<Polyline>()
	for (let _id of layerAPI.LayerPointersEncoding.Polylines) {
		let _polyline = frontRepo.map_ID_Polyline.get(_id)
		if (_polyline != undefined) {
			layer.Polylines.push(_polyline!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Polygones)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Polygones);
		return;
	}

	layer.Polygones = new Array<Polygone>()
	for (let _id of layerAPI.LayerPointersEncoding.Polygones) {
		let _polygone = frontRepo.map_ID_Polygone.get(_id)
		if (_polygone != undefined) {
			layer.Polygones.push(_polygone!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Paths)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Paths);
		return;
	}

	layer.Paths = new Array<Path>()
	for (let _id of layerAPI.LayerPointersEncoding.Paths) {
		let _path = frontRepo.map_ID_Path.get(_id)
		if (_path != undefined) {
			layer.Paths.push(_path!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.Links)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.Links);
		return;
	}

	layer.Links = new Array<Link>()
	for (let _id of layerAPI.LayerPointersEncoding.Links) {
		let _link = frontRepo.map_ID_Link.get(_id)
		if (_link != undefined) {
			layer.Links.push(_link!)
		}
	}
	if (!Array.isArray(layerAPI.LayerPointersEncoding.RectLinkLinks)) {
		console.error('Rects is not an array:', layerAPI.LayerPointersEncoding.RectLinkLinks);
		return;
	}

	layer.RectLinkLinks = new Array<RectLinkLink>()
	for (let _id of layerAPI.LayerPointersEncoding.RectLinkLinks) {
		let _rectlinklink = frontRepo.map_ID_RectLinkLink.get(_id)
		if (_rectlinklink != undefined) {
			layer.RectLinkLinks.push(_rectlinklink!)
		}
	}
}
