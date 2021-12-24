Quection 1.

a. Create a func with the identifier "foo" that returns an int

b.  Create a func with the identifier "bar" that returns an int and a string

c. Call both funcs

d. Print out their results



Question 2.
a. Create a func with the identifier "foo" that;
   * takes in a variadic parameter of type int
   * pass in a value of tyoe []int into your func (unfurl the []int)
   * return the sum of all values of type int passed in

b. Create a func with the identifier "bar" that;
   * takes in a parameter of type []int
   * return the sum of all values of type int passed in



Question 3.
Use the "defer" keyword to show that a deferred func runs after the func containing it exits.



Question 4.
a. Create a user defined struct with
   * the identifier "person"
   * the fields:
     first
     last
     age

b. Attach a method of type person with 
   * the identifier "speak"
   * the method should have the person say their name and age

c. Create a value of type person

d. call the method from the value of type person



Question 5.
a. Create a type SQUARE

b. Create a type CIRCLE

c. Attach a method to each that calculates AREA and return it
   * Circle area = pie * r2
   * Square area = l * b

d. Create a type SHAPE which dwfines an interface as anything which has the AREA method

e. Create a func INFO which takes type shape and then prints the area

f. Create a value of type square

g. Create a value of type circle

h. Use func info to print the area of square

i. Use func info to print the area of circle



Question 6.
Build and use an anonymous func



Question 7.
Assign a func to a variable and call that func



Question 8.
a. Create a func which returns a func

b. Assign the return func to a variable

c. Call the returned func



Question 9.
A "callback" is when we pass a func as an arguement. For this exercise,
* Pass a func into a func as an arguement.



Question 10.
Closure is when we have "enclosed" the scope of a variable in some code block.

For this hands-on exercise, create a func which "encloses" the scope of a variable.