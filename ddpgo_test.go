package ddpgo_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	ddpgo "github.com/milindmadhukar/ddp-go"
)

func generateInt8() byte {
	return byte(rand.Intn(256))
}

func TestController(t *testing.T) {

	fmt.Println("Testing Controller")

	plusClient, err := ddpgo.DefaultDDPConnection("192.168.1.41", 4048)

	if err != nil {
		log.Fatal(err)
	}

	defer plusClient.Close()

	crossClient, err := ddpgo.DefaultDDPConnection("192.168.1.42", 4048)

	if err != nil {
		log.Fatal(err)
	}

	defer crossClient.Close()

	pixelCount := 720

	fps := 1
	delay := 1000 / fps

	ticker := time.NewTicker(time.Millisecond * time.Duration(delay))

	// Start the ticker
	for range ticker.C {

		pixelData := make([]byte, pixelCount*3)
		for i := 0; i < pixelCount; i++ {
			r := generateInt8()
			g := generateInt8()
			b := generateInt8()

      brightness := 0.2

			fmt.Println(r, g, b)

			pixelData[i*3] = byte(float64(r) * brightness)
			pixelData[i*3+1] = byte(float64(g) * brightness)
			pixelData[i*3+2] = byte(float64(b) * brightness)
		}

		go func() {
			plusClient.Write(pixelData)
			plusClient.WriteOffset(pixelData, 360*3)
			crossClient.Write(pixelData)
			crossClient.WriteOffset(pixelData, 360*3)

		}()

	}
}
