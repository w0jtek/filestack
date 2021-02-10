package transformer

// Exif removes exif metadata
type Exif struct {
}

// NewExif constructor
func NewExif() *Exif {
	return &Exif{}
}

// Handle removes exif metadata
func (t *Exif) Handle(localPath string) (err error) {
	img, imgType, err := DecodeImage(localPath)
	// rewriting image does the job of removing exif metadata
	return RewriteImage(img, imgType, localPath)
}
