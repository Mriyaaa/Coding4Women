package practice

//import "fmt"

type Celsius float32
type Fahrenheit float32

func toFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit((t * 9 / 5) + 32)

}

// func main() {
// 	celsiusTemperature := Celsius(30)
// 	fahrenheitTemperature := toFahrenheit(celsiusTemperature)

// 	fmt.Printf("%.2f C is %.2f F\n", celsiusTemperature, fahrenheitTemperature)
// }
