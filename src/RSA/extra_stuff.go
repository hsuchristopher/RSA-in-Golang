package main

import(
    "fmt"
    "github.com/qeesung/image2ascii/convert"
)

type Options struct{
    Ratio float64
    FixedWidth int
    FixedHeight int
    FitScreen bool
    StretchedScreen bool
    Colored bool
    Reversed bool
}


func printAsciiArt(){
    convertOptions := convert.DefaultOptions
    convertOptions.FixedWidth = 100
    convertOptions.FixedHeight = 40

    converter := convert.NewImageConverter()
    fmt.Print(converter.ImageFile2ASCIIString("../Images/antoni.jpeg", &convertOptions))
    fmt.Print(converter.ImageFile2ASCIIString("../Images/wan_quan.jpg", &convertOptions))

}

func printNaruto(){
    convertOptions := convert.DefaultOptions
    convertOptions.FixedWidth = 100
    convertOptions.FixedHeight = 40

    converter := convert.NewImageConverter()
    fmt.Print(converter.ImageFile2ASCIIString("../Images/naruto.jpeg", &convertOptions))

}
