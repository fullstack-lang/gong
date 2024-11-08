// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AnimateAPIs []*AnimateAPI

	CircleAPIs []*CircleAPI

	EllipseAPIs []*EllipseAPI

	LayerAPIs []*LayerAPI

	LineAPIs []*LineAPI

	LinkAPIs []*LinkAPI

	LinkAnchoredTextAPIs []*LinkAnchoredTextAPI

	PathAPIs []*PathAPI

	PointAPIs []*PointAPI

	PolygoneAPIs []*PolygoneAPI

	PolylineAPIs []*PolylineAPI

	RectAPIs []*RectAPI

	RectAnchoredPathAPIs []*RectAnchoredPathAPI

	RectAnchoredRectAPIs []*RectAnchoredRectAPI

	RectAnchoredTextAPIs []*RectAnchoredTextAPI

	RectLinkLinkAPIs []*RectLinkLinkAPI

	SVGAPIs []*SVGAPI

	TextAPIs []*TextAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, animateDB := range backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB {

		var animateAPI AnimateAPI
		animateAPI.ID = animateDB.ID
		animateAPI.AnimatePointersEncoding = animateDB.AnimatePointersEncoding
		animateDB.CopyBasicFieldsToAnimate_WOP(&animateAPI.Animate_WOP)

		backRepoData.AnimateAPIs = append(backRepoData.AnimateAPIs, &animateAPI)
	}

	for _, circleDB := range backRepo.BackRepoCircle.Map_CircleDBID_CircleDB {

		var circleAPI CircleAPI
		circleAPI.ID = circleDB.ID
		circleAPI.CirclePointersEncoding = circleDB.CirclePointersEncoding
		circleDB.CopyBasicFieldsToCircle_WOP(&circleAPI.Circle_WOP)

		backRepoData.CircleAPIs = append(backRepoData.CircleAPIs, &circleAPI)
	}

	for _, ellipseDB := range backRepo.BackRepoEllipse.Map_EllipseDBID_EllipseDB {

		var ellipseAPI EllipseAPI
		ellipseAPI.ID = ellipseDB.ID
		ellipseAPI.EllipsePointersEncoding = ellipseDB.EllipsePointersEncoding
		ellipseDB.CopyBasicFieldsToEllipse_WOP(&ellipseAPI.Ellipse_WOP)

		backRepoData.EllipseAPIs = append(backRepoData.EllipseAPIs, &ellipseAPI)
	}

	for _, layerDB := range backRepo.BackRepoLayer.Map_LayerDBID_LayerDB {

		var layerAPI LayerAPI
		layerAPI.ID = layerDB.ID
		layerAPI.LayerPointersEncoding = layerDB.LayerPointersEncoding
		layerDB.CopyBasicFieldsToLayer_WOP(&layerAPI.Layer_WOP)

		backRepoData.LayerAPIs = append(backRepoData.LayerAPIs, &layerAPI)
	}

	for _, lineDB := range backRepo.BackRepoLine.Map_LineDBID_LineDB {

		var lineAPI LineAPI
		lineAPI.ID = lineDB.ID
		lineAPI.LinePointersEncoding = lineDB.LinePointersEncoding
		lineDB.CopyBasicFieldsToLine_WOP(&lineAPI.Line_WOP)

		backRepoData.LineAPIs = append(backRepoData.LineAPIs, &lineAPI)
	}

	for _, linkDB := range backRepo.BackRepoLink.Map_LinkDBID_LinkDB {

		var linkAPI LinkAPI
		linkAPI.ID = linkDB.ID
		linkAPI.LinkPointersEncoding = linkDB.LinkPointersEncoding
		linkDB.CopyBasicFieldsToLink_WOP(&linkAPI.Link_WOP)

		backRepoData.LinkAPIs = append(backRepoData.LinkAPIs, &linkAPI)
	}

	for _, linkanchoredtextDB := range backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB {

		var linkanchoredtextAPI LinkAnchoredTextAPI
		linkanchoredtextAPI.ID = linkanchoredtextDB.ID
		linkanchoredtextAPI.LinkAnchoredTextPointersEncoding = linkanchoredtextDB.LinkAnchoredTextPointersEncoding
		linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText_WOP(&linkanchoredtextAPI.LinkAnchoredText_WOP)

		backRepoData.LinkAnchoredTextAPIs = append(backRepoData.LinkAnchoredTextAPIs, &linkanchoredtextAPI)
	}

	for _, pathDB := range backRepo.BackRepoPath.Map_PathDBID_PathDB {

		var pathAPI PathAPI
		pathAPI.ID = pathDB.ID
		pathAPI.PathPointersEncoding = pathDB.PathPointersEncoding
		pathDB.CopyBasicFieldsToPath_WOP(&pathAPI.Path_WOP)

		backRepoData.PathAPIs = append(backRepoData.PathAPIs, &pathAPI)
	}

	for _, pointDB := range backRepo.BackRepoPoint.Map_PointDBID_PointDB {

		var pointAPI PointAPI
		pointAPI.ID = pointDB.ID
		pointAPI.PointPointersEncoding = pointDB.PointPointersEncoding
		pointDB.CopyBasicFieldsToPoint_WOP(&pointAPI.Point_WOP)

		backRepoData.PointAPIs = append(backRepoData.PointAPIs, &pointAPI)
	}

	for _, polygoneDB := range backRepo.BackRepoPolygone.Map_PolygoneDBID_PolygoneDB {

		var polygoneAPI PolygoneAPI
		polygoneAPI.ID = polygoneDB.ID
		polygoneAPI.PolygonePointersEncoding = polygoneDB.PolygonePointersEncoding
		polygoneDB.CopyBasicFieldsToPolygone_WOP(&polygoneAPI.Polygone_WOP)

		backRepoData.PolygoneAPIs = append(backRepoData.PolygoneAPIs, &polygoneAPI)
	}

	for _, polylineDB := range backRepo.BackRepoPolyline.Map_PolylineDBID_PolylineDB {

		var polylineAPI PolylineAPI
		polylineAPI.ID = polylineDB.ID
		polylineAPI.PolylinePointersEncoding = polylineDB.PolylinePointersEncoding
		polylineDB.CopyBasicFieldsToPolyline_WOP(&polylineAPI.Polyline_WOP)

		backRepoData.PolylineAPIs = append(backRepoData.PolylineAPIs, &polylineAPI)
	}

	for _, rectDB := range backRepo.BackRepoRect.Map_RectDBID_RectDB {

		var rectAPI RectAPI
		rectAPI.ID = rectDB.ID
		rectAPI.RectPointersEncoding = rectDB.RectPointersEncoding
		rectDB.CopyBasicFieldsToRect_WOP(&rectAPI.Rect_WOP)

		backRepoData.RectAPIs = append(backRepoData.RectAPIs, &rectAPI)
	}

	for _, rectanchoredpathDB := range backRepo.BackRepoRectAnchoredPath.Map_RectAnchoredPathDBID_RectAnchoredPathDB {

		var rectanchoredpathAPI RectAnchoredPathAPI
		rectanchoredpathAPI.ID = rectanchoredpathDB.ID
		rectanchoredpathAPI.RectAnchoredPathPointersEncoding = rectanchoredpathDB.RectAnchoredPathPointersEncoding
		rectanchoredpathDB.CopyBasicFieldsToRectAnchoredPath_WOP(&rectanchoredpathAPI.RectAnchoredPath_WOP)

		backRepoData.RectAnchoredPathAPIs = append(backRepoData.RectAnchoredPathAPIs, &rectanchoredpathAPI)
	}

	for _, rectanchoredrectDB := range backRepo.BackRepoRectAnchoredRect.Map_RectAnchoredRectDBID_RectAnchoredRectDB {

		var rectanchoredrectAPI RectAnchoredRectAPI
		rectanchoredrectAPI.ID = rectanchoredrectDB.ID
		rectanchoredrectAPI.RectAnchoredRectPointersEncoding = rectanchoredrectDB.RectAnchoredRectPointersEncoding
		rectanchoredrectDB.CopyBasicFieldsToRectAnchoredRect_WOP(&rectanchoredrectAPI.RectAnchoredRect_WOP)

		backRepoData.RectAnchoredRectAPIs = append(backRepoData.RectAnchoredRectAPIs, &rectanchoredrectAPI)
	}

	for _, rectanchoredtextDB := range backRepo.BackRepoRectAnchoredText.Map_RectAnchoredTextDBID_RectAnchoredTextDB {

		var rectanchoredtextAPI RectAnchoredTextAPI
		rectanchoredtextAPI.ID = rectanchoredtextDB.ID
		rectanchoredtextAPI.RectAnchoredTextPointersEncoding = rectanchoredtextDB.RectAnchoredTextPointersEncoding
		rectanchoredtextDB.CopyBasicFieldsToRectAnchoredText_WOP(&rectanchoredtextAPI.RectAnchoredText_WOP)

		backRepoData.RectAnchoredTextAPIs = append(backRepoData.RectAnchoredTextAPIs, &rectanchoredtextAPI)
	}

	for _, rectlinklinkDB := range backRepo.BackRepoRectLinkLink.Map_RectLinkLinkDBID_RectLinkLinkDB {

		var rectlinklinkAPI RectLinkLinkAPI
		rectlinklinkAPI.ID = rectlinklinkDB.ID
		rectlinklinkAPI.RectLinkLinkPointersEncoding = rectlinklinkDB.RectLinkLinkPointersEncoding
		rectlinklinkDB.CopyBasicFieldsToRectLinkLink_WOP(&rectlinklinkAPI.RectLinkLink_WOP)

		backRepoData.RectLinkLinkAPIs = append(backRepoData.RectLinkLinkAPIs, &rectlinklinkAPI)
	}

	for _, svgDB := range backRepo.BackRepoSVG.Map_SVGDBID_SVGDB {

		var svgAPI SVGAPI
		svgAPI.ID = svgDB.ID
		svgAPI.SVGPointersEncoding = svgDB.SVGPointersEncoding
		svgDB.CopyBasicFieldsToSVG_WOP(&svgAPI.SVG_WOP)

		backRepoData.SVGAPIs = append(backRepoData.SVGAPIs, &svgAPI)
	}

	for _, textDB := range backRepo.BackRepoText.Map_TextDBID_TextDB {

		var textAPI TextAPI
		textAPI.ID = textDB.ID
		textAPI.TextPointersEncoding = textDB.TextPointersEncoding
		textDB.CopyBasicFieldsToText_WOP(&textAPI.Text_WOP)

		backRepoData.TextAPIs = append(backRepoData.TextAPIs, &textAPI)
	}

}
