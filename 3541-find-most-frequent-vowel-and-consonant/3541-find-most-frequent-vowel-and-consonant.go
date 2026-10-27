func maxFreqSum(s string) int {
    vowel := make(map[rune]int)
    consonant := make(map[rune]int)

    for _, c := range s {
        if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
            vowel[c]++
        }else {
            consonant[c]++
        }
    }

    a := 0
    b := 0

    for _, v := range vowel {
        if a <= v {
            a = v
        }
    }

    for _, v := range consonant {
        if b <= v {
            b = v
        }
    }

    return a + b
}