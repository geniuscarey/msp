
// Different usages of variables
//  DECLARATIONS
//  DEFINITIONS
//  ASSIGNMENTS
//  REFERENCES

package ast

import "fmt"

// ====================
// Variable declaration:: var a;
// ====================
type Declaration struct {
  Var *Variable
}

func (d Declaration) Execute() interface{} {
  symbol.Declare(d.Var.VariableName)
  return d.Var.VariableName
}

func (d Declaration) Print(p string) {
  fmt.Println(p + "Declaration")
  d.Var.Print(p + "| ")
}

// ====================
// Variable Definition:: var a = 1
// Definitions are essentially just typedefed assignments in this language,
// But the Execute() function is different
// ====================
type Definition struct {
  Decl *Declaration
  AssignNode *Assignment
}

func (d Definition) Execute() interface{} {
  // TODO Declare Variable
  // TODO Assign Variable
  return nil
}

func (d Definition) Print(p string) {
  fmt.Println(p + "Definition")
  d.Decl.Print(p + "| ")
  d.AssignNode.Print(p + "| ")
}

// ====================
// Equals, Assignment:: a    =  1
//                      LHS     RHS
// ====================
type Assignment struct {
  Lhs *Variable
  Rhs Node
}

func (a Assignment) Execute() interface{} {
  // Reset the value at lhs to be rhs
  // Return nothing
  return nil
}

func (a Assignment) Print(p string) {
  fmt.Println(p + "Assign")
  a.Lhs.Print(p + "| ")
  a.Rhs.Print(p + "| ")
}

// ====================
// Variable reference:: var something = myvar;
// ====================
type Reference struct {
  Var *Variable
  Value interface{}
}

func (vr VarReference) Execute() interface{} {
  // TODO: GET Value of variable
  return vr.Value
}

func (vr VarReference) Print(p string) {
  switch vr.Value.(type) {
    case int:
      fmt.Println(p + vr.Var.VariableName + "=" + string(vr.Value.(int)))
      break
    case string:
      fmt.Println(p + vr.Var.VariableName + "=" + vr.Value.(string))
      break
  }
}
