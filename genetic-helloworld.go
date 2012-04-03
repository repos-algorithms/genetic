// http://play.golang.org/p/Xg9f0lM40n 

package main

import (
    "fmt"
    "math"
    "math/rand"
    "sort"
    "time"
)

const GA_POPSIZE int = 2048
const GA_MAXITER int = 16384
const GA_ELITRATE float64 = 0.10
const GA_MUTATIONRATE float64 = 0.25

// ideally GA_MUTATION would be an int32, but that requires truncating,
// which we can't do at initialization time
const GA_MUTATION float64 = math.MaxInt32 * GA_MUTATIONRATE
const GA_TARGET string = "Hello world!"

type ga_struct struct {
    str     []byte
    fitness uint
}

type ga_vector []ga_struct

func init_population(population, buffer *ga_vector) {
    tsize := len(GA_TARGET)

    *population = make(ga_vector, GA_POPSIZE)
    for i := 0; i < GA_POPSIZE; i++ {
        citizen := &(*population)[i]

        citizen.fitness = 0

        citizen.str = make([]byte, tsize)
        for j := 0; j < tsize; j++ {
            citizen.str[j] = byte(rand.Intn(90) + 32)
        }
    }
    *buffer = make(ga_vector, GA_POPSIZE)
}

func calc_fitness(population ga_vector) {
    target := GA_TARGET
    tsize := len(target)

    for i := 0; i < GA_POPSIZE; i++ {
        fitness := uint(0)
        for j := 0; j < tsize; j++ {
            fitness += uint(math.Abs(float64(population[i].str[j] - target[j])))
        }
        population[i].fitness = fitness
    }
}

func (s ga_vector) Len() int {
    return len(s)
}

func (s ga_vector) Less(i, j int) bool {
    return s[i].fitness < s[j].fitness
}

func (s ga_vector) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func sort_by_fitness(population ga_vector) {
    sort.Sort(population)
}

func elitism(population, buffer ga_vector, esize int) {
    copy(buffer[:esize], population[:esize])
}

func mutate(member ga_struct) {
    tsize := len(GA_TARGET)
    ipos := rand.Intn(tsize)
    delta := byte(rand.Intn(90) + 32)

    member.str[ipos] = ((member.str[ipos] + delta) % 122)
}

func mate(population, buffer ga_vector) {
    esize := int(math.Trunc(float64(GA_POPSIZE) * GA_ELITRATE))
    tsize := len(GA_TARGET)

    elitism(population, buffer, esize)

    // Mate the rest
    for i := esize; i < GA_POPSIZE; i++ {
        i1 := rand.Intn(GA_POPSIZE / 2)
        i2 := rand.Intn(GA_POPSIZE / 2)
        spos := rand.Intn(tsize)

        if len(buffer[i].str) < tsize {
            buffer[i].str = make([]byte, tsize)
        }
        copy(buffer[i].str, population[i1].str[:spos])
        copy(buffer[i].str[spos:], population[i2].str[spos:])

        if float64(rand.Int31()) < GA_MUTATION {
            mutate(buffer[i])
        }
    }
}

func print_best(gav ga_vector) {
    fmt.Printf("Best: %s (%d)\n", gav[0].str, gav[0].fitness)
}

func swap(population, buffer *ga_vector) {
    *population, *buffer = *buffer, *population
}

func main() {
    rand.Seed(time.Now().Unix())

    var population, buffer ga_vector
    init_population(&population, &buffer)

    for i := 0; i < GA_MAXITER; i++ {
        calc_fitness(population)
        sort_by_fitness(population)
        print_best(population)

        if population[0].fitness == 0 {
            break
        }

        mate(population, buffer)
        swap(&population, &buffer)
    }
}

