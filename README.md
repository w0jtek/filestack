# README

This aplication is a webservice which allows basic image manipulation:

- removing exif metadata
- cropping

In order to start the application, go to subdirectory:
`src/services/api` and type command `go run main.go`. The webserver should be launched and start listening on port `10000`.

The application provides one endpoint, which accepts json payloads sent using POST method. It is available under `http://localhost:10000/accept`

# Transformations

Available transformation commands:

- exif (removes exif metadata)
- crop (applies cropping - see examples below)

# Example payloads

The most basic payload will fetch and return the original image without any conversions.

```
{
	"imageUrl": "https://www.exiv2.org/include/img_1771.jpg",
	"transformations": []
}
```

Adding `exif` command will remove the metadata and create a copy of the original image:

```
{
	"imageUrl": "https://www.exiv2.org/include/img_1771.jpg",
	"transformations": ["exif"]
}
```

Cropping is possible with `crop` command and it requires four parameters (x, y, width, height) - so the starting point and the size of the subimage:

```
{
	"imageUrl": "https://www.exiv2.org/include/img_1771.jpg",
	"transformations": ["crop=60,60,400,300"]
}
```

Transformations can be defined as a sequence and each one takes an image from the previous result:

```
{
	"imageUrl": "https://www.exiv2.org/include/img_1771.jpg",
	"transformations": [
        "crop=60,60,400,300",
        "crop=100,100,300,200"
    ]
}
```
