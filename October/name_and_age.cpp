#include "std_lib_facilities.h"

int main()
{
	cout << "Please enter your first name, last name and age!\n";

	string first_name, last_name;
	int age;

	cin >> first_name >> last_name >> age;

	cout << "Hello, " << first_name << " " << last_name << " ! (age: " << age << ")\n";

	return 0;
}
