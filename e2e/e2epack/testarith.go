// Autogenerated file. DO NOT EDIT.
// Version: 0.2.1
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\testArith.sc


[31m ********* Error in testarith :***************[0m


@a int
@b int

c = a+b [31m <<<<<<<<<<<<<<<<< types  and  are mismatched for PLUS [0m
-a*b/a

RETURN c

[31m ********* Error in testarith :***************[0m


@a int
@b int

c = a+b-a [31m <<<<<<<<<<<<<<<<< types  and  are mismatched for MINUS [0m
*b/a

RETURN c

[31m ********* Error in testarith :***************[0m


@a int
@b int

c = a+b-a*b [31m <<<<<<<<<<<<<<<<< types  and  are mismatched for MULTI [0m
/a

RETURN c

[31m ********* Error in testarith :***************[0m


@a int
@b int

c = a+b-a*b/a [31m <<<<<<<<<<<<<<<<< types  and  are mismatched for DIV [0m


RETURN c
import (
	"github.com/xavier268/goscrapper/rt"
)


type Input_testarith struct {
}


type Output_testarith struct {
}


// 
// @a int
// @b int
// 
// c = a+b-a*b/a
// 
// RETURN c
func Do_testarith(_in Input_testarith) (_res []Output_testarith, _err error) {
_res = append(_res, Output_testarith{})
return _res, _err
}
