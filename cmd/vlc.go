package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath) // Открытие файла
	if err != nil {
		handleErr(err)
	}
	defer r.Close() //Обязательно закрываем

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	// packed := Encode(data)
	packed := "" //+ string(data) //TODO: remove
	fmt.Println(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 8644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) string {
	// /path/to/file/myFile.txt -> myFile.vlc
	fileName := filepath.Base(path) //Получаем полное название файла, к примеру: myFile.txt
	// ext := filepath.Ext(fileName)   // Получаем расшиерние файла, .txt
	// baseName := strings.TrimSuffix(fileName, ext) //Отрезаем расширение от названия файла

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension // Добавляем свое расширение для файла
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
