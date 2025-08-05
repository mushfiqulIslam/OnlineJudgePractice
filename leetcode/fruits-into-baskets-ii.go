func numOfUnplacedFruits(fruits []int, baskets []int) int {
    placedBaskets := make([]bool, len(baskets))
    notPlacedCount := 0
    for _, fCount := range fruits {
        placed := false 
        for index, capacity := range baskets {
            if !placedBaskets[index] && capacity >= fCount {
                placedBaskets[index] = true
                placed = true
                break
            } 
        }

        if !placed {
            notPlacedCount += 1
        }
    }
    return notPlacedCount
}
