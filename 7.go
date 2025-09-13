//Бедрин и Жирохов
package main

import (
    "fmt"
    "time"
)

func SetGet(settingChan chan int) {
    go func() {
        var newSetting int
        for {
            fmt.Print("Введите новое значение настроек: ")
            if _, err := fmt.Scanln(&newSetting); err != nil {
                fmt.Println("Ошибка ввода:", err)
                continue
            }
            settingChan <- newSetting
        }
    }()
}

func APP(settingChan chan int) {
    currentSettings := 1
    for {
        select {
        case s := <-settingChan:
            currentSettings = s
        default:
            fmt.Printf("Приложение работает на настройках %v\n", currentSettings)
            time.Sleep(time.Second * 1)
        }
        fmt.Print("\033[H\033[2J")
    }
}

func main() {
    settingChan := make(chan int)
    SetGet(settingChan)
    APP(settingChan)
}
