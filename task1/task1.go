package main

import ("fmt")


type Human struct{
	Age int
	Name string
	Surname string
}//Declare type Human


func (h Human) Personality() string{
	return fmt.Sprintf("Name - %s, Surname - %s, Age - %d", h.Name, h.Surname, h.Age)
}//Declare Human function
func (h Human) Human_id() string{
	return fmt.Sprintf("I am Human")
}//Declare Human function

type Action struct{
	Done bool
	Human
}//Declare type Action

func (a Action) is_Done() bool{
	return a.Done

}//Declare Action function

func (a *Action) Do_action(repeat int){
	a.Done = true
	fmt.Printf("Action done %d time\n", repeat)
}//Declare Action function

func main(){
	h := Human{
		Age : 22,
		Name : "Alexandr",
		Surname : "Pukha",
	}//Create Human  instance

	a := Action{
		Done : false,
		Human : h,
	}//Create Action instance
	fmt.Println(a.Human_id())//Use Human function from Action
	fmt.Println(a.Personality())//Use Human function from Action

	a.Do_action(1)//Use Action function
	fmt.Println(a.is_Done())//Use Action function
}
