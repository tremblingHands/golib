package mystring

func Palindrome(s string) string {
        l := len(s)
        t := make([]byte, l)
        copy(t, s)
        for i ,j := 0, l-1; i < j;{
                t[j] = t[i]
                i++
                j--
        }
        return string(t)
}
