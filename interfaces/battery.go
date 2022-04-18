package main

import "fmt"

type battery struct {
	batteryСharge string
	batteryCh [10]byte
}

func (b *battery) makeBattery() {
	spaceCount := 0

	for _, elem := range b.batteryСharge {
		if elem == '0' {
			b.batteryCh[spaceCount] = ' ';
			spaceCount++
		}
	}
	for ; spaceCount < len(b.batteryCh); spaceCount++ {
		b.batteryCh[spaceCount] = 'X';
	}
}

func (b* battery) String() string {
	return fmt.Sprintf("[%s]", b.batteryCh);
}

func main() {
	var batteryСharge string
	fmt.Scan(&batteryСharge)

	batteryForTest := new(battery)
	batteryForTest.batteryСharge = batteryСharge
	batteryForTest.makeBattery()
	//batteryForTest.batteryCh = make([]byte, 10, 10) // можно на срезе байт
	fmt.Println(batteryForTest)
}



/***** другое решение *****/
/*
	type battery struct{
		level string
		empty int}

	func (batteryForTest battery) String() string{
			return strings.Replace("[XXXXXXXXXX]","X"," ",batteryForTest.empty)}
			
	func main() {
		batteryForTest:=new(battery)
		fmt.Scan(&batteryForTest.level)
		batteryForTest.empty=strings.Count(batteryForTest.level,"0")
}
*/