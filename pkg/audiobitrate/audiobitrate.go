package audiobitrate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func Convert16to24Wav(path, pathOut string) error {
	fmt.Printf("Converting %v to 16 bit depth...", path)
	err := ffmpeg_go.Input(path).Output(pathOut, ffmpeg_go.KwArgs{"c:a": "pcm_s16le"}).OverWriteOutput().ErrorToStdOut().Run()
	return err
}

func ConvertFolder24to16Wav(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatalf("directory %s does not exist", path)
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Converting %v to 16 bit depth...", path)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Check if the file is an audio file
		mime, err := mimetype.DetectFile(filepath.Join(path, file.Name()))
		if err != nil {
			log.Printf("error detecting MIME type for file %v: %v", file.Name(), err)
			continue
		}
		// Create the output directory if it doesn't exist
		outputDir := "./16bit"
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				log.Fatal(err)
			}
		}
		if strings.HasPrefix(mime.String(), "audio") {
			// Set the input and output file paths
			inputPath := filepath.Join(path, file.Name())
			outputPath := fmt.Sprintf("./16bit/%v", file.Name())

			// Convert the bit depth to 16 bits
			if err := ffmpeg_go.
				Input(inputPath).
				Output(outputPath, ffmpeg_go.KwArgs{"c:a": "pcm_s16le"}).
				OverWriteOutput().
				ErrorToStdOut().
				Run(); err != nil {
				return err
			}
		}
	}

	return nil

}
