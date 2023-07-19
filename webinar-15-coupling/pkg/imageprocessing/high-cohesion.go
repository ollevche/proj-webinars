package imageprocessing

import "image"

type EditableImage struct {
	original image.Image
}

func NewEditableImage(i image.Image) *EditableImage {
	return &EditableImage{
		original: i,
	}
}

// ConvertToGray applies grayscale filter
func (i *EditableImage) ConvertToGray() {
	// processing logic
}

// Resize resize image to w, h size
func (i *EditableImage) Resize(width, height int) {
	// resizing logic
}

func (i *EditableImage) ApplyBlackAndWhiteFilter() {
	// logic
}

func (i *EditableImage) Finalize() image.Image {
	// defined earlier logic
	return i.original
}
