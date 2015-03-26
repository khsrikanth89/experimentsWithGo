package main

import (
	//"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
)

var xmlData = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<boo>
<command>update</command>
<view>parent</view>
<count>100</count>
<shares>
 <share>
  <vidaoid>me</vidaoid>
  <permissions>
   <read>true</read>
  </permissions>
 </share>
 <share>
  <vidaoid>me</vidaoid>
  <permissions>
   <read>true</read>
  </permissions>
 </share>
</shares>
<description>
  wow
</description>
</boo>
`)

type XMLStuff map[string]interface{}
type xmlShareStuff map[string]interface{}

func (this *xmlShareStuff) UnmarshalXML (decoder *xml.Decoder, starter xml.StartElement) error {

}

func (this *XMLStuff) UnmarshalXML (decoder *xml.Decoder, starter xml.StartElement) error {
   var err error
   name := starter.Name.Local
   var start xml.Token
   start,err = decoder.Token()
   for err != io.EOF {
      fmt.Printf("%v\n",start)
      switch start.(type) {
         case xml.StartElement :
	   key := start.(xml.StartElement).Name.Local
	   fmt.Println("Processing start element")
           switch key {
             case "command","description","view":
	       var value string
	       var newstart xml.StartElement = start.(xml.StartElement)
	       err=decoder.DecodeElement(&value,&newstart)
	       fmt.Printf("%v\n",value)
	       (*this)[key]=value
	     case "count","offset":
	       var value int
	       var newstart xml.StartElement = start.(xml.StartElement)
	       err=decoder.DecodeElement(&value,&newstart)
	       fmt.Printf("%v\n",value)
	       (*this)[key]=value
   	     case "share":
	       var newstart xml.StartElement = start.(xml.StartElement)
		if err := xml.Unmarshal(xmlData, &s); err != nil {
			panic(err)
		}
	       value = unMarshal(decoder) map[string]interface{}
	       fmt.Printf("%v\n",value)
	       (*this)[key]=value
           }
	 case xml.EndElement :
	   if start.(xml.EndElement).Name.Local == name {
	     err = io.EOF
	   } 
      }
      if err !=io.EOF{
      	start,err = decoder.Token()
      }
   }
   return nil
}

func main() {
	s := make(XMLStuff,0)
	if err := xml.Unmarshal(xmlData, &s); err != nil {
		panic(err)
	}
	fmt.Printf("s: %#v\n", s)



}
