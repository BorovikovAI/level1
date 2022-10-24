// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
// }

// func main() {
//  someFunc()
// }

// 1. В данном случае глобальная переменная изменяется в функции, что не хорошо,
//		например, если вдруг программа станет пакетом и будет использоваться несколькими пользователями,
//		то из-за наличия перезаписываемой глобальной переменной будут проблемы.
// 2. Сборщик мусора также не будет корректно работать, так как переменные внутри функций удаляются после
//		завершения работы функции, а глобальные - нет.
// 3. Вероятно, лучше использовать руны, чтобы не возникло проблем с кодировкой.

package main

import (
	"fmt"
)

func someFunc() {
	var justStringOfSamurai string
	var number int = 100

	YamamotoKsunatomo := "雄弁の芸術では、主なものは黙っている能力です。 いくつかのケースでは、あなたが話すことなく行うことができると思われる場合は、言葉を言わずに仕事をしてください。 あなたは言葉がいくつかのケースで言わなければならないことがわかった場合は、簡単かつ明確に話します"
	fmt.Println(YamamotoKsunatomo, len(YamamotoKsunatomo))

	slice := YamamotoKsunatomo[:number]
	YamamotoKsunatomoCopy := make([]rune, len([]rune(slice)), len([]rune(slice)))
	copy(YamamotoKsunatomoCopy, []rune(slice))

	justStringOfSamurai = string(YamamotoKsunatomoCopy)
	fmt.Println(justStringOfSamurai, len(justStringOfSamurai))
}

func main() {
	someFunc()
}
