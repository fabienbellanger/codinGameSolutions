package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math"
import "strconv"

 type defibrilatorType struct {
    ID          int
    Name        string
    Address     string
    PhoneNumber string
    Longitude   float64
    Latitude    float64
 }

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var longitudeStrFrom string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&longitudeStrFrom)
    longitudeFrom := strToFloat(longitudeStrFrom)
    
    var latitudeStrFrom string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&latitudeStrFrom)
    latitudeFrom := strToFloat(latitudeStrFrom)
    
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    var nearestDefibrilator defibrilatorType
    distanceMin := 1000000.0

    for i := 0; i < N; i++ {
        scanner.Scan()
        data := scanner.Text()

        array := strings.Split(data, ";")
        if len(array) == 6 {
            id, _ := strconv.Atoi(array[0])

            defibrilator := defibrilatorType{
                ID:          id,
                Name:        array[1],
                Address:     array[2],
                PhoneNumber: array[3],
                Longitude:   strToFloat(array[4]),
                Latitude:    strToFloat(array[5]),
            }

            distance := getDistance(longitudeFrom, latitudeFrom, defibrilator.Longitude, defibrilator.Latitude)

            if distance < distanceMin {
                distanceMin = distance
                nearestDefibrilator = defibrilator
            }
        }
    }
    
    fmt.Println(nearestDefibrilator.Name)
}

func degreeToRadian(d float64) float64 {
    return d * math.Pi / 180
}

func strToFloat(s string) float64 {
    f, _ := strconv.ParseFloat(strings.Replace(s, ",", ".", -1), 64)

	return f
}

func getDistance(longitudeFrom, latitudeFrom, longitudeTo, latitudeTo float64) (d float64) {
    longitudeFrom = degreeToRadian(longitudeFrom)
    latitudeFrom = degreeToRadian(latitudeFrom)
    longitudeTo = degreeToRadian(longitudeTo)
    latitudeTo = degreeToRadian(latitudeTo)

    x := (longitudeTo - longitudeFrom) * math.Cos((latitudeFrom + latitudeTo) / 2)
    y := latitudeTo - latitudeFrom

    return 6371 * math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
