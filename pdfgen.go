package pdfgen

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"reflect"

	"github.com/nguyenthenguyen/docx"
)

// Write to IO Writer
func Write(template string, data interface{}, w io.Writer) {
	Save(template, data, "./temp.pdf")

    pdf, err := ioutil.ReadFile("./temp.pdf")
	if (err != nil) {
		log.Fatal(err)
	}     
	w.Write(pdf)
	os.Remove("./temp.pdf")
}

// Save to outpath
func Save(template string, data interface{}, outpath string) {
	// Fix document
	arg0 := "lowriter"
	arg1 := "--invisible" //This command is optional, it will help to disable the splash screen of LibreOffice.
	arg2 := "--convert-to"
	arg3 := "docx:MS Word 2007 XML"
	arg4 := "--outdir"
	arg5 := "./temp"
	dpath := template
	_, err := exec.Command(arg0,arg1,arg2,arg3, arg4, arg5 ,dpath).CombinedOutput()
	if err != nil {
        log.Fatal(err)
	}

	processDocument("./temp/"+path.Base(template), data)

	// Convert to pdf
	arg0 = "lowriter"
	arg1 = "--invisible" //This command is optional, it will help to disable the splash screen of LibreOffice.
	arg2 = "--convert-to"
	arg3 = "pdf:writer_pdf_Export"
	arg4 = "--outdir"
	arg5 = path.Dir(outpath)
	dpath = "./temp/temp.docx"
	_, err = exec.Command(arg0,arg1,arg2,arg3, arg4, arg5 ,dpath).Output()
	if (err != nil) {
		log.Fatal(err)
	} 
	os.Rename(arg5 + "/temp.pdf", outpath)
	os.RemoveAll("./temp")
}

func processDocument(tpath string, data interface{}) {
	fields := reflect.TypeOf(data)
	values := reflect.ValueOf(data)
	num := fields.NumField()

	// Read from docx file
	r, err := docx.ReadDocxFile(tpath)
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	// loop over the fields in the provided struct and modify
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if (value.Kind() == reflect.String) {
			val := value.String()
			variable := "var#" + field.Name
			docx1.ReplaceRaw(variable, val, -1)

		}

	}
	
	docx1.WriteToFile(path.Dir(tpath) + "/temp.docx")
	r.Close()
	
}

