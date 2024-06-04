package iteration

func Repeat(c string, multiplier int) (repeated string) {
    for i := 0; i < multiplier; i++ {
        repeated += c
    }
    return
}
