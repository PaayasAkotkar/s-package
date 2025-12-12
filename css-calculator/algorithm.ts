export namespace CSScalculator {

    // calcIncreaseBy returns the new widht and height increase by perc
    // currently it is only implemented for percentage
    // increaseBy: must be in perc
    // width and height must be in px
    // example: input: 310*106 increaseBy 10%
    //          output: 341*117
    export function calcIncreaseBy(width: number, height: number, increaseBy: string): { width: number, height: number } {
        var valid = /%/.test(increaseBy)
        var w: number = 0, h: number = 0
        if (valid) {
            var perc = parseFloat(increaseBy.replace('%', ''))
            var x = 100 + perc
            var y = x / 100 // give back what taken the commmon rule of math
            var w = Math.round(width * y)
            var h = Math.round(height * y)
        }
        return { width: w, height: h }
    }

}


// this library is created to work with figma
// you dont have to do any further ccalualtion just put the x*y in the formula
export namespace CSSMaths {

    // IncreaseBy returns the value for further investigation based on percentage
    // perc: percentage to calculate on
    // size: current size meaning the full screen is 100 so 100
    export function IncreaseBy(perc: number, size: number = 100) {
        return 1 + (perc / size)
    }

    // DecreaseBy returns the value for further investigation based on percentage
    // perc: percentage to calculate on
    // size: current size meaning the full screen is 100 so 100
    export function DecreaseBy(perc: number, size: number = 100) {
        return 1 - (perc / size)
    }

    // RaiseBy returns the new width and height on privded pixelss
    // perc: it is cruical to either use the IncreaseBy method or DecreaseBy method to get perfect perc or pass calculated percentage
    export function RaiseBy(width: number, height: number, perc: number): { w: number, h: number } {
        var co = IncreaseBy(perc)
        return { w: width * co, h: height * co }
    }

    // RaiseBy returns the new width and height on privded pixelss
    // perc: it is cruical to either use the IncreaseBy method or DecreaseBy method to get perfect perc or pass calculated percentage
    export function ReduceBy(width: number, height: number, perc: number): { w: number, h: number } {
        var co = DecreaseBy(perc)
        return { w: width * co, h: height * co }
    }
}