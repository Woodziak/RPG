package main

import "fmt"
import "io/ioutil"
import "os"
import "os/exec"
import "time"
import "bufio"

func CallMenu(b bool){

	Menu:
		for {

			fmt.Println("Co robimy?\n1. Rozpocznij nową grę\n2. Wczytaj grę\n3. Wyjdź z gry")
			var i int


	/*		if b == true {
				fmt.Println("\nWybierz poprawną wartość\n")
			}*/
			fmt.Scanf("%d", &i)

			switch i {
			case 1:
				b=false
				fmt.Println("Wybrałeś nową grę.")
				time.Sleep(time.Second * 1)
				clrscr()
				CharPick(b)

				break Menu
			case 2:
				b=false
				fmt.Println("Trwa wczytywanie...")
				time.Sleep(time.Second * 1)
				clrscr()
				break Menu
			case 3:
				b=false
				fmt.Println("Wychodzę z gry, żegnaj Wędrowcu...")
				time.Sleep(time.Second * 1)
				os.Exit(0)
			default:
				b=true
				clrscr()
			}
		}
		b=false
}

type PlayerAtrib struct{
	PlayerName string
	PlayerSex string
	PlayerStrengh uint8
	PlayerPerception uint8
	PlayerEndurance uint8
	PlayerCharisma uint8
	PlayerIntelligence uint8
	PlayerAgility uint8
	PlayerLuck uint8
}
func (P PlayerAtrib)Player() {
	/*PN:=P.PlayerName
	PSex:=P.PlayerSex
	PS:=P.PlayerStrengh
	PP:=P.PlayerPerception
	PE:=P.PlayerEndurance
	PC:=P.PlayerCharisma
	PI:=P.PlayerIntelligence
	PA:=P.PlayerAgility
	PL:=P.PlayerLuck
*/


}

func PredefinedChars(){
Predef:
		for {
				fmt.Println("Wybierz postać:\n1. Nolan (wyświetl statystyki)\n2. Aika (wyświetl statystyki)\n3. Earl (wyświetl statystyki)")

				var i int
				fmt.Scanf("%d", &i)
				switch i {
				case 1:
					Txt, err := ioutil.ReadFile("txtsrc/Nolan.txt")
					if err != nil{
					panic(err)
					}
					Str := string(Txt)
					fmt.Println(Str)
					break Predef
				case 2:
					Txt, err := ioutil.ReadFile("txtsrc/Aika.txt")
					if err != nil{
					panic(err)
					}
					Str := string(Txt)
					fmt.Println(Str)
					break Predef
				case 3:
					Txt, err := ioutil.ReadFile("txtsrc/Earl.txt")
					if err != nil{
					panic(err)
					}
					Str := string(Txt)
					fmt.Println(Str)
					break Predef
				default:
					//clrscr()
				}


		}
}
func clrscr(){
	cmd:=exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CharPick(b bool){
Chars:
	for {
			fmt.Println("Wybierz:\n1. Predefiniowaną postać \n2. Stwórz własną!\n3. Wróć do menu")
			var i int
	/*		if b==true {
					fmt.Println("\nWybierz poprawną wartość\n")
			}*/
			fmt.Scanf("%d", &i)

			switch i {
			case 1:
				b=false
				fmt.Println("Oto postacie jakie możemy zaproponować.")
				time.Sleep(time.Second * 1)
				PredefinedChars()
				break Chars
			case 2:
				b=false
				fmt.Println("Przechodzę do ekranu tworzenia nowej postaci...")
				time.Sleep(time.Second * 1)
				clrscr()
				break Chars
			case 3:
				b=false
				fmt.Println("Wracam do menu")
				time.Sleep(time.Second * 1)
				clrscr()
				CallMenu(b)
				break Chars

			default:
				b=true
				clrscr()


			}
	}
	b=false
}

func main(){

	clrscr()
	fmt.Println("Radioaktywny opad: Post-nuklearna gra role play")
	fmt.Print("Luźna adaptacja świata i gry Fallout produkcji Interplay... Wszystkie prawa zachowane. Mam nadzieję.\n\n")
	fmt.Println("Naciśnij ENTER, aby rozpocząć.")
	//tutaj wstawić War, war never changes

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clrscr()
	b:=false
	CallMenu(b)
	b=false
		//funkcja ma za zadanie wybrac postac gracza


	fmt.Println("Kończymy. Wojna się nie zmieniła")
	bufio.NewReader(os.Stdin).ReadBytes('\n')




}
