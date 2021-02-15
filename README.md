# README

This aplication is a webservice which allows basic image manipulation:

- removing exif metadata
- cropping
- rotation

In order to start the application, go to subdirectory:
`src/services/api` and type command `go run main.go`. The webserver should be launched and start listening on port `10000`.

The application provides one endpoint, which accepts json payloads sent using POST method. It is available under `http://localhost:10000/accept`

# Transformations

Available transformation commands:

- exif (removes exif metadata)
- crop (applies cropping - see examples below)
- rotate (rotates image by multiple of 90)

# Example payloads

The payload should include at least one conversion. The conversions are performed from the source image and the collection of converted files will be returned as one zip file.

Example payload including all possible transformations:
```
{
	"imageUrl": "https://www.exiv2.org/include/img_1771.jpg",
	"transformations": ["exif", "crop=0,0,200,80", "rotate=270"]
}
```