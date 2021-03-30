GO library for nipo

# sample

    package main

    import (
    	"nipo"
    	"fmt"
    )

    func main() {
        config := nipo.CreateConfig("TOKEN", "IP of SERVER", "PORT")
        SetResult,Setok := nipo.Set(config, "KEY", "VALUE")
        if !Setok {
            fmt.Println("Error at set")    
        } else {
            fmt.Println("Set OK")
            fmt.Println(SetResult)
        }

        GetResult,Getok := nipo.get(config, "KEY")
        if !Getok {
            fmt.Println("Error at get")    
        } else {
            fmt.Println("get OK")
            fmt.Println(GetResult)
        }
    }
    