# Line Drawing :: Midpoint Algorithm

## Links

1. https://en.wikipedia.org/wiki/Linear_equation
2. https://en.wikipedia.org/wiki/Bresenham

## Exercise

Implement a program that contains `DrawLine(*ebiten.Image, x1, y1, x2, y2 int, c color.Color)`
function that draws (x1, y1), (x2, y2) lines using the midpoint algorithm.

Your implementation must use the floating point arithmetics and utilize the equation `Ax+Bx+C=0`
without any integer optimizations. The goal of this exercise is to let you understand the equation
and maths behind the midpoint algorithms and how it affects the output depending on the slope of
the line you attempt to draw.
