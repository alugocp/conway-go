package main
import(
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math/rand"
	"log"
)
const scale=2
const WIDTH=160
const HEIGHT=120
var black color.RGBA=color.RGBA{75,139,190,255}//95,95,95
var white color.RGBA=color.RGBA{255,232,115,255}//233,233,233
var grid [WIDTH][HEIGHT]uint8=[WIDTH][HEIGHT]uint8{}
var buffer [WIDTH][HEIGHT]uint8=[WIDTH][HEIGHT]uint8{}
var counter int=0

// Logic
func update() error{
	for x:=1;x<WIDTH-1;x++ {
		for y:=1;y<HEIGHT-1;y++ {
			buffer[x][y]=0
			n:=grid[x-1][y-1]+grid[x-1][y+0]+grid[x-1][y+1]+grid[x+0][y-1]+grid[x+0][y+1]+grid[x+1][y-1]+grid[x+1][y+0]+grid[x+1][y+1]
			if grid[x][y]==0 && n==3 {
				buffer[x][y]=1
			}else if n>3 || n<2 {
				buffer[x][y]=0
			}else{
				buffer[x][y]=grid[x][y]
			}
		}
	}
	temp:=buffer
	buffer=grid
	grid=temp
	return nil
}

// Main
func render(screen *ebiten.Image){
	screen.Fill(white)
	for x:=0;x<WIDTH;x++ {
		for y:=0;y<HEIGHT;y++ {
			if grid[x][y]>0{
				for x1:=0;x1<scale;x1++ {
					for y1:=0;y1<scale;y1++ {
						screen.Set((x*scale)+x1,(y*scale)+y1,black)
					}
				}
			}
		}
	}
}
func frame(screen *ebiten.Image) error{
	counter++
	var err error=nil
	if counter==20 {
		err=update()
		counter=0
	}
	if !ebiten.IsDrawingSkipped(){
    render(screen)
  }
  return err
}
func main() {
	for x:=1;x<WIDTH-1;x++ {
		for y:=1;y<HEIGHT-1;y++ {
			if(rand.Float32()<0.5){
				grid[x][y]=1
			}
		}
	}
	if err:=ebiten.Run(frame,WIDTH*scale,HEIGHT*scale,2,"Conway's Game of Go");err!=nil{
    log.Fatal(err)
	}
}
