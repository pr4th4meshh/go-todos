package main

func main() {
	todos := Todos{}
	todos.add("Buy groceries")
	todos.add("Start studying")

	todos.toggleCompleted(0)
	todos.print()
}
