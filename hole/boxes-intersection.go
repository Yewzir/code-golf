package hole

import (
    "strconv"
    "strings"
    "math/rand"
)

type bbox struct {
    x int
    y int
    w int
    h int
}

func strconvbox(box bbox) (out string) {
    var outs []string
    outs = append(outs, strconv.Itoa(box.x))
    outs = append(outs, strconv.Itoa(box.y))
    outs = append(outs, strconv.Itoa(box.w))
    outs = append(outs, strconv.Itoa(box.h))
    return strings.Join(outs, " ")
}

func unbox(b bbox) (tlx, tly, brx, bry int) {
    tlx = b.x
    tly = b.y
    brx = b.x + b.h
    bry = b.y + b.w

    return
}

// til that go doesn't have built-in max/min for int
func minint(a, b int) int {
    if a < b {
        return a}
    return b
}

func maxint(a, b int) int {
    if a > b {
        return a}
    return b
}

func calculateIntersection(b1, b2 bbox) int {
    tlx1, tly1, brx1, bry1 := unbox(b1)
    tlx2, tly2, brx2, bry2 := unbox(b2)

    // find top-left and bottom-right intersection coordinates
    itlx := maxint(tlx1, tlx2)
    itly := maxint(tly1, tly2)
    ibrx := minint(brx1, brx2)
    ibry := minint(bry1, bry2)

    // calculate intersection dimensions
    ih := itlx - ibrx
    iw := itly - ibry

    if iw < 0 || ih < 0 || iw > b1.w + b2.w || ih > b1.h + b2.h {
        return 0
    }

    return iw*ih
}

// generator of random non-null boxes (i.e. with  area != 0)
func boxGen() bbox {
    return bbox{x:rand.Intn(101), y:rand.Intn(101), w:rand.Intn(50)+1, h:rand.Intn(50)+1}
}

func boxesIntersection() (args []string, out string) {

    var outs []string

    //// default cases
    // define two non overlapping 1x1 boxes
    b1 := bbox{x:0, y:0, w:1, h:1}
    b2 := bbox{x:0, y:0, w:2, h:2}
    b3 := bbox{x:3, y:3, w:1, h:1}

    // b1 and b2 partially overlap
    //-> intersection area is 1 squared pixels
    args = append(args, strconvbox(b1)+" "+strconvbox(b1))
    outs = append(outs, strconv.Itoa(calculateIntersection(b1, b2)))

    // b1 and b3 do not overlap -> intersection area is 0
    args = append(args, strconvbox(b1)+" "+strconvbox(b2))
    outs = append(outs, strconv.Itoa(calculateIntersection(b1, b3)))

    //// generate 98 more random cases
    zeros := 0
    non_zeros := 0
    for zeros + non_zeros < 98 {
        b1 = boxGen()
        b2 = boxGen()
        intersection := calculateIntersection(b1, b2)

        // we want 90 fun (i.e. non zero) cases
        if intersection > 0 && non_zeros < 90 {
            args = append(args, strconvbox(b1)+" "+strconvbox(b2))
            outs = append(outs, strconv.Itoa(intersection))
            non_zeros += 1
        // 8 more zeroes to be sure edge case is tested
        } else if intersection == 0 && zeros < 8 {
            args = append(args, strconvbox(b1)+" "+strconvbox(b2))
            outs = append(outs, strconv.Itoa(intersection))
            zeros += 1
        }
    }

    // this should shuffle args and outputs in the same way
	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
		outs[i], outs[j] = outs[j], outs[i]
	})

    // there are 100 test cases in total
    out = strings.Join(outs, "\n")
	return
}
