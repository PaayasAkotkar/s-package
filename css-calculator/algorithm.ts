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