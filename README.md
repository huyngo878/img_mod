# image processing in golang

#### This GoLang project provides a set of tools to perform basic image processing tasks. It's designed to be modular and easy to expand upon.

## Modules:

### 1. GetPic

#### Downloads an image from a given URL.

#### Sample usage: GetPic.DownloadImage("https://example.com/image.jpg", "downloaded_image.jpg")

### 2. Colors

#### Processes an image and prints the RGB values of each pixel.

#### Sample usage: colors.ProcessImage("path_to_image.jpg")

#### Downloads an image from a given URL. Sample usage: GetPic.DownloadImage("https://example.com/image.jpg", "downloaded_image.jpg")

### 3. Grayscale

#### Converts an image to grayscale.

#### Sample usage: grayscale.ConvertToGrayscale("path_to_image.jpg", "output_image.png")

### 4. Text

#### Adds text to an image using the GoFont Regular font.

#### Sample usage: text.AddText("path_to_image.png")
