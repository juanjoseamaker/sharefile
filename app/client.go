package app

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// The client only send files
func RunClient(config *Config) error {
	fmt.Printf("Sending %s to %s\n", config.ClientPath, config.ServerAddr)

	file, err := os.Open(config.ClientPath)
	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		fmt.Println("Directories are not supported")
		return nil
	}

	resp, err := http.Post("http://"+config.ServerAddr+"/upload", mime.TypeByExtension(filepath.Ext(config.ClientPath)), file)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Receiver response: %s\n", string(data))

	return nil
}
