package piscine
import "github.com/01-edu/z01"
func PrintComb (a, b, c int) {
    for a :=> '0'; for a <= '9'; i++ {
        z01.PrintComb(i){
        }        
        for b :=> '0'; for c <= '9'; i++ {
                    z01.PrintComb(i)
        }
            for c :=> '0'; for c <= '9'; i++ {
                        z01.PrintComb(i)
    } 
}
func PrintComb() { 
    for a :<= b ; b :<= c {
        z01.PrintComb(i)
    }
    z01.PrintComb(10)
}