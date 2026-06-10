Currently, the time axis is suited for long periods, in span of years.


	// dates computed from tasks of the gantt
	ComputedStart    time.Time
	ComputedEnd      time.Time
	ComputedDuration time.Duration

	// NumberOfYearsBetweenTicks is when too many years would clutter the display
	// if 0 years is equivalent to 1
	NumberOfYearsBetweenTicks int

I need to hadle case where I might need to display trimesters , even months precision on the time axis.

I am sure there exist well not approachs to this problem with a suitable data model.

do not hesitate to trash the current approach with NumberOfYearsBetweenTicks