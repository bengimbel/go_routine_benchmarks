## Description

Performance testing go-routines. Also have a custom http package I made to make network requests easier

## Results

### Go-routines

London temperature is 32.100000 degrees farenheit.
Sydney temperature is 54.200000 degrees farenheit.
Rio De Janeiro temperature is 67.900000 degrees farenheit.
Chicago temperature is 23.600000 degrees farenheit.
New York temperature is 7.500000 degrees farenheit.
Tel Aviv temperature is 54.200000 degrees farenheit.
Tokyo temperature is 31.100000 degrees farenheit.
Nairobi temperature is 69.600000 degrees farenheit.
Mexico City temperature is 62.400000 degrees farenheit.

This concurrent operation took: 437.705875ms

### Single Threaded

Chicago temperature is 23.600000 degrees farenheit.
New York temperature is 7.500000 degrees farenheit.
Mexico City temperature is 62.400000 degrees farenheit.
Rio De Janeiro temperature is 67.900000 degrees farenheit.
Tel Aviv temperature is 54.200000 degrees farenheit.
Tokyo temperature is 31.100000 degrees farenheit.
Sydney temperature is 54.200000 degrees farenheit.
London temperature is 32.100000 degrees farenheit.
Nairobi temperature is 69.600000 degrees farenheit.

This single threaded operation took: 1.307422708s
