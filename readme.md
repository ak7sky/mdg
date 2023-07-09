# Tool to generate MD template

Instead of a thousand words...  

```
Tool to generate MD template

Usage:
  mdg [flags]

Flags:
  -h, --help            help for mdg
  -s, --nofs int        number of sections in MD template (default 1)
  -t, --noft int        number of topics in each section (default 1)
  -f, --tofile string   filename to write result (if no file is specified, the MD template will be written to stdout)

```

Example:  
`mdg -s=2 -t=5 -f=s2t5smpl`  

Result: [s2t5smpl](s2t5smpl.md)  


