# gcsdl

Simple way to download objects from Google Cloud Storage with environments default credentials.

Motivation behind the software is to have small and standalone binary instead of installing full Google Cloud SDK in order to download objects from storage.

## Configuration

[Cloud Storage Authentication](https://cloud.google.com/storage/docs/authentication#libauth)

If you run this tool outside of Google Cloud where the credentials are already set in the environment, you just have to point environment variable `GOOGLE_APPLICATION_CREDENTIALS` to service accounts key.

## Usage

    gcsdl gs://my-bucket/my/object/path /file/destination
    
## Author

Jussi Lindfors
