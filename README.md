# go-picturebook-flickr

Create a PDF file (a "picturebook") from a folder containing images, with support for images exported from Flickr.

## Important

Work in progress. Documentation to follow.

## Tools

To build binary versions of these tools run the `cli` Makefile target. For example:

```
$> make cli
go build -mod vendor -o bin/picturebook cmd/picturebook/main.go
```

### picturebook

Create a PDF file (a "picturebook") from a folder containing images, with support for images exported from Flickr.

```
> ./bin/picturebook -h
  -border float
    	The size of the border around images. (default 0.01)
  -caption string
    	A valid caption.Caption URI. Valid schemes are: filename, flickr, none, orthis
  -debug
    	DEPRECATED: Please use the -verbose flag instead.
  -dpi float
    	The DPI (dots per inch) resolution for your picturebook. (default 150)
  -exclude value
    	A valid regular expression to use for testing whether a file should be excluded from your picturebook. DEPRECATED: Please use -filter regexp://exclude/?pattern={REGULAR_EXPRESSION} flag instead.
  -filename string
    	The filename (path) for your picturebook. (default "picturebook.pdf")
  -fill-page
    	If necessary rotate image 90 degrees to use the most available page space.
  -filter value
    	A valid filter.Filter URI. Valid schemes are: any, flickr, orthis, regexp
  -height float
    	A custom width to use as the size of your picturebook. Units are currently defined in inches. This fs.overrides the -size fs. (default 11)
  -include value
    	A valid regular expression to use for testing whether a file should be included in your picturebook. DEPRECATED: Please use -filter regexp://include/?pattern={REGULAR_EXPRESSION} flag instead.
  -ocra-font
    	Use an OCR-compatible font for captions.
  -orientation string
    	The orientation of your picturebook. Valid orientations are: 'P' and 'L' for portrait and landscape mode respectively. (default "P")
  -pre-process value
    	DEPRECATED: Please use -process {PROCESS_NAME}:// flag instead.
  -process value
    	A valid process.Process URI. Valid schemes are: halftone, null, rotate
  -size string
    	A common paper size to use for the size of your picturebook. Valid sizes are: [please write me] (default "letter")
  -sort string
    	A valid sort.Sorter URI. Valid schemes are: modtime, orthis
  -source-uri string
    	A valid GoCloud blob URI to specify where files should be read from. By default file:// URIs are supported.
  -target string
    	Valid targets are: cooperhewitt; flickr; orthis. If defined this flag will set the -filter and -caption flags accordingly. DEPRECATED: Please use specific -filter and -caption flags as needed.
  -target-uri string
    	A valid GoCloud blob URI to specify where your final picturebook PDF file should be written to. By default file:// URIs are supported.
  -verbose
    	Display verbose output as the picturebook is created.
  -width float
    	A custom height to use as the size of your picturebook. Units are currently defined in inches. This fs.overrides the -size fs. (default 8.5)
```

### Notes

The `picturebook` tool does not download photos from Flickr. That is left to another process to do using the [Flickr API](https://www.flickr.com/services/api/).

It also makes the following assumptions about file names and file structures:

* The original photo has been downloaded. This is the photo with the `_o.{EXTENSION}` suffix.
* That the results of the [flickr.photos.getInfo](https://www.flickr.com/services/api/flickr.photos.getInfo.html) API method have been captured and stored in a file with the same ID and secret as the original photo, ending in `_i.json`, in the same directory as the original photo itself.

For example:

```
$> ls -la 168662172
total 160
-rw-r--r--  1 user  staff   1342 Jul 21  2013 7168662172_7464be4db8_i.json
-rw-r--r--  1 user  staff  23111 May 11  2012 7168662172_7464be4db8_o.jpg
-rw-r--r--  1 user  staff   4078 Jul 21  2013 7168662172_b8018fcd92_t.jpg
-rw-r--r--  1 user  staff  42735 May 11  2012 7168662172_b8018fcd92_z.jpg
```

## See also

* https://github.com/aaronland/go-picturebook
* https://www.flickr.com/services/api/