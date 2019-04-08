package main

import (
//	"fmt"
	"image"
	"image/jpeg"
	"os"
	"image/color"
//	"math"
)
func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func convolution(img *image.Image, matrice [][]int) *image.NRGBA {
    imageRGBA := image.NewNRGBA((*img).Bounds())
    w := (*img).Bounds().Dx()
    h := (*img).Bounds().Dy()

    done := make(chan bool)

 
    for y := 0; y < h; y++ {
        for x := 0; x < w; x++ {
            go func(){
                    
                    sumR := 0
                    sumB := 0
                    sumG := 0


            for i := -1; i <= 1; i++ {
                for j := -1; j <= 1; j++ {
                    var r uint32
                    var g uint32
                    var b uint32

                    var imageX int
                    var imageY int

                    imageX = x + i
                    imageY = y + j

                    r, g, b, _ = (*img).At(imageX, imageY).RGBA()
                    sumG = (sumG + (int(g) * matrice[i+1][j+1]))
                    sumR = (sumR + (int(r) * matrice[i+1][j+1]))
                    sumB = (sumB + (int(b) * matrice[i+1][j+1]))
                }
            }

            imageRGBA.Set(x, y, color.NRGBA{
                uint8(min(sumR/9, 0xffff) >> 8),
                uint8(min(sumG/9, 0xffff) >> 8),
                uint8(min(sumB/9, 0xffff) >> 8),
                255,
            })
            done <- true
            return

        }()



        }
    }
    for i := 0; i< h*w; i++ {
        <-done
    }

    return imageRGBA

}
    
func img(file string) *image.Image {
	existingImageFile, _ := os.Open(file)
	defer existingImageFile.Close()

	existingImageFile.Seek(0, 0)
	loadedImage, _ := jpeg.Decode(existingImageFile)
	return &loadedImage
	//outputFile, err := os.Create("test.jpg")
	//jpeg.Encode(outputFile, loadedImage, nil)

	//fmt.Println(loadedImage)
}

func main(){
	file := "fred.jpg"
	kernel := [][]int{}
	row := []int{1, 1, 1}
    kernel = append(kernel, row)
    kernel = append(kernel, row)
    kernel = append(kernel, row)
	myimg := img(file)
	final := convolution(myimg, kernel)
	outputFile, _ := os.Create("test2.jpg")
	jpeg.Encode(outputFile, final, nil)

}





