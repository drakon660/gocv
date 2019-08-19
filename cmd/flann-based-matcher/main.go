package main


import (
	"fmt"
	"sync"
	"time"
	"gocv.io/x/gocv"
)

func main() {
 	flann:=gocv.NewFlannBasedMatcher()

var wg sync.WaitGroup
	//window := gocv.NewWindow("Hello")
	startTime := time.Now()


	wg.Add(1)
	go func(){
		fmt.Println("robie 1")
		score1:=matchImage("G:\\GitHub\\gocv\\images\\doc.png", "G:\\GitHub\\gocv\\images\\pattern.png")
		fmt.Println(score1)
		wg.Done()
	}()

	wg.Add(1)
	go func(){
		fmt.Println("robie 2")
		score2:=matchImage("G:\\GitHub\\gocv\\images\\doc.png", "G:\\GitHub\\gocv\\images\\pattern.png")
		fmt.Println(score2)
		wg.Done()
	}()

	wg.Wait()
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime.Seconds())
 	flann.Close()
}

func matchImage(imageFile string, patternFile string) (score int){
	patternfile :=  gocv.IMRead(patternFile,gocv.IMReadGrayScale)

	imagefile := gocv.IMRead(imageFile,gocv.IMReadGrayScale)

	kaze:=gocv.NewKAZE()
	_, patternMat := kaze.DetectAndCompute(patternfile, gocv.NewMat())
	_, imageMat := kaze.DetectAndCompute(imagefile, gocv.NewMat())

	matcher:= gocv.NewFlannBasedMatcher()
	matches:= matcher.KnnMatch(patternMat,imageMat,2)

	for i := range matches {
		m:=matches[i][0].Distance
		n:=matches[i][1].Distance
		if m < 0.6 * n {
			score++
		}
	}
	return score
}