package main
import (
        "fmt"
        "os"
        "bufio"
        "regexp"
)

func StringsFromFile(fpath string) []string {
        f, _ := os.Open(fpath)
        defer f.Close()

        var s []string
        sc := bufio.NewScanner(f)
        for sc.Scan() {
                s = append(s, sc.Text())
        }

        return s
}

func CharTuples() []string {
        a := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

        s := []string{}
        for j := 0 ; j < len(a) - 1 ; j++ {
                for i := j + 1 ; i < len(a) ; i++ {
                        s = append(s, a[j] + a[i], a[i] + a[j])
                }
        }

        return s
}

func MatchTuples(tuples []string, words []string) []string {
        s := []string{}

        for j := 0 ; j < len(tuples) ; j++ {
                re := regexp.MustCompile(tuples[j])
                for i := 0 ; i < len(words) ; i++ {
                        if re.MatchString(words[i]) {
                                s = append(s, tuples[j])
                                break
                        }
                }
        }
        return s
}

func main() {
        str := MatchTuples(CharTuples(), StringsFromFile("corncob.dat"))
        for i := 0 ; i < len(str) ; i++ {
                fmt.Println(str[i])
        }
        fmt.Println(len(str))
}
