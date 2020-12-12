
# Create an API with golang and MongoDB
### Appointy 
==============================

In this task , request is sent by http, and fetching is done with json according to request. This repository also contain main.go file,exe file, screen shot to support it. For storation, mongoDB is used which store user and contact details . Timestamp is stored at the time of user creation.Postman is used to see the immediate effect of request as shown in screen shot.
### 1 . Create a user :-
In this, when http://localhost:8000/users link is send in post form , the golang store the json in mongodb and also represent log of that json. It conatain all required attribute. Time stamp is auto generated  at the user creation time. 
required format : 
{ "Name":"some",
"DOB":"23/1/2001",
"Email":"some@gamil.com",
"Phone": "9138938940"
}

### 2 . Fetch all and Fetch perticular ID :- 
In this,when http://localhost:8000/users/4563 link is sent in get form then it return details of particular user ID.

In this,when http://localhost:8000/users link is sent in get form then it return details of all the user IDs which are stored.

### 3 . Create contact details of user :-
In this, when http://localhost:8000/contacts link is sent in post form , the golang store the json in mongodb and also represent log of that json. It conatain all required attributes.Time stamp is auto generated at the user creation time.   
required format:
{
"UserIdTwo": "35",
"UserIdOne": "56"
}

### 4 . Fetch data of 14 Days with particular user ID and infection time :-
In this ,when ( http://localhost:8000/contacts/User-ID&Infection timeStamp ) link is sent in get form, this retrive data of the particular user with given ID in 14 days prior to the given infection time. Time stamp will be given in particular ISO formate " YYYYMMDDHHMMSS " .

example : http://localhost:8000/contacts/87&2020121182490
