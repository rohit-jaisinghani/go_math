package calc

import "errors"

func Add(a,b int) (int,error){
	if a!=0{
		return a+b,nil
	}else{
		return 0,errors.New("value of a is zero")
	}

}
