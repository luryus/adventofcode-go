package main

import(
    "os"
    "fmt"
    "log"
    "bufio"
    "regexp"
    "strconv"
)

const ScreenWidth int = 50
const ScreenHeight int = 5

type Screen [ScreenHeight][ScreenWidth]byte

const RectPattern = `rect (\d+)x(\d+)`
const RotateColPattern = `rotate column x=(\d+) by (\d+)`
const RotateRowPattern = `rotate row y=(\d+) by (\d+)`

func enableRect(screen *Screen, width int, height int) {
    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            (*screen)[i][j] = 1;
        }
    }
}

func rotateRow(screen *Screen, y int, by int) {
    (*screen)[y] = (*screen)[y][ScreenWidth-by:] + (*screen)[y][:ScreenWidth-by]
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Invalid args. Syntax:", os.Args[0], "<inputfile>")
        os.Exit(1)
    }

    inputfile := os.Args[1]
    var screen Screen

    file, err := os.Open(inputfile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    rectRegex := regexp.MustCompile(RectPattern)
    rotateColRegex := regexp.MustCompile(RotateColPattern)
    rotateRowRegex := regexp.MustCompile(RotateRowPattern)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        if m := rectRegex.FindStringSubmatch(line); m != nil {
            width, _ := strconv.Atoi(m[1])
            height, _ := strconv.Atoi(m[2])
            enableRect(&screen, width, height)
        } else if m := rotateRowRegex.FindStringSubmatch(line); m != nil {
            y, _ := strconv.Atoi(m[1])
            by, _ := strconv.Atoi(m[2])
            rotateRow(&screen, y, by)
        } else if m := rotateColRegex.FindStringSubmatch(line); m != nil {
            x, _ := strconv.Atoi(m[1])
            by, _ := strconv.Atoi(m[2])
            rotateCol(&screen, x, by)
        }
    }
}
