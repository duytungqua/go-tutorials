package structure

/* 
interface definning set of methods signatures, interface is a powerful feature to achive polymorphism
Key featueres of interfaces: 
1. Methods: Interfaces define a set of method that a type must implement
2. Dynamic Typing: Interfaces allow for dynamic typing, meaning that a variable of an interface can hold values 
3. Empty Interface: The empty interface `interface{}` can hold values of any type  var anyValue interface{}
4. Type Assertion: You can use type assertion to retrieve the underlying value
*/

type Greeter interface{
	greeter();
}
//how to implement interface
type People struct {
	Name string
	Age  int
}		

func (p People) Greet() string { // Value receiver
	return "Hello, " + p.Name
}

//value receiver vs pointer receiver
func (p People) GreetPointer() string { // Pointer receiver
	return "Hello, " + p.Name
}	

func (p *People) UpdateName(newName string) {
	p.Name = newName
}	
