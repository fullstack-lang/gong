import * as gongsvg from 'gongsvg'

export interface SegmentConf {

    // horizontal & vertical cutoff
    HC: number
    VC: number
    Orientation: gongsvg.OrientationType
}

function isBetweenZeroAndOne(value: number): boolean {
    return value >= 0.0 && value <= 1.0;
}

function ensureBetweenZeroAndOne(value: number): number {
    if (value < 0.0) {
        return 0.0;
    } else if (value > 1.0) {
        return 1.0;
    } else {
        return value;
    }
}


// computeSegmentConf takes an input SegmentConf
// and returns it unless
//
// there is a need for a orientation swap
// - Orientation is horizontal and VC is outside the [0, 1] range
// - Orientation is vertical and HC is outside the [0, 1] range
// It performs the orientation swaps and put HC or VC back into the  [0, 1] range
//
// HC and VC are both outside the [0, 1] range
// - it sets valid to false
//
export function computeSegmentConf(input: SegmentConf): { valid: boolean, conf: SegmentConf } {

    let res = { valid: true, conf: structuredClone(input) }

    if (!isBetweenZeroAndOne(input.HC) && !isBetweenZeroAndOne(input.VC)) {
        res.valid = false
        return res
    }

    if (input.Orientation == gongsvg.OrientationType.ORIENTATION_HORIZONTAL) {
        if (!isBetweenZeroAndOne(input.VC)) {
            // swap orientation
            res.conf.Orientation = gongsvg.OrientationType.ORIENTATION_VERTICAL
        }
    }
    if (input.Orientation == gongsvg.OrientationType.ORIENTATION_VERTICAL) {
        if (!isBetweenZeroAndOne(input.HC)) {
            // swap orientation
            res.conf.Orientation = gongsvg.OrientationType.ORIENTATION_HORIZONTAL
        }
    }
    res.conf.VC = ensureBetweenZeroAndOne(res.conf.VC)
    res.conf.HC = ensureBetweenZeroAndOne(res.conf.HC)

    return res
}