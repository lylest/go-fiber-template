package files

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
)

func UploadFile(context *fiber.Ctx) (uploadError error, data any, message string) {
	var fileUniqueId = context.Query("fileUniqueId")
	form, err := context.MultipartForm()

	if err != nil {
		return err, nil, "Unable to parse form"
	}

	files := form.File["myFile"]

	fmt.Println(files, "files")
	var uploadLoopErr error
	for _, file := range files {
		filename := filepath.Join("./assets/files", fileUniqueId+"-"+file.Filename)

		if err := context.SaveFile(file, filename); err != nil {
			uploadError = err
		}
	}

	if uploadLoopErr != nil {
		return uploadLoopErr, nil, "Failed to upload file"
	}

	return nil, nil, "File uploaded"
}
