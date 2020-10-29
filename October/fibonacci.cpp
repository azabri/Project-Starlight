#include "std_lib_facilities.h"

int main () {

int a, c = 0;
int b  = 1;

  while (c < 11) {
    c = a + b;
    cout << c << endl;
    a = b;
    b = c;
    }

return 0;
}
