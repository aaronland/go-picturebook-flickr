# go-picturebook-flickr

Create a PDF file (a "picturebook") from a folder containing images, with support for images exported from Flickr.

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
  -bleed float
    	An additional bleed area to add (on all four sides) to the size of your picturebook.
  -border float
    	The size of the border around images. (default 0.01)
  -caption value
    	Zero or more valid caption.Caption URIs. Valid schemes are: exif://, filename://, flickr://, json://, modtime://, multi://, none://.
  -dpi float
    	The DPI (dots per inch) resolution for your picturebook. (default 150)
  -even-only
    	Only include images on even-numbered pages.
  -filename string
    	The filename (path) for your picturebook. (default "picturebook.pdf")
  -fill-page
    	If necessary rotate image 90 degrees to use the most available page space. Note that any '-process' flags involving colour space manipulation will automatically be applied to images after they have been rotated.
  -filter value
    	A valid filter.Filter URI. Valid schemes are: any://, flickr://, regexp://.
  -height float
    	A custom width to use as the size of your picturebook. Units are defined in inches by default. This flag overrides the -size flag when used in combination with the -width flag.
  -margin float
    	The margin around all sides of a page. If non-zero this value will be used to populate all the other -margin-(N) flags.
  -margin-bottom float
    	The margin around the bottom of each page. (default 1)
  -margin-left float
    	The margin around the left-hand side of each page. (default 1)
  -margin-right float
    	The margin around the right-hand side of each page. (default 1)
  -margin-top float
    	The margin around the top of each page. (default 1)
  -max-pages int
    	An optional value to indicate that a picturebook should not exceed this number of pages
  -ocra-font
    	Use an OCR-compatible font for captions.
  -odd-only
    	Only include images on odd-numbered pages.
  -orientation string
    	The orientation of your picturebook. Valid orientations are: 'P' and 'L' for portrait and landscape mode respectively. (default "P")
  -process value
    	A valid process.Process URI. Valid schemes are: colorspace://, colourspace://, contour://, halftone://, null://, rotate://.
  -progress-monitor-uri string
    	A registered aaronland/go-picturebook/progress.Monitor URI (default "progressbar://")
  -size string
    	A common paper size to use for the size of your picturebook. Valid sizes are: "a3", "a4", "a5", "letter", "legal", or "tabloid". (default "letter")
  -sort string
    	A valid sort.Sorter URI. Valid schemes are: exif://, modtime://.
  -source-uri string
    	A valid GoCloud blob URI to specify where files should be read from. Available schemes are: file://, flickr://, flickrapi://. If no URI scheme is included then the file:// scheme is assumed.
  -target-uri string
    	A valid GoCloud blob URI to specify where files should be read from. Available schemes are: file://. If no URI scheme is included then the file:// scheme is assumed. If empty then the code will try to use the operating system's 'current working directory' where applicable.
  -text string
    	A valid text.Text URI. Valid schemes are: json://.
  -tmpfile-uri string
    	A valid GoCloud blob URI to specify where files should be read from. Available schemes are: file://. If no URI scheme is included then the file:// scheme is assumed. If empty the operating system's temporary directory will be used.
  -units string
    	The unit of measurement to apply to the -height and -width flags. Valid options are inches, millimeters, centimeters (default "inches")
  -verbose
    	Display verbose output as the picturebook is created.
  -width float
    	A custom height to use as the size of your picturebook. Units are defined in inches by default. This flag overrides the -size flag when used in combination with the -height flag.
```

### Notes

The `picturebook` tool does allow for photos to be downloaded from Flickr using the `flickrapi://` source, described below. The _default_ behaviour however is to assume a set of photos already downloaded from Flickr which adhere to the following conventions:

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

## Handlers

The `picturebook` application supports a number of "handlers" for customizing which images are included, how and whether they are transformed before inclusion and how to derive that image's caption.

In addition to the default `picturebook` handlers this package exports the following:

### Buckets (sources)

#### flickrapi://

This handler will gather photos to include in a picturebook from the Flickr API using any of the following Flickr API methods:

* [flickr.galleries.getPhotos](https://www.flickr.com/services/api/flickr.galleries.getPhotos.html)
* [flickr.groups.pools.getPhotos](https://www.flickr.com/services/api/flickr.groups.pools.getPhotos.html)
* [flickr.photos.getContactsPhotos](https://www.flickr.com/services/api/flickr.photos.getContactsPhotos.html)
* [flickr.photos.getContactsPublicPhotos](https://www.flickr.com/services/api/flickr.photos.getContactsPublicPhotos.html)
* [flickr.photos.getWithGeoData](https://www.flickr.com/services/api/flickr.photos.getWithGeoData.html)
* [flickr.photos.getWithoutGeoData](https://www.flickr.com/services/api/flickr.photos.getWithoutGeoData.html)
* [flickr.people.getPhotos](https://www.flickr.com/services/api/flickr.people.getPhotos.html)
* [flickr.people.getPhotosOf](https://www.flickr.com/services/api/flickr.people.getPhotosOf.html)
* [flickr.photos.search](https://www.flickr.com/services/api/flickr.photos.search.html)
* [flickr.photosets.getPhotos](https://www.flickr.com/services/api/flickr.photosets.getPhotos.html)

For example:

```
$> ./bin/picturebook \
	-source-uri 'flickrapi://?client_uri=oauth1%3A%2F%2F%3Fconsumer_key%3D...%26consumer_secret%3D...%26oauth_token%3D...%26oauth_token_secret%3D...'
	-target-uri /usr/local/src/go-picturebook-flickr \
	'method=flickr.photos.search&user_id=35034348999@N01&tags=flickrhq'
```

Or:

```
$> ./bin/picturebook \
	-source-uri 'flickrapi://?client_uri=oauth1%3A%2F%2F%3Fconsumer_key%3D...%26consumer_secret%3D...%26oauth_token%3D...%26oauth_token_secret%3D...'
	-target-uri /usr/local/src/go-picturebook-flickr \
	-filename ca.pdf \
	'method=flickr.photosets.getPhotos&user_id=35034348999@N01&photoset_id=72177720319945125'
```

The convention for Flickr API "bucket" source URIs is:

```
"flickr://?client_uri=" + {GO_FLICKR_API_CLIENT_URI}
```

Where "{GO_FLICKR_API_CLIENT_URI}" is a valid [aaronland/go-flickr-api](https://github.com/aaronland/go-flickr-api) client URI. As of this writing the need to define the entirety of that client URI as a URL-escape strings is a tiresome chore which will be simplified in time.

Rather than passing one or more paths on the local disk to crawl for file to add to your picturebook when using the `flickrapi://` bucket source you pass a series of URL query parameter strings mapping to one or more Flickr API call. For example, to return all of Flickr user [straup's photos tagged flickhq](https://flickr.com/photos/straup/tags/flickrhq) you would do this:

```
method=flickr.photos.search&user_id=35034348999@N01&tags=flickrhq
```

The `flickrapi://` handler works but should still be considered experimental.

### Captions

#### flickr://

This handler will derive the title for a Flickr photo using data stored in a `{PHOTO_ID}_{SECRET}_i.json` file, alongside an image. The data in the file is expected to be the out of a call to the [flickr.photos.getInfo](https://www.flickr.com/services/api/flickr.photos.getInfo.html) API method.

#### flickrapi://

This handler will derive the caption for an image using the [flickr.photos.getInfo]() Flickr API method. Captions will be formated as:

"{TITLE}" (or "untitled")
{PHOTOGRAPHER REALNAME} ({PHOTOGRAPHER_USERNAME})
{PHOTO DATE TAKEN YYYY-MM-DD}

For example, given the Flickr photo [https://live.staticflickr.com/8124/8659998886_aea0810d40_o_d.jpg](https://flickr.com/photos/straup/8659998886) the derived caption would be::

```
(untitled)
Aaron Straup Cope (straup)
2013-04-17
```

The convention for Flickr API "caption" URIs is:

```
"flickr://?client_uri=" + {GO_FLICKR_API_CLIENT_URI}
```

Where "{GO_FLICKR_API_CLIENT_URI}" is a valid [aaronland/go-flickr-api](https://github.com/aaronland/go-flickr-api) client URI. As of this writing the need to define the entirety of that client URI as a URL-escape strings is a tiresome chore which will be simplified in time.

### Filters

#### flickr://

This handler will ensure that only images whose filename matches `o_\..*$` are included.

## See also

* https://github.com/aaronland/go-picturebook
* https://github.com/aaronland/go-flickr-api
* https://www.flickr.com/services/api/