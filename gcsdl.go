package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: gcsdl gs://storage-name/object /file/destination")
	}
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	sourceURL := os.Args[1]
	bucket, object, err := parseSource(sourceURL)
	if err != nil {
		fmt.Println("Source error:", err)
		os.Exit(1)
	}
	destination := os.Args[2]

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = download(client, bucket, object, destination)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Downloaded %s to %s", sourceURL, destination)
}

func download(client *storage.Client, bucket string, object string, destination string) error {
	fp, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer fp.Close()

	ctx := context.Background()
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = io.Copy(fp, rc)
	if err != nil {
		return err
	}

	return nil
}

func parseSource(rawurl string) (bucket string, object string, err error) {
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return "", "", err
	}
	if url.Scheme != "gs" {
		return "", "", fmt.Errorf("Url scheme is invalid, expected 'gs' got '%s'", url.Scheme)
	}
	if len(url.Path) <= 1 {
		return "", "", errors.New("Url does not contain object")
	}
	return url.Host, url.Path[1:], nil
}
