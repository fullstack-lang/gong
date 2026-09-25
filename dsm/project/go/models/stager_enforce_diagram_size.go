package models

import "math"

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {

		width := 0.0
		height := 0.0

		// parse all concrete shapes in the diagram that are not links
		//  to compute the size of the diagram
		// the size of the diagram is the max of the position of the shapes + their size
		for _, shape := range diagram.Product_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}

		if !diagram.IsTimeDiagram {
			for _, shape := range diagram.Task_Shapes {
				if shape.X+shape.Width > width {
					width = shape.X + shape.Width
				}
				if shape.Y+shape.Height > height {
					height = shape.Y + shape.Height
				}
			}
		}

		for _, shape := range diagram.Note_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}
		for _, shape := range diagram.Resource_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}

		// add a margin to the diagram size
		margin := 100.0
		width += margin
		height += margin

		if diagram.IsTimeDiagram {
			maxRight := diagram.XRightMargin

			hasDuration := false
			if diagram.HideWeekendsPeriod {
				hasDuration = diagram.workDurationBetween(diagram.ComputedStart, diagram.ComputedEnd) > 0
			} else {
				hasDuration = diagram.ComputedDuration > 0
			}

			if hasDuration {
				root := stager.getRootLibrary()
				nbPixPerChar := 8.0
				if root != nil && root.NbPixPerCharacter > 0 {
					nbPixPerChar = root.NbPixPerCharacter
				}

				for _, taskGroupShape := range diagram.TaskGroupShapes {
					if taskGroupShape.IsHidden || taskGroupShape.TaskGroup == nil {
						continue
					}
					taskGroup := taskGroupShape.TaskGroup

					for _, task := range taskGroup.Tasks {
						taskShape, ok := diagram.map_Task_TaskShape[task]
						if !ok || taskShape.IsHidden {
							continue
						}

						textWidth := float64(len(task.Name)) * nbPixPerChar

						if task.IsMilestone {
							lineX := diagram.DateToX(task.Start)
							diamondWidth := 18.0
							dummyX := lineX + diamondWidth/2.0
							diamondRight := lineX + diamondWidth/2.0

							var textRight float64
							switch task.TextPosition {
							case TEXT_POSITION_TOP, TEXT_POSITION_BOTTOM, TEXT_POSITION_CENTER:
								textRight = dummyX + diagram.XLeftText + task.XOffset + textWidth/2.0
							case TEXT_POSITION_LEFT:
								textRight = dummyX + diagram.XLeftText + task.XOffset
							default: // default or TEXT_POSITION_RIGHT
								textRight = dummyX + diagram.XLeftText + task.XOffset + textWidth
							}

							rightX := math.Max(diamondRight, textRight)
							if rightX > maxRight {
								maxRight = rightX
							}
						} else {
							var taskToDisplay = *task
							endToDisplay := taskToDisplay.End
							if taskToDisplay.IsAllDay && !taskToDisplay.IsMilestone {
								endToDisplay = endToDisplay.AddDate(0, 0, 1)
							}
							if diagram.UseManualStartAndEndDates {
								if task.Start.Before(diagram.ManualStart) {
									taskToDisplay.Start = diagram.ManualStart
								}
								if endToDisplay.After(diagram.ManualEnd) {
									endToDisplay = diagram.ManualEnd
								}
							}
							startX := diagram.DateToX(taskToDisplay.Start)
							endX := diagram.DateToX(endToDisplay)

							var textRight float64
							switch task.TextPosition {
							case TEXT_POSITION_RIGHT:
								textRight = endX + diagram.XLeftText + task.XOffset + textWidth
							case TEXT_POSITION_TOP, TEXT_POSITION_BOTTOM:
								textRight = (startX+endX)/2.0 + task.XOffset + textWidth/2.0
							default:
								textRight = endX
							}

							rightX := math.Max(endX, textRight)
							if rightX > maxRight {
								maxRight = rightX
							}
						}
					}
				}
			}

			if width < maxRight+margin {
				width = maxRight + margin
			}

			var nbVisibleTaskGroups int
			for _, taskGroupShape := range diagram.TaskGroupShapes {
				if !taskGroupShape.IsHidden {
					nbVisibleTaskGroups++
				}
			}

			timeDiagramHeight := diagram.YTopMargin + diagram.LaneHeight*float64(nbVisibleTaskGroups)
			timeDiagramHeight += diagram.DateYOffset + margin // margin for date labels
			if height < timeDiagramHeight {
				height = timeDiagramHeight
			}
		}

		if width != diagram.Width {
			diagram.Width = width
			needCommit = true
		}

		if height != diagram.Height {
			diagram.Height = height
			needCommit = true
		}
	}

	return
}
