package controllers

import "github.com/revel/revel"
import "log"
import "os"
import "io"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Upload() revel.Result {
    //update this var with absolute path of your uploads dir. Make sure your dir is present and the
    //path string has trailing slash -/
    upload_dir := "/Users/svalleru/Desktop/golang/src/fileuploadapp/uploads/" 
    m := c.Request.MultipartForm
    var msg string
    for fname, _ := range m.File {

		fheaders := m.File[fname]
		for i, _ := range fheaders {
			//for each fileheader, get a handle to the actual file
			file, err := fheaders[i].Open()
			defer file.Close() //close the source file handle on function return
			if err != nil {
			   log.Print(err)
			   msg = "upload failed.."
			}
			//create destination file making sure the path is writeable.
			dst_path := upload_dir + fheaders[i].Filename
			dst, err := os.Create(dst_path)
			defer dst.Close() //close the destination file handle on function return
			defer os.Chmod(dst_path, (os.FileMode)(0644)) //limit access restrictions
			if err != nil {
				log.Print(err)
				msg = "upload failed.."
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				log.Print(err)
				msg = "upload failed.."
			}
		}
		//display success message.
		log.Print(fname, "upload successful..")
		msg = "upload successful.."
	}
    return c.Render(msg)

}

