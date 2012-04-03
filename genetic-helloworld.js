// http://jsfiddle.net/zENKM/

var pop_size = 256; // 2048;
var max_iterations = 512; // 16384;
var elite_rate = 0.10;
var mutation_rate = 0.25;
var mutation = function() {
    return mutation_rate;
}
var target = 'Hello world!';

var struct = function( str, fitness ) {
    return {
        str: str,
        fitness: fitness
    };
}

var init_population = function ( population, buffer ) {
    var size = target.length;

    for ( var i = 0; i < pop_size; i++ ) {
        var citizen = struct('', 0);
        var str = [];

        for ( j = 0; j < size; j++ ) {
            str[j] = String.fromCharCode(Math.floor(Math.random() * 256));

        }

        citizen.str = str.join('');

        population.push(citizen);
        buffer.push(citizen);
    }
}

var calc_fitness = function ( population ) {
    var size = target.length;

    for( var i = 0; i < pop_size; i++ ) {
        var fitness = 0;
        var str = population[i].str;

        for ( j = 0; j < size; j++ ) {
            fitness += Math.abs( str.charCodeAt(j) - target.charCodeAt(j) );
        }

        population[i].fitness = fitness;
    }
}

var fitness_sort = function ( x, y ) {
    return x.fitness - y.fitness;
}

var sort_by_fitness = function ( population ) {
    population.sort(fitness_sort);
}

var elitism = function ( population, buffer, esize )
{
    for ( var i = 0; i < esize; i++ ) {
        buffer[i].str = population[i].str;
        buffer[i].fitness = population[i].fitness;
    }
}

var mutate = function ( member ) {
    var size = target.length;
    var ipos = Math.floor(Math.random() * size);
    var delta = Math.floor(Math.random() * 256);
    var str = member.str.split('');

    str[ipos] = String.fromCharCode((member.str.charCodeAt(ipos) + delta) % 255 );

    member.str = str.join('');
}

var mate = function ( population, buffer ) {
    var esize = pop_size * elite_rate;
    var tsize = target.length;

    elitism ( population, buffer, esize );

    for ( var i = Math.floor(esize); i < pop_size; i++ )
    {
        var i1 = Math.floor(Math.random() * (pop_size / 2));
        var i2 = Math.floor(Math.random() * (pop_size / 2));
        var spos = Math.floor(Math.random() * tsize);

        buffer[i].str = population[i1].str.substr(0, spos) +
            population[i2].str.substr(spos, tsize - spos);

        if ( Math.random() < mutation() ) {
            mutate(buffer[i]);
        }
    }
}

var print_best = function ( population, i ) {
    console.log('Generation ' + i + ' Best: ' + population[0].str + ' (' + population[0].fitness + ')');
}

var swap = function ( population, buffer ) {
    var temp = population;

    population = buffer;
    buffer = temp;
}

var main = function () {
    var population = [];
    var buffer = [];
    var i = 0;

    init_population(population, buffer);

    while ( i < max_iterations ) {

        calc_fitness(population);
        sort_by_fitness(population);
        print_best(population, i);

        if ( population[0].fitness == 0 ) {
            break;
        }

        mate(population, buffer);
        swap(population, buffer);

        i++;
    }
}

main();
