package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func unzipFiles(pair string, interval string, year string, month string) (err error) {
	if strings.Contains(month, "-") {
		mRange := getMonthRange(month)
		for i := mRange[0]; i <= mRange[1]; i++ {
			if err = unzipKlineFile(pair, interval, year, strconv.Itoa(i)); err != nil {
				return err
			}
		}
	} else {
		if err = unzipKlineFile(pair, interval, year, month); err != nil {
			return err
		}
	}
	return nil
}

func unzipKlineFile(pair string, interval string, year string, month string) error {
	csvFolderPath := getCsvFolderPath(pair, interval, year)
	zipFilePath := getZipFilePath(pair, interval, year, month)
	if err := createFolderIfNotExists(csvFolderPath); err != nil {
		return err
	}
	return unzip(zipFilePath, csvFolderPath)
}

func getCsvFolderPath(pair string, interval string, year string) string {
	return FOLDER_CSV + "/" + fmt.Sprintf("binance-%s-%s-kline-%s", pair, interval, year)
}

func getCsvFileName(pair string, interval string, year string, month string) string {
	return fmt.Sprintf("%s-%s-%s-%02s.csv", pair, interval, year, month)
}

func getCsvFilePath(pair string, interval string, year string, month string) string {
	return getCsvFolderPath(pair, interval, year) + "/" + getCsvFileName(pair, interval, year, month)
}

func unzip(src string, dst string) (err error) {
	archive, err := zip.OpenReader(src)
	if err != nil {
		return
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err = io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
		fmt.Printf(" - File '%s' has been decompressed successfully\n", filePath)
	}
	return
}
