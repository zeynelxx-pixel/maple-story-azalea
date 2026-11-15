package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.1" )

func h0STeYqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOVpqZ2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 683eCHaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyfMSSuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRtSRfbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvbytiI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sh4hK2iFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyJT22XaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LCF4JKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N92UbuHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKPJJzLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3IbRlzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFDy7C6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HnRL0mttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIS9dsHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfyhaQEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVVD8S8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IrfOj8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1sXHA1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFzSNggKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5raCKvZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eI4iXWzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTfr4wOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQwG6NIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMGzz0S9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86eODf41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fO1UTrPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5OrEXTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9hcIepzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kb2ktjXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQNlHABdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcRXG5KsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kx967NqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uR4SVfFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2mg4J1IBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAR499rjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHbTKEqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUXjbcOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIsQ7qMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reCaSBRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVuYOT8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfGMmiViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvGqiR2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGsoRGX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMk0CPCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eJs9I3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySksv0PdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCUulRqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhjZZ5cdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGOiJuWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6jd5eCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HalCuPvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yswZ5gG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sjq2d3rbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WX3Pi4pqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mP4pWTW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfGem41gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUXhqPyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GlHEg0VmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SPKywLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 812PxeVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJIgwCStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tGU8G4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMqViP8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func af80kdZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A92zxiOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFYTQRwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STL2VyfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqI9OQUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k42QILKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qca0NmSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqEADLDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5I9kMrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUgljhYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZy1gswLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etBNAFlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0KHVckzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3EatFSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hF0WmvluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVB3XmGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XkTM3zrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPXP7yFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCs1CWrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgdouMKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2nyfDaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CD0LOMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEzgOuehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNirKERgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VweMwR1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yp2XYtTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPVBFAeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFKulTFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ml46MTTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wJjGedlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lp2eIugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsQ5vKnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snlzsmuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2qNnNK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBkRvIQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nk0FULOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5xo7LtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4M7R6AkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQs8uSP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dm95qaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvVc9CHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFNzEPKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PpR9J2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMUooERiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqzARpfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8aWUyZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Db8HcIl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPb9x5YBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yE2ofazCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBQLMGmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVwQayoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3W3b5AYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGjvIHgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oG5d882zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uokZSoEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyJHWPv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u74ZTsMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeSUHxRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5zs2tcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YajZT9AcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19E7edUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kipRNsRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWfz79WHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57Nzmea6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skQHBisrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHMntLqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func az1iEaoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOCTCzT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYgCKbu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUjc68LKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYcKITOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZjRqA91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65v1l54RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7lAm9TbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RFUdq3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIGz86O1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QY87v1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2TW19eiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AptieGBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IVJQN4fvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HW0kZeBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C0d7xYrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wz39573rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1wTkEgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTNZCP93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5J7Bv2u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vpKYc8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSF6AInfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QggXdNy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Opz6Xr2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLTSOYPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkvJh6XQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjOpm7CIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jk16kx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gz6kNbQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucvmi7bAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOdPQG3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEwSsB1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4ZsF0cfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VusxroYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxBKu0mGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVwDz5CVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eifINtH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlDkD4EUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpxYMj0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJkyQq3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ECq4gqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHpgJlXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxvRxKYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSbVuONMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3LezQfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MkIZtys9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7ka4vUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4jWF00nTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYYVFqPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WGKNF70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PP7rSHzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UPlGWVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAs2zUHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ol9Sur2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKj77bZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LNR2PUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mqmEnQPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsU9kVDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LC6HdIuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZhGHCDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QH6nQ4TwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xQLe22uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MU3EYkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFeSi4FVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gqRftQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDNBThxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQE09aE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41ZQUzf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsosoBdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwzHmqOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xbqp6jhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4G4HWSZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNA4KyoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seVLcTURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xkm4T0LQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WLCeNX6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNE5AwC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PNrKVYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uu2du80xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VgzzfBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRWaajWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNc0ysjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMJAEspJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwtN4i3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vJpmJxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hilKcyJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CB0zrSunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsMgxSH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXRYfzUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WS5Cqa9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pslFHMicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lU5INqf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HWqG24VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FkcRea8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsuNN3ztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7a6kNy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouwdrCAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQUoF23wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbTfmXeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8H4hpMUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMhojPgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDzLUl7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtkQ6shOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04SM8eynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyB7d63cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4NXQvSgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IApFg3lcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8p20cgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func abaQpPqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmvnO7KDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrGVnkBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAzc68xUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVdcFcnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2XCkc8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XA5ReFhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjZk7gtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrwyOaqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wly6hYBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PjDtNoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRr8OfeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXpJXEeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqnx0EAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gp5moFJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kr2geoIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95uZuZZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDvY0WcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEaEdjw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGY586cMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNCDnLg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mG6zgQ2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kY8cmYjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwMGpC42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqwge7kZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GifrxAMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92hW7dstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpCma3o0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVjZW7yZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUCBEsIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJHRkEYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tb3yFzC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5iOMD7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWymXpwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXQ3kXMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qzor3it5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyam8jRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezrfqTftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3sPZ7PaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIb6F2NiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 818b82wjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFnXWJbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lmJCnhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRLDdIuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyFmEeWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQ76zQdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wSvhn0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzvMspCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aT8ZDdblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uy662GAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqXhAyhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HX2iRF0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLXULakYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0Q3vROTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPKjepVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qu4YKH77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqaFTBoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gM7BPdYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1j8bG854Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6YQAaIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGuSkDodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8Cs1thZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNG8I4xNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bR7T29hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSoeU7jaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwOYMlHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBCroQNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cJuLcEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0S6qHFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1MR5CpgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIgVkEF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvoOfE3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhTHjHvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYmDXcvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhDLXiuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyASjvnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nV4nE77oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61YytGGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrIJ29KeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaDG0rzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIm4LxHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EloDIQDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWTuH7fFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xp9R1JBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDkhvxK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLS907gOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkqfdVpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lh5qZB24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKpgmMqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGTzPDEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zAcRCMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQUVzOJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdmKR6wpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RasXNTxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vbP3DDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXdw9t5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMAu8uAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uG3ElEoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eE6VvE3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WhqwkWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0f8MdxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnDVhtpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtDJmFB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6Vcqh2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O41qDyxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7nkGuI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLrGZiX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsKSF5XDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBpYoVOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YG3OpUK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6ov8QIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkVShkmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vhs128tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYySbMjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXCistJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcdEPve3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sdsAQQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMRungnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6S4RchC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5sKeOwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKtuPMySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAJKKWibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJxzspnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoaJa6zEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eg4RunrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MmSbKQvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKandtXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZSYd6WbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGAhcLKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3cIieW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHsyrotxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTyBu9qdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipTXPgdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4K5xjHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MzsZaewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Er4ZygDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCEctf2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rrg3qwrgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wf4VQWPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLDRPqJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1K2uuhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yvy6mlJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fhM95P6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ReW4XjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njpkBKJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcqIywtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6W27RfMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzXO4y8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Lm5kUnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1c29O473Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmSZ8Sm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjwNIufMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86y6via1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ddl7s2CRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GZe4uHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9uCKQu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DzcDLTWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtHmWmPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDHc0YtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nShzHGzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxSDWwP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wziwp2olWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXWdnAgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cv0okuQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItWdKjkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEnSCDYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5q6PHSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enSh5gASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwDjoYNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Io0tFcsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5RpqNSWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtL0c415Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CTOEeKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skarTM23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbOUrOvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4uoWiC8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKOKTNZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Jo0XZ7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOsGXTwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpPABdkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yt6iph0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBMDiN0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RpGd74ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvJfbStuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GdWAGENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmcyJFi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2IprXrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88PhocnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cF6GCQdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eONrvYQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLBB9YrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrX8NGhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWuoICOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUc7VsmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3DbyoYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SQOZRYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5pK9Ab0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFJQLAfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZM9dUFQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cwx9ZdRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlYFx0IwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkWMfddPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQFiHLXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wh8Q7F1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nedl7KZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMiJwNgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSi7z1m6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxlGCirpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRKq6sPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfBAoPJmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66FFLPArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9d2vOWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gr47vKkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIwp46SgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0qjW0UHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYU4IwbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pjuqWZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDTw2Kg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGZRxG1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JigEg2YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AO1PQb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cE9nZDsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCqJJMnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLFA2FHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogsoaWG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYiBC3tmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcqDgtuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZEv1prNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPDEiJPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zA4VLKn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PN0PJkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4B5UnyNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAjHyeE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiuBOiEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yga2drQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PISNmdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCERRS22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEWICUiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwv9wQcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPcQe4cpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRX6UMQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KtyTGjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAYZo80TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMVtKn96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eUoM023Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2ERqdOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xZXmkhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVQFP21uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLHQW5dIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNvIjhNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKJvLXjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oo6FjDP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzjbpFzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0o8TAVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTBCXcXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzBxk0zfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27AXzZqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldZ12RNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFDQU8CtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckHqDy41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJDv7iypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVLcgKeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05cGfpopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zj5jmpXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4crWRmmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tp1kULoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvOssKDIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwERBvi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVTLOD54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLIWPgCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkmZepaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpZQ0xBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pavuHD1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Besu9YeqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSsuEqA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBQjncznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0V2NACDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6KDtofcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3NK5iHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aC7abHW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmur0DXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPBS5oJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haxhp95OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmREkKTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func td3bq8CPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dztjR1L3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuOdDCdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OA2h8sXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUzj16YNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEQ0amSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KW1y963Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2qKwTGmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZWhrMugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1YLSLgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEHJ5SsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nc78ybVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQwbOwPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzqQXvMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybECSn8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1i0qZxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKypNuwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pLKb6k8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cl02YP9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Zx1X9YcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4bW7lQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMOdV6yiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21RxbFzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRQIQKw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iygXQnCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORLa9qF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPqUxFSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojEygxguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AkpMJEi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VP9gvNT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQ8RPdEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2M3REixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJMI90azWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPTnhkluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SV1meh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAMbDnw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRZ0N9xPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JE6Xkbh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZO4FaeDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vABI745OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ujester3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MXlHLcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33V4mlsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iWid3t6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooqHZkh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HwzOXrPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGx5a23fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzczW8MbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4ShbdBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koEuNrpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LmSmTDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9dHnKzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxzdkumNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUAHYBIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwnpIEijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJIFqmjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O30KhvzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMBr4N9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xaynWojNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YN4oOvS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeE7S71wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZLYA6UhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3W5iE3VWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFHm4oy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlBa9nhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSDBJw1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVYKsF9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uft8Xx8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPVsjrxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXNbnDzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOZQg3dOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwbSLjcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAl8TmJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJWFSlimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLKV8q2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSq445z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4MCSrIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L09uEr0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZg4Eh0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tiiSpCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgc9hpmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3E1Uk4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5LzFqQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yvg83AQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suEecOcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlhSWGGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQhOkXPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdJH2OYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yBUCYVqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A289Q2hpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YalLKZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxqWP17cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57VF2rweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjXI0QimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXINcL8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cwMuf4TwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiRdIYGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYfCOqGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUhoxQebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VnDJOLMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgqpcfIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7fP4ONhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azhIMkheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PAnmJA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBUK7lmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yi5wiC5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DKKZqXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQKXfylAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xv03LpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vlcEnXpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLLu3xOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wV8EKYLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ibuOxL0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1CBJDtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKaCGtOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbkGtf3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjPSVeslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rM3M1ZweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUydxo0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osmNrMcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQwbX9iVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqX9oyOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUvx6kTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3Bl0lnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsQ74eeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1AqkN3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0uG0xbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpSkNdlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkJ6NTq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dzmuAc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isN3r08pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuam96qGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTC9hhw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrDj5eerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhOKyb7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

