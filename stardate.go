package main

import (
"os"
"time"
"fmt"
)

func main(){
    if os.Args[1] == "now" {
        sd_origin := time.Date(1987, time.July, 15, 0, 0, 0, 0, time.UTC)
        dif := time.Since(sd_origin).Seconds() * 1000
        sd := (dif / (1000 * 60 * 60 * 24 * 0.036525)) + 410000
        sd_fin := float64(sd) / 10.0
        fmt.Printf("%6.2f\n", sd_fin)
        }
}

//var now = new Date();
//var then = new Date("July 15, 1987");
//var stardate = now.getTime() - then.getTime();
//document.write(stardate);
