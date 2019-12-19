package main

import (
	"fmt"
	"github.com/linzhoulxyz/unioffice/document"
	"log"
	"reflect"
	"strconv"
)

func main() {
	parseDocxFile()
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			//fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// parse docx file
func parseDocxFile() {
	doc, err := document.Open("demo1.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}
	for _, para := range doc.Paragraphs() {
		fmt.Println("================================")
		ppr := para.Properties()
		px := ppr.X()
		display("", reflect.ValueOf(px))

		for _, run := range para.Runs() {
			//fmt.Print(run.Text())
			fmt.Println("--------------------------------")
			rpr := run.Properties()
			x := rpr.X()
			//fmt.Print(rpr, x)
			display(run.Text(), reflect.ValueOf(x))
			fmt.Println("Bold:", rpr.IsBold(), rpr.BoldValue(), "Italic:", rpr.IsItalic(), rpr.ItalicValue(), "Color:", rpr.Color().X().ValAttr)

			//l := run.DrawingInline()
			//if len(l) > 0 {
			//	//fmt.Println(len(l))
			//	for _, i := range l {
			//		if ref, ok := i.GetImage(); ok {
			//			imgPath := ref.Path()
			//			//fmt.Println(imgPath)
			//			f, err := os.Open(imgPath)
			//			defer f.Close()
			//			if err == nil {
			//				imgData, err1 := ioutil.ReadAll(f)
			//				if err1 == nil {
			//					//fmt.Println(imgData)
			//					// save to file
			//					imgName := filepath.Base(imgPath)
			//					tmpf, tmpe := os.Create(imgName)
			//					defer tmpf.Close()
			//					if tmpe == nil {
			//						tmpf.Write(imgData)
			//						fmt.Println("[IMG] ", imgName)
			//					} else {
			//						log.Printf("save tmp image: %v fail\n", imgName)
			//					}
			//				} else {
			//					log.Printf("read image: %v fail\n", imgPath)
			//				}
			//			}
			//		}
			//	}
			//}

		}
		fmt.Println()
	}
}

// cut paper
func cutPaper() {

}