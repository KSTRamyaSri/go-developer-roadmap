package main;
import "fmt";
import "github.com/kstramyasri/go-developer-roadmap/01-fundamentals/fundamentals"


func main(){
	fmt.Println("Hello, World!");
	fmt.Println(investmentCalc(1000, 0.05, 1));
	fmt.Println("Global Variable:", fundamentals.GlobalVar);
	fundamentals.variables();
}