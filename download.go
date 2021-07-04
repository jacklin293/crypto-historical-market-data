package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func downloadFiles(pair string, interval string, year string, month string) (err error) {
	if month == "all" {
		for i := 1; i <= 12; i++ {
			if err = downloadKlineFile(pair, interval, year, strconv.Itoa(i)); err != nil {
				return err
			}
		}
	} else {
		if err = downloadKlineFile(pair, interval, year, month); err != nil {
			return err
		}
	}
	return nil
}

func downloadKlineFile(pair string, interval string, year string, month string) error {
	fileName := getZipFileName(pair, interval, year, month)
	filePath := getZipFilePath(pair, interval, year, month)
	downloadUrl := fmt.Sprintf("%s/%s/%s/%s", BINANCE_PUBLIC_DATA_URL, pair, interval, fileName)

	// Check if file exists
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(" - Failed to download '%s', status code: %d", downloadUrl, resp.StatusCode)
	}

	// Create folder if not exists
	folderPath := getZipFolderPath(pair, interval, year)
	if err = createFolderIfNotExists(folderPath); err != nil {
		return err
	}

	// Check if file exists
	if checkIfFileExists(filePath) {
		fmt.Printf(" - File '%s' has already existed\n", filePath)
		return nil
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}
	fmt.Printf(" - File '%s' has been downloaded successfully\n", filePath)
	return nil
}

func getZipFolderPath(pair string, interval string, year string) string {
	return FOLDER_DOWNLOAD + "/" + fmt.Sprintf("binance-%s-%s-kline-%s", pair, interval, year)
}

func getZipFileName(pair string, interval string, year string, month string) string {
	return fmt.Sprintf("%s-%s-%s-%02s.zip", pair, interval, year, month)
}

func getZipFilePath(pair string, interval string, year string, month string) string {
	return getZipFolderPath(pair, interval, year) + "/" + getZipFileName(pair, interval, year, month)
}

func createFolderIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		return err
	}
	return os.MkdirAll(path, os.ModePerm)
}

func checkIfFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
