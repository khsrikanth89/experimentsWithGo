package main

import (
    "fmt"
    "sort"
    "time"
)

type data struct {
    id   	string
    timesent    time.Time
    tbool       bool
}

type timeSlice []data

func (p timeSlice) Len() int {
    return len(p)
}

func (p timeSlice) Less(i, j int) bool {
    return p[i].timesent.Before(p[j].timesent)
}

func (p timeSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func main() {
    var data_map = make(map[string]data)
    data_map["1"] = data{timesent: time.Now().Add(12 * time.Hour)}
    data_map["2"] = data{timesent: time.Now()}
    data_map["3"] = data{timesent: time.Now().Add(24 * time.Hour)}
    
    date_sorted := make(timeSlice, 0)
    for _, d := range data_map {
        date_sorted = append(date_sorted, d)
    }
 
fmt.Println(date_sorted)
   
 sort.Sort(date_sorted)
    
fmt.Println(date_sorted)
}
