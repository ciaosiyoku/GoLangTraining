Quection 1.
Create your own type "person" which will have an underlying type of "struct" so that it can store the following data:

a. first name
b. last name
c. favorite ice cream flavors

Create two VALUES of TYPE person.

Print out the values, ranging over the elements in the slice which stores the favorite flavors.



Question 2.
Take the code from the previous exercise, then store the values of the type person in a map with the key of last name.

Access each value in the map.

Print out the values, ranging over the slice.



Question 3.
a. Create a new type vehicle.
   * The underlying type is struct.
   * The fields are:
     ** doors
     ** color

b. Create two new types: truck and sedan.
   * The underlying type of each of these new types is a struct.
   * Embed the "vehicle" type in both truck and sedan.
   * Give truck the field "fourWheel" which will be set to bool.
   * Give sedan the field "luxury" which will be set bool.

c. Using the vehicle, truck and sedan structs:
   * using a composite literal, create a value of type truck and assign values to the fields;
   * using a composite literal, create a value of type sedan and assign values to the fields;

d. Print out each of these values.

e. Print out a single field from each of these values.



Question 4.
Create and use an anonymous struct.