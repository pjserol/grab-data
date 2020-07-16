# Aeropoint

Command line tool that given a base station ID, start time and end time downloads the day blocks from a given time and merges them into a single RINEX observation file.

One of these networks in the NOAA CORS network in the United States. RINEX observation data from each of their base stations is published in 1 hour blocks on their FTP server at 'ftp://www.ngs.noaa.gov/cors/rinex'. You can read more about the network here: 'http://geodesy.noaa.gov/CORS/'

This program has 3 consumers to downloads the files from the ftp server to improve the performance.

## Requirements (installed)

- gunzip
- teqc 'https://www.unavco.org/software/data-processing/teqc/teqc.html'

## Golang dependencies (file go.mod)

- "github.com/jlaffaye/ftp"

## Environment Variables

- GUNZIP_CMD
- TEQC_CMD  

export env GUNZIP_CMD="gunzip"  
export env TEQC_CMD="teqc"

## Test locally

- go test -tags integration
- go test -v

## Run the app

- go run . nypb 2020-07-12T23:11:22Z 2020-07-14T01:33:44Z

## Improvements

- Download the file by 1 hour blocks when available instead of downloading only the files of the full day
- Respect order of the file download from the server when running the command teqc
- Add files processed in a report file (success or not)

## Limit

Number of consumers can't be higher (server ftp limit the number of simultaneous connections)
