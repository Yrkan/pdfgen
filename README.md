# pdfgen

A library to process docx files and convert them to pdf

## Dependencies

Uses loconvert, so having libreoffice installed is required

## Installation

```
$ go get github.com/Yrkan/pdfgen

```

## Usage

## pdfgen.Save()

Process a docx template and saves it to outpath

```go
$ pddgen.Save(template string, data interface{}, outpath string)
```

#### Params

- **template (string):** path to the docx template
- **data (struct):** a struct instance with variables names and vales (see example bellow)
- **outpath (string):** path for saving the generated pdf

#### Example

```go
package  main
import  (
	"os"
	"github.com/Yrkan/pdfgen"
)

type Data struct {
Date string
}

func  main()  {
	t2 := Data{Date: "11/11/1111"}
	pdfgen.Save("./input/Template.docx", t2 "./output/conv.pdf")
}
```

## pdfgen.Write()

Process a docx template and writes to io write

```go
$ Write(template string, data interface{}, w io.Writer)
```

#### Params

- **template (string):** path to the docx template
- **data (struct):** a struct instance with variables names and vales (see example bellow)
- **w (string):** IO writer instance to write to

#### Example

```go
package  main
import  (
	"os"
	"github.com/Yrkan/pdfgen"
)

type Data struct {
Date string
}

func  main()  {
	f, _ := os.Create("./test.pdf")
	t1 := Data{Date:  "11/11/1111"}
	pdfgen.Write("./input/Template.docx", t1, f)
	f.Close()
}
```
