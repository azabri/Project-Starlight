#include "std_lib_facilities.h"

//vektor
const vector<string> months =
{
	"January"
	"February"
	"March"
	"April"
	"May"
	"June"
	"July"
	"August"
	"September"
	"October"
	"November"
	"Dezember"
};

enum class Month {
	jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec
};

Month operator++ (Month& m)
{
	m = (m == Month::dec) ? Month::jan : Month(int(m) + 1);
	//ha m = december, akkor legyen január, egyébként + 1
	return m;
}

ostream& operator<< (ostream& os, Month m)
//ez az operátor: ' << ' hogyan működjön, ha bal oldalán ostream referencia van pl. cout, a jobb oldalán pedig a Month
{
	return os << months[int(m)];
	//months aktuális elemét írja ki
}

/*
ostream& operator<< (ostream& os, int i)
{
	//ezt már a C++ fejlesztői megírták, hogy mit csináljon az int i -vel, tehát ha ilyet akarunk kiíratni, az már működik, nekünk nem kell definiálni
}
*/

int main()
{
	
	Month m = Month::jan;
	m = Month::apr;

	++m;

	cout << m << endl;
	//bal érték egy cout operátor, a jobb érték egy month

	return 0;

}
