package main

import (
	"driver"
    "fmt"
)


func initMain(){
    driver.ElevInit()
    driver.ElevSetMotorDirection(1)
    var previousFloor = -1
    
    for{
        if (driver.ElevGetFloorSensorSignal()>=0) {
            driver.ElevSetMotorDirection(0)
            previousFloor = driver.ElevGetFloorSensorSignal()
            fmt.Println("I've reached floor" , previousFloor)
            break
        }
    }
}


func main(){
    initMain()    
    
}

