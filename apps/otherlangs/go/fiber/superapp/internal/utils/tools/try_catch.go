package tools

import (
	"log"
)

type Block struct {
	Try     func()
	Catch   func(interface{})
	Finally func()
  }
  
  type Exception interface{}
  
  func Throw(up interface{}) {
	//panic(up)
  }
  
  func (tcf Block) Do() {
	if tcf.Finally != nil {
  
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
  }
  
func Do_try_catch(do func() ,on_error func()) {
	Block{
		Try: func() { do() },
		Catch: func(e interface{}) { 
			if on_error==nil { log.Println("TRY_CATCH",e)
			}else{ on_error() }
		},
		Finally: func() { //fmt.Println("Finally...")
		}, 	}.Do()
}

func Try_catch(do func()) error { //,on_error func()) err error {
	var err interface{}
	Block{
		Try: func() { do() ; },
		Catch: func(e interface{}) { err=e; }, //log.Println("TRY_CATCH",e)
		Finally: func() {  }, 
	}.Do() //fmt.Println("Finally...")
	switch v := err.(type) { case error: return err.(error); default: _=v; return nil }
} ; var TryCatch=Try_catch

// func main(){
// 	result := []string{}
// 	log.Println("start")
// 	err := Try_catch(func(){ result=[]string{"ok"} })
// 	if err != nil { log.Println(err) }
// 	log.Println(result)
// }