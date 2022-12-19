package display

var bitmapImages map[string]IBitmapImage

func GetBitmapImage(name string, filePath string) IBitmapImage {
	if bitmapImages == nil {
		bitmapImages = make(map[string]IBitmapImage)
	}

	_, isExist := bitmapImages[name]

	if !isExist {
		bitmapImages[name] = newBitmapImage(name, filePath)
	}

	return bitmapImages[name]
}
