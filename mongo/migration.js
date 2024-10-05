db = db.getSiblingDB('mydb');

db.createCollection("users");

db.users.insertMany([
     { name: "John", age: 25 },
     { name: "Jane", age: 22 },
     { name: "Doe", age: 30 }
    ]);