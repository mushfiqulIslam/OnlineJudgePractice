SELECT Person.firstName, Person.lastName, Address.city, Address.state 
from Person 
LEFT JOIN Address ON Person.personId = Address.personId;
