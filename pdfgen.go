package pdfgen

import (
	"log"
	"os/exec"
	"path"
	"reflect"

	"github.com/nguyenthenguyen/docx"
)

// Save tt
func Save(template string, data interface{}) {
	processDocument(template, data)
	arg0 := "lowriter"
	arg1 := "--invisible" //This command is optional, it will help to disable the splash screen of LibreOffice.
	arg2 := "--convert-to"
	arg3 := "pdf:writer_pdf_Export"
	arg4 := "--outdir"
	arg5 := path.Dir(template)
	path := path.Dir(template) + "/temp.docx"
	_, err := exec.Command(arg0,arg1,arg2,arg3, arg4, arg5 ,path).Output()
	if (err != nil) {
		log.Fatal(err)
	} 

}

// SaveDoc Save doc
func processDocument(dpath string, data interface{}) {
	fields := reflect.TypeOf(data)
	values := reflect.ValueOf(data)
	num := fields.NumField()

	// Read from docx file
	r, err := docx.ReadDocxFile(dpath)
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
			variable := "{{." +  field.Name + "}}"
			docx1.Replace(variable, val, -1)

		}

	}
	
	docx1.WriteToFile(path.Dir(dpath) + "/temp.docx")
	r.Close()
	
}
