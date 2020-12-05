package pdfgen

import (
	"os/exec"

	"github.com/nguyenthenguyen/docx"
)

type variables struct {
	variables interface{}
}
func saveDoc(vars interface{}) {
	println(vars)
}
func test() {
		// Read from docx file
		r, err := docx.ReadDocxFile("./Template.docx")
		// Or read from memory
		// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)
		if err != nil {
			panic(err)
		}
		docx1 := r.Editable()
		// Replace like https://golang.org/pkg/strings/#Replace
		docx1.Replace("{{.PG1Name}}", "Hello Upwork", -1)
		
		docx1.Replace("{{.Date}}", "5/12/2020", -1)
		docx1.WriteToFile("./temptest.docx")
	
		// Or write to ioWriter
		// docx2.Write(ioWriter io.Writer)
	
		r.Close()
		
		arg0 := "lowriter"
		arg1 := "--invisible" //This command is optional, it will help to disable the splash screen of LibreOffice.
		arg2 := "--convert-to"
		arg3 := "pdf:writer_pdf_Export"
		path := "./temptest.docx"
		exec.Command(arg0,arg1,arg2,arg3,path).Output()
		
		//docx1.ReplaceHeader("out with the old", "in with the new")
		//docx1.ReplaceFooter("Change This Footer", "new footer")
		//docx1.WriteToFile("./new_result_1.docx")
	
		//docx2 := r.Editable()
		//docx2.Replace("old_2_1", "new_2_1", -1)
		//docx2.Replace("old_2_2", "new_2_2", -1)
		//docx2.WriteToFile("./new_result_2.docx")
	
		// Or write to ioWriter
		// docx2.Write(ioWriter io.Writer)
	
		r.Close()	
}