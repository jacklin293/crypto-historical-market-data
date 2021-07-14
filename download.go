package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func downloadFiles(pair string, interval string, year string, month string) (err error) {
	if strings.Contains(month, "-") {
		mRange := getMonthRange(month)
		for i := mRange[0]; i <= mRange[1]; i++ {
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

func downloadKlineFile(pair string, interval string, year string, month string) (err error) {
	fileName := getZipFileName(pair, interval, year, month)
	filePath := getZipFilePath(pair, interval, year, month)
	downloadUrl := fmt.Sprintf("%s/%s/%s/%s", BINANCE_PUBLIC_DATA_URL, pair, interval, fileName)

	// Create folder if not exists
	folderPath := getZipFolderPath(pair, interval, year)
	if err = createFolderIfNotExists(folderPath); err != nil {
		return
	}

	// Check if file exists
	if checkIfFileExists(filePath) {
		fmt.Printf(" - File '%s' has been downloaded already\n", filePath)
		return
	}

	// Check if file exists
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(" - Failed to download '%s', status code: %d", downloadUrl, resp.StatusCode)
	}

	// Create a file
	out, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer out.Close()

	// Copy body to the file
	if _, err = io.Copy(out, resp.Body); err != nil {
		return
	}
	fmt.Printf(" - File '%s' has been downloaded successfully\n", filePath)

	return
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
