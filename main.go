package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

const (
	nextStepTime = 7 * time.Second
	frameSize    = 960 * 720 * 3
)

func main() {
	drone := tello.NewDriver("8888")
	keys := keyboard.NewDriver()
	// window := opencv.NewWindowDriver()

	work := func() {
		/// video

		// mplayer := exec.Command("mplayer", "-fps", "25", "-")
		// mplayerIn, _ := mplayer.StdinPipe()
		// if err := mplayer.Start(); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// drone.On(tello.ConnectedEvent, func(data interface{}) {
		// 	fmt.Println("Connected")
		// 	drone.StartVideo()
		// 	drone.SetVideoEncoderRate(4)
		// 	gobot.Every(100*time.Millisecond, func() {
		// 		drone.StartVideo()
		// 	})
		// })
		// drone.On(tello.VideoFrameEvent, func(data interface{}) {
		// 	pkt := data.([]byte)
		// 	if _, err := mplayerIn.Write(pkt); err != nil {
		// 		fmt.Println(err)
		// 	}
		// })

		/// opencv

		// ffmpeg := exec.Command("ffmpeg", "-i", "pipe:0", "-pix_fmt", "bgr24", "-vcodec", "rawvideo",
		// 	"-an", "-sn", "-s", "960x720", "-f", "rawvideo", "pipe:1")
		// ffmpegIn, _ := ffmpeg.StdinPipe()
		// ffmpegOut, _ := ffmpeg.StdoutPipe()
		// if err := ffmpeg.Start(); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// go func() {
		// 	for {
		// 		buf := make([]byte, frameSize)
		// 		if _, err := io.ReadFull(ffmpegOut, buf); err != nil {
		// 			fmt.Println(err)
		// 			continue
		// 		}

		// 		img := gocv.NewMatFromBytes(720, 960, gocv.MatTypeCV8UC3, buf)
		// 		if img.Empty() {
		// 			continue
		// 		}
		// 		window.ShowImage(img)
		// 		window.WaitKey(1)
		// 	}
		// }()

		// drone.On(tello.ConnectedEvent, func(data interface{}) {
		// 	fmt.Println("Connected")
		// 	drone.StartVideo()
		// 	drone.SetExposure(1)
		// 	drone.SetVideoEncoderRate(4)

		// 	gobot.Every(100*time.Millisecond, func() {
		// 		drone.StartVideo()
		// 	})
		// })

		// drone.On(tello.VideoFrameEvent, func(data interface{}) {
		// 	pkt := data.([]byte)
		// 	if _, err := ffmpegIn.Write(pkt); err != nil {
		// 		fmt.Println(err)
		// 	}
		// })

		/// work

		drone.TakeOff()
		gobot.After(nextStepTime, func() {
			drone.RightFlip()
		})
		gobot.After(nextStepTime, func() {
			drone.Land()
		})

		/// keyboad

		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)
			switch key.Key {
			case keyboard.T:
				drone.TakeOff()
			case keyboard.L:
				drone.Land()
			}
		})
	}

	robot := gobot.NewRobot(
		"tello",
		[]gobot.Connection{},
		[]gobot.Device{drone, keys},
		work,
	)
	robot.Start()
}
