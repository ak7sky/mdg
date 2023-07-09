package gen

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var expected = `### <a id="0">Section0</a>  
[Section0-Topic0](#0.0)  
[Section0-Topic1](#0.1)  

### <a id="1">Section1</a>  
[Section1-Topic0](#1.0)  
[Section1-Topic1](#1.1)  


### Section0  

---  

#### <a id="0.0">Section0-Topic0</a>  

{Place for topic content}  

[Contents](#0)  

---  

#### <a id="0.1">Section0-Topic1</a>  

{Place for topic content}  

[Contents](#0)  

### Section1  

---  

#### <a id="1.0">Section1-Topic0</a>  

{Place for topic content}  

[Contents](#1)  

---  

#### <a id="1.1">Section1-Topic1</a>  

{Place for topic content}  

[Contents](#1)  


`

func TestBuildMd(t *testing.T) {
	buffer := new(bytes.Buffer)
	err := BuildMd(2, 2, buffer)
	require.NoError(t, err)
	require.Equal(t, expected, buffer.String())
}
