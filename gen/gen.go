package gen

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

const mdTmpl = `{{range $sIdx, $s := .}}## <a id="{{$sIdx}}">Section{{$sIdx}}</a>  
{{range $tIdx, $t := .Topics}}[Section{{$sIdx}}-Topic{{$tIdx}}](#{{$sIdx}}.{{$tIdx}})  
{{end}}
{{end}}
{{range $sIdx, $s := .}}## Section{{$sIdx}}  

{{range $tIdx, $t := .Topics}}---  

### <a id="{{$sIdx}}.{{$tIdx}}">Section{{$sIdx}}-Topic{{$tIdx}}</a>  

{Place for topic content}  

[Contents](#{{$sIdx}})  

{{end}}{{end}}
`

type Section struct {
	Topics []struct{}
}

func GetWriterForMd(file string) (io.Writer, error) {
	wr := os.Stdout
	if file != "" {
		file, err := os.Create(fmt.Sprintf("%s.md", file))
		if err != nil {
			return nil, err
		}
		wr = file
	}
	return wr, nil
}

func BuildMd(nOfS int, nOfT int, wr io.Writer) error {
	tmpl := template.Must(template.New("").Parse(mdTmpl))
	topics := make([]struct{}, nOfT)
	sections := make([]Section, nOfS)

	for i := range sections {
		sections[i] = Section{Topics: topics}
	}

	return tmpl.Execute(wr, sections)
}
