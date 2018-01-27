package equality

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	equality,
}

var equality = &typewriter.Template{
	Name: "Equality",
	Text: `
{{if .Pointer}}
func (this *{{.Name}}) Equals (other interface{}) bool {
	if other == nil || this == nil {
	    return this == other
	}
	
	if other, ok := other.(*{{.Name}}); ok {
	    if this == other {
	        return true
	    }

	    if *this == *other {
	        return true
	    }
	} 
	
	if other, ok := other.({{.Name}}); ok {
	    if this == &other {
	        return true
	    }

	    if *this == other {
	        return true
	    }
	}
	
	return false
}
{{else}}
func (this {{.Name}}) Equals (other interface{}) bool {
	if other == nil {
	    return false
	}

    if other, ok := other.({{.Name}}); ok {		
        if this == other {
	        return true
	    }
	}
	
	if other, ok := other.(*{{.Name}}); ok {
	    if this == *other {
	        return true
	    }
	}

	return false
}	
{{end}}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}

