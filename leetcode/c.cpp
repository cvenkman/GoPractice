#include <stdio.h>

int main() {
	int a, b, c, d;
	a= 0;
	b= 1;
	c=2;
	d=3;
	printf("%d,%d,%d,%d\n", a,b,c,d);
	a, b = c, d;
	printf("%d,%d,%d,%d", a,b,c,d);
}